# PR-AF — canonical project plan

Status: ACTIVE / RUNTIME_DEBUG
Updated: 2026-08-31
Owner/workstream: PR-AF / `n0namer/pr-af`
Source of truth: this file (`PLAN.md`)
BMad route: `[BH] bmad-help -> [QQ] bmad-quick-dev`, freeform brownfield, plan-code-review

## North Star

A real pull request is reviewed end-to-end by PR-AF in the permanent AgentField DEV runtime with the expected 17-reasoner execution surface, evidence-grounded findings, bounded resource use, observable failures, and a reproducible exact source identity; accepted runtime fixes are captured durably by AgentField SourceLoop without using Git/redeploy as the inner debugging transport.

North-Star proof requires all of:

1. PR-AF process is running and healthy in permanent DEV.
2. AgentField discovery shows node `pr-af` and the expected reasoner surface.
3. A bounded deterministic/functional canary passes.
4. One real semantic PR review completes and returns useful evidence-grounded output or an explicit justified failure.
5. The exact runtime source/binary identity used for the proof is recorded.
6. Any accepted live source delta is durably represented in `n0namer/pr-af:dev` through SourceLoop, with no required container-only delta left behind.

## Operating contract / anti-drift

### Authority

- CURRENT runtime owns claims about what is actually running now.
- This `PLAN.md` owns project intent, decisions, gates, backlog and last verified state.
- Git source owns accepted durable code, not transient debugging state.
- Historical runtime evidence is evidence only; it does not override a newer CURRENT readback.

### Container-first inner loop

For PR-AF debugging use:

```text
observe CURRENT runtime
-> bounded live patch in persistent DEV source workspace
-> targeted regression
-> rebuild/reload/restart only the PR-AF process when actually required
-> functional canary
-> semantic E2E
-> iterate
-> accepted delta -> SourceLoop -> durable dev SHA
```

Do NOT use this as the inner loop:

```text
edit GitHub -> commit -> redeploy -> wait -> test -> repeat
```

Git is the durable checkpoint for an accepted increment. It is not the mandatory transport for every debug hypothesis.

### Mutation boundaries

- Target project/code: `n0namer/pr-af` only.
- Permanent DEV only for runtime debugging.
- Production is out of scope unless separately authorized.
- Never expose secret values in logs, documents or chat.
- Before runtime mutation: read CURRENT state, record exact source/process/config identity, make one bounded change, then read back and verify.
- Do not restart/redeploy unrelated AgentField components.
- Do not change SourceLoop/control-plane/fleet semantics merely to make PR-AF pass.

## Verified repository state

### Branch/source state — 2026-08-31

- Default branch: `main`.
- Common executable-code baseline of current `main`/`dev`: `48ae7eeb4f07779004db6354728d49ca7b36dbc3` (`fix: default to deepseek/deepseek-v4-flash-0731`).
- `main` additionally contains documentation commit `3d4a728aeb37978233e8657ed1dd004fc052433e` adding `ERRORS.md`.
- `dev` additionally contains SourceLoop canary commit `7901b5a41baf2ad46d23361ba80716d85883e485` created by AgentField Runtime Capture from base `48ae7eeb...`.
- `main` and `dev` are therefore diverged by one documentation/canary commit each; no newer executable PR-AF code delta has been identified between them.
- `PLAN.md` did not exist on either `main` or `dev` before this plan bootstrap.

### SourceLoop contract already proven

AgentField SourceLoop watches PR-AF at:

```text
/src/pr-af
```

The runtime-capture watcher is proposal-only. Stable source deltas are captured to `runtime-capture/pr-af/<base-sha>`, validated, and only then fast-forwarded to `pr-af:dev` when the runtime base is still current. The existing `7901b5a...` canary proves the PR-AF repository round-trip transport worked at least once; it does NOT prove PR-AF semantic runtime correctness.

### Current PR-AF package contract

Current maintained node is the Go package:

