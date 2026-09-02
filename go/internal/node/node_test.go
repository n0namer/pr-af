package node

import (
	"testing"

	"github.com/Agent-Field/pr-af/go/internal/config"
)

func TestHarnessConfigProviderAwareBinary(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	cases := []struct {
		name string
		conf config.AIIntegrationConfig
		want string
	}{
		{"codex uses SDK default", config.AIIntegrationConfig{Provider: "codex", OpencodeBin: "C:/bin/opencode-custom"}, ""},
		{"opencode uses configured binary", config.AIIntegrationConfig{Provider: "opencode", OpencodeBin: "C:/bin/opencode-custom"}, "C:/bin/opencode-custom"},
		{"aforge uses SDK default", config.AIIntegrationConfig{Provider: "aforge", OpencodeBin: "C:/bin/opencode-custom"}, ""},
		{"generic override selects aforge", config.AIIntegrationConfig{Provider: "aforge", HarnessBin: "C:/bin/aforge-custom"}, "C:/bin/aforge-custom"},
		{"generic override wins", config.AIIntegrationConfig{Provider: "codex", OpencodeBin: "C:/bin/opencode-custom", HarnessBin: "C:/bin/provider-custom"}, "C:/bin/provider-custom"},
		{"generic override wins for opencode", config.AIIntegrationConfig{Provider: "opencode", OpencodeBin: "C:/bin/opencode-custom", HarnessBin: "C:/bin/provider-custom"}, "C:/bin/provider-custom"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := harnessConfig(tc.conf).BinPath; got != tc.want {
				t.Errorf("BinPath = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestHarnessConfigPreservesExistingFields(t *testing.T) {
	xdg := t.TempDir()
	t.Setenv("XDG_DATA_HOME", xdg)
	for _, key := range []string{"ANTHROPIC_API_KEY", "GOOGLE_API_KEY", "GH_TOKEN"} {
		t.Setenv(key, "")
	}
	t.Setenv("OPENROUTER_API_KEY", "legacy-dummy")
	t.Setenv("OPENAI_API_KEY", "openai-key")
	t.Setenv("OPENAI_BASE_URL", "http://fcm.internal:19280/v1")
	conf := config.AIIntegrationConfig{
		Provider: "opencode", HarnessModel: "openai/fcm", MaxTurns: 17,
		OpencodeBin: "C:/bin/opencode-custom",
	}
	got := harnessConfig(conf)
	if got.Provider != conf.Provider || got.Model != "compat/fcm" || got.MaxTurns != conf.MaxTurns || got.PermissionMode != "auto" || got.BinPath != conf.OpencodeBin {
		t.Errorf("harnessConfig fields = %+v", got)
	}
	if got.Env["OPENAI_API_KEY"] != "openai-key" || got.Env["OPENAI_BASE_URL"] != "http://fcm.internal:19280/v1" {
		t.Fatalf("OpenAI-compatible env = %#v", got.Env)
	}
	if _, ok := got.Env["OPENROUTER_API_KEY"]; ok {
		t.Fatal("legacy OPENROUTER_API_KEY must not be forwarded to opencode")
	}
	if got.Env["OPENCODE_CONFIG_CONTENT"] == "" {
		t.Fatal("OPENCODE_CONFIG_CONTENT missing")
	}
	if got.Env["XDG_DATA_HOME"] != xdg || got.Env["AGENTFIELD_AFORGE_COMMAND"] != "exec" {
		t.Errorf("supporting env = %#v", got.Env)
	}
}

// TestBuildAgentFromEnv is the main.go smoke: BuildAgent resolves node identity
// from the environment (with the pr-af / 8007 defaults), constructs the agent
// without a control plane or LLM key, and RegisterAll wires the full surface.
func TestBuildAgentFromEnv(t *testing.T) {
	cases := []struct {
		name       string
		env        map[string]string
		wantNodeID string
		wantServer string
		wantListen string
	}{
		{
			name: "defaults when env unset",
			env: map[string]string{
				"NODE_ID": "", "PORT": "", "AGENTFIELD_SERVER": "",
				"OPENROUTER_API_KEY": "", "OPENAI_API_KEY": "", "OPENAI_BASE_URL": "",
			},
			wantNodeID: "pr-af",
			wantServer: "http://localhost:8080",
			wantListen: ":8007",
		},
		{
			name: "env overrides",
			env: map[string]string{
				"NODE_ID":            "pr-af-canary",
				"PORT":               "9107",
				"AGENTFIELD_SERVER":  "http://cp.internal:8080",
				"OPENROUTER_API_KEY": "",
				"OPENAI_API_KEY":     "",
				"OPENAI_BASE_URL":    "",
			},
			wantNodeID: "pr-af-canary",
			wantServer: "http://cp.internal:8080",
			wantListen: ":9107",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.env {
				t.Setenv(k, v)
			}
			n, err := BuildAgent("pr-af", "8007", "AI-Native Pull Request Review Agent")
			if err != nil {
				t.Fatalf("BuildAgent: %v", err)
			}
			if n.App == nil {
				t.Fatal("BuildAgent returned a nil App")
			}
			if n.NodeID != tc.wantNodeID {
				t.Errorf("NodeID = %q, want %q", n.NodeID, tc.wantNodeID)
			}
			if n.AgentFieldServer != tc.wantServer {
				t.Errorf("AgentFieldServer = %q, want %q", n.AgentFieldServer, tc.wantServer)
			}
			if n.ListenAddress != tc.wantListen {
				t.Errorf("ListenAddress = %q, want %q", n.ListenAddress, tc.wantListen)
			}

			n.RegisterAll()
			if got := len(n.RegisteredNames()); got != 17 {
				t.Errorf("registered %d reasoners, want 17", got)
			}
		})
	}
}

// TestBuildAgentWithLLMKey proves AIConfig attaches on the same
// OpenAI-compatible key/base/model contract used by the harness. A legacy
// OpenRouter admission key may exist but must not select the runtime provider.
func TestBuildAgentWithLLMKey(t *testing.T) {
	t.Setenv("NODE_ID", "")
	t.Setenv("PORT", "")
	t.Setenv("AGENTFIELD_SERVER", "")
	t.Setenv("OPENROUTER_API_KEY", "legacy-dummy")
	t.Setenv("OPENAI_API_KEY", "fcm-key")
	t.Setenv("OPENAI_BASE_URL", "http://fcm.internal:19280/v1")
	t.Setenv("PR_AF_AI_MODEL", "openai/fcm")

	n, err := BuildAgent("pr-af", "8007", "desc")
	if err != nil {
		t.Fatalf("BuildAgent with OpenAI-compatible config: %v", err)
	}
	if n.App == nil {
		t.Fatal("nil App")
	}
}

func TestBuildAgentRejectsPartialOpenAICompatibleConfig(t *testing.T) {
	for _, tc := range []struct {
		name string
		key  string
		base string
	}{
		{name: "key without base", key: "fcm-key"},
		{name: "base without key", base: "http://fcm.internal:19280/v1"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("OPENROUTER_API_KEY", "legacy-dummy")
			t.Setenv("OPENAI_API_KEY", tc.key)
			t.Setenv("OPENAI_BASE_URL", tc.base)
			if _, err := BuildAgent("pr-af", "8007", "desc"); err == nil {
				t.Fatal("BuildAgent should reject a partial OpenAI-compatible provider configuration")
			}
		})
	}
}

func TestAIModelForAPIStripsProviderRoutingPrefix(t *testing.T) {
	if got := aiModelForAPI("openai/fcm"); got != "fcm" {
		t.Errorf("openai prefixed: got %q, want fcm", got)
	}
	if got := aiModelForAPI("openrouter/moonshotai/kimi-k2.5"); got != "moonshotai/kimi-k2.5" {
		t.Errorf("openrouter prefixed: got %q, want moonshotai/kimi-k2.5", got)
	}
	if got := aiModelForAPI("minimax/minimax-m2.5"); got != "minimax/minimax-m2.5" {
		t.Errorf("unprefixed: got %q, want unchanged", got)
	}
}
