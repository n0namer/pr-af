# PR-AF — canonical project plan

Status: ACTIVE / RUNTIME_DEBUG
Updated: 2026-08-31
Owner/workstream: PR-AF / `n0namer/pr-af`
Source of truth: this file (`PLAN.md`)
BMad route: `[BH] bmad-help -> [QQ] bmad-quick-dev`, freeform brownfield, plan-code-review

## North Star

A real pull request is reviewed end-to-end by the **maintained PR-AF implementation** in the permanent AgentField DEV runtime with the expected 17-reasoner execution surface, evidence-grounded findings, bounded resource use, observable failures, and a reproducible exact source identity; accepted runtime fixes are captured durably by AgentField SourceLoop without using Git/redeploy as the inner debugging transport.

North-Star proof requires all of:

1. The maintained PR-AF process is running and healthy in permanent DEV.
2. AgentField discovery shows node `pr-af` and the expected reasoner surface.
3. A bounded deterministic/functional canary passes.
4. One real semantic PR review completes and returns useful evidence-grounded output or an explicit justified failure.
5. The exact runtime implementation/source/binary identity used for the proof is recorded.
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

- Target project/code: `n0namer/pr-af` only unless runtime evidence proves the defect is orchestration-owned; do not patch PR-AF to compensate for an orchestration defect.
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

The runtime-capture watcher is proposal-only. Stable source deltas are captured to `runtime-capture/pr-af/<base-sha>`, validated, and only then fast-forwarded to `pr-af:dev` when the runtime base is still current. The existing `7901b5a...` canary proves the PR-AF repository round-trip transport worked at least once; it does NOT prove PR-AF semantic runtime correctness or which PR-AF implementation is executing.

### PR-AF has two installable implementations

The repository root and `go/` deliberately expose the same node id `pr-af` but different implementations.

**Maintained Go node (`go/`):**

- package/node id: `pr-af`
- build: `./cmd/pr-af`
- start: `bin/pr-af`
- healthcheck: `/health`
- default port: `8007`
- required model credential: `OPENROUTER_API_KEY`
- `GH_TOKEN` optional for public-read reviews; needed for private repos/posting/unthrottled GitHub use
- default provider: `aforge`
- default model: `deepseek/deepseek-v4-flash-0731`
- expected surface: top-level `review` plus 16 child reasoners

**Root local-path Python node:**

- entrypoint: `python -m pr_af.app`
- healthcheck: `/health`
- default port: `8004`
- root manifest declares `superseded_by: https://github.com/Agent-Field/pr-af//go`
- that redirect applies to Git installs; the manifest explicitly states that installing a cloned checkout/local path deliberately installs the Python node instead
- root manifest still declares `GH_TOKEN` required

### Verified static integration drift — 2026-08-31

Current `universal-solver:main` permanent/isolated DEV workforce materialization does all of the following:

```text
ensure_exact pr-af 7901b5a41baf2ad46d23361ba80716d85883e485 /src
...
af install /src/pr-af
```

The same Compose leaves PR-AF install-only by default and does not auto-start it.

Because `/src/pr-af` is a **local-path install**, the PR-AF root manifest says this selects the Python implementation rather than following `superseded_by` to `go/`.

Therefore the static deployment contract currently implies:

```text
SourceLoop watched source = /src/pr-af
fleet install command      = af install /src/pr-af
installed implementation   = Python root package (by manifest contract)
maintained implementation  = Go package under /src/pr-af/go
```

This is a verified orchestration/source-selection inconsistency. It does **not** yet prove which implementation the currently raised PR-AF process is executing, because a later manual install/start may have replaced the installed package.

**H1 — implementation identity drift**