- package/node id: `pr-af`
- source subtree: `go/`
- build: `./cmd/pr-af`
- start: `bin/pr-af`
- healthcheck: `/health`
- default port: `8007` in the package manifest; permanent fleet may override the runtime port
- required model credential: `OPENROUTER_API_KEY`
- `GH_TOKEN` is optional for public-read reviews and needed for private repositories/posting reviews/unthrottled GitHub use
- default provider: `aforge`
- default model: `deepseek/deepseek-v4-flash-0731`
- expected registered surface: top-level `review` plus 16 child reasoners

## Historical runtime evidence — do not treat as CURRENT

The last canonical reconciliation evidence from 2026-08-24/25 observed PR-AF installed in AgentField workforce but stopped and absent from discovery. It found no durable downstream PR-AF source delta: the installed historical source matched its historical upstream baseline, while later upstream changes were already present in the fork.

Historical runtime paths included:

```text
/afhome/packages/pr-af
/afhome/packages/pr-af/bin/pr-af
/afhome/logs/pr-af.log
```

Historical blocker was orchestration/runtime validation, not a proven component-source defect.

The current user state on 2026-08-31 is that PR-AF has now been raised in permanent DEV and needs debugging. This supersedes the historical `stopped` statement as intent/current report, but runtime health, registration, source identity and semantic behavior remain `UNVERIFIED_CURRENT` until direct readback.

## Current phase

BMad phase: implementation / brownfield debugging.

Current PDCA state:

- **Value:** prove PR-AF actually reviews real PRs correctly in permanent DEV, not merely that source/tests/deployment exist.
- **State:** repository/source contract and SourceLoop capture path are known; user reports PR-AF raised; CURRENT runtime evidence is not yet read back in this session.
- **Gap:** no current evidence chain `source -> binary -> process -> registration -> execution -> semantic result`.
- **One constraint:** preserve the fast container-first loop; do not convert each hypothesis into Git/redeploy.

## SMART goal

Within the next bounded debug cycle, establish one evidence chain for the raised PR-AF runtime and close the first reproducible failure without a full-stack redeploy between hypotheses.

Success means:

- exact source/process identity recorded;
- one failure reproduced at least twice when nondeterminism permits, or classified from deterministic runtime evidence;
- smallest targeted regression goes red -> green for any code defect;
- PR-AF-only reload/rebuild used only if code execution requires it;
- functional canary passes;
- semantic review is run or the exact blocking gate is proven;
- accepted code delta, if any, is captured durably with exact `dev` SHA.

## Current hypothesis

**If** PR-AF is debugged directly from the persistent DEV source/workspace with targeted tests and PR-AF-only process reloads, **then** a validated fix can be reached with fewer lifecycle operations than Git/redeploy iteration, **because** feedback comes from the actual runtime before durable promotion; **metric:** zero full-stack redeploys per exploratory hypothesis and one accepted evidence-backed increment; **deadline:** current 48–72 h debug cycle.

## Immediate 30-minute batch — CURRENT runtime localization

This is the nearest mandatory step. Do not patch component code before it is complete.

### Critical task 1 — prove runtime identity

Output:

```text
source SHA/worktree state
-> built/installed binary identity
-> running PR-AF process
-> node registration
```

Done when all four links are evidenced from CURRENT permanent DEV, including whether the active source workspace is `/src/pr-af` and whether the running binary was built from that workspace/generation.

Risk: a healthy/stale process may not correspond to the watched SourceLoop source tree.

### Critical task 2 — capture first reproducible failure

Output: one bounded test/review execution with execution id, failing phase, relevant bounded logs, and expected-vs-actual behavior.

Done when the first failure is localized to one of:

- startup/process supervision;
- AgentField registration/callback;
- provider/model/harness;
- repository intake/GitHub access;
- PR-AF orchestration/reasoner;
- evidence/coverage/synthesis;
- external dependency/configuration.

Risk: a large real PR can make the feedback loop too slow. Prefer the smallest representative canary first.

### Support tasks

1. Verify SourceLoop is watching `/src/pr-af` and has no blocked/stale capture state.
2. Record only presence/config shape for credentials; never print secret values.
3. Read `ERRORS.md` before any substantial code mutation and append only verified recurring/material errors.

## Implementation rule after localization

For a proven PR-AF code defect:

