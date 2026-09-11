# PR-AF — Canonical Plan and Current State

> Status: ACTIVE — B4 SWE-AF EVIDENCE BENCHMARK / VALID-R1 NEXT
> Updated: 2026-09-11
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

PR-AF performs **all of its intended PR-review functions as well as practical, with quality demonstrated by repeatable measurements rather than anecdote**. The intended functional contract is the product behavior documented by canonical PR-AF source/docs: ingest a PR/diff and repository context; understand change anatomy and blast radius; build a task-specific review plan; run focused semantic/mechanical/systemic review; ground findings in exact code evidence and change causality; challenge weak claims; synthesize compound risks; close meaningful coverage gaps; verify obligations; apply severity/blocking policy; and return actionable review output without false-safe approval when required review evidence is missing.

SWE-AF, CloudSecurity, any individual defect class, the current broker/model, latency optimization, runtime plumbing, and a particular benchmark are **means/evidence**, not the North Star. Engineering reliability requirements (fail-closed behavior, provider compatibility, provenance, runtime stability) are enabling constraints and measurable quality dimensions, not substitutes for product purpose.

## Current phase

Phase: **B4 measurable intended-function quality baseline / model differential**.

### Phase Goal

Turn the intended PR-AF functional contract into a compact evidence-backed scorecard and establish a trustworthy baseline on the exact current runtime/model using representative natural, seeded, clean-negative, and holdout cases. Preserve already-proven runtime/fail-closed safeguards, but optimize or remove complexity only when fresh evidence shows it is unnecessary for the intended-function quality gates.

## Operating contract

`OBSERVE → LOCALIZE → PATCH IN /src/pr-af → TARGETED VERIFY → FULL VERIFY → PR-AF-ONLY RELOAD → RUNTIME PROOF → CANONICALIZE → VERIFY SHA → WRITE-BACK`

Hard rules:
- no GitHub-first application coding/redeploy loop;
- preserve SWE / Deep Research and unrelated runtime state;
- secrets presence-only in evidence;
- primary reviewer failures remain fail-closed;
- additive coverage reviewer failures may stop further coverage expansion without discarding already-proven primary findings;
- work in bounded ~30-minute BMAD batches with explicit DoD and 80/20 priority.

## CURRENT evidence — 2026-09-10

