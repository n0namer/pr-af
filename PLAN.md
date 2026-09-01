# PR-AF — Canonical Plan and Current State

> Status: ACTIVE — B1 PASS / B2 RUNTIMEFILEACI-ACTIVATION + EXECUTION-BRIDGE BLOCKERS
> Updated: 2026-09-01
> Canonical owner: `n0namer/pr-af:dev/PLAN.md`
> Active development branch: `dev`
> Runtime topology owner: `n0namer/universal-solver`
> BMAD lane: `bmad-quick-dev` / implementation-debugging

## Authority and anti-drift

- This file owns PR-AF current phase, Phase Goal, bounded batches, DoD, progress, blockers, and ONE next move.
- `README.md` and `docs/ARCHITECTURE.md` own product/architecture intent.
- `AGENTS.md` owns repository engineering rules.
- `docs/LLM_PROVIDER_SECURITY_CONTRACT.md` owns PR-AF-specific LLM/provider/security semantics and links to the cross-component contract in `universal-solver`.
- CURRENT runtime/readback owns actual loaded state.
- `universal-solver` owns the permanent AgentField DEV topology and SourceLoop plumbing; its project PLAN is not PR-AF current-state SoT.
- `BMAD-MNNZ` is the workflow/skill rulebook, not project SoT.
- After every material state/evidence change: verify → update this file → replan from CURRENT evidence.
- Never claim `DONE` without all required DoD evidence.

## North Star

PR-AF is an open-source AgentField PR-review node for deep, evidence-grounded code review: turn an incoming pull request into a task-specific review plan, run focused reviewer work, ground findings in code evidence, challenge weak claims, synthesize compound risks, and return useful review output with strong recall and low false positives.

For this workstream the engineering North Star is the fastest path to a **verified-running maintained Go PR-AF** in permanent DEV, with real semantic review proof and exact source/runtime provenance.

## Current phase

Phase: implementation / runtime enablement and provider-path verification.

### Phase Goal

Run the maintained Go PR-AF from the existing permanent AgentField DEV workspace without GitHub-first debugging or per-hypothesis redeploy, prove registration + functional canary + one real semantic PR review, fix only evidence-proven Go defects directly in `/src/pr-af`, then let SourceLoop capture the exact accepted delta to `pr-af:dev`.

## Operating contract

`OBSERVE → LOCALIZE → ROUTE → PATCH → TARGETED VERIFY → ITERATE → FULL VERIFY → RUNTIME PROOF → CANONICALIZE → POST-DEPLOY VERIFY → WRITE-BACK`

Hard rules:
- Runtime-bound debugging uses the existing permanent DEV workspace first.
- PR-AF code edits happen directly in persistent `/src/pr-af`; do not use “edit GitHub → redeploy → observe → repeat”.
- GitHub is the canonicalization/release boundary, not the inner debug loop.
- SourceLoop captures only an accepted verified runtime delta.
- Preserve unrelated runtime state.
- Do not patch business logic while the failure is still explained by install/start/config/orchestration.
- After timeout or ambiguous mutation, inspect post-state before retrying.
- Production live editing is forbidden by default.
- Work in bounded ~30-minute BMAD batches with explicit DoD and 80/20 priority.

## CURRENT evidence — 2026-09-01

