# AGENTS.md — PR-AF Engineering Contract

Status: active repository-local engineering instructions  
Repository: `n0namer/pr-af`  
Active development branch for this project: `dev`  
Global standard: `n0namer/server-ops:main/docs/standards/FAST_VERIFIED_ENGINEERING.md`

## 1. Mission and authority

Optimize for **time-to-verified-running-change**, not time-to-patch or time-to-merge.

Authority for the current PR-AF workstream:
- Project North Star / product intent: this repo's `README.md` and `docs/ARCHITECTURE.md`.
- Current phase, Phase Goal, bounded batches, DoD, blockers, anti-drift status: this repo's `PLAN.md`.
- Durable PR-AF architecture: this repo's `README.md`, `docs/ARCHITECTURE.md`, and component manifests.
- PR-AF source owner: this repository.
- Permanent AgentField DEV topology and SourceLoop plumbing owner: `n0namer/universal-solver`.
- CURRENT runtime/readback overrides historical runtime-evidence documents for actual state.
- `BMAD-MNNZ` is the workflow/skill rulebook, not project SoT.

Before material mutation:
1. read this file;
2. read root `ERRORS.md` if it exists (currently absent on `dev`);
2. resolve North Star → Phase Goal → current gate/DoD → next bounded move from the project SoT;
4. observe exact source identity, dirty state, and target runtime identity;
5. identify rollback/recovery and the evidence required for completion.

Do not invent commands, targets, credentials, services, or acceptance evidence.

## 2. Fast Verified Engineering loop

Use the smallest evidence-complete route:

`OOBSERVE Ⱂ LOCALIZE → ROUTE → PATCH → TARGETED VERIFY → ITERATE → FULL VERIFY → RUNTIME PROOF → CANONICALIZE → DEPLOY → POST-DEPLOY VERIFY → WRITE-BACK`

Skip stages only when they are irrelevant to the task and the remaining route still satisfies DoD.