| Claim | CURRENT evidence | Verdict |
|---|---|---|
| Durable application source | accepted live app tree was canonicalized to `dev`; squash commit `1967bb2275855d8f7626806169b2a274b379c9e0` was independently compared against accepted runtime commit `5a0f3b2b2c6c37d5cecab140cd2a0938c1715b7f`: `go/` diff empty, 9/9 intended blobs MATCH | PASS |
| Maintained package | PR-AF callback topology was repaired with a PR-AF-only stop/start using `AGENT_CALLBACK_URL=http://workforce:8007`; exact authenticated 1-second runtime execution `exec_20260910_173155_ttme514j` / `run_20260910_173155_42gj08yn` reached the real control-plane callback and failed explicitly with `Review time budget exceeded (max_duration_seconds=1) before anatomy`. No approve-equivalent output was synthesized. `BUDGET_FALSE_SAFE` is therefore runtime-proven closed. The subsequent `QUICK_META_FUSION` candidate was implemented container-first and reinstalled: live/installed `phases.go` both SHA-256 `6070ac483880b10a337b092f9baae45b96637759b36433c5822a38a8601dafc8`; live/installed `prompts/meta.go` both SHA-256 `f6825f7866da5ed46441d11a997c95abd09c0d50d4bb4def6f4e0f8bad087d76`. | FAIL-CLOSED RUNTIME PASS / FUSION IDENTITY MATCH |
| Runtime/load state | Current fused runtime is PID `213365` on port `8007`; local `/health` returns `{"status":"ok"}`. Non-secret process readback confirms `AGENT_CALLBACK_URL=http://workforce:8007`, `AGENTFIELD_SERVER=http://control-plane:8080`, `PR_AF_PROVIDER=opencode`, `PR_AF_MODEL=openai/fcm`, and `PR_AF_HARNESS_BIN=/src/pr-af/go/scripts/opencode-stream.sh`. Full `make check` on the exact fusion source passed build + vet + all Go tests before install. A clean-negative CloudSecurity quality run against immutable range `985234b...6c1133c` is CURRENT running as `exec_20260910_174415_zqjzg97y` / `run_20260910_174415_l3rh4nu3`. | FUSION RUNTIME HEALTHY / QUALITY ACCEPTANCE RUNNING |
| Generic provider contract | live PR-AF uses `OPENAI_BASE_URL=http://fcm-dev-internal:19280/v1`, `OPENAI_API_KEY` present, `PR_AF_MODEL=openai/fcm`, `PR_AF_AI_MODEL=openai/fcm` | PASS |
| Broker transport | direct `/chat/completions` probe with current broker key returned HTTP 200; tool-calling probe returned valid tool calls | PASS |
| OpenCode transport | external `openai/<model>` is internally adapted to dedicated `compat/<model>` via `@ai-sdk/openai-compatible`; runtime process trees proved `compat/fcm → fcm/fcm`; no unintended OpenRouter provider selection | PASS |
| Direct `.ai()` path | uses the same OpenAI-compatible key/base/model identity; partial key/base config rejected | PASS |
| Weak-model tolerance | meta planning uses lean structured schema → deterministic JSON recovery → one plain-text line-protocol fallback → fail-closed; operational ids/budgets/defaults are deterministic PR-AF policy | SOURCE + TEST PASS |
| False-safe prevention | unrecoverable meta output no longer becomes `dimensions=0 → Looks Good` | PASS |
| Meta runtime | repeated real FCM canaries completed `meta_semantic`, `meta_mechanical`, and `meta_systemic` 3/3 | PASS |
| Downstream DAG | clean real-file canary `run_20260903_114900_74zj0kf2` completed intake, anatomy, meta 3/3, primary review, coverage gate, coverage-added review, obligation verification, adversary, and root review | PASS |
| Coverage-only failure policy | `runCoverageLoop` preserves existing primary findings and stops further coverage expansion if a coverage-added reviewer fails; primary review path remains fatal. Targeted regression PASS; accepted clean canary also completed the coverage-added reviewer without regressing root behavior | PASS |
| Deterministic validation | Fresh 2026-09-10 `vps-terminal-dev` validation on current live source: mediation allows canonical `make check`, but the runtime PATH omits Go. A temporary stale-safe Makefile-only PATH adaptation invoked `/usr/local/go/bin/go` for build/vet/test; `make check` then exited `0` with build + vet + all Go tests PASS. The Makefile was immediately restored byte-for-byte to SHA `74da56...`; both temporary SourceLoop captures were marked REJECTED, so no canonical tooling delta remains. | PASS / CURRENT SOURCE |
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
- [x] Current B4 fail-closed candidate installed/running/registered on `8007` — 2026-09-10 `vps-terminal-dev` lifecycle route installed exact `/src/pr-af/go`; live/installed `phases.go` SHA-256 match, `/afhome/installed.yaml` records PID `200930` on `8007`, and fresh SDK logs show successful node registration.
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
Status: **ACTIVE — BUDGET FALSE-SAFE RUNTIME PASS / QUICK_META_FUSION LOADED / CLEAN-NEGATIVE QUALITY ACCEPTANCE RUNNING**.

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
33. Fresh infrastructure/control evidence localizes the current interruption above PR-AF: Hostinger reports VPS `1904412` running with continuous uptime and no recent host restart action; the shared AgentField DEV Coolify application `edshqtkwskg3lrczekhcmd71` remains `running:unknown`. At the latest read Coolify marked the server reachable/usable and had refreshed Docker version `29.7.2`, while bounded application log reads still returned HTTP 504. VPS Terminal `health`/`ready` can report gateway/Docker connected while target stats, container listing, exec, and target-file reads still time out on Docker `/containers/json`. AgentField gateway health can be `healthy` while node/capability reads return Bad Gateway. The latest available Hostinger sample during the degraded interval still shows CPU `100%` with RAM ~`5.2 GiB` on the 8 GiB host; this is correlation, not a proven root cause. Classify this as **CONTROL_PLANE_OBSERVABILITY_DEGRADED**, not PR-AF application failure. Do not restart/redeploy the shared AgentField DEV app or reboot the host merely to recover this PR-AF batch.
34. This batch remained mutation-free at application/runtime level after OBSERVE: no repeated `af install`, no PR-AF reload, no shared-app restart, and no host reboot were attempted because post-state could not be read reliably. The only durable change is this PLAN anti-drift write-back. `BUDGET_FALSE_SAFE` source repair remains the next candidate to load once scoped readback is stable.
35. The measured quick-stage cost now gives a concrete 80/20 latency hypothesis for the next code batch, **not yet an accepted design**: `intake (~31.9s) + anatomy (~153.4s) + three serialized meta calls (~225.6s + 360.3s + 356.7s)` consumes the whole 900s budget before primary review. The first candidate after fail-closed runtime proof is `QUICK_META_FUSION`: one bounded quick-only meta session that must cover semantic + mechanical + systemic risk and emit at most the same small review-plan surface, while standard/deep retain separate lenses. Rationale: remove two expensive harness/model round trips without dropping the three risk classes. Accept only if a clean real-repo run reaches substantive primary review inside budget and XOR recall is preserved; otherwise reject and replan.
36. Supplementary external-skill scan was performed for this batch. Useful guidance from `skills.sh` `terminal-ops` and `verify-behavior`, plus GitHub `debugging-methodology`, converges on the same operating rule already selected here: execute in the real repo/runtime, reproduce/localize before changing behavior, make one narrow change, and require raw executed proof before completion. `skills.com` returned no directly relevant indexed skill in the current search, so no guidance was imported from it. These external skills are advisory only; BMAD + this PLAN remain the workflow/SoT authority.
37. Applying that guidance exposed and closed a validation-transport issue without changing product behavior: `vps-terminal-dev` mediation explicitly recognizes canonical `make check`, but the target PATH omits `go`. A temporary stale-safe Makefile-only path adaptation allowed the exact canonical gate to run and PASS, then the Makefile was restored byte-for-byte and both temporary SourceLoop captures were rejected. This strengthens evidence for the live fail-closed source while preserving zero tooling delta.
38. Capability-gap diagnosis is now source-backed, not inferred from tool behavior alone. CURRENT `vps-terminal-dev` deployment reports mediator source commit `a49061a6687ee8d8cc6f890711cfe79e6ded0bd6`. Its `operation-mediation.mjs` normalizes a small set of exact verification/read/lifecycle commands and otherwise classifies arbitrary executables/paths as `OPAQUE_EXEC`; no AgentField package-install/start route is registered. The deployed target registry gives `agentfield-dev-workforce` terminal/process/stats/live_patch capabilities with `/src` as the live-patch root, but no reload capability and no `/afhome` writable root. A search of the current default branch also finds no `af install`, `INSTALL_PACKAGE`, or `PACKAGE_INSTALL` route. Therefore **DEV_PACKAGE_LIFECYCLE_CAPABILITY_GAP is confirmed**; bypassing mediation or widening `/afhome` access would be a new operator-security scope, not PR-AF work.
39. `bmad-testarch-atdd` is now used as the acceptance design layer for the next runtime gate without creating a second spec file. Red/green acceptance remains in this PLAN: (a) stale installed package must be replaced by the exact tested `/src/pr-af/go` source; (b) only PR-AF may be started/restarted; (c) loaded source/hash + callback/provider/model/harness + health must be independently read back; (d) a forced/exact budget-exhaustion case must terminate explicitly, never synthesize APPROVE; (e) only after that runtime proof may `QUICK_META_FUSION` be implemented and measured against clean-negative + XOR-positive + repeat. This prevents the capability blockage from silently relaxing the product gate.
40. Fresh skill research adds OpenAI Agents' `runtime-behavior-probe` and Vercel's `verification` guidance: runtime behavior must be demonstrated on the exact live destination/candidate, and end-to-end verification must cross all relevant boundaries rather than treating source/tests or configuration as runtime proof. This matches the current ATDD gate, so no workflow change is needed; the useful addition is to record a compact runtime case matrix (loaded identity → health → budget-fail-closed → clean-negative → XOR-positive → repeat) once PR-AF lifecycle becomes callable.
41. The first authenticated runtime gate localized a topology defect before touching review logic. Local `127.0.0.1:8007/health` returned `{"status":"ok"}`, but the exact secret-safe 1-second dry-run returned HTTP failure payload `agent_unreachable`: control plane attempted `http://localhost:8007/reasoners/review`. Fresh PID `200930` environment contains provider/model/harness/server but no `AGENT_CALLBACK_URL`. Historical accepted topology used `AGENT_CALLBACK_URL=http://workforce:8007`. Classify **PR_AF_CALLBACK_TOPOLOGY_DRIFT**; this run is transport evidence only and does not pass/fail the budget semantic gate.
42. A PR-AF-only callback recovery adapter is now container-first implemented in `vps-terminal-dev`: exact `af stop pr-af --json` and exact `/usr/bin/env AGENT_CALLBACK_URL=http://workforce:8007 /opt/af/bin/af run pr-af --port 8007 --json`; wrong callback/extra variants remain opaque and the DEV/PROD scope boundary is preserved. `node_check` PASS for mediator/test and targeted mediation suite PASS 28/28. SourceLoop checkpoints are CANDIDATE on existing terminal PR #144 at Git head `68d86e6f678f1b637ca0924dec17a638deef70cf`; terminal deployment is requested but CURRENT runtime still reports prior source `76976cea...`, so recovery commands must not run until loaded identity changes to `68d86e6f...`.
43. Callback recovery then passed live: PR-AF alone was stopped and restarted with `AGENT_CALLBACK_URL=http://workforce:8007`; PID/env/health/registration were independently reread. Exact authenticated runtime canary `exec_20260910_173155_ttme514j` / `run_20260910_173155_42gj08yn` traversed control-plane → workforce → PR-AF and terminated with explicit `Review time budget exceeded (max_duration_seconds=1) before anatomy`. This is runtime evidence that the fail-closed repair prevents budget exhaustion from becoming approve-equivalent output; **BUDGET_FALSE_SAFE is closed**.
44. With that ATDD gate green, `QUICK_META_FUSION` was implemented live as the smallest 80/20 latency repair: quick mode schedules only the semantic meta entrypoint, whose bounded quick prompt now explicitly covers semantic + mechanical + systemic risk and emits 1–3 fused review dimensions; standard/deep retain the existing separate three-lens behavior. Regressions prove quick calls exactly one fused planner while standard remains multi-lens. Full `make check` PASS; temporary Makefile Go-path adaptation was restored byte-for-byte and its SourceLoop captures rejected. The exact candidate was installed and PR-AF-only restarted as PID `213365`; live/installed source hashes match and health/provider/model/harness/callback readback PASS.
45. Clean-negative quality acceptance is now executing on the fused runtime against immutable CloudSecurity range `985234b5ed79323afd2a8a7b5d975c893ac4394f..6c1133cae087f98887c6dfbe7b8248e414b5c02f` as `exec_20260910_174415_zqjzg97y` / `run_20260910_174415_l3rh4nu3`. Intake completed in ~35.6s, anatomy in ~91.8s, and logs show exactly one `meta_semantic` start with no mechanical/systemic meta start. The synchronous HTTP client timed out at 90s while the asynchronous root continued, so final verdict remains pending root completion/log evidence.

