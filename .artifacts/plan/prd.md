---
title: PR-AF MVP completion and truthful runtime acceptance
status: final
planning_status: ready-for-downstream-review-not-PLAN_ACCEPTED
created: 2026-10-20
updated: 2026-10-20
source_of_truth: PLAN.md
inspected_head: 093d11980442304fb925e6bfbcdf5fcbfd5e15a0
branch: mission/swe-full-pr-af-20261020
---

# PR-AF completion requirements

## 1. Purpose and decision

Deliver one useful, evidence-grounded review of a real pull request using the existing maintained **Go** PR-AF implementation in the permanent AgentField DEV runtime. Establish the chain from current source to installed artifact, process, registration, callback, execution, semantic result and SourceLoop durability. Preserve the existing pipeline and all existing project code. No new architecture is requested.

This document refines `PLAN.md`; it does not replace its gates or create another project tracker. Downstream architecture work is limited to mapping requirements onto existing components. Sprint planning must reuse existing issues when discovered, build dependencies and assign exact write ownership before implementation. The PRD is ready for downstream review, not a statement that implementation or runtime gates passed.

**FACT:** The immediate deliverable of this planning invocation is this PRD. Product repairs and their scoped commits are requirements for the subsequent execution stage, not changes claimed by this document.

**FACT:** The inspected checkout is branch `mission/swe-full-pr-af-20261020`, HEAD `093d11980442304fb925e6bfbcdf5fcbfd5e15a0`. Initial Git status contained untracked `.codex/`, `.graphify-code/` and `graft/`; these are unrelated local state and must not be staged, deleted or rewritten. The local `.artifacts/plan/issues/` directory contained no files at inspection. Remote issue inventory has not been verified.

**Constraint:** Current instructions prohibit push, publishing/opening a PR, comments, outbound messages, shared-system restart and deployment. They permit local source/test work and local Git commits. Runtime readback is read-only. A live patch, process reload, installation change, SourceLoop promotion or model/GitHub execution with external effects is not implicitly authorized by this PRD. Where a required proof conflicts with these boundaries, record a blocker and continue independent local work.

## 2. Current state grounded in source

