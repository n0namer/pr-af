# PR-AF — LLM Provider & Security Profile

Status: canonical repo-local provider profile
Cross-component contract: `n0namer/universal-solver:main/docs/architecture/llm-provider-security-contract.md`
Current project SoT: `PLAN.md`

## Scope

This profile defines PR-AF-specific LLM/provider/security behavior. It does not replace `PLAN.md`, `README.md`, `docs/ARCHITECTURE.md`, or `AGENTS.md`, and it does not contain secret values.

## Maintained runtime

The maintained PR-AF implementation is the Go package under `go/`.

Relevant package/runtime facts:
- node id: `pr-af`
- default port: `8007`
- build: `./cmd/pr-af`
- start: `bin/pr-af`
- `GH_TOKEN` is a separate GitHub capability; it is not an LLM credential.

## Provider surfaces

Relevant provider/runtime variables include:

- `OPENAI_API_KEY`
- `OPENAI_BASE_URL`
- `OPENROUTER_API_KEY`
- `ANTHROPIC_API_KEY`
- `GOOGLE_API_KEY`
- PR-AF provider/model selection variables/configuration used by the Go runtime

The permanent AgentField DEV topology currently injects OpenAI-compatible key/base variables by name. Environment presence alone is not proof that a PR-AF review call used that provider.

## Canonical OpenAI-compatible transport contract

PR-AF standardizes on one provider transport contract for OpenAI-compatible backends, including Gonka:

```text
model/provider = openai/<intended model>
OPENAI_API_KEY = <provider credential from runtime secret store>
OPENAI_BASE_URL = <intended OpenAI-compatible endpoint>
```

`OPENAI_API_KEY` and `OPENAI_BASE_URL` are the canonical transport variable names. PR-AF MUST NOT introduce `AI_BASE_URL` as a parallel public/fleet contract. Adapter-specific aliases may exist only as explicitly documented compatibility shims and must not become a second source of truth.

The credential, endpoint/base, and selected model/provider MUST survive together to the final harness/model call. A key without the intended base URL, or a model without the intended key/base pair, is an invalid provider state.

Deep Research is the proven reference pattern: model/key overrides must preserve the configured custom `api_base`. PR-AF should preserve the same invariant at its OpenCode/harness boundary.

## Current source evidence and required PR-AF delta

`go/internal/config/ai.go::ProviderEnv()` currently forwards `OPENAI_API_KEY` but does not forward `OPENAI_BASE_URL`.

Required minimal PR-AF implementation delta:
- add `OPENAI_BASE_URL` to the environment forwarded by `ProviderEnv()`;
- add a focused regression in `go/internal/config/config_test.go` proving exact `OPENAI_BASE_URL` propagation and absence when unset;
- do not add `AI_BASE_URL` to PR-AF as another canonical variable;
- after the source-level regression passes, prove the real model path and absence of unintended OpenRouter/default-OpenAI fallback before semantic acceptance.

This is now an approved design decision for PR-AF, not merely a naming preference. Runtime `BASE_URL_LOSS` remains unproven until a real provider call is observed, so semantic/runtime PASS still requires execution evidence.

## Future SWE convergence map

SWE is intentionally not modified by this PR-AF task. When SWE is migrated to the same fleet-wide contract, inspect and update these owners together:

- `n0namer/swe-af:dev/docs/LLM_PROVIDER_SECURITY_CONTRACT.md` — change the public/canonical transport wording to `OPENAI_API_KEY + OPENAI_BASE_URL + openai/<model>` and mark `AI_BASE_URL` compatibility-only if still supported;
- `n0namer/swe-af:dev/go/internal/node/node.go` — verify `resolveAIConfig()` still produces the intended key/base/model triple after the transport-name migration;
- `n0namer/swe-af:dev/go/internal/node/aiconfig_test.go` — migrate/add regression coverage for `OPENAI_BASE_URL` as canonical input and explicitly test any temporary `AI_BASE_URL` compatibility fallback;
- `n0namer/agentfield:main/sdk/go/ai/config.go` — current SDK-level `AI_BASE_URL` reader is the compatibility boundary. Preferred migration is `OPENAI_BASE_URL` first, optional `AI_BASE_URL` fallback for a bounded deprecation window, with tests proving precedence and no silent fallback;
- `n0namer/universal-solver` AgentField DEV topology — ensure SWE receives only the canonical fleet transport values at the deployment boundary; avoid maintaining two independently configurable endpoint variables.

SWE migration acceptance should use the same invariant as PR-AF/Deep Research: intended `openai/<model>` + intended `OPENAI_API_KEY` + intended `OPENAI_BASE_URL` reach the actual client together, and no unintended default OpenAI/OpenRouter fallback occurs.

The current Go package manifest also remains OpenRouter-oriented at bootstrap/credential-contract level. If maintained PR-AF supports an OpenAI-compatible lane but package admission rejects it, classify that as `BOOTSTRAP_ADMISSION` drift rather than forcing a real OpenRouter dependency.

## GitHub vs LLM credentials

Treat these as separate capabilities and acceptance gates:

### LLM provider capability
Responsible for:
- model selection;
- provider endpoint;
- review/reasoning quality.

### GitHub capability (`GH_TOKEN`)
Responsible for:
- private repository access when required;
- posting review/comments when enabled.

A successful LLM review without GitHub posting may still prove review correctness. Conversely, successful GitHub API access does not prove LLM/provider correctness.

## Security requirements

- Never commit, print, or document raw LLM/GitHub secret values.
- Record only variable names, presence/config state, secret-store owner, and redacted provider identity.
- Do not use credentials in URLs/CLI arguments where process/log exposure is possible.
- Installer/runtime diagnostics must report secret presence, not values.
- Provider fallback MUST be explicit. Unexpected fallback to OpenRouter/Anthropic/another provider is a functional failure.
- Real PR-review E2E evidence must mask repository/private data and secrets.
- SourceLoop/runtime-capture artifacts must pass secret scanning before durable promotion.

## Acceptance ladder

1. Exact PR-AF source/runtime identity known.
2. Maintained `/src/pr-af/go` package installed and started.
3. `pr-af` registered in AgentField; expected reasoner discoverable.
4. Intended provider/model selection resolved explicitly.
5. For OpenAI-compatible routing, `OPENAI_API_KEY` and `OPENAI_BASE_URL` reach the final harness/model client.
6. Minimal real reasoner/model call succeeds.
7. Execution/log evidence shows the intended provider/model and no unintended fallback.
8. Deterministic/functional PR-AF canary succeeds.
9. One bounded real PR review produces a semantically useful result with code-grounded findings (or a correct no-finding result) and no obvious hallucinated findings.
10. If code changed, the smallest relevant regression is red→green and broader required Go validation passes.
11. Accepted runtime delta is captured by SourceLoop only after verification; durable `pr-af:dev` identity is reconciled and verified.

Health, HTTP 200, node registration, or execution `succeeded` alone are insufficient for provider/semantic PASS.

## Failure classes

Use the cross-component classes:
- `BOOTSTRAP_ADMISSION`
- `MODEL_RESOLUTION`
- `ENV_PROPAGATION`
- `BASE_URL_LOSS`
- `AUTH`
- `FALLBACK`
- `TRANSPORT`
- `SEMANTIC`

Patch the first failing layer only.

## Current execution rule

For runtime-bound defects, use the existing permanent DEV workspace and edit PR-AF code directly in `/src/pr-af` only after runtime evidence proves a code defect. Do not use GitHub-first coding/redeploy as the inner loop.

After material evidence or state changes, update `PLAN.md` in place.
