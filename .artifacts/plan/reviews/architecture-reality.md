# Architecture source reality review

- Verdict: RATIFIED WITH PRECISION NOTES. Architecture API signatures and key boundaries reviewed against maintained Go source. This is a document review only; no product tests, runtime, spend provider, semantic review, or SourceLoop acceptance were executed or established.
- Baseline: expected `dc8d42e366952ff717331ad8a8e33f4e8e297133`. Git baseline/status could not be independently confirmed in this worker because Git rejected the workspace as dubious ownership; source reads succeeded. Parent reported this baseline.

## Verified contracts

- Go module pins are `go 1.21` and AgentField SDK `v0.1.130` (`go/go.mod:5,8`). No technology selection or upgrade is warranted.
- `orch.App`, `LocalCaller`, `Deps`, `New`, and `Run` match the declarations in `go/internal/orch/orchestrator.go:47-101,333-354`; `github.Client` matches `go/internal/github/client.go`.
- `node.BuildAgent` and `Node.RegisterAll` match `go/internal/node/node.go:153` and `go/internal/node/register.go:33`. Production sets `localCaller: app` (`node.go:203`), `reviewHandler` forwards it in `orch.Deps.Local` (`register.go:137`), and `orch.New` selects `callLocalSeams` for non-nil Local (`orchestrator.go:333-354`). `viaLocal` maps typed input with `afx.ToMap`, invokes `CallLocal`, then validates the map result (`orchestrator.go:304-330`).
- The 17 listed registered names match the `review` handler and 16 constants in `go/internal/node/register.go:33-62` and `go/internal/reasoners/names.go:5-22`. Registration tags and test parity are source-backed; actual control-plane discovery/callback reachability remain external runtime gates.
- `ResolveBudgetCaps` and `ReviewConfig.FromInput` signatures and precedence match `go/internal/config/budget.go:82-103` and `review.go:126-134`: explicit values, then environment, then $2/3600 defaults. The configured concurrency/coverage/depth defaults and per-input override fields are correctly identified in `budget.go` and `review.go`.
- Output posting guard is exactly `postToGitHub && !o.input.DryRun && strp(o.input.PrURL) != ""` (`go/internal/orch/output.go:~100`); the review correctly requires real output-path reachability and a fake positive control. Review handler resolves the repo before orchestration/config (`go/internal/node/register.go:~110-132`), so dry-run is not a network-off switch.
- Error behavior is consistent with `register.go:104-155`: bind/config failures are 400; repo-resolution failure is 500 before the mapped orchestration path; `ErrBadInput` maps to raw 400; other orchestration errors note and return the prefixed 500.
- Budget caution is correctly stated: `TestCostStaysInert` proves no normal cost accumulation in current reasoner results (`go/internal/orch/budget_test.go:18-32`); `TestWallClockBudgetTrips` tests a synthetic elapsed clock (`:34-65`). Source phase checks do not interrupt every in-flight provider call.

## Corrections / precision notes

1. The architecture inventory lists `go/internal/schemas/unmarshal.go`; parent has already corrected this to `go/internal/schemas/defaults.go`, where `ReviewInput.UnmarshalJSON` and default seeding live. Keep that corrected filename in the architecture.
2. Phrase duration behavior as “checked at phase boundaries/entry; does not establish cancellation of an in-flight external call.” This is more exact than implying a global deadline interrupts execution. Architecture already communicates this limitation; retain it in implementation handoff.
3. Treat any mention of the default $2/3600 values as configuration defaults, not spend/deadline enforcement. Source supports the effective configuration and synthetic wall-clock guard, while provider-side spend enforcement remains unproven and cost accumulation is inert by default.
4. The function `BuildAgent` attaches `AIConfig` only when `OPENROUTER_API_KEY` is present, while the harness configuration is built from `AIConfigFromEnv`; do not infer that successful construction proves usable AI/provider credentials. Existing architecture currently avoids that inference.

No further interface/signature correction found. The document's statements that runtime identity, callback, actual mandatory test events, positive semantic acceptance, and SourceLoop durability are unverified remain necessary; this review supplies none of that evidence.
