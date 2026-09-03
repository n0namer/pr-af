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

Phase: semantic acceptance/durability complete → quality baseline/hardening.

### Phase Goal

Establish a small evidence-backed quality baseline on the durable Go PR-AF: preserve the proven generic OpenAI-compatible broker path, fail-closed primary review semantics, and weak-model recovery while measuring review usefulness across a bounded mix of buggy and clean real-file fixtures before any broader optimization.

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
| Durable application source | accepted live app tree was canonicalized to `dev`; squash commit `1967bb2275855d8f7626806169b2a274b379c9e0` was independently compared against accepted runtime commit `5a0f3b2b2c6c37d5cecab140cd2a0938c1715b7f`: `go/` diff empty, 9/9 intended blobs MATCH | PASS |
| Maintained package | `/afhome/installed.yaml`: `pr-af.source_path=/src/pr-af/go`, status running, port `8007` | PASS |
| Latest loaded process after coverage-policy reload | PID `105691` from `/src/pr-af/go` | PASS |
| Generic provider contract | live PR-AF uses `OPENAI_BASE_URL=http://fcm-dev-internal:19280/v1`, `OPENAI_API_KEY` present, `PR_AF_MODEL=openai/fcm`, `PR_AF_AI_MODEL=openai/fcm` | PASS |
| Broker transport | direct `/chat/completions` probe with current broker key returned HTTP 200; tool-calling probe returned valid tool calls | PASS |
| OpenCode transport | external `openai/<model>` is internally adapted to dedicated `compat/<model>` via `@ai-sdk/openai-compatible`; runtime process trees proved `compat/fcm → fcm/fcm`; no unintended OpenRouter provider selection | PASS |
| Direct `.ai()` path | uses the same OpenAI-compatible key/base/model identity; partial key/base config rejected | PASS |
| Weak-model tolerance | meta planning uses lean structured schema → deterministic JSON recovery → one plain-text line-protocol fallback → fail-closed; operational ids/budgets/defaults are deterministic PR-AF policy | SOURCE + TEST PASS |
| False-safe prevention | unrecoverable meta output no longer becomes `dimensions=0 → Looks Good` | PASS |
| Meta runtime | repeated real FCM canaries completed `meta_semantic`, `meta_mechanical`, and `meta_systemic` 3/3 | PASS |
| Downstream DAG | clean real-file canary `run_20260903_114900_74zj0kf2` completed intake, anatomy, meta 3/3, primary review, coverage gate, coverage-added review, obligation verification, adversary, and root review | PASS |
| Coverage-only failure policy | `runCoverageLoop` preserves existing primary findings and stops further coverage expansion if a coverage-added reviewer fails; primary review path remains fatal. Targeted regression PASS; accepted clean canary also completed the coverage-added reviewer without regressing root behavior | PASS |
| Deterministic validation | `go test ./internal/orch` PASS and full live-source `make check` PASS (`build + vet + all tests`) on the exact accepted source later proven byte-identical to durable `dev` | PASS |
| Acceptance fixture discipline | historical off-tree auth/payments synthetic fixture is retired for acceptance because live workspace context contaminated planner/reviewer routing; acceptance now uses real existing-file diffs or isolated valid fixtures | RESOLVED |
| Clean semantic canary `run_20260903_114900_74zj0kf2` | real diff over existing `go/internal/schemas/defaults.go` changed `MaxDurationSeconds: 60 → 0`; root succeeded in ~537s and post-completion payload reported a critical blocking finding with evidence on the changed default and real consumers, producing merge-blocking output | PASS |
| B4 clean-negative canary `run_20260903_131346_m6jyj65p` | comment-only diff over the same real file completed the full review path in ~416.5s; terminal payload returned `findings=[]`, `blocking_count=0`, no severities, event `APPROVE`, and `Looks Good / Safe to merge`; one degraded dimension was reported but produced no fabricated finding | PASS |
| B4 recall canary `run_20260903_132450_lcvyp3i6` | real `go/internal/node/node.go` diff inverted the key/base XOR guard; planner generated an exact dimension to verify pair enforcement, but both primary `review_dimension` children returned `findings=[]`, `schema_parse_failed=false`; root returned `APPROVE` with 0 findings. Finding was not lost in scoring/output; reviewer reasoning itself missed the semantic regression | FAIL / REVIEWER RECALL GAP |

## Accepted application delta

