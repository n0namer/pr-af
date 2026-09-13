package orch

import (
	"context"
	"sync"
	"testing"

	"github.com/Agent-Field/pr-af/go/internal/config"
	"github.com/Agent-Field/pr-af/go/internal/reasoners"
	"github.com/Agent-Field/pr-af/go/internal/schemas"
)

func TestFullPipelineInProcess(t *testing.T) {
	cfg := config.DefaultReviewConfig()
	cfg.Comments.PolishEnabled = false
	cfg.Comments.MergeGateEnabled = false
	o := New(Deps{App: &fakeApp{}}, schemas.ReviewInput{DryRun: true}, cfg)
	o.cleanupFn = func() {}

	var mu sync.Mutex
	calls := map[string]int{}
	mark := func(name string) {
		mu.Lock()
		calls[name]++
		mu.Unlock()
	}

	o.runIntakeFn = func(context.Context) (schemas.IntakeResult, error) {
		mark("intake")
		return schemas.IntakeResult{PrSummary: "change app", ReviewDepth: "standard"}, nil
	}
	o.runAnatomyFn = func(context.Context, schemas.IntakeResult) (schemas.AnatomyResult, error) {
		mark("anatomy")
		return schemas.AnatomyResult{Clusters: []schemas.ChangeCluster{{ID: "c1", Files: []string{"app.go"}}}}, nil
	}
	o.resolveDepthFn = func(schemas.IntakeResult) string { return "standard" }

	meta := func(lens string) func(context.Context, reasoners.Deps, reasoners.MetaInput) (map[string]any, error) {
		return func(context.Context, reasoners.Deps, reasoners.MetaInput) (map[string]any, error) {
			mark("meta_" + lens)
			return metaResultMap(lens, "d", "app.go"), nil
		}
	}
	o.rfns.metaSemantic = meta("semantic")
	o.rfns.metaMechanical = meta("mechanical")
	o.rfns.metaSystemic = meta("systemic")
	o.rfns.reviewDim = func(context.Context, reasoners.Deps, reasoners.ReviewDimensionInput) (map[string]any, error) {
		mark("review")
		return map[string]any{
			"findings": []any{map[string]any{
				"file_path": "app.go", "line_start": 1,
				"severity": "important", "title": "regression",
				"body": "breaks contract", "evidence": "app.go:1",
				"confidence": 0.9,
			}},
			"sub_reviews": []any{},
		}, nil
	}
	o.rfns.adversary = func(context.Context, reasoners.Deps, reasoners.AdversaryInput) (map[string]any, error) {
		mark("adversary")
		return map[string]any{"results": []any{}}, nil
	}
	o.rfns.coverageGate = func(context.Context, reasoners.Deps, reasoners.CoverageGateInput) (map[string]any, error) {
		mark("coverage")
		return map[string]any{"fully_covered": true, "confident": true, "gap_descriptions": []any{}}, nil
	}
	o.rfns.extractOblig = func(context.Context, reasoners.Deps, reasoners.ExtractObligationsInput) (map[string]any, error) {
		mark("consistency")
		return map[string]any{"obligations": []any{}}, nil
	}
	o.rfns.evidenceVerify = func(context.Context, reasoners.Deps, reasoners.EvidenceVerifierInput) (map[string]any, error) {
		mark("evidence")
		return map[string]any{"verified_findings": []any{map[string]any{
			"title": "regression", "verified": true,
			"revised_severity": "important", "revised_confidence": 0.9,
		}}}, nil
	}
	o.generateOutputFn = func(_ context.Context, scored []schemas.ScoredFinding, _ schemas.IntakeResult, _ schemas.AnatomyResult, _ schemas.ReviewPlan, _ bool) (schemas.ReviewResult, error) {
		mark("output")
		return schemas.ReviewResult{Summary: schemas.ReviewSummary{TotalFindings: len(scored)}}, nil
	}

	if _, err := o.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{
		"intake", "anatomy", "meta_semantic", "meta_mechanical",
		"meta_systemic", "review", "evidence", "adversary",
		"coverage", "consistency", "output",
	} {
		if calls[name] == 0 {
			t.Fatalf("architecture step %q not reached: %v", name, calls)
		}
	}
}
