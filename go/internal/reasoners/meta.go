package reasoners

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/Agent-Field/agentfield/sdk/go/harness"

	"github.com/Agent-Field/pr-af/go/internal/harnessx"
	"github.com/Agent-Field/pr-af/go/internal/prompts"
	"github.com/Agent-Field/pr-af/go/internal/schemas"
)

// The three meta-dimension selectors (meta_semantic / meta_mechanical /
// meta_systemic) share one implementation parameterized by lens: build the
// shared context, spill it to <repo>/.pr-af-context/meta_<lens>_context.json
// when it exceeds 8000 characters (code points), run the lens prompt, then
// FORCE the lens field regardless of what the model returned.
//
// Output keys (§B.2): lens, dimensions, confidence, rationale. Meta selectors
// use incremental schema mode and fail closed when structured output cannot be
// produced; an empty seeded fallback can otherwise become a false "Looks Good".

// MetaSemantic ports meta_semantic.
func MetaSemantic(ctx context.Context, deps Deps, in MetaInput) (map[string]any, error) {
	return runMetaLens(ctx, deps, in, "semantic", prompts.MetaSemanticPrompt)
}

// MetaMechanical ports meta_mechanical.
func MetaMechanical(ctx context.Context, deps Deps, in MetaInput) (map[string]any, error) {
	return runMetaLens(ctx, deps, in, "mechanical", prompts.MetaMechanicalPrompt)
}

// MetaSystemic ports meta_systemic.
func MetaSystemic(ctx context.Context, deps Deps, in MetaInput) (map[string]any, error) {
	return runMetaLens(ctx, deps, in, "systemic", prompts.MetaSystemicPrompt)
}

type metaDraftDimension struct {
	Name         string   `json:"name"`
	ReviewPrompt string   `json:"review_prompt"`
	TargetFiles  []string `json:"target_files"`
}

type metaDraftResult struct {
	Lens       string               `json:"lens"`
	Dimensions []metaDraftDimension `json:"dimensions"`
	Confidence float64              `json:"confidence"`
	Rationale  string               `json:"rationale"`
}

func runMetaLens(
	ctx context.Context,
	deps Deps,
	in MetaInput,
	lens string,
	buildPrompt func(context, repoPath, depth string) string,
) (map[string]any, error) {
	metaContext := prompts.MetaContext(in.Intake, in.Anatomy, []prompts.StrPair(in.DiffPatches), in.ReviewerFeedback)
	// The builder embeds a file reference under the same condition; the write
	// itself is this reasoner's job (Python _write_context_file).
	if in.RepoPath != "" && utf8.RuneCountInString(metaContext) > 8000 {
		if _, err := writeContextFile(metaContext, "meta_"+lens+"_context.json", in.RepoPath); err != nil {
			return nil, err
		}
	}

	prompt := buildPrompt(metaContext, in.RepoPath, in.Depth)
	parsed, hres, err := harnessx.Run[metaDraftResult](ctx, deps.Harness, prompt, harness.Options{
		Cwd:              in.RepoPath,
		SchemaMode:       "incremental",
		SchemaMaxRetries: 2,
	})
	if err != nil {
		return nil, err
	}

	changedFiles := metaChangedFiles(in.DiffPatches)
	if parsed != nil && hres != nil && hres.Parsed != nil && !hres.IsError {
		if result, err := normalizeMetaDraft(lens, *parsed, changedFiles); err == nil {
			return dumpMap(result)
		}
	}

	// Some models ignore the schema/file protocol but still emit usable JSON in
	// normal text (possibly fenced or wrapped in prose). Recover that response
	// deterministically before spending another model call.
	if hres != nil && strings.TrimSpace(hres.Result) != "" {
		if draft, ok := decodeMetaDraft(hres.Result); ok {
			if result, err := normalizeMetaDraft(lens, draft, changedFiles); err == nil {
				return dumpMap(result)
			}
		}
	}

	// Last compatibility rung for models that are capable of reasoning but are
	// poor at JSON/schema adherence: one unstructured call with a deliberately
	// simple line protocol. Operational policy (ids, priority, budgets, context
	// files) is still supplied by PR-AF, not by the model.
	fallbackPrompt := prompt + `

STRUCTURED OUTPUT RECOVERY MODE
The previous structured-output attempt could not be validated.
Do NOT return JSON. Do NOT write files. Return only plain text in this format:
DIMENSION: <short human-readable name>
PROMPT: <specific review task>
FILES: <comma-separated changed file paths, or * for all changed files>
END
Repeat DIMENSION..END for each review dimension.
Then optionally add:
CONFIDENCE: <number from 0 to 1>
RATIONALE: <one concise line>
`
	raw, rawErr := deps.Harness.Harness(ctx, fallbackPrompt, nil, nil, harness.Options{
		Cwd:      in.RepoPath,
		MaxTurns: 2,
	})
	if rawErr != nil {
		return nil, rawErr
	}
	if raw != nil {
		if draft, ok := parseMetaLineProtocol(raw.Result, changedFiles); ok {
			if result, err := normalizeMetaDraft(lens, draft, changedFiles); err == nil {
				return dumpMap(result)
			}
		}
	}

	if hres == nil {
		return nil, fmt.Errorf("meta %s harness returned no structured result and text recovery failed", lens)
	}
	return nil, fmt.Errorf(
		"meta %s output recovery failed: failure_type=%s model=%q is_error=%t message=%q",
		lens,
		hres.FailureType,
		hres.Model,
		hres.IsError,
		hres.ErrorMessage,
	)
}

