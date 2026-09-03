# PR-AF — Canonical Plan and Current State

> Status: ACTIVE — B1 PASS / B2 PROVIDER + MODEL-TOLERANCE PASS / B3 SEMANTIC ACCEPTANCE + DURABILITY PASS
> Updated: 2026-09-03
> Canonical owner: `n0namer/pr-af:dev/PLAN.md`
> Active development branch: `dev`
> Runtime topology owner: `n0namer/universal-solver`
> BMAD lane: `bmad-help` → `bmad-quick-dev` / implementation-debugging

## Authority and anti-drift

- This file owns PR-AF phase, Phase Goal, bounded batches, DoD, progress, blockers, and ONE next move.
- `README.md` / `docs/ARCHITECTURE.md` own product architecture; `AGENTS.md` owns repository engineering rules.
- CURRENT runtime/readback owns actual loaded state.
- Application code is edited and verified directly in persistent DEV `/src/pr-af`; GitHub is canonical write-back/release, not the inner debug loop.
- `universal-solver` owns permanent AgentField DEV topology and SourceLoop plumbing.
- `BMAD-MNNZ` is the workflow/skill rulebook, not project SoT.
- After every material state/evidence change: VERIFY → update this file → replan from CURRENT evidence.
- Never claim DONE without semantic payload evidence and durable source capture.

## North Star

PR-AF is an AgentField PR-review node for deep evidence-grounded review: convert a pull request into a task-specific review plan, run focused reviewers, ground findings in code evidence, challenge weak claims, synthesize compound risks, close coverage gaps, verify obligations, and return useful review output with strong recall and low false positives.

Engineering North Star for this workstream: a verified-running maintained Go PR-AF in permanent DEV that works through a generic OpenAI-compatible broker env contract, tolerates weak/non-JSON models without false-safe output, completes a real semantic review DAG, and has exact source/runtime provenance.

## Current phase

Phase: implementation / semantic acceptance / durability.

### Phase Goal

Finish one bounded real-model semantic review on the maintained Go PR-AF using the generic broker env contract, preserve fail-closed semantics for primary reviewer failures, degrade safely only for additive coverage work, then SourceLoop-capture the exact accepted live delta to `pr-af:dev`.

## Operating contract

`OBSERVE → LOCALIZE → PATCH IN /src/pr-af → TARGETED VERIFY → FULL VERIFY → PR-AF-ONLY RELOAD → RUNTIME PROOF → CANONICALIZE → VERIFY SHA → WRITE-BACK`

Hard rules:
- no GitHub-first application coding/redeploy loop;
- preserve SWE / Deep Research and unrelated runtime state;
- secrets presence-only in evidence;
- primary reviewer failures remain fail-closed;
- additive coverage reviewer failures may stop further coverage expansion without discarding already-proven primary findings;
- work in bounded ~30-minute BMAD batches with explicit DoD and 80/20 priority.

## CURRENT evidence — 2026-09-03