| Area | Existing implementation and evidence | Consequence for scope |
|---|---|---|
| Canonical intent | `PLAN.md`, G0–G5 and 2026-10-20 G0 handoff | G0 is BLOCKED until authorized CURRENT readback; history cannot advance it. |
| Go package | `go/agentfield-package.yaml`: `language: go`, build `./cmd/pr-af`, start `bin/pr-af`, node `pr-af`, default port 8007 | Maintain this implementation; do not port or recreate it. |
| Legacy Python package | Root manifest starts `python -m pr_af.app`, default port 8004; Git installs redirect through `superseded_by`, local-path root installs do not | Node name and health alone cannot distinguish implementations. Root Compose is Python; Go add-on uses node id `pr-af-go` and is not permanent DEV proof. |
| Registration and callback | `go/internal/node/node.go:BuildAgent`, `register.go:RegisterAll`, `go/internal/reasoners/names.go` | All 17 handlers already exist; callback is `AGENT_CALLBACK_URL` mapped to SDK `PublicURL`. Empty value falls back to node-local localhost. Observe actual registered URL and network namespace. |
| Execution graph | `go/internal/orch/phases.go`, `calllocal_test.go`, `go/internal/node/calllocal_wiring_test.go`; Go README explains SDK `CallLocal` | Tracked phase dispatch already exists. Preserve it; discovery of 17 handlers does not require every handler to execute on every review. |
| Pipeline | `docs/ARCHITECTURE.md`, `go/internal/orch/`, `go/internal/reasoners/` | Intake, anatomy, dynamic planning, concurrent dimensions, evidence, adversary, obligations, compound synthesis, coverage, scoring and output already exist. Do not invent replacement phases. |
| Public request and result | `go/internal/schemas/input.go`, `output.go`, node `reviewInputSchema` | Preserve nullable budget overrides, raw diff/local repository/PR URL intake, snake_case output and defaults. `dry_run` defaults to false; all verification calls must explicitly set true. |
| Posting boundary | `go/internal/orch/output.go:generateOutput` posts only when `postToGitHub && !DryRun && PrURL != ""`; webhook has separate dispatch behavior | Test the real dry-run route with recording fakes. Never invoke a webhook for verification, or use README examples without adding dry_run. |
| Existing regressions | `go/internal/orch/resolve_test.go`, `degradation_test.go`, `output_test.go`, `budget_test.go`; config, node, reasoner and prompt tests | Workspace isolation/reused PR checkout, all-dimension parse failure, result parity and budget semantics already have tests. Reuse them and add only demonstrated missing assertions. |
| Model configuration | `go/internal/config/ai.go`; Go README versus package manifest | Code fallback is `minimax/minimax-m2.5`; manifest/container default is `deepseek/deepseek-v4-flash-0731`; provider default is `aforge`. This deliberate distinction is not a defect. Record effective harness AND AI model. |
| Resource limits | `config.ResolveBudgetCaps`, `ReviewConfig.FromInput`, `orch/budget_test.go` | Explicit input > environment > defaults ($2 / 3600 seconds). Reviewer concurrency, coverage iterations and depth are bounded. `TestCostStaysInert` proves current recorded cost is not provider spend accounting. Do not claim a hard dollar ceiling from configured caps or zero reported cost. |
| Local verification | `go/Makefile`, `.github/workflows/ci.yml` | Required Go checks are build, vet, tests, gofmt. CI also lints Python and builds the root Python image; these do not prove maintained Go DEV runtime. |
| Functional tests | `go/test/functional/main_test.go`, registration/health/reasoner API tests | Build-tagged suite starts and removes an isolated Compose project. Docker unavailable or readiness timeout can produce skips and zero exit. It is not a read-only probe of permanent DEV. |
| Fixture E2E | `go/test/e2e/run.sh`, `go/test/mockcli/` | Harness is mocked, but `.ai()` still calls OpenRouter; absent credential can skip with exit zero. Script starts a control plane and has lifecycle effects. It is neither fully offline nor semantic real-PR proof. Do not run under current restrictions. |
| Historical benchmark | `benchmark/martian-code-review-bench/{README.md,problems.json,results/,RESULTS.md}` | Existing real PRs, goldens and historical judgments can guide bounded selection. Historical Python/GLM-5.2 runs, uncapped benchmark runs and published recall do not prove current Go behavior. Do not rerun campaign/judge scripts with external calls in this scope. |
| SourceLoop | Canonical `PLAN.md` describes watcher `/src/pr-af`, proposal capture, validation and stale-base fast-forward checks | Historical canary `7901b5a...` proves a prior transport round trip only. Current watch/base/pending proposal/dev SHA and accepted-diff equality still require readback. |

Expected discovery set, with node prefixes normalized only when the API supplies them:

```text
review
intake_phase
anatomy_phase
planning_phase
meta_semantic
meta_mechanical
meta_systemic
review_dimension
compound_finder_phase
post_worthiness_gate
compound_dedup_phase
evidence_verifier
adversary_phase
deepen_findings
extract_obligations
verify_obligation
coverage_gate
```

## 3. Observation receipts and blockers

These are planning-session observations, not runtime acceptance:

