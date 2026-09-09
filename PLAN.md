# PR-AF — Canonical Plan and Current State

> Status: ACTIVE — B1 PASS / B2 PROVIDER + MODEL-TOLERANCE PASS / B3 SEMANTIC ACCEPTANCE + DURABILITY PASS
> Updated: 2026-09-09
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

## CURRENT evidence — 2026-09-09

| Claim | CURRENT evidence | Verdict |
|---|---|---|
| Durable application source | accepted live app tree was canonicalized to `dev`; squash commit `1967bb2275855d8f7626806169b2a274b379c9e0` was independently compared against accepted runtime commit `5a0f3b2b2c6c37d5cecab140cd2a0938c1715b7f`: `go/` diff empty, 9/9 intended blobs MATCH | PASS |
| Maintained package | `/afhome/installed.yaml`: `pr-af.source_path=/src/pr-af/go`, status running, port `8007` | PASS |
| Runtime/load state | Last accepted runtime identity was PID `156016` with health PASS and the quick-meta serialization build. Fresh OBSERVE later found **no PR-AF process** and `127.0.0.1:8007/health` refused connection. A subsequent scoped `af install /src/pr-af/go --force --json` attempt timed out at the transport boundary; post-state is not yet readable because Portwing/Docker target resolution is timing out. Do **not** assume either stopped or loaded until post-state readback succeeds. | AMBIGUOUS POST-MUTATION / READBACK BLOCKED |
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
Status: **ACTIVE — REAL-REPO E2E PASS / CHANGE-CAUSALITY REPAIR DETERMINISTIC PASS / BUDGET FALSE-SAFE REPAIR DETERMINISTIC PASS / RUNTIME ACCEPTANCE NEXT**.

