# AGENTS.md

## Repository entry point

PR-AF owns its AgentField implementation, project plan, package contract, tests and benchmark behavior.

Read `PLAN.md`, `README.md`, `ERRORS.md`, `agentfield-package.yaml`, and the relevant `docs/`, `src/`, `tests/` and `benchmark/`. Treat model/reasoner execution status, deterministic checks and user-visible PR outcome as separate acceptance layers.

## Fast Verified Engineering

Canonical standard: `n0namer/server-ops:docs/standards/FAST_VERIFIED_ENGINEERING.md`.

Optimize for **time-to-verified-running-change**. Before mutation resolve `Project North Star -> Phase Goal -> gate/DoD -> next bounded move`, observe exact source/dirty state and runtime identity when relevant, and define scope/rollback/evidence.

Keep design SoT, source, loaded runtime, execution, deterministic validation and functional/semantic outcome separate. Code-on-disk != loaded runtime; tests != deploy proof; health/HTTP 200 or `succeeded` != acceptance.

Route: runtime-bound defect -> authorized permanent DEV loop with bounded stale-safe patch, affected check, reload/canary/log evidence. Source-bound or multi-file work -> exact-SHA isolated Coding Station. GitHub/CI/deploy -> canonicalization/release boundary, not the default inner debug loop.

Preserve the exact verified delta across lanes; do not rewrite it. Validation ladder: `syntax/static -> affected tests -> related regression -> full required suite -> runtime smoke/integration -> semantic/business/E2E`.

Diagnose narrow-to-broad and correlate logs with source/runtime and execution/request identity. Production live editing is forbidden by default. Preserve unrelated state; inspect post-state after timeout/ambiguity before retry.

Final status: `DONE | PARTIAL | BLOCKED | FAILED | EVIDENCE_MISSING`; DONE requires every DoD evidence item.