- **FACT:** Enabled tool inventory exposed no AgentField DEV connector, graph/Graft navigation tools or native issue advisor. Thus no graph generation/coverage or advisor result can be asserted. Exact source reads were used; local index directories are not evidence of callable services.
- **FACT:** `Get-Command af -ErrorAction SilentlyContinue` returned no command; Go, uv and Docker executables were present.
- **FACT:** `docker context show` returned `desktop-linux`; sandboxed `docker ps --format '{{.ID}} {{.Names}} {{.Image}} {{.Status}}'` failed with permission denied on `npipe:////./pipe/dockerDesktopLinuxEngine`. This is a local access failure, not proof that DEV is stopped, Python, Go, healthy or unhealthy.
- **FACT:** A separately elevated read-only Docker inventory returned no runtime data before cancellation (exit 1). It supplies no additional CURRENT identity evidence; do not retry it blindly or claim a successful inventory.
- **FACT:** From `go/`, with `GOCACHE` set to a temporary directory and `GOPROXY=off`, `GOSUMDB=off`, `GOWORK=off`, `go test ./internal/config ./internal/node ./internal/orch -count=1` reported config PASS (2.834s), node/orch setup FAIL because dependency lock files under `D:\Tools\Go\cache\mod\cache\download` were inaccessible. Combined command exited 1. No node/orch test-pass claim is allowed.
- **FACT:** A separate config rerun passed (7.868s). Read-only `gofmt -l internal/config` listed eight existing config Go files; the formatting acceptance predicate is not green in this checkout. No formatting or source change was made; do not broadly reformat them for a planning deliverable.
- **FACT:** A native explorer launch first failed because the inherited model identifier was unavailable. Retrying with an explicitly available model admitted an agent but its provider returned 429/retry exhaustion. Root continued exact source investigation instead of treating either failure as a product defect. Independent QA must be reported separately from implementer review.
- **FACT:** Repository-local `_bmad/scripts/resolve_customization.py` and `memlog.py` are absent. Installed BMad Quick Dev planning and PRD headless instructions were used with defaults; do not invent workflow execution receipts or add another scheduler/log service.
- **FACT:** Independent planning QA is recorded in `review-rubric.md`: PASS WITH BOUNDARIES for the PRD, product acceptance BLOCKED. A deterministic document check found 15 unique AC IDs and verified all 15 distinct cited test names exist. This is document/source verification only.

`G0=BLOCKED`; G1 through G4 remain dependent on current evidence. SourceLoop closure remains unverified. Runtime source SHA, worktree, binary/process identity, reasoner discovery, callback, health and watch state are **UNVERIFIED_CURRENT**. Successful readback later must identify its timestamp and exact environment; local Windows checkout HEAD must never be substituted for `/src/pr-af` HEAD.

## 4. Users, value and minimum outcome

The primary user is the DEV operator who needs a review they can trust and an exact reproduction chain. The downstream engineering agents need bounded tasks with executable checks and no guesswork about ownership or permissions.

The MVP is one pinned real PR, one bounded Go review, a local returned review payload, evidence-grounded findings checked independently, and current SourceLoop readback. No GitHub posting is needed for usefulness. A faithful failed review or abstention is an acceptable **diagnostic result**, but it does not meet the positive useful-review North Star. Empty findings are not automatically successful; they must agree with the selected PR's independent evidence.

Measure separately: local checks actually executed; current identity completeness; executed/non-skipped canary count; semantic fixture obligations met; unsupported emitted claims; wall time; effective limits; provider spend observability; required runtime delta remaining uncaptured. Do not optimize finding count, agent count or benchmark recall at the expense of unsupported claims or external effects.

## 5. Functional requirements

### FR-1 — Current runtime identity and ownership

Before runtime execution or repair, collect a sanitized, timestamped read-only evidence bundle from the **existing permanent DEV**: installed `pr-af` entry from `/afhome/installed.yaml`; `/src/pr-af` HEAD/status and diff identity; installed manifest and package layout under `/afhome/packages/pr-af`; process argv/executable/listener; binary hash and build identity when available; control-plane identity/discovery; callback/PublicURL; node health from the appropriate namespace; SourceLoop watch/base/proposal state and durable `pr-af:dev` readback. Bind all observations to the same runtime/container/process generation. Record overrides rather than assuming default ports or names prove identity.

Classify implementation as Go, Python or unknown from executable/entrypoint plus installed source, not node id. A binary/source relationship without build metadata requires equivalent existing build/capture receipts; if the relationship cannot be proved, exact identity remains blocked. If current Python is confirmed, classify implementation selection as orchestration-owned. Do not add a PR-AF shim or change root manifests/Compose to compensate. Record the owning-system blocker; changing it is outside current authority.

### FR-2 — Callback and tracked execution

Resolve the actual callback for that process and discovery record. Cross-container callback cannot be localhost in the wrong namespace. Use read-only reachability from the control-plane namespace only through an authorized existing interface. A direct Windows health response does not prove callback reachability. Preserve `CallLocal` dispatch and capture the review's actual parent/child execution ids and terminal states. Do not require all 17 conditional reasoners to run; require the observed phases to match the selected request and pipeline path.

### FR-3 — Independently useful local work