## Current blocker

`PR_AF_CALLBACK_TOPOLOGY_DRIFT` and `BUDGET_FALSE_SAFE` runtime acceptance are closed. The product-quality gate is now **SWE-AF EVIDENCE BENCHMARK**: measure PR-AF against the latest actual SWE-AF brownfield delta rather than treating CloudSecurity as the primary quality target. CloudSecurity remains only a small historical calibration/negative fixture. CURRENT `/src/swe-af` is a large moving worktree (46 tracked changed files, +2843/-350 lines plus an untracked Python test) over base `58c4e0d19081bc52363c120b7963a34cebb1e894`; canonical Go `make check` on that exact worktree is PASS and `git diff --check` is PASS. Python validation is presently blocked by missing `pytest` in this runtime and is not classified as an SWE failure. The benchmark must establish its oracle before looking at PR-AF outputs, adjudicate findings against source/contracts/tests, and separate natural defects, seeded defects, clean negatives, and holdout cases. No scoring/prompt retune is justified until this evidence exists.

Repository governance remains unchanged: do not rewrite/reset `main`; component work continues through `dev` and SourceLoop only after live proof. SWE-AF is the read-only benchmark specimen for this batch: observe/review its current delta, but do not mutate, reset, clean, install dependencies into, or canonicalize SWE-AF while constructing the oracle.

