# PR-AF Error Ledger

## 2026-09-13 — mockcli meta output drifted from the live private meta schema

**Symptom.** `TestReviewHandlerWithExternalMockHarness` failed in `meta_semantic` with structured-output/schema recovery errors even though the mock returned the expected top-level keys.

**Cause.** Production `runMetaLens` now validates harness output against the private `metaDraftResult/metaDraftDimension` shape, but `test/mockcli.roleMeta` still emitted the larger public `schemas.MetaDimensionResult/ReviewDimension` shape. The extra nested fields (`id`, `context_files`, `priority`, `budget`) made the external harness output invalid for the live private draft schema.

**Fix.** `test/mockcli.roleMeta` now emits an exact private-compatible mock shape: each dimension contains only `name`, `review_prompt`, and `target_files`; the result contains `lens`, `dimensions`, `confidence`, and `rationale`.

**Prevention.** When a reasoner changes its private harness destination type, run an external-harness integration path, not only direct reasoner unit tests. Mock output structs must mirror the actual `harnessx.Run[T]` destination shape rather than a related public pipeline model.

**Evidence / verification.** The failing node-level integration reproduced the drift before the fix. After the fix, `go test ./internal/node -run TestReviewHandlerWithExternalMockHarness -count=1 -v`, `go test ./internal/node ./test/mockcli -count=1`, and `go test ./... -count=1` all pass on the exact live source. Canonical fix: `82659fa6b0e31ed6bcf6cd3204cd1f961c67a625`.
