# PR-AF Error Ledger

## 2026-09-13 — mockcli meta output drifted from the live private meta schema

**Symptom.** `TestReviewHandlerWithExternalMockHarness` failed in `meta_semantic` with structured-output/schema recovery errors even though the mock returned the expected top-level keys.

**Cause.** Production `runMetaLens` now validates harness output against the private `metaDraftResult/metaDraftDimension` shape, but `test/mockcli.roleMeta` still emitted the larger public `schemas.MetaDimensionResult/ReviewDimension` shape. The extra nested fields (`id`, `context_files`, `priority`, `budget`) made the external harness output invalid for the live private draft schema.

**Fix.** `test/mockcli.roleMeta` now emits an exact private-compatible mock shape: each dimension contains only `name`, `review_prompt`, and `target_files`; the result contains `lens`, `dimensions`, `confidence`, and `rationale`.

**Prevention.** When a reasoner changes its private harness destination type, run an external-harness integration path, not only direct reasoner unit tests. Mock output structs must mirror the actual `harnessx.Run[T]` destination shape rather than a related public pipeline model.

**Evidence / verification.** The failing node-level integration reproduced the drift before the fix. After the fix, `go test ./internal/node -run TestReviewHandlerWithExternalMockHarness -count=1 -v`, `go test ./internal/node ./test/mockcli -count=1`, and `go test ./... -count=1` all pass on the exact live source. Canonical fix: `82659fa6b0e31ed6bcf6cd3204cd1f961c67a625`.

## 2026-09-14 — declared output formats were inert at runtime

**Symptom.** `ReviewInput.output_format` advertised `github | json | sarif | markdown`, but the node always returned `ReviewResult` and allowed the normal GitHub-posting path regardless of requested non-GitHub format.

**Cause.** `OutputFormat` was defined in schemas/defaults/input API but had no runtime consumer in `reviewHandler` or output transport adaptation.

**Fix.** `reviewHandler` now validates/normalizes the format, forces non-GitHub modes to dry-run, rejects unknown formats before execution, and adapts the completed result to Markdown or SARIF while preserving structured output for GitHub/JSON.

**Prevention.** For every public input mode declared by architecture/schema, require at least one handler-level test that proves the mode changes executable behavior rather than only parsing/defaulting successfully.

**Evidence / verification.** Focused tests first RED on all non-GitHub modes and invalid format, then GREEN after the live patch. `go test ./internal/node -count=1`, `go test ./... -count=1`, and `go test -vet=all ./... -run '^$' -count=1` all pass on the exact live source. Canonical implementation commits: `739b3367b779db756fba58f029df1f5af39ce47a`, `27a5dd84baaa3fc96926265275486960b9acbb7c`, `9e8fbb13e386f1d333df1a1edd0fbcbc57d4dca7`, `5e6406b3a236e41480a3f9abc887c3e80a83bdbb`, `ec5c958e2afdfca3532bc63ac08745309a120743`.