## Active BMAD batch — B4 fail-closed budget + quick-latency gate

Method: `bmad-help` → `bmad-quick-dev`, evidence-first. North Star Drift Check: **CONTINUE** — preserve the proven provider/callback/streaming path; do not add infrastructure, raise the timeout as a workaround, retune scoring broadly, or touch SWE-AF's moving working tree.

DoD:
1. Re-observe live `/src/pr-af`, current PR-AF PID/config/health, and any active OpenCode/AgentField work before mutation; if transport is unavailable, remain read-only and record the blocker.
2. Load the already-tested fail-closed budget candidate from `/src/pr-af/go`; PR-AF-only reload must preserve callback/provider/model/harness identity and health.
3. Prove budget exhaustion can no longer synthesize approve-equivalent output: exact CloudSecurity replay must either reach substantive primary review or terminate explicitly as budget exhausted.
4. Reduce `depth=quick` pre-review cost with the smallest evidence-backed change so the same clean real-repo benchmark reaches substantive primary review within its declared budget; do not remove semantic/mechanical/systemic coverage without an explicit quality tradeoff decision.
5. After a trustworthy clean negative (`0` unrelated findings with primary review actually executed), run XOR-positive and one repeat on the same verified runtime; require an evidence-grounded XOR/XNOR finding both times.
6. Canonicalize only the proven tracked delta via SourceLoop after runtime acceptance; exclude runtime/debug artifacts and rejected OpenCode config candidates.

## SWE-AF evidence benchmark — active BMAD test-design / trace gate

Method: `bmad-testarch-test-design` for risk-first coverage and `bmad-testarch-trace` for requirement/risk → case → execution → evidence. Advisory skills already selected (`verification-before-completion`, `systematic-debugging`, runtime-behavior-probe) reinforce the same rules: establish the oracle before the system-under-test output, vary one hypothesis at a time, and accept only fresh executed evidence on the exact candidate.