Current ~30-minute BMAD batch (2026-09-09):
1. `bmad-help` was activated from canonical `BMAD-MNNZ`; target project is `n0namer/pr-af`, phase is implementation/debugging, and `bmad-quick-dev` is the selected specialist lane. No duplicate BMAD/spec document was created; this `PLAN.md` remains the project SoT.
2. anti-drift OBSERVE confirmed persistent DEV source is still detached at base `5a0f3b2b2c6c37d5cecab140cd2a0938c1715b7f`. Current tracked dirty candidate surface is `go/agentfield-package.yaml`, `go/internal/reasoners/meta.go`, `go/internal/reasoners/reasoners_test.go`, and `go/internal/reasoners/reviewdim.go`; runtime/debug artifacts remain untracked and excluded from canonicalization.
3. bootstrap admission drift was repaired live in `go/agentfield-package.yaml`: provider admission now supports one of the supported provider keys instead of hard-requiring OpenRouter; `af show-requirements` parsed the manifest and `af install /src/pr-af/go --force` rebuilt the maintained package. SourceLoop capture: `vtchg_ab44785edb194ab5b9f77d773f8749be` (pending canonical write-back).
4. previous belief that OpenCode was absent was false: canonical wrapper `/afhome/bin/opencode` and pinned runtime `/afhome/opencode-runtime/v1.17.15/opencode` both exist. The wrapper owns the workforce OpenCode config/model adaptation; PR-AF now uses the wrapper rather than the raw binary.
5. CURRENT runtime after PR-AF-only reload is PID `99844`, health PASS, `AGENT_CALLBACK_URL=http://workforce:8007`, `PR_AF_PROVIDER=opencode`, `PR_AF_MODEL=openai/fcm`, `PR_AF_HARNESS_BIN=/afhome/bin/opencode`. Shared workforce and SWE were not restarted.
6. the pre-fix direct `meta_mechanical` canary `exec_20260909_091208_nupf9vnh` failed before semantic output. Wrapper diagnostics showed `cwd=/src/pr-af`, while OpenCode attempted `/go/internal/node/node.go`, classified it as external-directory access, and auto-rejected the read; no schema output file was created. The wrapper also does not parse the SDK incremental-schema wording into its optional output shim, but that shim is not required when the model follows the file protocol.
7. an independent OpenCode micro-canary using the same wrapper/model with explicit absolute repository path successfully read `/src/pr-af/go/internal/node/node.go` and wrote the requested JSON artifact. This proved the harness/model/write-tool path itself works.
8. live `go/internal/reasoners/meta.go` now appends explicit repository-root/path-resolution guidance to all meta lenses; regression `TestMetaSelectorAnchorsRepositoryRelativePaths` was added to `go/internal/reasoners/reasoners_test.go`. Targeted `go test ./internal/reasoners` PASS, full `go test ./...` PASS, `go vet ./...` PASS, and `go build ./...` PASS.
9. after reinstall + PR-AF-only reload, direct `meta_mechanical` `exec_20260909_113330_qrla1drq` **SUCCEEDED** in ~95.6s and returned an evidence-grounded mechanical dimension referencing the real repository file. This closes the immediate L2 meta execution blocker and proves repository investigation no longer fails on `/go/...` path drift when `repo_path=/src/pr-af` is supplied.
10. first full XOR root attempt `exec_20260909_113531_3bj0xhg9` intentionally exercised the raw-diff production path without `repo_path`. It reached intake/anatomy/meta but failed after ~460s with `CLI command made no progress for 360s` in `meta_semantic`. CURRENT process evidence showed the harness was invoked with `--dir /afhome/packages/pr-af`; that package directory does not own the live source tree, so this run is **runtime-context failure, not reviewer-recall evidence**.
11. ReviewInput already supports `diff_text` and `repo_path` together: raw diff remains the reviewed patch while `repo_path` is retained as repository context for anatomy/meta/evidence. Therefore the smallest safe correction is input/runtime context, not another code patch.
12. corrected XOR root canary `exec_20260909_114330_ce3vdgsf` used the same diff plus `repo_path=/src/pr-af`; CURRENT child-process readback proved OpenCode was invoked with `--dir /src/pr-af`. It still failed after 360s in `meta_semantic`, so repository-root routing alone was not sufficient.
13. direct semantic isolation `exec_20260909_115407_kcd3mxuz` reproduced the failure with compact precomputed intake/anatomy and the correct `/src/pr-af` cwd. OpenCode created the required output file and wrote a valid partial object containing `lens=SEMANTIC` and `confidence=0.95`, then made no progress until the 360s watchdog.
14. a first bounded-quick semantic prompt candidate was implemented live and regression-tested. Direct canary `exec_20260909_123512_ax4w5721` still terminated at the 360s SDK watchdog, but per-run OpenCode logs proved the model obeyed the bounded protocol: it read only `go/internal/node/node.go`, correctly identified the `!= -> ==` XOR/XNOR inversion, and wrote a complete evidence-grounded JSON finding to the exact requested AgentField output path before termination.
15. the remaining failure is therefore not semantic quality or repository exploration. `/afhome/bin/opencode` launches the pinned OpenCode child with `stdout`/`stderr` redirected to per-run files and only replays them after the child exits. The AgentField SDK idle watchdog observes the wrapper process, not those files, so legitimate child progress is invisible and the whole process group is killed after 360s. This is **WRAPPER_WATCHDOG_DRIFT**. Do not increase the watchdog or change reviewer logic to mask it.
16. PR-AF-only raw pinned OpenCode experiment was executed by setting `PR_AF_HARNESS_BIN=/afhome/opencode-runtime/v1.17.15/opencode` and restarting only PR-AF as PID `122691`. Direct semantic `exec_20260909_131126_p3n48ooy` failed in ~53.7s with no output file; an isolated raw-CLI micro-canary under the exact PR-AF process environment exited `RC=1` with OpenCode `UnknownError / Unexpected server error` (`ref=err_c3c17334`). Therefore the raw binary is **not** a valid substitute for the workforce wrapper: the wrapper also owns required OpenCode config/model adaptation.
17. A PR-AF-owned streaming wrapper was implemented live at `go/scripts/opencode-stream.sh`. Differential micro-canary proved the wrapper succeeds when it preserves the exact workforce OpenCode config/model adaptation while streaming child stdout/stderr directly; the same canary produced the requested JSON after reading the real source file. The earlier PR-AF-minimal OpenCode config was insufficient and is not accepted as a runtime contract.
18. Direct semantic canary `exec_20260909_134403_4o2doeqy` through the streaming wrapper **SUCCEEDED** in ~112.7s with confidence `0.95`, one focused semantic dimension, and the correct XOR/XNOR diagnosis of `!= -> ==`. This closes the isolated semantic planner/output gate.
19. Full XOR root `exec_20260909_134655_8xrg9iw5` still **FAILED** after ~600s. Intake (~33s) and anatomy (~96s) passed, then all three quick meta lenses launched concurrently. `meta_semantic` and `meta_mechanical` both hit the 360s no-progress watchdog; therefore the remaining root blocker is not isolated semantic correctness but the cost/concurrency behavior of broad quick meta prompts under three simultaneous FCM/OpenCode calls.
20. The smallest next repair was to make quick mode cheap across **all three** meta lenses. `MetaMechanicalPrompt` and `MetaSystemicPrompt` received bounded quick-specific protocols matching the proven semantic quick lane; targeted prompt/reasoner tests, full `go test ./...`, `go vet ./...`, and `go build ./...` all PASS. PR-AF was reinstalled and restarted only as PID `129010`, health PASS, with `AGENT_CALLBACK_URL=http://workforce:8007`, `PR_AF_PROVIDER=opencode`, `PR_AF_MODEL=openai/fcm`, and the PR-AF streaming wrapper.
21. Real-repository acceptance was moved to clean `cloudsecurity-af` using committed range `985234b5ed79323afd2a8a7b5d975c893ac4394f..6c1133cae087f98887c6dfbe7b8248e414b5c02f`. The diff changes only `tests/test_config.py`: `QUICK == 10 -> 20` and adds `STANDARD == 30`; `DEPTH_PROVER_CAPS` is already `{QUICK:20, STANDARD:30, THOROUGH:10000}` in both base and head source, so this is a stale-test correction with no production-code change.
22. Root review `exec_20260909_144921_135drh7t` **SUCCEEDED** end-to-end. It returned 7 findings, 0 blocking; 3 `important`, 4 `suggestion`. Every final finding had `diff_line=null`. Findings were factually grounded but PR-irrelevant/pre-existing: missing Python CI, historical stale-test drift, undocumented cap rationale, and an unrelated hard-coded fallback in `reasoners/phases.py`. No finding identified a defect introduced by the selected test-only change. This is a **PRECISION / CHANGE-CAUSALITY GAP**, not a runtime failure.
23. The 80/20 repair was restricted to PR relevance and implemented live in `/src/pr-af/go`: `review_dimension` now has a mandatory PR change-causality gate; evidence verification now applies the same causality rule, runs for **all** findings including suggestions, and drops `verified=false` findings before adversary/scoring. Regression `TestEvidenceVerificationCoversSuggestionsAndDropsUnverified` was added; targeted prompt/orchestrator tests PASS, full `go test ./...` PASS, `go vet ./...` PASS, and `go build ./...` PASS. The exact candidate was reinstalled and PR-AF-only restarted; last independently verified runtime was PID `144062`, health PASS.
24. Exact `cloudsecurity-af` replay `exec_20260909_153550_ifztqssl` on the causality-repair runtime **FAILED before reviewer/evidence-verifier acceptance could run**. Root duration was ~1054.8s; `meta_semantic`, `meta_mechanical`, and `meta_systemic` all failed with `CLI command made no progress for 360s`. Their requested `.agentfield_output.json` artifacts were absent. Therefore this replay is **runtime/meta-fan-out failure, not precision evidence**; the causality repair remains unaccepted at runtime.
25. Real-repo smoke `exec_20260909_160214_zy5qk18m` against immutable committed SWE-AF range `f9aec2111d084ab0204d5612f2ea00a562226ac7..0c64fe7cc4fc216f4d32d0b855015509750eb4aa` **SUCCEEDED** with `findings=[]`. Because SWE-AF is concurrently edited elsewhere and the historical touched `README.md` differs from current HEAD, this is diagnostic smoke only, not strict acceptance; it does prove the current PR-AF can finish a substantial real-repo review without mutating SWE-AF.
26. `QUICK_META_CONCURRENCY_STARVATION` repair was implemented live and kept narrow: `runMetaSelectors` now serializes semantic → mechanical → systemic only when `depth=quick`; standard/deep retain the existing errgroup fan-out. Regression `TestMetaSelectorQuickModeSerializesLenses` proves quick max concurrency = 1 while the existing adversarial-order test continues to exercise standard concurrent fan-out. `go test ./internal/orch`, full `go test ./...`, `go vet ./...`, and `go build ./...` all PASS.
27. The exact candidate was installed from `/src/pr-af/go`. Because `af install` started PR-AF without the callback override, that process was immediately stopped and PR-AF alone was restarted with `AGENT_CALLBACK_URL=http://workforce:8007`. CURRENT runtime is PID `156016`, health PASS, `PR_AF_PROVIDER=opencode`, `PR_AF_MODEL=openai/fcm`, `PR_AF_HARNESS_BIN=/src/pr-af/go/scripts/opencode-stream.sh`.
28. Exact `cloudsecurity-af` negative replay `exec_20260909_165047_udbz3tcc` on PID `156016` **SUCCEEDED** with `findings=[]`, but this is **NOT a precision PASS**. Total root duration was ~1128s despite `max_duration_seconds=900`; serialized quick meta alone consumed essentially the whole budget (semantic ~225.6s, mechanical ~360.3s, systemic ~356.7s after intake/anatomy), so substantive primary review did not run. PR-AF nevertheless synthesized `Looks Good / Safe to merge`. This is **BUDGET_FALSE_SAFE**: exhausted budget could become an approve-equivalent result.
29. The smallest fail-closed repair was implemented live in `/src/pr-af/go`: `runParallelReview` now returns a budget-exhausted error when review starts or a dimension is scheduled after the budget is exhausted, instead of silently returning no findings. Regression `TestPrimaryReviewBudgetExhaustionFailsClosed` was added. `go test ./internal/orch`, full `go test ./...`, `go vet ./...`, and `go build ./...` all PASS on the exact candidate.
30. The fail-closed budget repair is **not yet runtime-accepted or canonicalized**. Do not run XOR quality acceptance until that exact source is loaded and independently read back: a zero-finding result on the prior runtime can be caused by budget exhaustion rather than reviewer recall.
31. Fresh OBSERVE after the earlier timed-out install attempt produced one authoritative application-state read: `pgrep -af /afhome/packages/pr-af/bin/pr-af` exited `1` with no process. The VPS itself remains `running` and Hostinger shows no host restart during this interval; however Portwing/Docker target resolution is intermittently timing out (`/containers/json` deadline / `upstream request timed out`). Therefore the current application state is **PR-AF NOT OBSERVED RUNNING + INSTALL POST-STATE AMBIGUOUS**. Do not repeat install/reload blindly until package/source post-state can be read.
32. The next 80/20 product decision remains: quick mode must fit the declared review budget without buying speed by skipping substantive review. Raising `max_duration_seconds` is not a fix. First restore a stable control/readback window and load the already-tested fail-closed source; only then optimize quick pre-review latency from measured stage cost.