G0 does not block offline test/document improvements explicitly authorized by the user. BMad Quick Dev must investigate the exact HEAD, reproduce a local defect or identify a missing assertion, apply the smallest change and obtain independent QA. No product-code change is justified solely by historical runtime symptoms. Existing tests must not be rewritten to call known implemented behavior missing.

The highest-confidence local gap is **acceptance visibility**, not a replacement pipeline: existing functional/E2E skips can return zero. Make downstream verification reject skipped or zero-executed mandatory checks using existing test JSON/output where possible. Prefer a small test-runner/test correction over a new framework. A second candidate is a narrow recording-fake regression proving dry-run emits zero GitHub writes through the real output path if the equivalent test is absent. A third is reconciling effective model/cost documentation without changing intentional defaults. Exact writer files are selected only after discovery and test deduplication.

### FR-4 — Failure localization and bounded repair

After G0, one bounded request must yield a request/execution id, expected versus actual, phase and bounded redacted logs. Assign the failure to implementation selection, supervision, callback/registration, provider/model/harness, repository intake, PR-AF reasoner/orchestrator, evidence/coverage/synthesis or external configuration. Use the native issue advisor when actually available, with exact command/tool identity and response receipt; if unavailable record `ADVISOR_UNAVAILABLE` and make an evidence-based root diagnosis. Do not install or fabricate an advisor or send external messages.

Only a reproducible PR-AF defect releases a product writer. Runtime access, provider throttling and dependency-cache failures release infrastructure blockers, not speculative product fixes. Preserve the failed attempt before replanning. Do not retry an ambiguous execution until its execution id/terminal state and effects have been read back. Reproduce twice when feasible; a deterministic local failing assertion is sufficient evidence without two expensive model runs.

### FR-5 — Side-effect-free review result

Every canary/review request explicitly sets `dry_run=true`, leaves `post_pr_number` absent and bypasses webhooks and posting utilities. Test with fake GitHub writers and fail on any write call. Offline tests use local repository/diff fixtures and injected harness/AI seams. A real semantic run uses only a read-authorized pinned PR snapshot/current DEV and already authorized provider route. If external fetch/model traffic is prohibited or the snapshot/provider is unavailable, mark semantic execution BLOCKED rather than invoking the network.

Preserve `ReviewResult` top-level keys `review_id`, `pr_url`, `review`, `findings`, `summary`, `metadata` and existing nested schemas. A finding carries existing file/line/severity/title/body/evidence/confidence fields and actionable root cause; review body is nonempty. Unsupported findings must be pruned or explicitly excluded from acceptance. All-dimension schema failures must remain failures and cannot turn into clean approvals. Partial parse degradation must remain visible via existing metadata.

### FR-6 — Bounded resource and failure reporting

For the MVP canary use explicit maximum duration 300 seconds, max cost 0.25 USD, reviewer concurrency 1, coverage iterations 1 and max review depth 1. These bound intended work; a timeout is a diagnostic result rather than a reason to silently increase limits. For a representative semantic review use explicit duration 3600 seconds, max cost 2 USD, reviewer concurrency 2, coverage iterations 1 and max depth 1. Record actual effective configuration and elapsed time; cancel/reconcile at the deadline through the existing interface. Do not change shared configuration.

**ASSUMPTION / fixed execution constraint:** These conservative MVP limits may yield abstention on a complex PR; choose a smaller existing benchmark PR rather than raising them silently. Hard spend enforcement is unproven: without an existing provider-side cap or verified accounting, do not authorize a paid run under this budget contract. Return `BLOCKED_SPEND_ENFORCEMENT` with the source evidence. New accounting infrastructure is outside this MVP; a proven minimal correction may be separately planned only in the existing executor.

### FR-7 — Semantic useful-review proof

Choose one existing, retrievable real PR from `benchmark/martian-code-review-bench/problems.json`, deterministically by ascending changed-line count among entries with a recoverable PR URL and at least one independently reproducible correctness/security issue; break ties by lexical URL. If changed-line data/snapshots are unavailable, discovery is blocked, not replaced with a synthetic PR or fabricated ranking. Pin owner/repo/PR URL, base SHA, head SHA and snapshot/diff hashes before execution; record why that issue is testable.