Pareto risk map from CURRENT SWE delta (`git diff --numstat`):
- **R1 structured-output recovery / fail-closed behavior — P1:** `internal/harnessx/schema.go` +346/-2, `harnessx_test.go` +222, `run.go` refactor. Risk: malformed/no-progress provider output could be incorrectly promoted to a valid orchestration result or silently defaulted.
- **R2 autonomous coding/review loop — P1:** `internal/coding/loop.go` +62/-8 plus tests, coding role +246/-50. Risk: partial coder work, reviewer failure, blocking feedback, or retry semantics can become false success, lost work, or unsafe continuation.
- **R3 resolve/approval/HITL boundary — P1:** `orch/resolve.go` +192, HITL ~+376/-34, approval gate +44/-14. Risk: authorization or human-decision boundaries can be bypassed, stale, or interpreted fail-open.
- **R4 issue/build/git delivery — P1/P2:** build/gitops ~+192 plus tests. Risk: wrong worktree/commit/delivery identity or destructive/unintended git behavior.
- **R5 provider/runtime configuration — P2:** manifest/node/opencode config. Risk: current OpenAI-compatible provider identity or callback/runtime behavior drifts from intended contract.
- **R6 prompts/advisor/planning — P2:** reviewer/verifier/advisor/planning deltas. Risk: semantic quality regression without compile failure.

First gold-set target is deliberately small: 4–6 natural SWE cases concentrated on R1–R4, 6–8 seeded defects spanning fail-open/error propagation/approval/worktree/provider boundaries, and 3–4 clean negatives; reserve at least 3 cases as holdout. Each case records expected behavior and evidence before PR-AF output is inspected. Scorecard: critical/major recall, precision, change-causality, severity calibration, evidence completeness/actionability, stability on repeat, wall time/model-call cost, and instrument validity (substantive reviewer actually executed). Python-suite absence remains a validation blocker only; do not install dependencies merely to construct this benchmark.

### Oracle-first R1/R2 slice — frozen before PR-AF output

Independent source/test inspection establishes these expectations before the reviewer sees the cases:
1. **R1 watchdog recovery clean behavior:** a no-progress watchdog may be recovered only when assistant text or the captured output file contains an exact-schema-valid result; generic transport errors remain fail-closed. Existing tests directly cover both paths.
2. **R1 candidate natural concern — weakened by caller evidence:** when `Harness` returns no Go error but `Result.Parsed == nil`, `executeStructured` returns a default-seeded `T` with `err=nil` for non-fatal schema/no-output/provider-result failures and preserves the failing `Result`. Fresh caller inspection shows the important coding/planning/fast paths explicitly test `result == nil || result.Parsed == nil`, and CI documents deterministic fallback as an intentional parity contract. Therefore this is no longer presumed fail-open. It is now a precision trap for PR-AF: a finding is valid only if it demonstrates a concrete caller that consumes the seeded value without checking the preserved failure state and thereby changes behavior incorrectly.
3. **R1 candidate natural defect — concurrent output-capture collision:** `startStructuredOutputCapture` watches the shared `ProjectDir/.agentfield-out-*` namespace and permanently disables salvage when it sees more than one new output directory. `orch/plan.go` intentionally runs issue writers concurrently, while each `RunIssueWriter` passes the same repository path as `Cwd`/`ProjectDir`. Therefore two simultaneous issue-writer harness calls can make each watcher observe multiple new output directories and discard otherwise exact-schema-valid watchdog salvage. No concurrency regression for this capture path was found. This is source-evidenced and falsifiable; severity remains unassigned until targeted proof.
4. **R2 retry clean behavior:** coder failure after an actual git-worktree change retries in-place with explicit validation/repair feedback; unchanged-worktree provider failure remains unrecoverable. Existing git-backed test covers the changed-worktree path.
5. **R2 candidate natural defect — dirty-worktree fingerprint blind spot:** `gitWorktreeFingerprint` records only `HEAD` plus `git status --porcelain`. If a file is already modified before a coder iteration and the failed coder changes the contents of that same already-dirty file without changing its porcelain status line, the before/after fingerprints are identical. `gitWorktreeChanged` then returns false and the loop can classify useful partial work as unchanged/unrecoverable instead of retrying in place. No regression covering a pre-dirty same-path content change was found. This is a frozen source-evidenced hypothesis for the R2 review, not yet a confirmed defect.
6. **R2 reviewer invariant:** reviewer execution errors now propagate instead of synthesizing `approved=true`; a blocking but successfully returned review is fed back as `fix` and may recover on the next bounded iteration. Existing tests cover blocking-review recovery; reviewer-error propagation is source-evident and should be checked by PR-AF rather than assumed from compile success.