func metaChangedFiles(patches OrderedPatches) []string {
	files := make([]string, 0, len(patches))
	seen := map[string]struct{}{}
	for _, p := range patches {
		name := strings.TrimSpace(p.Key)
		if name == "" {
			continue
		}
		if _, ok := seen[name]; ok {
			continue
		}
		seen[name] = struct{}{}
		files = append(files, name)
	}
	return files
}

func normalizeMetaDraft(lens string, draft metaDraftResult, changedFiles []string) (schemas.MetaDimensionResult, error) {
	if len(draft.Dimensions) == 0 {
		return schemas.MetaDimensionResult{}, fmt.Errorf("meta %s returned no review dimensions", lens)
	}
	if draft.Confidence < 0 || draft.Confidence > 1 {
		return schemas.MetaDimensionResult{}, fmt.Errorf("meta %s confidence %.3f outside [0,1]", lens, draft.Confidence)
	}

	result := schemas.MetaDimensionResult{
		Lens:       lens,
		Dimensions: make([]schemas.ReviewDimension, 0, len(draft.Dimensions)),
		Confidence: draft.Confidence,
		Rationale:  strings.TrimSpace(draft.Rationale),
	}
	if result.Rationale == "" {
		result.Rationale = "Recovered from a model response that did not satisfy the structured-output protocol."
	}

	for i, d := range draft.Dimensions {
		name := strings.TrimSpace(d.Name)
		reviewPrompt := strings.TrimSpace(d.ReviewPrompt)
		if name == "" || reviewPrompt == "" {
			continue
		}
		files := compactStrings(d.TargetFiles)
		if len(files) == 0 || (len(files) == 1 && files[0] == "*") {
			files = append([]string(nil), changedFiles...)
		}

		var dim schemas.ReviewDimension
		_ = json.Unmarshal([]byte(`{}`), &dim) // apply schema-owned defaults
		dim.ID = fmt.Sprintf("%s-%02d", lens, i+1)
		dim.Name = name
		dim.ReviewPrompt = reviewPrompt
		dim.TargetFiles = files
		dim.ContextFiles = orEmptyStrs(dim.ContextFiles)
		result.Dimensions = append(result.Dimensions, dim)
	}
	if len(result.Dimensions) == 0 {
		return schemas.MetaDimensionResult{}, fmt.Errorf("meta %s returned no usable review dimensions", lens)
	}
	return result, nil
}

