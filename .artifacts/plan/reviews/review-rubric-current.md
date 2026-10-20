# Independent architecture review rubric

**Reviewer scope:** Independent document review of `.artifacts/plan/architecture.md` against `.artifacts/plan/prd.md`, canonical `PLAN.md`, and targeted maintained Go source evidence. Reviewed architecture decisions, AC-01–15 coverage, interfaces, boundaries, and remaining assumptions. No product source changes, commits, runtime access, or product tests were performed.

## Verdict

**PASS WITH BOUNDARIES for architecture consistency; product acceptance remains BLOCKED.** The design preserves the maintained Go/AgentField architecture, maps all 15 PRD criteria to concrete owners and evidence, binds interfaces to existing Go declarations, and does not claim that blocked runtime, functional, semantic, spend, or SourceLoop gates passed. It also correctly leaves implementation placement dependent on P0/L0 and canonical PLAN admission.

## Findings

### MEDIUM — make the baseline/HEAD relationship explicit

The architecture metadata and opening FACT identify `dc8d42e…` as `baseline_head`, while the reviewed checkout is now at `c261109…` (`docs: define verified Go PR-AF architecture contracts`), whose parent is `dc8d42e…`. This is consistent with an architecture document authored against the PRD commit, but a reader could mistake the baseline for the current checkout. Clarify that `dc8d42e…` is the pre-architecture inventory baseline and that the architecture commit is the current HEAD. Downstream P0 must still re-inventory state as already required.

### No blocking interface or requirement gaps found

- The existing `orch.App`, `LocalCaller`, `Deps`, `New`, `Run`, `github.Client`, `BuildAgent`, and `RegisterAll` contracts are consistent with the Go source. Graft `ask --source` returned the maintained node, orchestration, and functional registration sources; wiring graph check reported in sync, while semantic indexing remained pending. This was targeted source verification, not an exhaustive graph audit.
- AC-01–15 appear in the architecture acceptance map with appropriate evidence owners. In particular, AC-02/03 do not treat names or health as identity/callback proof; AC-05 has a recorder positive control; AC-09 rejects skip-shaped success; AC-10 requires a complete effect inventory; AC-11 requires frozen witnesses and independent adjudication; AC-13 distinguishes local commits from SourceLoop durability.
- The `go-test-events/v1`, `pr-af-evidence/v1`, and `pr-af-runtime/v1` conventions are explicitly evidence formats, not new product APIs or service interfaces. Their owners, fields, and blocked behavior are specified sufficiently to keep independent evidence reducers aligned while deferring implementation location to L0.
- Prior high-risk ambiguity around exposed tags, mandatory test-event aggregation, runtime effect evidence, artifact hashing, and per-finding adjudication is addressed in AD-3, AD-7, AD-8, and AD-9. Existing source-reality and seam receipts provide corroborating review evidence; this receipt independently checked their resolution against the architecture text.
- PLAN remains the status authority. The architecture does not imply PLAN_ACCEPTED, authorize runtime execution, alter the three-worker cap, or introduce a replacement pipeline, dependency, gateway, or runtime adapter.

## Required boundary

The architecture review cannot satisfy product AC-02–15. Current DEV identity/callback, mandatory test execution, functional canary, semantic real-PR acceptance, spend enforcement, and SourceLoop equality remain unverified or blocked as documented. The source graph's semantic layer is pending; exact targeted reads are the evidence used here.

**Independent QA receipt:** This reviewer wrote only `.artifacts/plan/reviews/review-rubric-current.md`. No product test/build/runtime result or external acceptance is asserted.

## Resolution review

**Final verdict: ACCEPT / RATIFIED.** The architecture now labels `dc8d42e…` as the original design baseline and records `c261109…` as the revalidated HEAD, including the relationship between them. AD-7 now fixes exact package identities for the node, orchestration, and functional event tuples, separates functional stack evidence from test-event evidence, and defines the shared reducer input/output and negative fixtures. AD-8 now specifies exact negative-control row columns and cardinality keyed by control and result hash, with missing/duplicate rows rejecting aggregate acceptance. These changes close the finding above and remove the remaining reviewed evidence-shape ambiguity. Product acceptance remains BLOCKED pending its runtime and execution gates.