This slice intentionally contains both likely-clean behavior and one falsifiable natural defect hypothesis so precision and recall can be judged together. Oracle labels remain `expected-clean`, `candidate-defect`, or `needs-runtime-proof` until adjudication; do not count a candidate hypothesis as a true positive merely because PR-AF repeats it.

Execution note: first R1 execution `exec_20260910_231158_qo0763rm` / `run_20260910_231158_14yqz6s6` completed in `441606 ms` and did reach substantive `review_dimension`, but the newly loaded exact `PR_AF_EXECUTION_READ` exposed a **fixture validity failure**: the stored anatomy parsed the compact hand-written diff as `0` additions, `0` deletions and no hunks, then returned `findings=[] / APPROVE`. That result is **INVALID AS QUALITY EVIDENCE**; it proves only runtime/instrument execution and must not count as precision/recall PASS. The exact execution-result route itself is now live on `vps-terminal-dev` source `1571f8525af160d890805271e23cbf574ca8d101` with source/configured/coordinator identity aligned and `source_conflict=false`; its targeted container-first suite was 30/30 PASS. CURRENT startup selftest on that terminal build reports `selftest_mediation_policy_failed`, so operator qualification is PARTIAL even though the exact read route is functionally proven. A second same-malformed-fixture R1 run `exec_20260911_000834_3canmv8r` / `run_20260911_000834_ncr30vho` was already accepted before this fixture diagnosis and is still running; treat it only as stability/instrument evidence and do not start another benchmark until it terminates.

External-skill refresh reinforces the benchmark design: current GitHub `qa-methodology` guidance recommends risk-based testing, independent verification, mutation-guided test hardening and diff-aware mutation review; current `code-review` guidance emphasizes concrete file/symbol evidence and actionable BLOCKER/MAJOR findings. These are advisory and fit the existing BMAD test-design/trace lane; they do not authorize broader review scope or mutation.

### R1 valid-fixture retry — CURRENT

The previously duplicated malformed-fixture repeat `exec_20260911_000834_3canmv8r` / `run_20260911_000834_ncr30vho` is now terminal: root completed in `1095058 ms` after intake `31.2s`, anatomy `83.2s`, fused meta `349.5s`, two substantive reviewers `306.2s` / `352.9s`, and evidence verifier `272.5s`. Because its fixture was already classified malformed, this remains instrument/stability evidence only; its `0/2 blocking` merge-gate log must not enter quality scoring.

A replacement R1 probe was then accepted as `exec_20260911_064550_em3nx00s` / `run_20260911_064550_cecq1hub` using `repo_path=/src/swe-af` and a syntactically valid numeric unified hunk over the real `go/internal/harnessx/schema.go` path with five explicit additions. This run is CURRENT in progress from intake. Per the frozen gate, no finding will be scored unless returned execution/anatomy evidence first confirms the intended hunk was parsed as non-zero additions. No PR-AF retune occurred between malformed and valid-fixture runs.

The exact execution-read route is currently blocked by `required_context_provider_degraded` even though operator mediation classifies it ALLOW; the AgentField connector independently still returns `Bad Gateway`. Therefore result retrieval is an observation-plane blocker, not application evidence. Do not retry the same read mutation blindly; use runtime logs until Required Context recovers, then perform one exact execution read.

### Smarter-model full-path differential — CURRENT

User reports the broker now routes to a materially stronger model. Treat this as an external experimental-variable change, not as proven PR-AF source change: the PR-AF package remains the same installed runtime (`PID 213365`, started 2026-09-10) and its public model contract is still `openai/fcm`, so any quality/latency improvement must be demonstrated by fresh execution rather than inferred from configuration. The valid R1 run `exec_20260911_064550_em3nx00s` has now completed end-to-end in `1088062 ms`; it reached fused meta and two substantive reviewers, but exceeded the nominal 900s budget, so its payload still requires adjudication before it can be quality evidence.

To isolate the model change with minimal confounding, a fresh full end-to-end PR-AF review has been started on the same immutable SWE range previously used as diagnostic smoke: `f9aec2111d084ab0204d5612f2ea00a562226ac7..0c64fe7cc4fc216f4d32d0b855015509750eb4aa`, `repo_path=/src/swe-af`, `depth=quick`, correctness focus, dry-run, 900s. New execution is `exec_20260911_105118_g3dm8y4b` / `run_20260911_105118_ydjdjp96`; intake has started. Historical same-range execution `exec_20260909_160214_zy5qk18m` succeeded with `findings=[]`, giving a useful before/after reference while holding repository range and PR-AF review mode fixed. This is the current 80/20 experiment for the reported model upgrade.

### Intended-function quality model — BMAD test-design / trace

