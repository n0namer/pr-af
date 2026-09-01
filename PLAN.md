# PR-AF — Canonical Plan and Current State

> Status: ACTIVE — B1 PASS / B2 EXECUTION-BRIDGE BLOCKER
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
| Go provider env propagation | `go/internal/config/ai.go::ProviderEnv()` forwards `OPENAI_API_KEY` but not `OPENAI_BASE_URL` | SOURCE FACT / RUNTIME HYPOTHESIS |
| Bootstrap lifecycle route | SWE-style workforce bootstrap now installs `/src/pr-af/go` and starts `pr-af:8007`; B1 no longer depends on a separate typed install capability | RESOLVED |
| External execution bridge | stale-stack binding was diagnosed and current control-plane is now on shared `coolify` network with unique alias `agentfield-current-control-plane`; gateway process has the correct unique upstream + upstream API-key injection, but external ingress remains unavailable because Coolify strips custom Traefik labels from the custom-service runtime and generic Traefik file-write is blocked by the current mediator | EXECUTION_BRIDGE_BLOCKER |

## Design/runtime drift register

1. **INSTALL_PATH_DRIFT** — topology installs `/src/pr-af` root while maintained Go runtime requires `/src/pr-af/go` for local-path installation.
2. **START_POLICY_DRIFT** — maintained Go PR-AF is not started in CURRENT permanent DEV, so registration/functional proof cannot exist yet.
3. **PROVIDER_CONTRACT_CANDIDATE** — workforce already exposes Gonka/OpenAI-compatible key + base, but Go `ProviderEnv()` currently omits `OPENAI_BASE_URL`; this is not promoted to a confirmed runtime defect until the maintained Go node is exercised.
4. **EXECUTION_CAPABILITY_GAP** — available DEV container mediator can read the workspace but cannot currently execute the exact `af install/run` mutation through a typed capability.

## Phase DoD

- [x] Canonical `PLAN.md` exists in the PR-AF repo and owns current PR-AF status.
- [x] Exact live `/src/pr-af` source identity read back.
- [x] Live application-code identity established at `c2953a...`; canonical `dev` is ahead only by documentation commits.
- [ ] Repository/base HEAD reconciled before any SourceLoop code capture.
- [x] Maintained Go package path and local-root escape-hatch semantics identified.
- [x] Current orchestration install/start drift identified.
- [x] Provider env/config source inspected without exposing secret values.
- [ ] Maintained Go PR-AF installed from `/src/pr-af/go` in permanent DEV without full-stack redeploy.
- [ ] `pr-af` process started on the intended port and loaded from the intended source.
- [ ] `pr-af` registered in CURRENT AgentField.
- [ ] Expected reasoner surface discoverable.
- [ ] Deterministic/functional canary PASS.
- [ ] Gonka/OpenAI-compatible path proven end-to-end if used: configured base survives to the actual model/tool call and no OpenRouter fallback occurs.
- [ ] One bounded real semantic PR-review E2E PASS with non-dummy provider credentials and outcome inspection.
- [ ] If code changed: smallest relevant regression red→green, then required broader Go validation.
- [ ] Accepted code delta captured by SourceLoop.
- [ ] Durable `pr-af:dev` SHA corresponds to the accepted runtime delta.
- [ ] No unintended container-only code delta remains.

## Bounded BMAD batches

### B1 — Correct maintained runtime path
Goal: move from “source present” to “maintained Go node actually running”.

Tasks:
1. Re-read live source identity and relevant install state.
2. Install `/src/pr-af/go` using the existing AgentField CLI in the same workforce.
3. Start only `pr-af` on port `8007`.
4. Read back installed state, process/log evidence, AgentField registration, and reasoner inventory.
5. Do not change PR-AF code unless Go startup itself proves a code/config defect.

DoD:
- maintained Go package installed;
- `pr-af` running and registered;
- reasoner discoverable;
- exact blocker recorded if any criterion cannot be executed.

Status: `DEPLOYING_RUNTIME_BOOTSTRAP_FIX`.
Current state: comparison with working SWE-AF proved that maintained package lifecycle belongs in the workforce bootstrap itself, not in a separate typed operator capability. Universal Solver target commit `99c45ce...` now installs `/src/pr-af/go` and starts `pr-af:8007` with explicit `opencode + openai/<Gonka>` selection. Deployment `jgqxk6jxe8krtxj9xzfqkr9u` is creating the new generation; B1 remains open until the new workforce readback proves install/start/registration.

### B2 — Provider path and minimal Go repair
Goal: prove or repair the real provider/model path only after B1 starts the maintained node.

Priority hypothesis:
- if Gonka/OpenAI-compatible routing is selected, verify whether omission of `OPENAI_BASE_URL` from `ProviderEnv()` causes endpoint fallback;
- fix directly in `/src/pr-af/go` only with reproduced evidence;
- add the smallest regression test;
- targeted test → related regression → `make check` as required;
- same-runtime PR-AF reload only.

Status: `BLOCKED_BY_B1`.

### B3 — Semantic acceptance and durability
Goal: prove useful PR review and canonicalize the exact accepted delta.

Tasks:
- deterministic/functional canary;
- bounded real PR review with outcome inspection;
- correlate source SHA, process/node identity, execution id, model/provider, logs, and result;
- reject HTTP/health/execution-success-only evidence as semantic proof;
- SourceLoop capture only after acceptance;
- verify final `pr-af:dev` SHA and absence of unintended runtime-only delta.

Status: `BLOCKED_BY_B1_B2`.

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

The project is not blocked on understanding the next technical action. The exact next runtime mutation is known and bounded, but the CURRENT DEV execution surface lacks a typed capability that can perform it. Generic command execution rejects `af install` as opaque, and the generic approval adapter is unavailable.

This is `CAPABILITY_GAP`, not missing user authorization and not evidence of a PR-AF code defect.

## Current decision / ONE next move

Use the first CURRENT authoritative typed route that can execute in the existing permanent workforce, then perform exactly:

`af install /src/pr-af/go` → `af run pr-af --port 8007 --detach=true` → installed/process/log readback → AgentField registration → reasoner discovery.

Do not patch PR-AF Go code before that runtime path is exercised. If B1 produces a provider-path failure, B2 edits only `/src/pr-af` directly and proves the smallest red→green delta before SourceLoop capture.

## Write-back rule

Update this file in place after every bounded batch. Do not create a second PR-AF plan/status document.
