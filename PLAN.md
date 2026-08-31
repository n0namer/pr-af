# PR-AF — Canonical Plan and Current State

> Status: ACTIVE
> Updated: 2026-08-31
> Canonical owner: `n0namer/pr-af:dev/PLAN.md`
> Active development branch: `dev`
> Runtime topology owner: `n0namer/universal-solver`
> BMAD lane: `bmad-quick-dev` / implementation-debugging

## Authority and anti-drift

- This file owns PR-AF current phase, Phase Goal, bounded batches, DoD, progress, blockers, and ONE next move.
- `README.md` and `docs/ARCHITECTURE.md` own product/architecture intent.
- `AGENTS.md` owns repository engineering rules.
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

## CURRENT evidence — 2026-08-31

| Claim | Evidence | Verdict |
|---|---|---|
| Canonical PR-AF branch | `n0namer/pr-af:dev` contains documentation-only `PLAN.md` / `AGENTS.md` commits after application-code commit `c2953a48792aed2bdf15fb31f38507e676f3fb41`; no PR-AF application-code delta is introduced by those documentation commits | FACT |
| Live source identity | permanent workforce `/src/pr-af/.git/HEAD` = application-code commit `c2953a48792aed2bdf15fb31f38507e676f3fb41` | FACT |
| Source/runtime code identity | live application code still matches the canonical application-code base at `c2953a...`, while repository HEAD is ahead by documentation-only commits; reconcile repository/base identity before SourceLoop code capture | DOC_ONLY_HEAD_DRIFT |
| Current DEV topology target | Coolify app `edshqtkwskg3lrczekhcmd71` is pinned to `universal-solver` commit `b78866efd17cfdca232019e66e902e16c778c152` | FACT |
| Maintained implementation | `go/` package; node id `pr-af`; default port `8007`; build `./cmd/pr-af`; start `bin/pr-af` | FACT |
| Local install semantics | local-path install of repo root keeps legacy Python package; maintained Go local path is `/src/pr-af/go` | FACT |
| Current orchestration install path | workforce Compose still runs `af install /src/pr-af` | DESIGN_RUNTIME_DRIFT |
| Current orchestration start policy | PR-AF is intentionally not auto-started after install | FACT |
| Provider wiring in workforce | topology injects `OPENAI_API_KEY`, `OPENAI_BASE_URL`, `OPENROUTER_API_KEY`, and `GH_TOKEN` by name; secret values are not recorded here | FACT |
| Go manifest provider contract | `go/agentfield-package.yaml` currently requires `OPENROUTER_API_KEY`; `GH_TOKEN` is optional | FACT |
| Go provider env propagation | `go/internal/config/ai.go::ProviderEnv()` forwards `OPENAI_API_KEY` but not `OPENAI_BASE_URL` | SOURCE FACT / RUNTIME HYPOTHESIS |
| CURRENT AgentField registration | current node inventory contains no `pr-af` node | FACT |
| Native install/start route | exact desired operation is `af install /src/pr-af/go` then `af run pr-af --port 8007 --detach=true` | DECISION |
| Current execution capability | DEV guidance allows the scoped operation, but `execContainer` rejects `af install` as `REVIEW_REQUIRED: opaque_or_unknown_mutation`; `prepareChange` returns `approval_capability_gap` because no executable generic approval adapter exists | CAPABILITY_GAP |

## Design/runtime drift register

1. **INSTALL_PATH_DRIFT** — topology installs `/src/pr-af` root while maintained Go runtime requires `/src/pr-af/go` for local-path installation.
2. **START_POLICY_DRIFT** — maintained Go PR-AF is not started in CURRENT permanent DEV, so registration/functional proof cannot exist yet.
3. **PROVIDER_CONTRACT_CANDIDATE** — workforce already exposes Gonka/OpenAI-compatible key + base, but Go `ProviderEnv()` currently omits `OPENAI_BASE_URL`; this is not promoted to a confirmed runtime defect until the maintained Go node is exercised.
4. **EXECUTION_CAPABILITY_GAP** — available DEV container mediator can read the workspace but cannot currently execute the exact `af install/run` mutation through a typed capability.

## Phase DoD

- [x] Canonical `PLAN.md` exists in the PR-AF repo and owns current PR-AF status.
- [x] Exact live `/src/pr-af` source identity read back.
- [x] Durable `pr-af:dev` identity matches live source.
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

Status: `PARTIAL_CAPABILITY_BLOCKER`.
Current blocker: no available typed execution capability can run the exact scoped `af install/run` operation in the existing workforce; shell bypass and full-stack redeploy are not accepted inner-loop substitutes.

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
