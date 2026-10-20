# Independent architecture compatibility review

**Scope:** Adversarial review of `.artifacts/plan/architecture.md` for interface compatibility, source accuracy, and evidence/test reducer ambiguity. This receipt records the follow-up review of amended AD-7 and AD-8. Exact Go source was treated as authority. This review did not change product code, runtime state, Git state, or the architecture document.

## Verdict

**ACCEPT.** Both prior contract concerns are resolved in the amended architecture. The Go signatures and public request/result contracts examined match source. The design preserves current ownership boundaries and explicitly avoids specifying replacement modules.

## Findings

### RESOLVED — Negative-control adjudication schema

**Location:** AD-8, semantic witness/adjudication contract.

The follow-up confirms AD-8 now fixes the exact columns, one-row-per-control/result cardinality, `(control_id, result_artifact_sha256)` join key, count scope, nonnegative counts, and PASS/FAIL/BLOCKED rules. This closes the ambiguity; findings from another control's result cannot be mixed into the count.

**Source/requirement basis:** PRD AC-11 requires independent negative controls and zero false critical/important reports. AD-8 already establishes the result artifact hash as the join key for finding rows.

**Disposition:** Resolved in AD-8. The documented exact schema is `control_id`, `result_artifact_sha256`, `expected_safe_behavior`, `observed_critical_count`, `observed_important_count`, `reviewer`, `verdict`, `reason`.

### RESOLVED — Test-evidence reducer identity set

**Location:** AD-7 and the AC-04/06/08/09 handoff.

Correction to the initial review: AD-7 already contained the `go-test-events/v1` header and row schema, process/artifact fields, and reducer PASS predicate; the earlier receipt overlooked that paragraph. The remaining concern was whether each AC froze its expected package/test identities. The follow-up confirms AD-7 now lists exact node, orchestrator, and functional package identities and requires expansion to explicit `(Package, Test)` tuples before reduction. It also separates functional transcript status from environment status and requires a distinct stack receipt for functional PASS.

**Source/requirement basis:** PRD AC-04/06 require exact test sets to execute without skips; AC-09 additionally requires actual functional pass events. The PRD’s cited Go tests are split across node and orchestrator packages, while functional checks have separate package/lifecycle prerequisites.

**Disposition:** Resolved in AD-7. File placement remains correctly deferred to L0; schema, test identity sets, reducer semantics, and event/environment distinction are now fixed.

## Source compatibility checks

- AD-1 `orch.App`, `LocalCaller`, `Deps`, `New`, `Run`, `github.Client`, `BuildAgent`, `RegisterAll`, `ResolveBudgetCaps`, and `ReviewConfig.FromInput` signatures match the inspected Go declarations.
- AD-2 `ReviewInput` field types/tags and `ReviewResult` top-level JSON shape match `go/internal/schemas/input.go` and `output.go`; default-seeding and pointer semantics are consistent with the source comments.
- AD-3’s 17-name inventory matches `RegisterAll` and the reasoner constants’ intended registration surface. The distinction between registered handlers and conditional child executions is correct.
- AD-4’s output guard matches `go/internal/orch/output.go`; the proposed positive control must call the real output path with valid PR state, as the architecture specifies.
- AD-6’s effective model/config and budget precedence remain delegated to existing config code, with spend-enforcement evidence separated from configured caps.

## Limits

The checkout’s Git commands were refused by Git’s dubious-ownership guard under the sandbox identity, so this review did not independently verify HEAD or dirty state. No runtime, provider, network, or test execution was performed. The architecture’s baseline and runtime acceptance statements are therefore not independently re-certified here.

## Receipt

**Final reviewer verdict:** ACCEPT. Both findings are resolved. No further compatibility, schema, or reducer ambiguities were identified in the amended AD-7/AD-8 scope. Git identity and runtime/test limitations above remain unchanged.
