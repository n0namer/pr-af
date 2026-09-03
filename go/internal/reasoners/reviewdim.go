package reasoners

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/Agent-Field/agentfield/sdk/go/harness"

	"github.com/Agent-Field/pr-af/go/internal/harnessx"
	"github.com/Agent-Field/pr-af/go/internal/prompts"
	"github.com/Agent-Field/pr-af/go/internal/schemas"
)

// ReviewDimension ports review_dimension: one focused reviewer over an
// assigned dimension, with optional sub-review spawning below max depth.
//
// Output keys (§B.2): findings, sub_reviews, current_depth. A schema parse
// failure logs loudly (indistinguishable from a clean review otherwise) and
// reports zero findings — Python prints and substitutes an empty
// _ReviewFindingsResult.
func ReviewDimension(ctx context.Context, deps Deps, in ReviewDimensionInput) (map[string]any, error) {
	canSpawn := in.CurrentDepth < in.MaxDepth

	// The prompt builder embeds file references for oversized diff/primed
	// sections under these exact conditions; the writes are the reasoner's job.
	if len(in.DiffPatches) > 0 {
		var parts []string
		for _, path := range in.TargetFiles {
			if patch, ok := in.DiffPatches[path]; ok && patch != "" {
				parts = append(parts, "### "+path+"\n```diff\n"+patch+"\n```")
			}
		}
		if len(parts) > 0 {
			patchesText := strings.Join(parts, "\n\n")
			if in.RepoPath != "" && utf8.RuneCountInString(patchesText) > 6000 {
				if _, err := writeContextFile(patchesText, "review_dimension_diff_patches.md", in.RepoPath); err != nil {
					return nil, err
				}
			}
		}
	}
	if in.PrimedCode != "" && in.RepoPath != "" && utf8.RuneCountInString(in.PrimedCode) > 6000 {
		if _, err := writeContextFile(in.PrimedCode, "review_dimension_primed_code.md", in.RepoPath); err != nil {
			return nil, err
		}
	}

	prompt := prompts.ReviewDimensionPrompt(prompts.ReviewDimensionOptions{
		ReviewPrompt:      in.ReviewPrompt,
		TargetFiles:       in.TargetFiles,
		ContextFiles:      in.ContextFiles,
		RepoPath:          in.RepoPath,
		CurrentDepth:      in.CurrentDepth,
		MaxDepth:          in.MaxDepth,
		PrNarrative:       in.PrNarrative,
		RiskSurfaces:      in.RiskSurfaces,
		IntakeSummary:     in.IntakeSummary,
		PrDescription:     in.PrDescription,
		DiffPatches:       in.DiffPatches,
		AllDimensionNames: in.AllDimensionNames,
		ReviewerFeedback:  in.ReviewerFeedback,
		PrimedCode:        in.PrimedCode,
	})
	deltaHints := ""
	if len(in.DiffPatches) > 0 {
		prompt += `

## Proposed-Diff Verification Rule
For changed lines, the supplied diff is the authoritative proposed program state. The repository checkout or primed code may still show the pre-change state and MUST NOT be used to dismiss an added-line regression.
Before returning no findings:
1. Compare every relevant removed expression/guard with its added replacement as OLD versus NEW semantics.
2. For boolean conditions, validation guards, boundary checks, or branch predicates, evaluate representative truth cases (including mixed/partial states) and verify which branch each version takes.
3. Check callers/consumers against the NEW behavior, not only the current checkout.
4. If OLD and NEW are not behaviorally equivalent, explain why the change is safe; otherwise report an evidence-grounded finding.
Do this analysis internally and return only the required review result schema.`
		deltaHints = semanticDeltaHints(in.TargetFiles, in.DiffPatches)
		if deltaHints != "" {
			prompt += "\n\n## Deterministic Semantic-Delta Hints\n" + deltaHints
		}
	}

	parsed, res, err := harnessx.Run[reviewFindingsResult](ctx, deps.Harness, prompt, harness.Options{Cwd: in.RepoPath})
	if err != nil {
		return nil, err
	}
	schemaParseFailed := res == nil || res.Parsed == nil
	if schemaParseFailed {
		// Schema parse failed entirely — don't silently report "0 findings",
		// which is indistinguishable from a clean review. Make it visible.
		errMsg := "None"
		if res != nil && res.ErrorMessage != "" {
			errMsg = "'" + res.ErrorMessage + "'"
		}
		fmt.Printf(
			"[PR-AF] review_dimension: schema parse failed — treating as 0 findings for this dimension (error=%s)\n",
			errMsg,
		)
	}
	result := *parsed
	if schemaParseFailed {
		result.Findings = nil
		result.SubReviews = nil
	}
	if deltaHints != "" && len(result.Findings) == 0 {
		recovered, err := verifySemanticDelta(ctx, deps.Harness, in, deltaHints)
		if err != nil {
			return nil, fmt.Errorf("review_dimension semantic-delta verification failed: %w", err)
		}
		result.Findings = recovered
		// A successful SAFE/FINDING verifier response is a valid alternate
		// protocol, so the dimension is no longer a schema-parse failure.
		schemaParseFailed = false
	}
	for i := range result.Findings {
		result.Findings[i].Tags = orEmptyStrs(result.Findings[i].Tags)
	}

	subReviewDicts := []any{}
	if canSpawn && len(result.SubReviews) > 0 {
		// Python: `for sr in parsed.sub_reviews[:2] if sr.review_prompt and
		// sr.target_files` — slice first, then filter.
		for _, sr := range capN(result.SubReviews, 2) {
			if sr.ReviewPrompt == "" || len(sr.TargetFiles) == 0 {
				continue
			}
			subReviewDicts = append(subReviewDicts, map[string]any{
				"reason":        sr.Reason,
				"review_prompt": sr.ReviewPrompt,
				"target_files":  orEmptyStrs(sr.TargetFiles),
				"context_files": orEmptyStrs(sr.ContextFiles),
				"priority":      sr.Priority,
			})
		}
	}

	findings, err := dumpSlice(result.Findings)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"findings":            findings,
		"sub_reviews":         subReviewDicts,
		"current_depth":       in.CurrentDepth,
		"schema_parse_failed": schemaParseFailed,
	}, nil
}