QA freezes an oracle before running the reviewed version: at least one expected root cause, file and line range in the pinned snapshot; a failing behavioral witness on the head; a minimal mitigation that makes that witness pass in a disposable copy; and at least one safe behavior/mitigation that must not be falsely reported. Do not feed goldens to PR-AF or mutate the canonical PR. Execute witnesses only in an authorized isolated test environment. Reuse historical goldens as candidates and verify them against the pinned source. If no safe automated witness is feasible, choose the next eligible PR; if none exists, report a selection blocker.

Semantic acceptance requires a current unmocked Go result that matches at least one frozen root cause and valid location, explains its evidence and impact, and emits no unsupported critical/important claim. Independent QA maps every emitted finding to source evidence and a witness or verified static proof; negative controls must have zero false critical/important reports. Serialize this judgment and witness command/output hashes. String matching a golden, process `succeeded`, a model judge alone, old benchmark scores and mocked findings cannot satisfy this gate. Failure/abstention must carry explicit reason and remains separate from positive semantic acceptance.

### FR-8 — SourceLoop durability and local commits

Read SourceLoop through its existing authorized interface; do not alter watcher/control-plane semantics. For an already accepted runtime delta, compare captured patch identity, validation receipt, current base, durable `dev` SHA and source tree, and prove no required container-only delta remains. A stale-base/conflicting proposal remains blocked. Under the current no-push/no-publish boundary do not trigger capture promotion or fast-forward; observe existing state only. If no runtime delta exists, G5 is `NOT_APPLICABLE_NO_RUNTIME_DELTA` with identity/readback evidence, never a fabricated capture PASS.

Local source/test improvements and planning artifacts receive focused local commits only after review and relevant checks. Stage exact paths, exclude local indexes/caches/secrets, preserve preexisting untracked state, record commit SHA and post-commit dirty state. A local commit is not SourceLoop closure, remote `dev` state or runtime proof.

## 6. Dependency plan and execution ownership

This is a refinement of `PLAN.md` tasks and gates, not a second authority. Maximum **three concurrent workers total**; reserve one slot for independent QA when candidates exist. Root owns integration and acceptance. Workers cannot approve their own changes. Read-only exploration may overlap; serialize writers touching the same file. No duration estimates or filler tasks are required.

| Task | Dependencies | Scope / ownership | Completion evidence |
|---|---|---|---|
| P0 inventory and advisor discovery | none | Root: HEAD/status, canonical plan/issues, exact tool availability; no remote issue mutations | Source citations, existing-capability dedupe, advisor available/unavailable receipt, next tasks bound to this HEAD. |
| R0 current DEV identity | P0 | Runtime observer: authorized read-only permanent DEV evidence, no source/config writes | FR-1/2 chain or precise G0 blocker. |
| L0 offline verification audit | P0; independent of R0 | Test owner: existing Go test/runner contract discovery, skip handling and no-post coverage; no product-code authority | Minimal reproducible assertion gap and exact proposed writer files, or no-change evidence. |
| L1 bounded local improvement | L0 | One writer owns approved exact test/script/doc files; root owns `PLAN.md` and PRD updates | Red/green test, focused/full relevant gates, candidate diff; no runtime acceptance claim. |
| QL independent local QA | L1 | Separate reviewer, read-only code; rerun affected checks, inspect final diff and scope | Accepted/rejected report with command receipts; repair returns to same owner. |
| R1 failure localization / canary | R0 with Go/callback/identity proven; limits and execution authorization available | Observer executes existing API only if side-effect boundaries permit | G1 classification and G3 actual canary evidence; otherwise explicit blocked prerequisite. |
| C0 conditional component repair | R1 proven component defect | Discovery first; bind one writer to exact failing implementation files/tests; no new architecture | Targeted regression then related checks and independent QA; local commit only. Runtime deployment/reload remains separately blocked. |
| S0 semantic oracle | P0 and available pinned snapshot; may overlap L0/R0 | Independent QA owns local oracle evidence; no PR-AF product writes or hints to reviewer | Frozen root cause/witness/negative control and pinned real PR identities. |
| R2 current semantic execution | G3, S0, current identity unchanged, permitted real provider access/spend bound | Existing DEV Go review with explicit dry_run | Current unmocked result, execution DAG, oracle checks; G4 accepted or diagnostic failure/blocked. |
| R3 SourceLoop readback | Accepted runtime delta OR explicit no-delta case; R0 | Read-only observer; no promotion | Exact diff/base/dev SHA/validation equality or blocked/no-delta status. |
| F0 reduce and handoff | QL and all attempted runtime tasks terminal/blocked | Root only: reconcile `PLAN.md` after evidence changes; scoped local commits | Separate local, functional, semantic, SourceLoop statuses, blockers and next bounded move. |

