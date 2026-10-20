# Architecture source reality review

Status: ACCEPT / RATIFIED AFTER CORRECTIONS
Reviewer lens: independent source-reality check of existing pins, boundaries, test names, functional skip behavior, and configuration precedence.
Scope: read-only inspection; no product source, runtime, dependency, commit, or network changes.

## Verified

- `go/go.mod` pins Go 1.21 and `github.com/Agent-Field/agentfield/sdk/go v0.1.130`, consistent with the architecture's no-upgrade decision.
- Existing owners and public signatures are present in `go/internal/orch/orchestrator.go`, `go/internal/github/client.go`, `go/internal/node/node.go`, `go/internal/node/register.go`, `go/internal/config/budget.go`, and `go/internal/config/review.go`. `BuildAgent` wires the SDK app through `Deps.Local`; `New` retains the deliberate offline seam.
- The request/config precedence stated in AD-6 matches `ResolveBudgetCaps`: explicit input, then environment, then 2.0 USD / 3600 seconds. `ReviewConfig.FromInput` maps concurrency and coverage values from input and clamps review depth to 3. The architecture correctly avoids claiming provider spend enforcement.
- All five node test names and all six orchestration test names listed in AD-7 exist at the named Go package boundaries. The all-dimension parse failure test verifies output is not reached; mixed/clean tests verify actual degradation counts.
- Functional tests are `TestHealth`, `TestRegistrationParity`, and `TestReviewErrorShape`. `go/test/functional/main_test.go` returns `m.Run()` on Docker unavailability and readiness timeout; each test's `requireStack` calls `t.Skipf`. Thus exit zero can contain skips, and AD-7's actual test-event requirement is necessary. Compose-up failure returns exit 1.
- Repository source confirmed the Graft findings; MCP graph/coverage tools were unavailable in this reviewer context. No claim of exhaustive graph coverage is made.

## Re-review of corrections

- The architecture now records `revalidated_head: c2611096cd41cc5a6404a8ee658cfe1bf6c6e744` and distinguishes the historical design baseline from the current checkout. This resolves the stale-baseline concern. Its refreshed inventory also truthfully records the inaccessible task output directory rather than claiming to have rechecked it.
- AD-7 now makes the event reducer and environment assessment separate. A mandatory skip remains FAIL for the deterministic event predicate; when preserved `TestMain` evidence shows Docker/readiness prerequisites unavailable, the owning functional gate is BLOCKED_ENV and explicitly not a product-defect finding. This preserves the PRD's permitted BLOCKED/FAIL distinction while preventing skipped tests from passing.
- AD-7 now requires a separate stack receipt tying the Compose project, image/build identity, readiness observations, endpoint identities, and test interval to the functional transcript. Test events alone cannot establish stack identity/readiness.

## Verdict

The architecture's Go interfaces, package boundaries, technology pins, test names, and config precedence are grounded in current source. The three prior findings are resolved; ratify the corrected architecture. This receipt does not establish PLAN acceptance, runtime identity, any test execution, functional acceptance, semantic acceptance, or SourceLoop durability.