| Claim | Evidence | Verdict |
|---|---|---|
| Canonical PR-AF branch | `n0namer/pr-af:dev` runtime source was clean-reconciled to `6245796a6c47a0f114dd0e8382f4abf63a89752f`; later PLAN-only commits advance branch metadata but do not change PR-AF application code | FACT |
| Live source identity | permanent workforce `/src/pr-af` was reconciled cleanly before current bootstrap; fresh SourceLoop capture still requires rereading HEAD/dirty state immediately before capture | FACT |
| Current DEV topology target | Coolify app `edshqtkwskg3lrczekhcmd71` is loaded on `universal-solver` commit `b7e2f00116358d78d01a73b77aa31d1c2bdfb9d5`; current control-plane, workforce and Deep Research are healthy | FACT |
| Maintained implementation | `go/` package; node id `pr-af`; default port `8007`; build `./cmd/pr-af`; start `bin/pr-af` | FACT |
| Runtime install path | current `/afhome/installed.yaml` shows `pr-af.source_path: /src/pr-af/go` | PASS |
| Runtime process | current `/afhome/installed.yaml` shows `pr-af` `status: running`, `desired_state: running`, port `8007`, pid `7706` | PASS |
| Runtime registration | PR-AF log previously recorded `node.register.complete`; current control-plane accepted callback `http://workforce:8007` and created the `pr-af` DID | PASS |
| Reasoner surface | exact Go source registers external `review` plus 16 internal review-DAG reasoners | SOURCE PASS |
| Provider wiring in workforce | topology injects `OPENAI_API_KEY`, `OPENAI_BASE_URL`, `OPENROUTER_API_KEY`, and `GH_TOKEN` by name; secret values are not recorded here | FACT |
| Go manifest provider contract | `go/agentfield-package.yaml` currently requires `OPENROUTER_API_KEY`; `GH_TOKEN` is optional | FACT |
| Go provider env propagation | `go/internal/config/ai.go::ProviderEnv()` forwards `OPENAI_API_KEY` but not `OPENAI_BASE_URL`; unified provider contract explicitly requires `OPENAI_API_KEY + OPENAI_BASE_URL + openai/<model>` | SOURCE FACT / APPROVED_CONFORMANCE_FIX |
| Bootstrap lifecycle route | SWE-style workforce bootstrap now installs `/src/pr-af/go` and starts `pr-af:8007`; B1 no longer depends on a separate typed install capability | RESOLVED |
| External execution bridge | stale-stack binding was diagnosed and current control-plane is now on shared `coolify` network with unique alias `agentfield-current-control-plane`; gateway process has the correct unique upstream + upstream API-key injection, but external ingress remains unavailable because Coolify strips custom Traefik labels from the custom-service runtime and generic Traefik file-write is blocked by the current mediator | EXECUTION_BRIDGE_BLOCKER |

## Design/runtime drift register

1. **INSTALL_PATH_DRIFT — RESOLVED**: workforce bootstrap installs maintained `/src/pr-af/go`; current installed-state readback proves it.
2. **START_POLICY_DRIFT — RESOLVED**: maintained Go PR-AF is running on `8007` and registered in the current control-plane.
3. **PROVIDER_CONTRACT_DRIFT — APPROVED_FIX**: user/design authority selected one fleet transport contract: `OPENAI_API_KEY + OPENAI_BASE_URL + openai/<model>`. `ProviderEnv()` omits `OPENAI_BASE_URL`, so PR-AF requires a bounded conformance patch plus regression coverage even before semantic/runtime failure reproduction. Real provider PASS still requires an actual review call.
4. **RUNTIMEFILEACI_ACTIVATION_BLOCKER — OPEN**: canonical `vps-terminal/config/targets.json` already contains logical target `agentfield-dev-workforce` with `/src` live-patch root, and read/preview work on the current workforce. Apply is falsely blocked by the currently loaded `vps-terminal-dev` runtime (`aaa2c5b...`) on `CANONICAL_OWNER_BEFORE_CONFIG`; the vps-terminal SoT records a validated source fix whose DEV activation is itself blocked without a non-Git activation capability. Generic shell mutation remains prohibited.
5. **EXECUTION_BRIDGE_BLOCKER — OPEN**: the former stale AgentField stack binding is resolved, but external `agentfield_actions` ingress is still unavailable because Coolify strips custom Traefik labels; direct internal API probing and Traefik file mutation remain mediator-blocked.

## Phase DoD

- [x] Canonical `PLAN.md` exists and owns PR-AF current state.
- [x] Exact runtime source/install path observed.
- [x] Maintained Go package installed from `/src/pr-af/go`.
- [x] `pr-af` process running on port `8007` from the maintained package.
- [x] `pr-af` registration proven by node-side and control-plane-side evidence.
- [x] Expected reasoner surface identified from exact Go source (`review` + 16 internal reasoners).
- [x] Provider env/config source inspected without exposing secret values.
- [ ] Execute the first current-control-plane `pr-af.review` canary.
- [ ] Gonka/OpenAI-compatible path proven end-to-end: custom base survives to the actual model/tool call and no unintended OpenRouter fallback occurs.
- [ ] Deterministic/functional canary PASS.
- [ ] One bounded real semantic PR-review E2E PASS with outcome inspection.
- [ ] If code changed: smallest relevant regression red→green, then required broader Go validation.
- [ ] Accepted code delta captured by SourceLoop.
- [ ] Durable `pr-af:dev` SHA corresponds to the accepted runtime delta.
- [ ] No unintended container-only code delta remains.

## Bounded BMAD batches

### B1 — Correct maintained runtime path
Goal: move from “source present” to “maintained Go node actually running”.

DoD:
- maintained Go package installed from `/src/pr-af/go` — **PASS**;
- `pr-af` running on `8007` — **PASS**;
- current control-plane registration — **PASS**;
- exact source/runtime identity read back — **PASS for current generation**.

Status: `PASS`.