Initial useful frontier is R0 + L0 + S0 (subject to actual available read-only interfaces/snapshots). Missing R0 access does not stop L0/L1/QL. Once a local candidate exists, prioritize QL over more discovery. Runtime blockers do not release C0 by assumption. If no local defect/gap remains, commit evidence/docs only and report no product change needed.

Replan immediately after failed admission, provider error, failed check, stale identity, contradictory callback, changed PR SHA or QA rejection. Preserve attempt id, logs, diff and command; update only affected dependencies. For environment failures do not consume product-repair attempts. For a code failure perform one evidence-based bounded repair, then escalate root diagnosis if it fails again with unchanged cause; do not blindly repeat. An unavailable advisor is a disclosed tool limitation, not a reason to create a competing orchestration service.

## 7. Executable acceptance criteria

Each assertion is independently reported PASS, FAIL, BLOCKED or NOT_APPLICABLE with evidence paths/hashes and command exit codes. Missing/empty evidence, skips or unavailable prerequisites cannot be PASS. Current planning-session status is given separately; these are downstream tests, not claims of execution.

| ID | Automated assertion / evidence predicate | Gate / current state |
|---|---|---|
| AC-01 | Resolve HEAD and initial status; assert baseline equals inspected SHA or require fresh inventory/replan. Compare final diff against exact task allowlist; no deletion/overwrite of preexisting edits, no changes to indexes, dependencies, manifests or deployment unless specifically justified and authorized. | Local; inventory observed. |
| AC-02 | Parse CURRENT sanitized runtime observations; require same DEV/container/process generation, observed source HEAD/status, installed source path/manifest, actual executable identity/hash, listener and current discovery. Assert executable is maintained Go and build/source link is proved. Unknown/missing link or Python => G0 not PASS. | G0 BLOCKED. |
| AC-03 | Normalize discovery names and assert set equality with the 17 names in section 2 (not just count); `review` untagged and child domain tags checked when exposed. Assert callback record equals current effective PublicURL and read-only reachability succeeds from CP namespace. If tags unavailable, disclose this narrower proof; callback unavailable blocks gate. | G0 BLOCKED. |
| AC-04 | `go test ./internal/node -run 'TestRegisterAllExactSurface|TestBuildAgentFromEnv|TestBuildAgentWiresLocalCallerToApp|TestReviewHandlerForwardsLocalCaller|TestReviewHandlerErrorMapping' -count=1` exits 0 with all five tests actually executed; exact names already exist. | Local BLOCKED dependency cache. |
| AC-05 | Through real `generateOutput`/review path with recording GitHub fakes, set `dry_run=true` and a PR URL; assert zero `PostReview` and zero other GitHub write calls; positive fake control proves recorder detects a posting attempt. Explicitly assert request has no post_pr_number and is not a webhook. Add only missing test cases after dedupe. | Local candidate; not yet executed. |
| AC-06 | `go test ./internal/orch -run 'TestRunFailsBeforeOutputWhenAllDimensionsFailSchemaParsing|TestMixedAndCleanDimensionsExposeOnlyActualDegradation|TestReviewResultKeySetParity|TestCheckoutReusedWorkspaceUpdatesToNewPRHead|TestResolveRepoKeysWorkspacePerPR|TestWallClockBudgetTrips' -count=1` exits 0 with all six tests actually run. All-parse-failure must not call output; mixed degradation count equals actual failures. | Local BLOCKED dependency cache. |
| AC-07 | `go test ./internal/config -count=1` exits 0; cap precedence, defaults and depth tests run. For any live request assert effective limits equal FR-6, elapsed/deadline recorded, and spend enforcement receipt exists before paid execution. Recorded cost=0 alone fails spend-proof predicate. | Config PASS; live/spend unverified. |
| AC-08 | From `go/`: `go build ./...`, `go vet ./...`, `go test -json ./... -count=1` all exit 0, and `gofmt -l .` returns empty output; parse test JSON and reject skipped mandatory tests or zero executed required tests. For changed Python/scripts also run existing Ruff command `ruff check src/ scripts/` and affected pytest tests using declared dev dependencies. | Full suite NOT RUN; focused dependency setup failed. |
| AC-09 | A mandatory functional command is successful only if TestHealth, TestRegistrationParity and TestReviewErrorShape each have a non-skipped passing test event and actual stack evidence. Inject unavailable-Docker/readiness conditions in the verification logic and assert BLOCKED/FAIL, never accepted PASS despite underlying exit 0. Do not start Compose in current scope; a later isolated-stack receipt remains separate from permanent DEV G3. | Functional not run; lifecycle restricted. |
| AC-10 | For current permanent DEV bounded canary, assert explicit dry_run and limits, recorded request/execution id, terminal successful result with nonempty review body and schema keys, coherent observed child parent ids and phase results; zero GitHub mutations. Any fatal reasoner/parse error or skipped run fails positive canary. Synthetic/mocked result cannot pass semantic AC-11. | G3 BLOCKED_BY_G0. |
| AC-11 | Require real pinned PR/base/head/diff hashes equal frozen oracle and execution inputs; current unmocked Go result maps >=1 finding to expected root cause/location; head witness fails and mitigation witness passes; negative controls have zero false critical/important findings; every emitted critical/important finding has independent evidence; independent QA receipt accepts. Check raw results, proof command exit codes and hashes, not only an acceptance boolean. | G4 BLOCKED; no semantic run accepted. |
| AC-12 | On failed/abstained run, assert report contains exact execution id when created, terminal error/abstention, owner-layer classification, redacted correlated logs, expected/actual, preserved attempt and affected-only replan. Assert G4 positive acceptance=false. Missing runtime access reports BLOCKED with no invented execution id. | Blocker handling required. |
| AC-13 | Accepted runtime delta case: compare source/captured/durable-dev patch hashes and tree contents; require current base validation, exact dev SHA readback and empty required residual delta. Stale base => BLOCKED. No-delta case requires explicit no-delta proof/status. Local Git commit alone must fail capture predicate. | G5 unverified; read-only only. |
| AC-14 | Scheduler evidence shows <=3 active workers, disjoint writes and independent reviewer identity. If native advisor callable, retain advice request/result receipt; otherwise retain ADVISOR_UNAVAILABLE. Each failure has preserved receipt and updated affected dependency/status before retry. | Native discovery attempted; one provider 429 observed. |
| AC-15 | Review `git diff --check`, staged exact-path diff and scan for secrets/unwanted artifacts before local commit; commit only intended reviewed paths. After commit assert `git show --stat <sha>` scope, repeat applicable verification and record dirty state. Handoff separately states local checks, functional G3, semantic G4 and SourceLoop G5; assert no push/PR/comment/restart/deploy/publication commands or GitHub writes occurred. | Required per execution increment. |