func compactStrings(values []string) []string {
	out := make([]string, 0, len(values))
	seen := map[string]struct{}{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	return out
}

func decodeMetaDraft(text string) (metaDraftResult, bool) {
	text = strings.TrimSpace(text)
	if text == "" {
		return metaDraftResult{}, false
	}
	var draft metaDraftResult
	if json.Unmarshal([]byte(text), &draft) == nil && len(draft.Dimensions) > 0 {
		return draft, true
	}
	for _, candidate := range balancedJSONObjects(text) {
		draft = metaDraftResult{}
		if json.Unmarshal([]byte(candidate), &draft) == nil && len(draft.Dimensions) > 0 {
			return draft, true
		}
		var wrapper map[string]json.RawMessage
		if json.Unmarshal([]byte(candidate), &wrapper) == nil {
			for _, key := range []string{"result", "output", "content", "response"} {
				raw, ok := wrapper[key]
				if !ok {
					continue
				}
				var nested string
				if json.Unmarshal(raw, &nested) == nil {
					if recovered, ok := decodeMetaDraft(nested); ok {
						return recovered, true
					}
				}
			}
		}
	}
	return metaDraftResult{}, false
}

func balancedJSONObjects(text string) []string {
	objects := []string{}
	depth, start := 0, -1
	inString, escaped := false, false
	for i, r := range text {
		if inString {
			if escaped {
				escaped = false
				continue
			}
			if r == '\\' {
				escaped = true
				continue
			}
			if r == '"' {
				inString = false
			}
			continue
		}
		if r == '"' {
			inString = true
			continue
		}
		switch r {
		case '{':
			if depth == 0 {
				start = i
			}
			depth++
		case '}':
			if depth == 0 {
				continue
			}
			depth--
			if depth == 0 && start >= 0 {
				objects = append(objects, text[start:i+1])
				start = -1
			}
		}
	}
	return objects
}

func parseMetaLineProtocol(text string, changedFiles []string) (metaDraftResult, bool) {
	var result metaDraftResult
	var current *metaDraftDimension
	field := ""
	flush := func() {
		if current == nil {
			return
		}
		if strings.TrimSpace(current.Name) != "" && strings.TrimSpace(current.ReviewPrompt) != "" {
			result.Dimensions = append(result.Dimensions, *current)
		}
		current = nil
		field = ""
	}
	for _, rawLine := range strings.Split(text, "\n") {
		line := strings.TrimSpace(rawLine)
		line = strings.TrimSpace(strings.TrimLeft(line, "-*#> "))
		if line == "" {
			continue
		}
		upper := strings.ToUpper(line)
		switch {
		case strings.HasPrefix(upper, "DIMENSION:"):
			flush()
			current = &metaDraftDimension{Name: strings.TrimSpace(line[len("DIMENSION:"):])}
			field = "name"
		case strings.HasPrefix(upper, "NAME:"):
			if current == nil {
				current = &metaDraftDimension{}
			}
			current.Name = strings.TrimSpace(line[len("NAME:"):])
			field = "name"
		case strings.HasPrefix(upper, "PROMPT:"):
			if current == nil {
				current = &metaDraftDimension{}
			}
			current.ReviewPrompt = strings.TrimSpace(line[len("PROMPT:"):])
			field = "prompt"
		case strings.HasPrefix(upper, "REVIEW_PROMPT:"):
			if current == nil {
				current = &metaDraftDimension{}
			}
			current.ReviewPrompt = strings.TrimSpace(line[len("REVIEW_PROMPT:"):])
			field = "prompt"
		case strings.HasPrefix(upper, "FILES:"):
			if current == nil {
				current = &metaDraftDimension{}
			}
			value := strings.TrimSpace(line[len("FILES:"):])
			if value == "*" {
				current.TargetFiles = append([]string(nil), changedFiles...)
			} else {
				current.TargetFiles = compactStrings(strings.Split(value, ","))
			}
			field = "files"
		case strings.HasPrefix(upper, "TARGET_FILES:"):
			if current == nil {
				current = &metaDraftDimension{}
			}
			value := strings.TrimSpace(line[len("TARGET_FILES:"):])
			current.TargetFiles = compactStrings(strings.Split(value, ","))
			field = "files"
		case strings.HasPrefix(upper, "CONFIDENCE:"):
			value := strings.TrimSpace(line[len("CONFIDENCE:"):])
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				result.Confidence = v
			}
			field = ""
		case strings.HasPrefix(upper, "RATIONALE:"):
			result.Rationale = strings.TrimSpace(line[len("RATIONALE:"):])
			field = "rationale"
		case upper == "END":
			flush()
		default:
			if current != nil && field == "prompt" {
				current.ReviewPrompt = strings.TrimSpace(current.ReviewPrompt + " " + line)
			} else if field == "rationale" {
				result.Rationale = strings.TrimSpace(result.Rationale + " " + line)
			}
		}
	}
	flush()
	return result, len(result.Dimensions) > 0
}
