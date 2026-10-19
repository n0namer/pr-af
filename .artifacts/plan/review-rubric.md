# Independent PRD review rubric

**Reviewer scope:** read-only planning QA for `.artifacts/plan/prd.md` against the user mission, `PLAN.md`, Go test sources, and the installed PRD checklist. No runtime, Docker, network, or source changes were performed.

## Verdict

**PASS WITH BOUNDARIES for planning quality; product acceptance remains BLOCKED.** The PRD preserves the existing Go/AgentField architecture, identifies the current G0/runtime and dependency-cache blockers, separates mocked/local evidence from semantic runtime proof, and gives executable predicates rather than treating skips or historical benchmark output as success.

## Critical/high findings

### HIGH — functional skip semantics must remain a hard acceptance boundary

`go/test/functional/main_test.go` deliberately skips when Docker is unavailable or readiness fails, and can therefore return exit 0 with no live test executed. The PRD correctly calls this out in AC-09, but any implementation of the proposed acceptance visibility work must parse `go test -json` (or equivalent explicit test events) and require non-skipped PASS events for `TestHealth`, `TestRegistrationParity`, and `TestReviewErrorShape`. A plain command exit code is insufficient.

### HIGH — E2E is not offline and has lifecycle effects

`go/test/e2e/run.sh` starts a real control plane and node, creates a fixture repository, and leaves the control plane running by default. It uses a mocked opencode path, but four `.ai()` gates still call OpenRouter and the script requires `OPENROUTER_API_KEY`; without it, exit 0 means SKIP. The PRD is factually correct to prohibit using this as semantic acceptance under current restrictions. No-post is bounded by `repo_path` + `dry_run=true` and an API-traffic assertion, but this does not make the run side-effect-free at the process/filesystem level.

### HIGH — AC-10/AC-11 correctly require unmocked current evidence

Mock CLI findings, process success, nonempty review body, and historical benchmark scores cannot satisfy semantic acceptance. The PRD correctly requires a pinned real snapshot, frozen oracle, witness/mitigation, negative control, and independent QA. This must remain a separate gate from local unit/functional checks.

## Factual verification

The exact tests cited by the PRD exist:

- AC-04: `TestRegisterAllExactSurface`, `TestBuildAgentFromEnv`, `TestBuildAgentWiresLocalCallerToApp`, `TestReviewHandlerForwardsLocalCaller`, `TestReviewHandlerErrorMapping`.
- AC-06: `TestRunFailsBeforeOutputWhenAllDimensionsFailSchemaParsing`, `TestMixedAndCleanDimensionsExposeOnlyActualDegradation`, `TestReviewResultKeySetParity`, `TestCheckoutReusedWorkspaceUpdatesToNewPRHead`, `TestResolveRepoKeysWorkspacePerPR`, `TestWallClockBudgetTrips`.

The command shape `cd go; go test ./internal/node ...` and the package-local `-run` expressions are valid. The current dependency-lock permission failure prevents claiming execution; it is not a test failure. `go/test/functional` is correctly build-tagged and requires Docker Compose. `go/test/e2e/run.sh` is Bash-oriented and is not a Windows-native offline check.

## Scope and architecture checks

- **No new architecture:** PASS. The PRD explicitly preserves `CallLocal`, the current reasoner surface, existing output/posting seams, and SourceLoop semantics; it prohibits replacement pipelines, new services, dependencies, and runtime shims.
- **No external effects:** PASS within the stated boundary. The PRD requires `dry_run=true`, absent `post_pr_number`, no webhook invocation, fake GitHub writers for local tests, and no posting/deploy/restart/publish actions. It accurately distinguishes this from process, fixture, Docker, and provider effects of the existing E2E/functional harnesses.
- **Acceptance completeness:** PASS for planning. AC-01–15 cover source state, identity/callback, local checks, functional skip handling, semantic oracle, resource reporting, failure localization, and SourceLoop readback. Runtime-dependent predicates remain explicitly BLOCKED rather than silently downgraded.
- **Overconstraints:** The fixed FR-6 budget values and FR-7 deterministic PR selection are intentionally conservative execution constraints. They may produce abstention; the PRD already says to report that rather than raise limits or substitute a synthetic PR. This is acceptable for the mission, provided no paid run is started without spend enforcement evidence.

## Required boundaries for downstream execution

1. Do not label `go test` or functional suite PASS when dependency setup fails, Docker is unavailable, or mandatory tests are skipped.
2. Do not treat the E2E script's missing-key exit 0 as success.
3. Do not claim semantic correctness from mock invocations, historical benchmark JSON, or a successful AgentField execution status.
4. Preserve the exact current runtime identity/callback and SourceLoop blockers; local test/document work may proceed independently.

**Independent QA receipt:** no files other than this rubric were edited; no runtime or external systems were touched. Product status is not accepted by this review.