Use Goal–Question–Metric discipline: every metric must answer whether an intended PR-AF function works, not merely whether a workflow completed. Trace each intended function → risk → benchmark case → oracle → execution → finding adjudication → gate. Core scorecard: critical/major defect recall; precision and false positives per clean case; change-causality accuracy; evidence correctness/completeness; severity/blocking calibration; coverage-gap closure; obligation/adversarial verification value; fail-closed correctness under missing/invalid evidence; repeat stability; and cost/latency/model-call efficiency. Report per-function results as well as aggregates so one strong capability cannot hide another broken one.

Benchmark construction remains mixed by design: natural real-repository cases for ecological validity, seeded/mutation cases for known ground truth, clean negatives for specificity, and an untouched holdout to limit benchmark overfitting. The current SWE worktree is one high-value source of cases, not the benchmark definition. Academic evidence supports this direction: contemporary code-review benchmark research reports fragmented evaluation and argues for broader task coverage, dynamic/runtime evaluation, and fine-grained capability assessment; mutation testing provides controlled fault-detection ground truth; GQM ties metrics to explicit goals/questions rather than convenient counters. External agent skills add operational detail only: verification-before-completion, falsification/systematic-debugging, diff-aware review, and mutation-guided test hardening.

### Stronger-model differential — RESULT

Fresh exact execution read for `exec_20260911_105118_g3dm8y4b` / `run_20260911_105118_ydjdjp96` is now available. The same immutable SWE range that historically returned `findings=[]` produced **6 findings** on the current broker/model: one `important`, three `suggestion`, two `nitpick`. The important finding identifies a concrete cross-language model-default drift hole: the PR adds Python Dockerfile/default coupling but canonical Go checks do not couple `openRouterAutoDefaultModel` to `go/Dockerfile`, allowing a stale baked `HARNESS_MODEL` to survive both suites and override runtime role defaults. Other findings identify missing codex deployer-intent contract tests, untested planning-path HARNESS_MODEL gating, silent legacy HARNESS_MODEL behavior change, and two maintainability/test-robustness issues. Anatomy confirms the real 11-file `194+/22-` commit diff was parsed, so this is valid semantic evidence rather than the earlier malformed-fixture artifact.

The stronger model therefore materially changes observed review usefulness on this case, but **does not solve latency**: total root duration was `1163001 ms` (~19.4 min) despite a 900s requested budget; intake `63.3s`, anatomy `324.3s`, fused meta `224.5s`, two substantive reviewers `366.3s` / `534.8s`. Treat model quality and execution-budget correctness as separate dimensions. Do not remove fail-closed/budget/runtime safeguards because semantic quality improved.

### Stronger-model finding adjudication — batch 1

BMAD trace/adjudication rule: a reviewer claim is TP only when exact changed-source + consumer/test evidence supports both behavior and PR causality; plausible missing tests without demonstrated bad behavior are tracked as coverage observations, not product defects.

Current adjudication from immutable SWE head `0c64fe7cc4fc216f4d32d0b855015509750eb4aa`:
- **f_important / Go Dockerfile default-chain drift — TP, important justified.** `go/Dockerfile` bakes `ENV HARNESS_MODEL=openrouter/deepseek/deepseek-v4-flash-0731` and explicitly says it MUST match Go `openRouterAutoDefaultModel`. Go tests assert the same literal in runtime resolution but do not couple the Dockerfile value to the Go constant. A one-sided Go default update can therefore leave the image stale while source tests still pass. This is change-causal to the model-default contract touched by the reviewed commit and actionable.
- **planning-path HARNESS_MODEL gating — coverage observation, not yet TP.** Go source/tests already prove HARNESS_MODEL is scoped to `open_code`, including `ResolveRuntimeModels` and `FastResolveModels`; Python tests cover `_default_planning_model` cascade. The review may be right that one exact planning entrypoint lacks an end-to-end contract test, but current evidence does not establish incorrect runtime behavior.
- **codex deployer-intent contract — coverage observation, not yet TP.** Both Python and Go tests cover codex auth-mode/default selection and HARNESS_MODEL non-interference. Missing a specific `SWE_DEFAULT_MODEL` + codex combination is useful test-hardening advice, not demonstrated product failure.
- **legacy HARNESS_MODEL behavior-change claim — needs historical-contract proof.** Current source intentionally documents HARNESS_MODEL as open_code-only; without an authoritative prior compatibility promise, classify the claim as unproven rather than a defect.
- **Dockerfile parser robustness nitpick — valid maintainability observation, non-defect.** Exact-string parsing can make a guard brittle under formatting changes, but it fails loudly rather than permitting false-safe behavior.
- **duplicated Go expected-model literal nitpick — valid maintainability observation, non-defect.** Duplication increases edit cost but deliberately retaining an independent literal can also strengthen drift detection; do not count toward defect recall/precision.

