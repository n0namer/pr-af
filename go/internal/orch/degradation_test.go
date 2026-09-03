package orch

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/Agent-Field/pr-af/go/internal/config"
	"github.com/Agent-Field/pr-af/go/internal/reasoners"
	"github.com/Agent-Field/pr-af/go/internal/schemas"
)

func degradationOrchestrator(t *testing.T) *Orchestrator {
	t.Helper()
	cfg := config.DefaultReviewConfig()
	cfg.Comments.PolishEnabled = false
	o := New(Deps{App: &fakeApp{}}, schemas.ReviewInput{DryRun: true}, cfg)
	o.prData = &schemas.GitHubPRData{}
	o.runIntakeFn = func(context.Context) (schemas.IntakeResult, error) {
		return schemas.IntakeResult{PrSummary: "summary", PrType: "feature", Complexity: "low"}, nil
	}
	o.runAnatomyFn = func(context.Context, schemas.IntakeResult) (schemas.AnatomyResult, error) {
		return schemas.AnatomyResult{}, nil
	}
	o.resolveDepthFn = func(schemas.IntakeResult) string { return "standard" }
	o.cleanupFn = func() {}
	return o
}

func degradationPlan(n int) schemas.ReviewPlan {
	dims := make([]schemas.ReviewDimension, n)
	for i := range dims {
		dims[i] = schemas.ReviewDimension{ID: string(rune('a' + i)), Name: "dimension", ReviewPrompt: "review", TargetFiles: []string{"a.go"}}
	}
	return schemas.ReviewPlan{Dimensions: dims}
}

func TestRunFailsBeforeOutputWhenAllDimensionsFailSchemaParsing(t *testing.T) {
	o := degradationOrchestrator(t)
	var outputCalls atomic.Int32
	o.rfns.reviewDim = func(context.Context, reasoners.Deps, reasoners.ReviewDimensionInput) (map[string]any, error) {
		return map[string]any{"schema_parse_failed": true}, nil
	}
	o.runReviewPhasesFn = func(ctx context.Context, _ schemas.IntakeResult, _ schemas.AnatomyResult, _, feedback string) (schemas.ReviewPlan, []schemas.ScoredFinding, error) {
		plan := degradationPlan(2)
		_, err := o.collectParallelReview(ctx, plan, feedback)
		return plan, nil, err
	}
	o.generateOutputFn = func(context.Context, []schemas.ScoredFinding, schemas.IntakeResult, schemas.AnatomyResult, schemas.ReviewPlan, bool) (schemas.ReviewResult, error) {
		outputCalls.Add(1)
		return schemas.ReviewResult{}, nil
	}

	_, err := o.Run(context.Background())
	if err == nil || !strings.Contains(err.Error(), "all review dimensions failed schema parsing (failed dimensions: 2)") {
		t.Fatal("Run error = %v", err)
	}
	if outputCalls.Load() != 0 {
		t.Fatalf("generateOutput calls = %d, want 0", outputCalls.Load())
	}
}

func TestCoverageGapReviewerFailurePreservesPrimaryFindings(t *testing.T) {
	o := degradationOrchestrator(t)
	o.config.Budget.MaxCoverageIterations = 1
	o.rfns.coverageGate = func(context.Context, reasoners.Deps, reasoners.CoverageGateInput) (map[string]any, error) {
		return map[string]any{
			"fully_covered":    false,
			"confident":        true,
			"gap_descriptions": []any{"Inspect uncovered auth edge cases"},
		}, nil
	}
	o.rfns.reviewDim = func(context.Context, reasoners.Deps, reasoners.ReviewDimensionInput) (map[string]any, error) {
		return nil, errors.New("reviewer made no progress for 360s")
	}

	primary := []schemas.ReviewFindings{{Title: "Primary finding", FilePath: "already.go"}}
	anatomy := schemas.AnatomyResult{Clusters: []schemas.ChangeCluster{{ID: "gap", Files: []string{"gap.go"}}}}
	got, _, err := o.runCoverageLoop(context.Background(), degradationPlan(1), anatomy, primary, nil)
	if err != nil {
		t.Fatalf("coverage-only reviewer failure must not fail root review: %v", err)
	}
	if len(got) != 1 || got[0].Title != "Primary finding" {
		t.Fatalf("primary findings changed after coverage-only failure: %#v", got)
	}
}

func TestMixedAndCleanDimensionsExposeOnlyActualDegradation(t *testing.T) {
	for _, tc := range []struct {
		name   string
		failed int
		want   string
	}{
		{name: "mixed", failed: 1, want: "> Degraded dimensions: 1"},
		{name: "clean empty", failed: 0, want: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			o := degradationOrchestrator(t)
			var calls atomic.Int32
			o.rfns.reviewDim = func(context.Context, reasoners.Deps, reasoners.ReviewDimensionInput) (map[string]any, error) {
				n := calls.Add(1)
				return map[string]any{
					"schema_parse_failed": n <= int32(tc.failed),
					"findings":            []any{},
					"sub_reviews":         []any{},
				}, nil
			}
			o.resetDimensionStats()
			plan := degradationPlan(3)
			if _, err := o.collectParallelReview(context.Background(), plan, ""); err != nil {
				t.Fatal(err)
			}
			result, err := o.generateOutput(context.Background(), nil, schemas.IntakeResult{PrSummary: "summary"}, schemas.AnatomyResult{}, plan, false)
			if err != nil {
				t.Fatal(err)
			}
			if result.Metadata.DegradedDimensions != tc.failed {
				t.Fatalf("degraded_dimensions = %d, want %d", result.Metadata.DegradedDimensions, tc.failed)
			}
			if tc.want != "" && !strings.Contains(result.Review.Body, tc.want) {
				t.Fatalf("review body missing %q", tc.want)
			}
			if tc.want == "" && strings.Contains(result.Review.Body, "Degraded dimensions:") {
				t.Fatalf("clean body unexpectedly degraded: %s", result.Review.Body)
			}
		})
	}
}