| Claim | CURRENT evidence | Verdict |
|---|---|---|
| Live application source | persistent `/src/pr-af`, base Git HEAD `6245796a6c47a0f114dd0e8382f4abf63a89752f`, accepted application delta still dirty and not SourceLoop-captured | FACT / DURABILITY PENDING |
| Maintained package | `/afhome/installed.yaml`: `pr-af.source_path=/src/pr-af/go`, status running, port `8007` | PASS |
| Latest loaded process after coverage-policy reload | PID `105691` from `/src/pr-af/go` | PASS |
| Generic provider contract | live PR-AF uses `OPENAI_BASE_URL=http://fcm-dev-internal:19280/v1`, `OPENAI_API_KEY` present, `PR_AF_MODEL=openai/fcm`, `PR_AF_AI_MODEL=openai/fcm` | PASS |
| Broker transport | direct `/chat/completions` probe with current broker key returned HTTP 200; tool-calling probe returned valid tool calls | PASS |
| OpenCode transport | external `openai/<model>` is internally adapted to dedicated `compat/<model>` via `@ai-sdk/openai-compatible`; runtime process trees proved `compat/fcm → fcm/fcm`; no unintended OpenRouter provider selection | PASS |
| Direct `.ai()` path | uses the same OpenAI-compatible key/base/model identity; partial key/base config rejected | PASS |
| Weak-model tolerance | meta planning uses lean structured schema → deterministic JSON recovery → one plain-text line-protocol fallback → fail-closed; operational ids/budgets/defaults are deterministic PR-AF policy | SOURCE + TEST PASS |
| False-safe prevention | unrecoverable meta output no longer becomes `dimensions=0 → Looks Good` | PASS |
| Meta runtime | repeated real FCM canaries completed `meta_semantic`, `meta_mechanical`, and `meta_systemic` 3/3 | PASS |
| Downstream DAG | real canaries executed multiple `review_dimension`, `coverage_gate`, `extract_obligations`, and several `verify_obligation` children | PASS / ROOT ACCEPTANCE NOT YET STABLE |
| Coverage-only failure policy | `runCoverageLoop` now preserves existing primary findings and stops further coverage expansion if a coverage-added reviewer fails; primary review path remains unchanged/fatal | SOURCE PASS; targeted regression PASS; runtime-specific proof pending |
| Deterministic validation | latest `go test ./internal/orch` PASS and full live-source `make check` PASS (`build + vet + all tests`) after coverage policy patch | PASS |
| Old synthetic acceptance fixture | auth/payments paths do not exist in the real repo and the live dirty workspace dominated planner/reviewer context; latest run assigned a primary reviewer to dirty `go/internal/orch/phases.go` | INVALID FIXTURE / DRIFT |
| Latest canary `run_20260903_084948_bfbvulrr` | intake/anatomy/meta 3/3 passed; first primary reviewer passed; second primary reviewer made no progress for 360s and root failed before coverage. This is correct fail-closed behavior, but not valid evidence against the coverage-only policy because the fixture was contaminated | FAILED / FIXTURE INVALID |

## Live dirty application delta

Tracked intended files currently include:
- `go/internal/config/ai.go`
- `go/internal/config/config_test.go`
- `go/internal/node/node.go`
- `go/internal/node/node_test.go`
- `go/internal/reasoners/contexts_test.go`
- `go/internal/reasoners/meta.go`
- `go/internal/reasoners/reasoners_test.go`
- `go/internal/orch/phases.go`
- `go/internal/orch/degradation_test.go`

Observed untracked runtime/test artifacts that must not be silently canonicalized:
- `go/analysis.json`
- `go/test/e2e/pr-af-review.json`

## Design/runtime drift register

1. **INSTALL_PATH_DRIFT — RESOLVED.** Maintained Go package is installed from `/src/pr-af/go`.
2. **PROVIDER_CONTRACT_DRIFT — RESOLVED.** Canonical app-level contract is generic `OPENAI_API_KEY + OPENAI_BASE_URL + openai/<model>`; FCM is only the current env-configured broker endpoint.
3. **OPENCODE_BUILTIN_OPENAI_MISMATCH — RESOLVED.** Built-in OpenCode `openai` used OpenAI-native behavior incompatible with the broker; dedicated `@ai-sdk/openai-compatible` adapter is proven.
4. **WEAK_MODEL_STRUCTURED_OUTPUT_DRIFT — RESOLVED FOR META.** Model output is treated as untrusted input; structured output is preferred, not required; recovery is deterministic and fail-closed.
5. **FALSE_SAFE_META_FALLBACK — RESOLVED.** Parse/provider failure cannot silently become an empty safe review.
6. **COVERAGE_ADDITIVE_FAILURE_POLICY — SOURCE RESOLVED / RUNTIME PROOF PENDING.** Coverage-only reviewer failure no longer discards primary findings.
7. **ACCEPTANCE_FIXTURE_DRIFT — OPEN.** Historical off-tree synthetic auth/payments diff is invalid in the dirty real workspace; acceptance must use a real existing-file diff or isolated valid fixture.
8. **PRIMARY_REVIEWER_NO_PROGRESS — CURRENT BLOCKER.** Latest contaminated canary hit a primary `review_dimension` no-progress timeout at 360s. Do not weaken primary fail-closed. First reproduce on a valid real-file acceptance fixture before patching reviewer/harness behavior.
9. **DURABILITY_DRIFT — OPEN.** Accepted live application delta is not yet SourceLoop-captured to canonical `pr-af:dev`.