The following intended files were accepted in runtime, canonicalized to `dev`, and independently verified byte-identical between accepted runtime commit `5a0f3b2b2c6c37d5cecab140cd2a0938c1715b7f` and durable squash commit `1967bb2275855d8f7626806169b2a274b379c9e0`:
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
6. **COVERAGE_ADDITIVE_FAILURE_POLICY — RESOLVED.** Coverage-only reviewer failure no longer discards primary findings; targeted regression PASS and clean accepted runtime preserved root behavior.
7. **ACCEPTANCE_FIXTURE_DRIFT — RESOLVED.** Historical off-tree auth/payments fixture is retired; acceptance uses real existing-file diffs or isolated valid fixtures.
8. **PRIMARY_REVIEWER_NO_PROGRESS — NOT REPRODUCED ON VALID FIXTURE.** The 360s timeout occurred on the contaminated historical fixture; the clean real-file canary completed primary review in ~80s. Primary fail-closed semantics remain unchanged.
9. **DURABILITY_DRIFT — RESOLVED.** Exact accepted app tree is durable on `dev` at squash commit `1967bb2275855d8f7626806169b2a274b379c9e0`; 9/9 intended blobs MATCH accepted runtime commit and `go/` diff is empty.

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
- [x] Clean semantic acceptance canary using a valid real-existing-file diff/fixture.
- [x] Primary reviewer no-progress did not reproduce on the valid fixture; primary fail-closed semantics remain unchanged.
- [x] Coverage-only degradation policy has targeted regression evidence; clean accepted runtime also completed the additive coverage reviewer without root regression.
- [x] Root review terminal success with inspected useful findings.
- [x] Exact accepted live delta canonicalized to `pr-af:dev` via PR #8.
- [x] Durable canonical SHA `1967bb2275855d8f7626806169b2a274b379c9e0` verified against accepted runtime delta with empty `go/` diff and 9/9 blob MATCH.
- [x] Untracked runtime/test artifacts excluded from canonicalization.

## Bounded BMAD batches

### B1 — Maintained runtime path
Status: **PASS**.

### B2 — Generic provider + model-tolerance contract
Status: **PASS — LIVE + DURABLE**.

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
Status: **PASS — LIVE + DURABLE**.

Delivered:
- clean real-existing-file semantic canary;
- primary reviewer completed without reproducing the contaminated-fixture timeout;
- full downstream review/coverage/obligation/adversary DAG completed;
- terminal payload inspected with a critical blocking finding on the injected duration-budget regression;
- exact accepted app tree canonicalized by PR #8;
- durable squash SHA `1967bb2275855d8f7626806169b2a274b379c9e0` independently verified byte-identical to the accepted runtime app tree.

### B4 — Quality baseline / low-false-positive hardening
Status: **ACTIVE — CLEAN NEGATIVE PASS / SECOND RECALL FIXTURE FAIL**.

Current ~30-minute sub-batch:
1. preserve the passed clean-negative baseline (`run_20260903_131346_m6jyj65p`: 0 findings, APPROVE);
2. repair the localized reviewer-reasoning recall gap from `run_20260903_132450_lcvyp3i6` without weakening fail-closed semantics;
3. keep the canonical Python/golden prompt builder unchanged; add the smallest review-dimension runtime instruction that makes the proposed diff authoritative for changed lines and requires explicit old/new predicate semantics checking before returning no findings;
4. add a targeted regression for that reviewer contract;
5. run targeted tests + full `make check` on exact live source;
6. reload only PR-AF and rerun the same XOR inversion fixture;
7. PASS requires an evidence-grounded blocking finding on the broken key/base invariant and no regression of the clean-negative behavior.

## Current blocker

No provider/runtime blocker is open. The current blocker is a **reviewer-reasoning recall gap**: on `run_20260903_132450_lcvyp3i6`, planning correctly requested verification of the `OPENAI_API_KEY` / `OPENAI_BASE_URL` pair invariant, both primary `review_dimension` calls completed with `schema_parse_failed=false`, but both returned `findings=[]`, and root incorrectly approved the inverted XOR guard. The finding was not lost in scoring, adversary, or output synthesis.

## ONE next move

In persistent DEV `/src/pr-af`, patch only the review-dimension runtime reasoning contract so reviewers must treat the supplied diff as authoritative for changed lines and explicitly compare removed vs added guard semantics before concluding `no findings`; validate with targeted tests + full `make check`, reload only PR-AF, then rerun the exact XOR recall canary. Do not revisit broker/provider work unless fresh CURRENT evidence regresses that layer.

## Write-back rule

Update this file in place after every bounded batch. Do not create a second PR-AF plan/status document. GitHub is canonical write-back/release only; application code continues to be edited and verified in the persistent DEV container first.