## Current blocker

The immediate blocker is now **CONTROL/READBACK INSTABILITY + PR-AF NOT OBSERVED RUNNING**; the product blocker behind it remains **QUICK_BUDGET_OVERSPEND + BUDGET_FALSE_SAFE acceptance**. Fail-closed behavior is deterministic PASS in live source but not runtime-proven. The safe sequence is: recover stable target readback → inspect package/source/install post-state → start/reload only PR-AF with the proven callback/provider/model/harness contract → verify health and loaded identity → run a bounded fail-closed budget canary. Only after that contract is proven should latency optimization and XOR recall acceptance resume.

Repository governance remains unchanged: do not rewrite/reset `main`; component work continues through `dev` and SourceLoop only after live proof. SWE-AF is concurrently edited elsewhere, so PR-AF must review immutable committed ranges and must not mutate or benchmark against its moving dirty working tree.

## Active BMAD batch — B4 fail-closed budget + quick-latency gate

Method: `bmad-help` → `bmad-quick-dev`, evidence-first. North Star Drift Check: **CONTINUE** — preserve the proven provider/callback/streaming path; do not add infrastructure, raise the timeout as a workaround, retune scoring broadly, or touch SWE-AF's moving working tree.

DoD:
1. Re-observe live `/src/pr-af`, current PR-AF PID/config/health, and any active OpenCode/AgentField work before mutation; if transport is unavailable, remain read-only and record the blocker.
2. Load the already-tested fail-closed budget candidate from `/src/pr-af/go`; PR-AF-only reload must preserve callback/provider/model/harness identity and health.
3. Prove budget exhaustion can no longer synthesize approve-equivalent output: exact CloudSecurity replay must either reach substantive primary review or terminate explicitly as budget exhausted.
4. Reduce `depth=quick` pre-review cost with the smallest evidence-backed change so the same clean real-repo benchmark reaches substantive primary review within its declared budget; do not remove semantic/mechanical/systemic coverage without an explicit quality tradeoff decision.
5. After a trustworthy clean negative (`0` unrelated findings with primary review actually executed), run XOR-positive and one repeat on the same verified runtime; require an evidence-grounded XOR/XNOR finding both times.
6. Canonicalize only the proven tracked delta via SourceLoop after runtime acceptance; exclude runtime/debug artifacts and rejected OpenCode config candidates.

## ONE next move

Recover CURRENT container/control-plane readback, then load the deterministic fail-closed budget patch before any further quality canary. The immediate acceptance question is binary: after declared budget is exhausted, PR-AF must fail explicitly instead of returning `Looks Good`; only then optimize quick-mode latency and resume precision/recall scoring.

## Write-back rule

Update this file in place after every bounded batch. Do not create a second PR-AF plan/status document. GitHub is canonical write-back/release only; application code continues to be edited and verified in the persistent DEV container first.