```text
one bounded patch in persistent DEV source
-> smallest relevant Go test(s)
-> broader package regression if justified
-> PR-AF-only rebuild/reload if required
-> functional canary
-> semantic canary/E2E
-> SourceLoop durable capture
-> exact dev SHA readback
```

A fix is NOT accepted if required behavior exists only in the container.

A SourceLoop capture is NOT semantic acceptance by itself.

## Acceptance gates

### G0 — CURRENT runtime identity
Status: OPEN

- exact source/worktree proven
- exact process/binary proven
- health proven
- registration/reasoner surface proven

### G1 — reproducible failure localization
Status: OPEN

- bounded execution captured
- expected vs actual stated
- failing layer classified
- logs/evidence sufficient to test one hypothesis

### G2 — targeted repair
Status: BLOCKED_BY_G1

- minimal patch only if component defect is proven
- targeted regression PASS
- no unrelated code/config mutation

### G3 — functional PR-AF canary
Status: BLOCKED_BY_G2_OR_NO_CODE_CHANGE

- review endpoint/call reaches PR-AF
- execution DAG/reasoners behave coherently
- deterministic/fixture canary completes

### G4 — semantic E2E
Status: BLOCKED_BY_G3

- one representative real PR review completes
- findings are grounded in code evidence
- obvious unsupported findings are pruned
- result is useful or failure/abstention is explicit and faithful

### G5 — durable SourceLoop closure
Status: BLOCKED_BY_ACCEPTED_DELTA

- captured diff matches accepted runtime delta
- safety/validation passes
- exact `pr-af:dev` SHA recorded
- no required container-only delta remains
- later fleet promotion is a separate acceptance step

## Metrics

Record per accepted batch:

- time to reproduce
- patch -> targeted-feedback latency
- PR-AF process reload count
- full-stack redeploy count
- targeted regression PASS/FAIL
- functional canary PASS/FAIL
- semantic E2E PASS/FAIL
- runtime delta -> durable dev SHA PASS/FAIL

North-star operating metric: accepted fixes per full-stack deployment cycle.

## Decisions

1. `PLAN.md` is the single PR-AF project SoT. Do not create `PLAN-v2`, sidecar handoffs or duplicate project plans.
2. Use BMad `[BH]` to recover state and `[QQ]` Quick Dev for bounded brownfield implementation batches unless a later verified contradiction requires `[CC]` Correct Course.
3. Container-first debugging is mandatory for the inner loop.
4. Runtime evidence outranks historical runbooks for current-state claims.
5. `main` documentation updates must not be confused with `dev` runtime source acceptance.
6. Do not remotely advance `pr-af:dev` merely to update this plan while a SourceLoop-watched runtime may still be based on its current exact SHA.

## Risks

- Running PR-AF may be sourced from a different generation than `/src/pr-af`.
- SourceLoop may reject a capture if runtime base and current `dev` moved independently.
- A process can be healthy while registration/callback or semantic review is broken.
- Unit/Go test success does not prove semantic review quality.
- Historical docs may still describe the earlier stopped/credential-gated state.

## Backlog

Ordered, not parallelized:

1. G0 CURRENT runtime identity.
2. G1 first reproducible failing execution.
3. G2 smallest justified repair.
4. G3 functional canary.
5. G4 semantic E2E.
6. G5 SourceLoop durable closure.
7. Only after DEV acceptance: reconcile any fleet/source-root/promotion state as a separate operation.

## Next three actions

1. Read CURRENT permanent DEV PR-AF process/source/registration state without mutation.
2. Run the smallest representative PR-AF canary and collect bounded execution/log evidence.
3. Rebuild the plan from that evidence; patch only the proven failing layer.

## Batch write-back contract

At the end of every material batch update this file in place with:

```text
Timestamp UTC:
Phase goal:
Observed exact source/runtime state:
Actions performed:
Evidence/tests/canaries:
Metric/variance:
Verified cause or remaining hypotheses:
Decision: standardize | continue | change | stop | scale
Changed files:
Durable commit/SHA if any:
Current constraint:
One next bounded move:
```

Do not mark a gate PASS from planned work, source inspection alone, HTTP health alone, or conversation memory.