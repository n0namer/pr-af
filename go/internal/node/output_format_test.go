package node

import (
	"context"
	"testing"

	"github.com/Agent-Field/pr-af/go/internal/config"
	"github.com/Agent-Field/pr-af/go/internal/orch"
	"github.com/Agent-Field/pr-af/go/internal/schemas"
)

func TestReviewHandlerOutputFormats(t *testing.T) {
	repo := t.TempDir()
	base := schemas.ReviewResult{
		ReviewID: "r1",
		Review:   schemas.GitHubReview{Body: "# Review\n", Event: "COMMENT"},
		Findings: []schemas.ScoredFinding{{
			ID: "f1", FilePath: "app.go", LineStart: 7, LineEnd: 8,
			Severity: "important", Title: "bad guard", Body: "mixed state is wrong", Confidence: .9,
		}},
	}
	cases := []struct {
		name, format string
		check        func(*testing.T, any)
	}{
		{"json", "json", func(t *testing.T, out any) {
			if _, ok := out.(schemas.ReviewResult); !ok {
				t.Fatalf("type=%T", out)
			}
		}},
		{"markdown", "markdown", func(t *testing.T, out any) {
			if out != "# Review\n" {
				t.Fatalf("out=%#v", out)
			}
		}},
		{"sarif", "sarif", func(t *testing.T, out any) {
			m, ok := out.(map[string]any)
			if !ok {
				t.Fatalf("type=%T", out)
			}
			if m["version"] != "2.1.0" {
				t.Fatalf("version=%v", m["version"])
			}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var seen schemas.ReviewInput
			n := &Node{
				NodeID:    "pr-af",
				reviewApp: &fakeApp{},
				runReview: func(_ context.Context, _ orch.Deps, in schemas.ReviewInput, _ config.ReviewConfig) (schemas.ReviewResult, error) {
					seen = in
					return base, nil
				},
			}
			out, err := n.reviewHandler(context.Background(), map[string]any{"repo_path": repo, "output_format": tc.format})
			if err != nil {
				t.Fatal(err)
			}
			if !seen.DryRun {
				t.Fatal("non-github format must force dry_run")
			}
			tc.check(t, out)
		})
	}
}

func TestReviewHandlerRejectsUnknownOutputFormat(t *testing.T) {
	n := &Node{
		NodeID:    "pr-af",
		reviewApp: &fakeApp{},
		runReview: func(context.Context, orch.Deps, schemas.ReviewInput, config.ReviewConfig) (schemas.ReviewResult, error) {
			t.Fatal("runReview must not run")
			return schemas.ReviewResult{}, nil
		},
	}
	_, err := n.reviewHandler(context.Background(), map[string]any{"repo_path": t.TempDir(), "output_format": "xml"})
	exec := asExecuteError(t, err)
	if exec.StatusCode != 400 {
		t.Fatalf("status=%d", exec.StatusCode)
	}
}