Keep these claims separate:
1. **Project/Design SoT** — what should exist.
2. **Source-on-disk** — what files are actually present in the intended worktree.
3. **Loaded runtime** — what process/container/config is actually executing.
4. **Concrete execution** — what happened in a specific request/run/execution.
5. *(Deterministic validation** — what syntax/static/tests passed on the exact intended source.
6. **Functional/semantic outcome** — whether PR-AF actually performed a correct review.

Therefore:
- code-on-disk != loaded runtime;
- test PASS != deploy/runtime proof;
- HTTP 200 / health / execution `succeeded` != functional or semantic acceptance;
- SourceLoop capture success != PR-AF semantic acceptance.

Final status is exactly one of:
`DONE | PARTIAL | BLOCKED | FAILED | EVIDENCE_MISSING`.

`DONE` requires all task-specific DoD evidence.

## 3. Repository-specific implementation boundary

The maintained PR-AF implementation is the Go module under `go/`.

Important install semantics:
- Bare repository URL install follows the root manifest redirect to the maintained Go package.
- A **local-path install of the repository root** deliberately does not follow `superseded_by` and keeps the legacy Python node.
- Therefore, in the current permanent DEV SourceLoop workspace, when the maintained Go node is intended, the local install target is `/src/pr-af/go`, not `/src/pr-af`.

Maintained Go package contract (`go/agentfield-package.yaml`):
- build: `./cmd/pr-af`;
- start: `bin/pr-af`;
- node id: `pr-af`;
- default port: `8007`;
- `OPENROUTER_API_KEY` required;
- `GH_TOKEN` optional: public-repository analysis can work without it, while private repository access / GitHub posting is degraded.

Do not change the legacy Python implementation to repair a failure proven to belong to the maintained Go node.

## 4. Coding lanes

### Runtime-first lane — preferred for current work

Use the already opted-in permanent DEV runtime when the defect depends on:
- AgentField registration;
- process loading/reload;
- provider/config/credential behavior;
- network/integration state;
- runtime workspace behavior;
- semantic E2E.

Current project topology is owned by `universal-solver`:
- permanent AgentField DEV workforce uses a writable persistent `/src`;
- PR-AF live checkout is `/src/pr-af`;
- SourceLoop/runtime-capture observes the same persistent source volume;
- accepted runtime deltas are captured only after validation;
- GitHub/redeploy is **not** the inner hypothesis loop.

Typical runtime loop:
`CURRENT runtime → bounded stale-safe live patch → affected check → same-runtime reload/restart if required → canary → execution/log evidence → iterate`.

Rules:
- production live editing is forbidden by default;
- preserve unrelated dirty/runtime state;
- do not redeploy the whole stack for each hypothesis;
- after timeout or ambiguous mutation, inspect post-state before retrying;
- retry an identical failed mutation at most once unless new evidence changes the diagnosis;
- accepted container-only deltas must not remain container-only.

### Exact-SHA isolated coding lane

Prefer isolated Coding Station for:
- source-bound logic;
- multi-file changes;
- refactors;
- dependency/build work;
- tasks that do not require real target state per iteration.

Start from the exact intended base SHA. Export/replay the **same verified delta** into the runtime/canonical repo; do not rewrite the fix from scratch.

### Canonical/release lane

GitHub/CI/deploy is the durable publication/release boundary, not the default debugging transport.

For the current SourceLoop workstream, code fixes are developed and verified in the permanent DEV workspace first, then SourceLoop captures the exact accepted delta into `pr-af:dev`. Do not use “edit GitHub → redeploy → observe → repeat” as the inner loop.

## 5. Validation commands and ladder

For the maintained Go module, run from `go/`.

Canonical commands proven by the repository:
- `make build` → `go build ./...`
- `make vet` → `go vet ./...`
- `make test` → `go test ./...`
- `make check` → build + vet + test
- `make run` → `go run ./cmd/pr-af`
- `make fmt` → `gofmt -w .`

Do not claim a command ran unless runtime evidence shows it ran.

Validation ladder:
1. syntax / formatting / static check;
2. directly affected tests;
2. related module/component regression;
4. full required Go suite (`make check` / equivalent exact commands);
5. runtime smoke/integration;
6. functional PR-AF acceptance;
7. semantic/business/E2E acceptance.

The tests under `go/test/functional/` cover node health, registration, and reasoner API behavior and are included by the Go test suite.

The deterministic E2E harness is `go/test/e2e/run.sh`. It:
- starts/proves a real AgentField control plane;
- builds the Go PR-AF node;
- uses a seeded local git fixture;
- mocks the external harness path;
- runs `pr-af.review`;
- asserts terminal success, findings, non-empty review body, expected phase activity, and zero GitHub writes.

Important limitation documented by that harness:
- PR-AF `.ai()` gates still require `OPENROUTER_API_KEY`; the harness is not proof of a fully offline semantic run.
- deterministic E2E is not sufficient proof of real-model review quality.

For current Phase DoD, semantic acceptance additionally requires one bounded real PR review with non-dummy provider credentials and outcome inspection.

## 6. Runtime proof and observability

Diagnose narrow-to-broad:
1. structured execution/status evidence;
2. execution-scoped events/logs;
3. bounded PR-AF runtime logs;
4. broader AgentField node/control-plane logs only if needed.

Correlate evidence when available with:
- source SHA / worktree identity;
- loaded process/container identity;
- node id / port;
- request/execution/workflow/correlation id;
- start/reload generation or timestamp;
- provider/model when relevant;
- result/failure class.

For a runtime-relevant change, prove:
`SOURCE SHA → BUILT/START ARTIFACT → RUNNING PROCESS → AGENTFIELD REGISTRATION → CONCRETE EXECUTION → FUNCTIONAL/SEMANTIC OUTCOME`.

A Healthy process or registered reasoner surface is necessary but not sufficient for semantic acceptance.

## 7. Current permanent DEV facts and anti-drift rule

Current project SoT (`universal-solver:main/PLAN.md`) records:
- active PR-AF source/durable branch identity at `7901b5a41baf2ad46d23361ba80716d85883e485`;
- the permanent workforce source mount `/src/pr-af`;
- SourceLoop running over the shared `/src`;
- no current active `pr-af` registration;
- current orchestration installs `/src/pr-af` root, selecting the stopped legacy local-path package instead of the maintained Go package;
- the nearest runtime action is to exercise the maintained Go install/start path without a full-stack redeploy.

Treat these as CURRENT only while confirmed by fresh readback. If actual state differs, record `DESIGN_RUNTIME_DRIFT` in the project SoT and replan from actual state.

Do not patch PR-AF business code while a reproduced failure is still explained by install/start/config/orchestration.

## 8. Safety, rollback, and state preservation

- No production live editing by default.
- Mask secrets; verify presence/configuration without printing values.
- Preserve unrelated files, containers, processes, volumes, and dirty worktrees.
- Scope restart/reload to PR-AF when sufficient; do not restart the whole AgentField/Coolify stack without demonstrated need.
- Before destructive rollback/delete, state what will and will not change.
- Revert only the exact delta owned by this work unless broader cleanup is explicitly authorized.
- Maintain a recoverable pre-change source/runtime identity.
- Missing runner/dependency/environment is `VALIDATION_BLOCKER`, not an application defect.

## 9. Canonicalization and write-back owners

Do not create parallel `v2`, `final`, or sidecar status documents.

Write back:
- current project phase/progress/DoD/blockers/next move → `n0namer/universal-solver:main/PLAN.md`;
- product North Star → `universal-solver/BRIEF.md`;
- durable PR-AF architecture/usage contract → existing PR-AF `README.md` / `docs/ARCHITECTURE.md` / manifests as appropriate;
- engineering operating rules → this `AGENTS.md`;
- accepted PR-AF source delta → exact SourceLoop-captured `pr-af:dev` delta;
- reusable verified error lesson → root `ERRORS.md` only when a canonical error ledger is required/created and the lesson is genuinely reusable.

Historical runtime evidence is evidence, not a substitute for current project status.

## 10. Current project DoD

For the active PR-AF permanent-DEV phase, completion requires:
- exact `/src/pr-af` source identity read back;
- maintained Go PR-AF installed/started without full-stack redeploy;
- `pr-af` registered in CURRENT AgentField;
- expected reasoner surface discoverable;
- deterministic/functional canary PASS;
- one bounded real semantic PR-review E2E PASS with non-dummy provider credentials;
- if code changed: smallest relevant regression red→green, then required broader Go validation;
- accepted code delta captured by SourceLoop;
- durable `pr-af:dev` SHA corresponds to the accepted runtime delta;
- no unintended container-only code delta remains.

If any criterion lacks evidence, do not report `DONE`.