The runtime predicates are evaluated by independent QA using the existing authorized readback interface. Its exact URI/tool command is deliberately not guessed: discovery must bind it to observed interface/version before task release. If that interface is absent, the assertion deterministically returns BLOCKED and all dependent execution remains unreleased. This is an external prerequisite with a decided fallback, not an unanswered product-design question.

## 8. Verification commands and prerequisites

Run repository-native commands from `go/` using PowerShell. Do not assume make, Bash or WSL. For offline discovery set a writable temporary `GOCACHE`, `GOWORK=off`, `GOPROXY=off`, `GOSUMDB=off`; do not alter machine/global Go settings. A dependency-cache lock denial requires an authorized writable populated `GOMODCACHE` or cache access repair. Do not download/install dependencies silently or weaken tests. Once prerequisites are restored, the exact next focused command is:

```powershell
go test ./internal/node ./internal/orch -count=1
```

Then use the AC-08 build/vet/test/format gate. Required Go version is at least 1.21 per `go.mod`; module dependencies are pinned in `go.mod`/`go.sum` including AgentField SDK v0.1.130. No dependency upgrade or new production dependency is required by this PRD.

The existing `go test -tags functional -json ./test/functional -count=1` requires a running Docker daemon, isolated Compose lifecycle authority, build/image prerequisites and unused functional ports 28080/28017. It is listed for a future permitted isolated lane only. Do not run it as permanent DEV readback. `go/test/e2e/run.sh` additionally requires Bash, af CLI and real `.ai()` provider calls; do not call it offline or infer PASS from missing-key skip. Prefer existing injected unit seams for local work; no new AI base-URL feature is demanded merely to make this script runnable.