func semanticDeltaHints(targetFiles []string, patches map[string]string) string {
	var hints []string
	for _, path := range targetFiles {
		patch := patches[path]
		if patch == "" {
			continue
		}
		lines := strings.Split(patch, "\n")
		for i := 0; i < len(lines); i++ {
			if !strings.HasPrefix(lines[i], "-") || strings.HasPrefix(lines[i], "---") {
				continue
			}
			oldLine := strings.TrimSpace(strings.TrimPrefix(lines[i], "-"))
			for j := i + 1; j < len(lines); j++ {
				if strings.HasPrefix(lines[j], "-") && !strings.HasPrefix(lines[j], "---") {
					break
				}
				if !strings.HasPrefix(lines[j], "+") || strings.HasPrefix(lines[j], "+++") {
					continue
				}
				newLine := strings.TrimSpace(strings.TrimPrefix(lines[j], "+"))
				if hint := operatorDeltaHint(path, oldLine, newLine); hint != "" {
					hints = append(hints, hint)
				}
				break
			}
		}
	}
	return strings.Join(hints, "\n\n")
}

func operatorDeltaHint(path, oldLine, newLine string) string {
	type opChange struct {
		oldOp string
		newOp string
		cases string
	}
	changes := []opChange{
		{oldOp: "!=", newOp: "==", cases: "boolean operands: (F,F) OLD=false NEW=true; (F,T) OLD=true NEW=false; (T,F) OLD=true NEW=false; (T,T) OLD=false NEW=true"},
		{oldOp: "==", newOp: "!=", cases: "boolean operands: (F,F) OLD=true NEW=false; (F,T) OLD=false NEW=true; (T,F) OLD=false NEW=true; (T,T) OLD=true NEW=false"},
		{oldOp: "&&", newOp: "||", cases: "boolean operands: (F,F) OLD=false NEW=false; (F,T) OLD=false NEW=true; (T,F) OLD=false NEW=true; (T,T) OLD=true NEW=true"},
		{oldOp: "||", newOp: "&&", cases: "boolean operands: (F,F) OLD=false NEW=false; (F,T) OLD=true NEW=false; (T,F) OLD=true NEW=false; (T,T) OLD=true NEW=true"},
		{oldOp: "<=", newOp: "<", cases: "boundary case lhs==rhs: OLD=true NEW=false"},
		{oldOp: "<", newOp: "<=", cases: "boundary case lhs==rhs: OLD=false NEW=true"},
		{oldOp: ">=", newOp: ">", cases: "boundary case lhs==rhs: OLD=true NEW=false"},
		{oldOp: ">", newOp: ">=", cases: "boundary case lhs==rhs: OLD=false NEW=true"},
	}
	for _, change := range changes {
		if strings.Contains(oldLine, change.oldOp) && strings.Contains(newLine, change.newOp) &&
			!strings.Contains(oldLine, change.newOp) && !strings.Contains(newLine, change.oldOp) {
			return fmt.Sprintf("File: %s\nOLD: %s\nNEW: %s\nDetected operator change: %s -> %s\nRepresentative cases: %s", path, oldLine, newLine, change.oldOp, change.newOp, change.cases)
		}
	}
	return ""
}