- Status: `VERIFIED_STATIC / UNVERIFIED_CURRENT_RUNTIME`.
- Hypothesis: the raised DEV PR-AF may be the Python implementation while the intended/current maintained implementation is Go.
- First discriminating evidence: `/afhome/installed.yaml` source/source_path + actual process argv/executable + installed package layout + listen port.
- Decision rule: if CURRENT confirms Python was installed by `af install /src/pr-af`, do not patch PR-AF product logic first; classify the mismatch as orchestration-owned and correct the install/start path to the maintained Go package through the appropriate source-of-truth/runtime path.

### Callback identity is the second startup discriminator

Go PR-AF maps `AGENT_CALLBACK_URL` to AgentField `PublicURL`. If unset, the SDK falls back to a localhost callback for the node port. A control plane in a different container cannot use a node-local localhost address as its cross-container callback.

The canonical workforce start pattern for other agents explicitly sets `AGENT_CALLBACK_URL=http://workforce:<port>` before `af run`. PR-AF is not auto-started by that Compose, so a manual start that omitted the callback is a plausible runtime/config failure mode.

**H2 — callback routing**

- Status: `SOURCE-CONTRACT VERIFIED / CURRENT VALUE UNKNOWN`.
- Check only after implementation identity: process environment presence/shape for `AGENT_CALLBACK_URL`, registered callback/public URL, and node reachability from the control plane.
- Do not print secret values while checking runtime environment.

## Historical runtime evidence — do not treat as CURRENT

The last canonical reconciliation evidence from 2026-08-24/25 observed PR-AF installed in AgentField workforce but stopped and absent from discovery. It found no durable downstream PR-AF source delta: the installed historical source matched its historical upstream baseline, while later upstream changes were already present in the fork.

Historical runtime paths included:

```text
/afhome/packages/pr-af
/afhome/packages/pr-af/bin/pr-af
/afhome/logs/pr-af.log
```

Historical blocker was orchestration/runtime validation, not a proven component-source defect.

The current user state on 2026-08-31 is that PR-AF has now been raised in permanent DEV and needs debugging. This supersedes the historical `stopped` statement as intent/current report, but runtime health, registration, implementation/source identity and semantic behavior remain `UNVERIFIED_CURRENT` until direct readback.

## Current phase

BMad phase: implementation / brownfield debugging.

Current PDCA state:

- **Value:** prove the maintained PR-AF actually reviews real PRs correctly in permanent DEV, not merely that source/tests/deployment exist.
- **State:** repository/source contract and SourceLoop capture path are known; static inspection has exposed a Python-vs-Go install-path inconsistency; user reports PR-AF raised; CURRENT runtime identity is not yet read back in this session.
- **Gap:** no current evidence chain `implementation -> source -> installed artifact -> process -> registration -> execution -> semantic result`.
- **One constraint:** preserve the fast container-first loop; do not convert each hypothesis into Git/redeploy.

## SMART goal

Within the next bounded debug cycle, establish one evidence chain for the raised PR-AF runtime, resolve implementation identity, and close the first reproducible failure without a full-stack redeploy between hypotheses.

Success means:

- exact implementation/source/process identity recorded;
- Python-vs-Go mismatch either ruled out or classified and corrected at its owning layer;
- one failure reproduced at least twice when nondeterminism permits, or classified from deterministic runtime evidence;
- smallest targeted regression goes red -> green for any component-code defect;
- PR-AF-only reload/rebuild used only if code execution requires it;
- functional canary passes;
- semantic review is run or the exact blocking gate is proven;
- accepted code delta, if any, is captured durably with exact `dev` SHA.

## Current hypothesis

**If** PR-AF is debugged directly from the persistent DEV source/workspace with implementation identity proven first, targeted tests and PR-AF-only process reloads, **then** a validated fix can be reached with fewer lifecycle operations and less misdiagnosis than Git/redeploy iteration, **because** feedback comes from the actual runtime and the correct implementation before durable promotion; **metric:** zero full-stack redeploys per exploratory hypothesis and one accepted evidence-backed increment; **deadline:** current 48–72 h debug cycle.

## Immediate 30-minute batch — CURRENT runtime localization