## 9. Scope, risks and completion rules

**Must have:** maintained Go preservation; G0 identity/callback proof or precise blocker; dependency-aware tasks with at most three workers; advisor availability truth; independent QA; narrow independently testable local progress where supported; explicit no-post contract; real pinned semantic oracle and current execution evidence; SourceLoop readback; scoped local commits; exact failure/replan handoff.

**Nice to have, only after MVP gates:** a second representative real PR/negative case, additional benchmark sampling, clearer operator docs and optional extra fault injection within existing seams. None justifies a new service or production dependency.

**Out of scope:** new architecture/pipeline/reasoners; replacing Go with Python; rewriting existing project code; whole-benchmark campaigns; claimed recall/cost improvements without measurements; frontend work; global model/default normalization; shared Factory/CrewSync task mutations or messages; installing tools; changing production/credentials/access/billing; orchestration-owned fixes in this repository; SourceLoop redesign; remote Git mutation, deployment, publishing, comments or shared restarts.

**Risks and decided responses:**

- Runtime access remains absent: retain G0 BLOCKED and finish eligible local checks/docs; owner is authorized runtime operator/interface, next move is sanitized CURRENT minimum readback.
- Go module cache unavailable: retain BLOCKED_ENV for affected checks; request authorized populated cache/access restoration; preserve config PASS and do not claim full tests.
- Healthy wrong node or stale binary: classify ownership, reject G0, preserve identities; no compensating product patch.
- Callback is namespace-local: record actual URL and failing vantage; orchestration/config owner must resolve under separate authority.
- Provider 429 or absent advisor: preserve admitted task/attempt where available, do not manufacture agent/QA results; root may perform read-only fallback, but independent acceptance stays missing until a distinct reviewer succeeds.
- Semantic test cannot run under outbound/spend restrictions: keep G4 BLOCKED; local oracle and fixture tests remain supporting evidence only.
- Semantics look plausible but witness fails: reject acceptance, preserve exact result, select affected diagnosis; never relax the oracle to match output.
- SourceLoop readback/promotion unavailable or stale: keep durability blocked; do not push or promote to satisfy this document.

**Definition of Done for planning:** this PRD exists at the requested path, reflects exact inspected code, names existing capabilities, has automatable acceptance predicates and a bounded DAG, and explicitly preserves runtime blockers. Independent review and local commit receipts are reported outside product acceptance. No PLAN_ACCEPTED token or G0–G5 status is manufactured.

**Definition of Done for the product goal:** independent QA accepts AC-01–15 as applicable, positive G4 is proved with the real current Go runtime and pinned semantic witness, G0/G3 identities agree, and G5 accepted-delta closure or evidenced no-delta case is recorded in canonical `PLAN.md`. Otherwise final product status is PARTIAL/BLOCKED/EVIDENCE_MISSING with separate local/functional/semantic/durability results. A faithful diagnostic failure is useful progress, not completed positive semantic E2E.

**Handoff:** Goal = useful Go semantic PR review; state = planning/local evidence with G0 blocked; facts = inspected HEAD and contracts above; decisions = preserve architecture, no external effects, independent local lane; changed files = planning artifacts only in this stage; commands/evidence = section 3 and review receipts; blockers = runtime interface, dependencies, provider/QA availability; do-not-repeat = no historical identity substitution, no skip-as-pass, no mocked semantic claim, no blind 429 retry; next move = authorize/provide read-only CURRENT DEV evidence while executing accepted local assertion gaps.
