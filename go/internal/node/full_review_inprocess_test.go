package node

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Agent-Field/pr-af/go/internal/schemas"
)

func TestReviewHandlerWithExternalMockHarness(t *testing.T) {
	goRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	tmp := t.TempDir()
	mockBin := filepath.Join(tmp, "opencode")
	build := exec.Command("/usr/local/go/bin/go", "build", "-o", mockBin, "./test/mockcli")
	build.Dir = goRoot
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build mock harness: %v\n%s", err, out)
	}

	repo := filepath.Join(tmp, "repo")
	if err := os.MkdirAll(repo, 0o755); err != nil {
		t.Fatal(err)
	}
	runGit := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = repo
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	runGit("init", "-q")
	runGit("config", "user.email", "pr-af-test@example.com")
	runGit("config", "user.name", "PR-AF Test")
	if err := os.WriteFile(filepath.Join(repo, "base.go"), []byte("package fixture\n\nfunc Value() int { return 1 }\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit("add", "base.go")
	runGit("commit", "-q", "-m", "base")
	if err := os.WriteFile(filepath.Join(repo, "base.go"), []byte("package fixture\n\nfunc Value() int { return 2 }\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit("add", "base.go")
	runGit("commit", "-q", "-m", "change")

	t.Setenv("PR_AF_PROVIDER", "opencode")
	t.Setenv("PR_AF_HARNESS_BIN", mockBin)
	t.Setenv("PR_AF_OPENCODE_BIN", mockBin)
	t.Setenv("PR_AF_MOCK_STATE_DIR", filepath.Join(tmp, "state"))
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("OPENAI_BASE_URL", "")
	t.Setenv("OPENROUTER_API_KEY", "")
	t.Setenv("GH_TOKEN", "")

	n, err := BuildAgent("pr-af-test", "0", "PR-AF in-process node probe")
	if err != nil {
		t.Fatal(err)
	}
	n.RegisterAll()
	if len(n.RegisteredNames()) != 17 {
		t.Fatalf("registered reasoners=%d, want 17", len(n.RegisteredNames()))
	}

	resultAny, err := n.reviewHandler(context.Background(), map[string]any{
		"repo_path": repo,
		"depth":     "standard",
		"dry_run":   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	result, ok := resultAny.(schemas.ReviewResult)
	if !ok {
		t.Fatalf("review result type=%T", resultAny)
	}
	if result.Summary.TotalFindings == 0 {
		t.Fatal("review returned zero findings")
	}
	if len(result.Metadata.PhasesCompleted) == 0 {
		t.Fatal("review reported no completed phases")
	}

	logBytes, err := os.ReadFile(filepath.Join(tmp, "state", "invocations.jsonl"))
	if err != nil {
		t.Fatalf("read mock harness invocation log: %v", err)
	}
	logText := string(logBytes)
	for _, role := range []string{"intake_fallback", "anatomy", "meta_semantic", "meta_mechanical", "meta_systemic", "review_dimension", "evidence_verifier", "adversary", "compound_finder", "extract_obligations"} {
		if !strings.Contains(logText, `"role":"`+role+`"`) {
			t.Fatalf("external harness role %q not observed\n%s", role, logText)
		}
	}
}