### B2 — Unified provider contract conformance + runtime proof
Goal: make PR-AF conform to the approved fleet transport contract `OPENAI_API_KEY + OPENAI_BASE_URL + openai/<model>`, then prove that contract on a real `pr-af.review` execution.

Priority order:
1. apply the approved minimal live-source delta directly in `/src/pr-af/go`: add `OPENAI_BASE_URL` to `ProviderEnv()` and focused regression coverage in `internal/config/config_test.go`;
2. run targeted `go test ./internal/config`, then required `make check` on the same live source;
3. reload only PR-AF in the same runtime;
4. restore/use one executable route to the CURRENT control-plane and run a bounded synthetic dry-run `pr-af.review` canary;
5. inspect execution evidence for `opencode + openai/<Gonka>`, exact custom base survival, and absence of unintended OpenRouter/default-OpenAI fallback.

DoD:
- source forwards `OPENAI_BASE_URL` unchanged when set and omits it when unset;
- focused regression PASS;
- `make check` PASS on the same source;
- same-runtime PR-AF reload PASS;
- real canary proves model/key/base as one provider identity;
- no unintended fallback.

Status: `ACTIVE / CODE_PATCH_BLOCKED_BY_LIVE_PATCH_CAPABILITY + EXECUTION_BRIDGE_BLOCKER`.

### B3 — Semantic acceptance and durability
Goal: prove useful PR review and canonicalize the exact accepted delta.

Tasks:
- deterministic/functional canary;
- bounded real PR review with outcome inspection;
- correlate source SHA, process/node identity, execution id, model/provider, logs, and result;
- reject HTTP/health/execution-success-only evidence as semantic proof;
- SourceLoop capture only after acceptance;
- verify final `pr-af:dev` SHA and absence of unintended runtime-only delta.

Status: `BLOCKED_BY_B2`.

## Validation ladder

From `go/`:
1. syntax/format/static;
2. directly affected tests;
3. related regression;
4. full required suite (`make check` = build + vet + test);
5. runtime smoke/integration;
6. functional PR-AF canary;
7. semantic PR-review E2E.

`go/test/e2e/run.sh` is deterministic E2E evidence but is not sufficient proof of real-model review quality.

## Current blocker

PR-AF itself is no longer blocked on install/start/registration. B2 has two independent execution-surface blockers:

1. **RUNTIMEFILEACI_ACTIVATION_BLOCKER** — the logical target `agentfield-dev-workforce` already exists in canonical `vps-terminal/config/targets.json`, resolves to the current workforce, and successfully supports `readTargetFile` + `previewTargetPatch` for `/src/pr-af`. `applyTargetPatch` is falsely re-blocked by the currently loaded `vps-terminal-dev` runtime (`aaa2c5b...`) with `CANONICAL_OWNER_BEFORE_CONFIG`, even after canonical config revision + target identity were freshly observed and guidance returned `ALLOW`. The requested prod-operator fallback was also checked: prod VPS Terminal is healthy and can inventory the DEV gateway + its SockGuard, but its loaded target registry does not contain `vps-terminal-dev-gateway` (`unknown_target`); registered `coolify-control` and `workbench` targets have no Docker socket, and the DEV gateway code is image-layer only (no bind/volume for `/app/gateway`). Therefore prod currently has no typed, policy-compliant path to patch the DEV gateway container. GitHub-source redeploy remains outside this task's coding constraint.
2. **EXECUTION_BRIDGE_BLOCKER** — stale `agentfield_actions` binding to the old AgentField stack was diagnosed and removed; CURRENT control-plane is healthy on shared `coolify` network with unique alias `agentfield-current-control-plane`, but the external Action ingress remains unavailable because Coolify strips the custom Traefik labels and generic Traefik file mutation is mediator-blocked.

Neither blocker is evidence of a PR-AF application failure, and neither is missing user authorization.

## Current decision / ONE next move

Wait for or use the first CURRENT-authoritative **non-Git activation path** for the already validated `vps-terminal-dev` RuntimeFileACI fix. Do not bypass mediation with shell/session mutation and do not activate it via GitHub-source redeploy under the current user constraint.

As soon as that activation exists, perform exactly this B2 batch in the existing workforce:

`ProviderEnv()` adds `OPENAI_BASE_URL` → focused `config_test.go` regression → `go test ./internal/config` → `make check` → reload only PR-AF → restore/use one CURRENT control-plane execution route → synthetic `pr-af.review` → prove `openai/<model> + OPENAI_API_KEY + OPENAI_BASE_URL` and no unintended fallback.

## Write-back rule

Update this file in place after every bounded batch. Do not create a second PR-AF plan/status document.