## Phase DoD

- [x] Canonical `PLAN.md` owns current PR-AF state.
- [x] Maintained Go package installed/running/registered on `8007`.
- [x] Generic OpenAI-compatible key/base/model contract proven live.
- [x] No unintended OpenRouter provider selection.
- [x] `@ai-sdk/openai-compatible` transport proven.
- [x] Weak/non-JSON model tolerance implemented and regression-tested for meta planning.
- [x] Meta fail-closed semantics proven.
- [x] 3/3 meta lenses completed through the current broker/model.
- [x] Downstream reviewer / coverage / obligation DAG proven to execute.
- [x] Coverage-only reviewer failure policy implemented + targeted regression PASS.
- [x] Full live-source `make check` PASS after latest patch.
- [ ] Clean semantic acceptance canary using a valid real-existing-file diff/fixture.
- [ ] Primary reviewer no-progress either does not reproduce on valid fixture or is repaired without fail-open semantics.
- [ ] If coverage-only failure is triggered, runtime proves primary findings survive and root continues.
- [ ] Root review terminal success with inspected useful findings.
- [ ] Exact accepted live delta SourceLoop-captured to `pr-af:dev`.
- [ ] Durable canonical SHA verified against accepted runtime delta.
- [ ] Untracked runtime/test artifacts excluded unless intentionally owned.

## Bounded BMAD batches

### B1 — Maintained runtime path
Status: **PASS**.

### B2 — Generic provider + model-tolerance contract
Status: **PASS IN LIVE SOURCE/RUNTIME; DURABILITY PENDING**.

Delivered:
- generic OpenAI-compatible env contract;
- current FCM endpoint only as env-configured broker;
- dedicated OpenCode compatible provider;
- direct `.ai()` same provider identity;
- no OpenRouter runtime fallback;
- lean meta schema + JSON/text recovery;
- meta fail-closed;
- targeted regressions + repeated `make check`;
- real 3/3 meta runtime proof.

### B3 — Semantic acceptance and durability
Status: **ACTIVE**.

Current sub-batch:
1. replace invalid off-tree auth/payments synthetic canary with a bounded diff over a real existing PR-AF file that does not overlap the current dirty patch;
2. rerun `pr-af.review` on the already validated binary/source;
3. if primary reviewer no-progress reproduces, localize reviewer/harness recovery while preserving primary fail-closed;
4. if root succeeds, inspect semantic findings and coverage/obligation behavior;
5. SourceLoop-capture only the exact intended tracked delta;
6. verify canonical `pr-af:dev` SHA and update this file with terminal evidence.

## Current blocker

Not FCM, token, OpenRouter, meta JSON, install, or registration.

The current blocker is **clean semantic acceptance**. The latest canary used an invalid off-tree synthetic diff and was contaminated by the live dirty workspace; it then hit a primary reviewer no-progress timeout. That failure correctly remained fatal, but it is not yet evidence of a product defect independent of the invalid fixture.

## ONE next move

Run one bounded `bmad-quick-dev` acceptance batch using a **real existing-file diff that does not overlap the current dirty application patch**. Observe whether primary reviewer no-progress reproduces.

- If it does **not** reproduce: inspect terminal findings, prove coverage behavior if triggered, then SourceLoop-capture the accepted delta.
- If it **does** reproduce: patch the smallest reviewer/harness recovery path while keeping primary failures fail-closed, validate with targeted tests + `make check`, reload only PR-AF, and rerun the same valid fixture.

Do not spend the next batch on broker/provider work unless fresh CURRENT evidence regresses that layer.

## Write-back rule

Update this file in place after every bounded batch. Do not create a second PR-AF plan/status document. GitHub is canonical write-back/release only; application code continues to be edited and verified in the persistent DEV container first.