This is the nearest mandatory step. Do not patch component code before it is complete.

### Critical task 1 — prove implementation/runtime identity

Output:

```text
installed implementation (Python or Go)
-> source/source_path + SHA/worktree state
-> installed artifact / binary or Python entrypoint
-> running PR-AF process argv/executable
-> node registration + callback + port
```

Minimum readback set:

- `/afhome/installed.yaml` entry for `pr-af` without secret values;
- `/src/pr-af` HEAD/status;
- installed package entrypoint/layout under `/afhome/packages/pr-af`;
- actual PR-AF process command/executable and listen port;
- AgentField discovery entry for `pr-af`, including reasoner count/surface and callback/public URL when exposed;
- SourceLoop state for `pr-af`.

Done when the chain identifies exactly **which implementation is executing** and whether it corresponds to the SourceLoop-watched source generation.

Risk: a healthy/stale process may not correspond to `/src/pr-af`; a local-path install may have selected Python while operators assume Go.

### Critical task 2 — capture first reproducible failure

Only after task 1.

Output: one bounded test/review execution with execution id, failing phase, relevant bounded logs, and expected-vs-actual behavior.

Done when the first failure is localized to one of:

- wrong implementation/source generation;
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

For a proven PR-AF component-code defect:

```text
one bounded patch in the source tree belonging to the proven running implementation
-> smallest relevant test(s)
-> broader package regression if justified
-> PR-AF-only rebuild/reload if required
-> functional canary
-> semantic canary/E2E
-> SourceLoop durable capture
-> exact dev SHA readback
```

If the failure is implementation selection/callback/fleet orchestration, fix the owning orchestration source rather than adding compensating PR-AF product code.

A fix is NOT accepted if required behavior exists only in the container.

A SourceLoop capture is NOT semantic acceptance by itself.

## Acceptance gates

### G0 — CURRENT implementation/runtime identity
Status: OPEN

- Python vs Go implementation proven
- exact source/worktree proven
- exact process/artifact proven
- health proven
- registration/reasoner surface proven
- callback/public URL routing proven

### G1 — reproducible failure localization
Status: BLOCKED_BY_G0

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
5. Prove Python-vs-Go implementation identity before interpreting PR-AF behavior or editing product code.
6. `main` documentation updates must not be confused with `dev` runtime source acceptance.
7. Do not remotely advance `pr-af:dev` merely to update this plan while a SourceLoop-watched runtime may still be based on its current exact SHA.
8. Do not patch PR-AF to conceal an `af install /src/pr-af` implementation-selection mistake; that defect belongs to orchestration.

## Risks

- Running PR-AF may be the legacy Python local-path package while the operator assumes maintained Go.
- Running PR-AF may be sourced from a different generation than `/src/pr-af`.
- A manual `af run` may have omitted a routable `AGENT_CALLBACK_URL`.
- SourceLoop may reject a capture if runtime base and current `dev` moved independently.
- A process can be healthy while registration/callback or semantic review is broken.
- Unit/Go test success does not prove semantic review quality.
- Historical docs may still describe the earlier stopped/credential-gated state.

## Backlog

Ordered, not parallelized:

1. G0 CURRENT implementation/runtime identity.
2. If Python-vs-Go mismatch is confirmed: correct the owning orchestration install/start path, then repeat G0.
3. G1 first reproducible failing execution.
4. G2 smallest justified repair.
5. G3 functional canary.
6. G4 semantic E2E.
7. G5 SourceLoop durable closure.
8. Only after DEV acceptance: reconcile any fleet/source-root/promotion state as a separate operation.

## Next three actions

1. Read CURRENT permanent DEV installed package/process/discovery and prove Python vs Go plus exact source identity.
2. Verify callback/public URL routing and SourceLoop state; then run the smallest representative PR-AF canary.
3. Rebuild the plan from that evidence and patch only the proven owning layer.

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