Interim semantic score for this six-finding case: **1 confirmed product defect / 6 reported findings**; 2 useful coverage observations, 2 maintainability observations, 1 unresolved compatibility claim. Precision is therefore reported at two levels rather than gamed: strict defect precision currently `1/6`; useful-review precision currently `5/6` if coverage/maintainability observations are accepted as useful and the unresolved compatibility claim is excluded pending proof. This single case is not enough for an aggregate PR-AF precision estimate.

Minimal intended-function trace matrix for B4: intake/anatomy → fixture-validity evidence; planning/coverage → risk-to-dimension coverage; primary review → critical/major recall + change causality; evidence verification/adversary → unsupported-claim rejection; severity/merge gate → severity/blocking calibration; failure paths → fail-closed correctness; repeat execution → stability; whole DAG → wall time/model calls/budget compliance. A case cannot score semantic quality when its fixture is malformed or substantive review did not execute.

### Paired gate P1 — oracle frozen before execution

BMAD ATDD + test-design/trace pre-registration, with external `verification-before-completion` / falsification guidance: both cases use the same current PR-AF/model, `repo_path=/src/swe-af`, `depth=quick`, correctness focus, dry-run, 900s; no PR-AF prompt/scoring/code change is permitted between cases.

**P1-positive — seeded major provider-precedence regression.** Real file: `go/internal/node/node.go`, current `resolveAIConfig`. Seeded diff changes only the guard `if strings.TrimSpace(os.Getenv("AI_BASE_URL")) == ""` to `!= ""`. Oracle: this reverses the intended precedence condition. When `AI_BASE_URL` is absent and `OPENAI_BASE_URL` is configured, the OpenAI-compatible endpoint is no longer copied into SDK config; when both are present, `OPENAI_BASE_URL` can overwrite the explicit SDK-native `AI_BASE_URL`. Expected reviewer behavior: report a change-causal provider/base-URL precedence defect with concrete evidence; severity at least important/major. APPROVE/no causally equivalent finding is a recall failure. The case is synthetic and must never be counted as a natural SWE defect.

**P1-negative — comment-only clarification.** Same real file/function; diff changes only explanatory comments around the existing OPENAI_BASE_URL fallback/timeout behavior and leaves executable Go unchanged. Oracle: no correctness/security/reliability defect is introduced. Expected reviewer behavior: no blocking/important defect finding; style/nitpick may be useful but does not count as defect TP. Any claimed runtime behavior regression must be evidence-rejected as FP.

Instrument gate for both: anatomy must report a real non-zero hunk/addition/deletion as appropriate and substantive primary review must execute. Malformed/zero-hunk or pre-review budget exhaustion invalidates semantic scoring rather than becoming PASS/FAIL.

### P1 execution — CURRENT

Oracle was durably frozen before either request. Both exact `PR_AF_QUALITY_PROBE` calls were then accepted on the unchanged runtime/model: P1-positive = `exec_20260911_113554_k12hbon5` / `run_20260911_113554_zidaoioo`; P1-negative = `exec_20260911_113602_l8ulrowi` / `run_20260911_113602_ef61mls9`. Runtime log readback proves both reached root `review` and started `intake_phase`. They were launched ~8s apart for wall-clock efficiency; therefore semantic recall/precision remain comparable to the frozen oracle, but their wall-time/model-latency measurements are **contention-confounded** and must not be used as clean latency benchmarks. AgentField connector observation is currently `Bad Gateway`; managed terminal sessions and runtime logs remain the authoritative CURRENT observation path.

### P1 observation checkpoint

The two client sessions both ended after the control-plane synchronous 90s HTTP timeout (`execution timeout after 1m30s`), but post-state inspection proves this was **client/control-plane timeout, not execution termination**. Runtime logs show P1-positive intake completed in `58776 ms`, anatomy completed in `259391 ms`, and fused `meta_semantic` started; P1-negative intake completed in `73807 ms` and anatomy started. No terminal root event exists yet in the latest log read. Current workforce resource readback is near its 2 GiB memory limit (`2140483584 / 2147483648` bytes) with `417` PIDs; because both benchmark cases overlap, do not attribute their current stage latency to the model or PR-AF alone. Do not retry or launch duplicates while these executions are still live.

## ONE next move

Observe only the two existing P1 runs until terminal root events appear; then retrieve exact results, validate non-zero anatomy/substantive-review gates and adjudicate against the frozen oracle. No new benchmark execution, PR-AF retune, or latency conclusion while this contention-confounded pair is active.

## Write-back rule

Update this file in place after every bounded batch. Do not create a second PR-AF plan/status document. GitHub is canonical write-back/release only; application code continues to be edited and verified in the persistent DEV container first.
