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
	parsed, hres, err := harnessx.Run[schemas.MetaDimensionResult](ctx, deps.Harness, prompt, harness.Options{
		Cwd:              in.RepoPath,
		SchemaMode:       "incremental",
		SchemaMaxRetries: 4,
	})
	if err != nil {
		return nil, err
	}
	if hres == nil || hres.Parsed == nil || hres.IsError {
		if hres == nil {
			return nil, fmt.Errorf("meta %s harness returned no result", lens)
		}
		return nil, fmt.Errorf(
			"meta %s structured output unavailable: failure_type=%s model=%q is_error=%t message=%q",
			lens,
			hres.FailureType,
			hres.Model,
			hres.IsError,
			hres.ErrorMessage,
		)
	}
	result := *parsed
	// Python forces the lens on both the parsed and the fallback result.
	result.Lens = lens
	if result.Dimensions == nil {
		result.Dimensions = []schemas.ReviewDimension{}
	}
	for i := range result.Dimensions {
		result.Dimensions[i].TargetFiles = orEmptyStrs(result.Dimensions[i].TargetFiles)
		result.Dimensions[i].ContextFiles = orEmptyStrs(result.Dimensions[i].ContextFiles)
	}
	return dumpMap(result)
}
