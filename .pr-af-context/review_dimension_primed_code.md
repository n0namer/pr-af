### go/internal/node/node.go
```
1: // Package node is the PR-AF wiring wave (T4.2): it constructs the shared
2: // *agent.Agent from the environment and registers the exact Python reasoner
3: // surface (design §B.1) so the Go node is a drop-in opt-in sibling of the Python
4: // pr-af node.
5: //
6: // node.go owns agent construction (env -> agent.Config, mirroring
7: // src/pr_af/app.py:26-50) plus the custom HTTP server that wraps the SDK handler
8: // with the /webhook/github route (app.py:365-367). register.go owns the
9: // per-reasoner registration (the 17 names of §B.1); webhook.go owns the GitHub
10: // @mention webhook handler.
11: package node
12: 
13: import (
14: 	"context"
15: 	"errors"
16: 	"fmt"
17: 	"log"
18: 	"net"
19: 	"net/http"
20: 	"os"
21: 	"os/signal"
22: 	"strings"
23: 	"syscall"
24: 	"time"
25: 
26: 	"github.com/Agent-Field/agentfield/sdk/go/agent"
27: 	"github.com/Agent-Field/agentfield/sdk/go/ai"
28: 
29: 	"github.com/Agent-Field/pr-af/go/internal/config"
30: 	"github.com/Agent-Field/pr-af/go/internal/github"
31: 	"github.com/Agent-Field/pr-af/go/internal/orch"
32: 	"github.com/Agent-Field/pr-af/go/internal/schemas"
33: )
34: 
35: // Node bundles the constructed agent with the resolved environment config and
36: // the collaborators the review handler and webhook thread through.
37: type Node struct {
38: 	// App is the SDK agent. It satisfies orch.App (Harness/AI/Pause/Note) and
39: 	// the reasoner Deps interfaces directly, so register.go points every Deps
40: 	// field at it and Serve mounts App.Handler() as the fallback route.
41: 	App *agent.Agent
42: 
43: 	// labelDedupe bounds duplicate label-triggered review dispatches. It is
44: 	// process-local by design; see webhookDedupe.claim for the limitation.
45: 	labelDedupe webhookDedupe
46: 
47: 	// webhookClient is nil in production (fireReview uses a bounded default).
48: 	// Tests inject a transport so webhook dispatches need no listening socket.
49: 	webhookClient *http.Client
50: 
51: 	// NodeID is the resolved node id (NODE_ID env, or the pr-af default).
52: 	NodeID string
53: 
54: 	// AgentFieldServer is the control-plane base URL (AGENTFIELD_SERVER). The
55: 	// webhook fires the async review at "{AgentFieldServer}/api/v1/execute/async/
56: 	// {NodeID}.review" and the HITL gate derives the approval webhook URL from it.
57: 	AgentFieldServer string
58: 
59: 	// ListenAddress is the ":port" the custom server binds (":"+PORT).
60: 	ListenAddress string
61: 
62: 	// reviewApp is the agent-capability seam the review handler feeds into
63: 	// orch.Deps.App AND emits the pipeline-failure note through. It defaults to
64: 	// App; the error-mapping tests override it with a fake that records notes.
65: 	reviewApp orch.App
66: 
67: 	// gh is the GitHub client injected into orch.Deps.GH. Tests that stub
68: 	// runReview leave it unused (nil).
69: 	gh github.Client
70: 
71: 	// localCaller is the tracked same-process invocation seam fed into
72: 	// orch.Deps.Local: production points it at App so every pipeline phase is
73: 	// reported to the control plane as a child execution (the review DAG).
74: 	// Tests that stub runReview or need direct-call phases leave it nil.
75: 	localCaller orch.LocalCaller
76: 
77: 	// runReview is the orchestrator-construct-and-run seam. Production builds and
78: 	// runs the real orchestrator; the error-mapping tests inject ErrBadInput /
79: 	// other failures without a live harness (design §F "seam for the orchestrator
80: 	// constructor").
81: 	runReview func(ctx context.Context, deps orch.Deps, in schemas.ReviewInput, cfg config.ReviewConfig) (schemas.ReviewResult, error)
82: 
83: 	// registered records every reasoner name passed through the single
84: 	// registration path, in order, so the parity test (V1) can assert the exact
85: 	// surface. tags records the tags registered per name (review -> nil; the 16
86: 	// internal reasoners -> ["review","pr"]).
87: 	registered []string
88: 	tags       map[string][]string
89: }
90: 
91: // RegisteredNames returns a copy of the reasoner names registered on this node,
92: // in registration order — the functional/unit parity source of truth (V1).
93: func (n *Node) RegisteredNames() []string {
94: 	return append([]string(nil), n.registered...)
95: }
96: 
97: // TagsFor returns a copy of the tags registered for name (nil when none).
98: func (n *Node) TagsFor(name string) []string {
99: 	return append([]string(nil), n.tags[name]...)
100: }
101: 
102: // defaultRunReview constructs the real orchestrator and runs it. Kept as a
103: // package function so BuildAgent can point Node.runReview at it and tests can
104: // swap in a stub.
105: func defaultRunReview(ctx context.Context, deps orch.Deps, in schemas.ReviewInput, cfg config.ReviewConfig) (schemas.ReviewResult, error) {
106: 	return orch.New(deps, in, cfg).Run(ctx)
107: }
108: 
109: // harnessConfig maps the resolved AI integration configuration to the SDK
110: // harness configuration. An empty BinPath lets the SDK select the configured
111: // provider's default executable.
112: func harnessConfig(c config.AIIntegrationConfig) *agent.HarnessConfig {
113: 	return &agent.HarnessConfig{
114: 		Provider:       c.Provider,
115: 		Model:          c.HarnessRuntimeModel(),
116: 		MaxTurns:       c.MaxTurns,
117: 		PermissionMode: "auto",
118: 		Env:            c.ProviderEnv(),
119: 		BinPath:        resolvedHarnessBin(c),
120: 	}
121: }
122: 
123: func resolvedHarnessBin(c config.AIIntegrationConfig) string {
124: 	if c.HarnessBin != "" {
125: 		return c.HarnessBin
126: 	}
127: 	if c.Provider == "opencode" {
128: 		return c.OpencodeBin
129: 	}
130: 	return ""
131: }
132: 
133: // BuildAgent constructs the PR-AF agent from the environment exactly as the
134: // Python entry point does (app.py:26-50):
135: //
136: //   - NODE_ID            default "pr-af" (the maintained package identity).
137: //   - AGENTFIELD_SERVER  default "http://localhost:8080".
138: //   - AGENTFIELD_API_KEY -> Config.Token (control-plane bearer).
139: //   - PORT               default "8007" -> ListenAddress ":8007".
140: //   - AGENT_CALLBACK_URL -> Config.PublicURL — the base URL the CP uses to reach
141: //     this node; unset falls back to the SDK's http://localhost:<port>.
142: //   - HarnessConfig / AIConfig — the selected harness + LLM credentials the
143: //     reasoners rely on. Every reasoner calls the harness with only Cwd set, so
144: //     the agent's default HarnessConfig Provider/Model must be present, and the
145: //     two .ai() gates (intake/coverage) need AIConfig. Mirrors app.py's
146: //     harness_config=/ai_config=.
147: //
148: // The direct .ai() path follows the same OpenAI-compatible contract as the
149: // harness: OPENAI_API_KEY + OPENAI_BASE_URL + PR_AF_AI_MODEL. A partial pair is
150: // rejected at boot so provider misconfiguration cannot silently fall back.
151: func BuildAgent(defaultNodeID, defaultPort, description string) (*Node, error) {
152: 	nodeID := envOr("NODE_ID", defaultNodeID)
153: 	server := envOr("AGENTFIELD_SERVER", "http://localhost:8080")
154: 	token := os.Getenv("AGENTFIELD_API_KEY")
155: 	port := envOr("PORT", defaultPort)
156: 
157: 	aiConf, err := config.AIConfigFromEnv()
158: 	if err != nil {
159: 		// Python constructs AIIntegrationConfig at module import, so a malformed
160: 		// numeric env var (e.g. PR_AF_MAX_TURNS=abc) crashes the node at boot.
161: 		return nil, err
162: 	}
163: 
164: 	cfg := agent.Config{
165: 		NodeID:        nodeID,
166: 		Version:       "0.1.0",
167: 		AgentFieldURL: server,
168: 		Token:         token,
169: 		ListenAddress: ":" + port,
170: 		PublicURL:     os.Getenv("AGENT_CALLBACK_URL"),
171: 		CLIConfig:     &agent.CLIConfig{AppDescription: description},
172: 		HarnessConfig: harnessConfig(aiConf),
173: 	}
174: 	openAIKey := os.Getenv("OPENAI_API_KEY")
175: 	openAIBase := os.Getenv("OPENAI_BASE_URL")
176: 	if (openAIKey == "") != (openAIBase == "") {
177: 		return nil, errors.New("OPENAI_API_KEY and OPENAI_BASE_URL must be configured together")
178: 	}
179: 	if openAIKey != "" {
180: 		cfg.AIConfig = &ai.Config{
181: 			Model:   aiModelForAPI(aiConf.AIModel),
182: 			APIKey:  openAIKey,
183: 			BaseURL: openAIBase,
184: 		}
185: 	}
186: 
187: 	app, err := agent.New(cfg)
188: 	if err != nil {
189: 		return nil, fmt.Errorf("create agent %q: %w", nodeID, err)
190: 	}
191: 
192: 	n := &Node{
193: 		App:              app,
194: 		NodeID:           nodeID,
195: 		AgentFieldServer: server,
196: 		ListenAddress:    ":" + port,
197: 		reviewApp:        app,
198: 		gh:               github.NewClient(""), // reads GH_TOKEN internally (app.py GitHubClient())
199: 		localCaller:      app,
200: 		runReview:        defaultRunReview,
201: 		tags:             map[string][]string{},
202: 	}
203: 	return n, nil
204: }
205: 
206: // Serve runs the custom HTTP server and registers with the control plane.
207: //
208: // It mounts /webhook/github on the PR-AF handler and delegates every other path
209: // (/health, /reasoners/, /execute, /discover, …) to the SDK's App.Handler(),
210: // mirroring Python's app.add_api_route("/webhook/github", …) grafted onto the
211: // SDK's own routes.
212: //
213: // Ordering (design §G): bind the listener BEFORE App.Initialize so the control
214: // plane's post-registration health check reaches a live server — the same
215: // startServer→Initialize order agent.Serve uses, reproduced here because the
216: // webhook route forces a bespoke mux instead of App.Serve.
217: func (n *Node) Serve(ctx context.Context) error {
218: 	mux := http.NewServeMux()
219: 	mux.HandleFunc("/webhook/github", n.webhookGitHub)
220: 	mux.Handle("/", n.App.Handler())
221: 
222: 	ln, err := net.Listen("tcp", n.ListenAddress)
223: 	if err != nil {
224: 		return fmt.Errorf("listen %s: %w", n.ListenAddress, err)
225: 	}
226: 	srv := &http.Server{Handler: mux}
227: 
228: 	serveErr := make(chan error, 1)
229: 	go func() {
230: 		if serr := srv.Serve(ln); serr != nil && !errors.Is(serr, http.ErrServerClosed) {
231: 			serveErr <- serr
232: 		}
233: 	}()
234: 
235: 	if err := n.App.Initialize(ctx); err != nil {
236: 		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
237: 		defer cancel()
238: 		_ = srv.Shutdown(shutdownCtx)
239: 		return fmt.Errorf("initialize node %q: %w", n.NodeID, err)
240: 	}
241: 
242: 	sigCh := make(chan os.Signal, 1)
243: 	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
244: 
245: 	select {
246: 	case <-ctx.Done():
247: 	case sig := <-sigCh:
248: 		log.Printf("pr-af: received signal %s, shutting down", sig)
249: 	case err := <-serveErr:
250: 		return fmt.Errorf("webhook server: %w", err)
251: 	}
252: 
253: 	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
254: 	defer cancel()
255: 	return srv.Shutdown(shutdownCtx)
256: }
257: 
258: // envOr returns the value of key, or def when the env var is unset or empty.
259: func envOr(key, def string) string {
260: 	if v := os.Getenv(key); v != "" {
261: 		return v
262: 	}
263: 	return def
264: }
265: 
266: // aiModelForAPI converts the configured LiteLLM/OpenCode-style model id into
267: // the model name sent to an OpenAI-compatible API. Provider prefixes are
268: // routing metadata and must not be sent as part of the upstream model name.
269: func aiModelForAPI(model string) string {
270: 	model = strings.TrimPrefix(model, "openai/")
271: 	return strings.TrimPrefix(model, "openrouter/")
272: }
```
_import/usage context:_ IMPORTS: import (
IMPORTED BY: none