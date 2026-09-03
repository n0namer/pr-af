### go/internal/orch/phases.go (showing first 400 of 1224)
```
1: package orch
2: 
3: // phases.go ports the finding-producing phases of orchestrator.py: intake,
4: // anatomy, meta-selectors, the streaming review+layer, evidence verification,
5: // parallel adversary, coverage loop, and consistency verify. The concurrency
6: // primitives map (design §D): asyncio.gather → order-preserving errgroup with a
7: // pre-indexed result slice; asyncio.Semaphore → semaphore.Weighted; the
8: // streaming asyncio.Queue → an (unbuffered) chan []ReviewFinding with the
9: // producer closing on completion and the consumer ranging it.
10: 
11: import (
12: 	"context"
13: 	"fmt"
14: 	"os"
15: 	"path/filepath"
16: 	"sort"
17: 	"strings"
18: 	"sync"
19: 
20: 	"golang.org/x/sync/errgroup"
21: 	"golang.org/x/sync/semaphore"
22: 
23: 	"github.com/Agent-Field/pr-af/go/internal/config"
24: 	"github.com/Agent-Field/pr-af/go/internal/diffengine"
25: 	"github.com/Agent-Field/pr-af/go/internal/evidence"
26: 	"github.com/Agent-Field/pr-af/go/internal/gates"
27: 	"github.com/Agent-Field/pr-af/go/internal/prompts"
28: 	"github.com/Agent-Field/pr-af/go/internal/reasoners"
29: 	"github.com/Agent-Field/pr-af/go/internal/schemas"
30: )
31: 
32: // ---- Phase 1: INTAKE ----
33: 
34: func (o *Orchestrator) runIntake(ctx context.Context) (schemas.IntakeResult, error) {
35: 	if o.budgetOrTimeoutExhausted("intake") {
36: 		return schemas.IntakeResult{}, budgetExhaustedErr(o.budgetExhaustedMessage("intake"))
37: 	}
38: 
39: 	switch {
40: 	case strp(o.input.PrURL) != "":
41: 		owner, repo, number, err := o.deps.GH.ParsePRURL(strp(o.input.PrURL))
42: 		if err != nil {
43: 			return schemas.IntakeResult{}, badInput(err.Error())
44: 		}
45: 		prData, err := o.deps.GH.FetchPR(ctx, owner, repo, number)
46: 		if err != nil {
47: 			return schemas.IntakeResult{}, err
48: 		}
49: 		o.prData = &prData
50: 	case strp(o.input.DiffText) != "":
51: 		diff := strp(o.input.DiffText)
52: 		parsed := diffengine.ParseUnifiedDiff(diff)
53: 		o.prData = &schemas.GitHubPRData{
54: 			Title:        "Local Diff Review",
55: 			Diff:         diff,
56: 			ChangedFiles: toChangedFiles(parsed),
57: 		}
58: 	case strp(o.input.RepoPath) != "":
59: 		diff, err := computeRepoDiff(ctx, strp(o.input.RepoPath), strp(o.input.BaseRef), strp(o.input.HeadRef))
60: 		if err != nil {
61: 			return schemas.IntakeResult{}, err
62: 		}
63: 		parsed := diffengine.ParseUnifiedDiff(diff)
64: 		number := 0
65: 		if o.input.PostPRNumber != nil {
66: 			number = *o.input.PostPRNumber
67: 		}
68: 		o.prData = &schemas.GitHubPRData{
69: 			Number:       number,
70: 			Title:        "Local Repository Review",
71: 			Diff:         diff,
72: 			ChangedFiles: toChangedFiles(parsed),
73: 		}
74: 	default:
75: 		return schemas.IntakeResult{}, badInput("One of pr_url, diff_text, or repo_path is required")
76: 	}
77: 
78: 	resultRaw, err := o.rfns.intake(ctx, o.reasonerDeps(), reasoners.IntakeInput{
79: 		PRData: *o.prData,
80: 		Depth:  o.input.Depth,
81: 	})
82: 	if err != nil {
83: 		return schemas.IntakeResult{}, err
84: 	}
85: 	o.incInvocations(1)
86: 	o.registerCost("intake", resultRaw)
87: 	return mapToStruct[schemas.IntakeResult](resultRaw), nil
88: }
89: 
90: // ---- Phase 2: ANATOMY ----
91: 
92: func (o *Orchestrator) runAnatomy(ctx context.Context, intake schemas.IntakeResult) (schemas.AnatomyResult, error) {
93: 	if o.budgetOrTimeoutExhausted("anatomy") {
94: 		return schemas.AnatomyResult{}, budgetExhaustedErr(o.budgetExhaustedMessage("anatomy"))
95: 	}
96: 	if o.prData == nil {
97: 		return schemas.AnatomyResult{}, errPRDataNotInitialized
98: 	}
99: 	resultRaw, err := o.rfns.anatomy(ctx, o.reasonerDeps(), reasoners.AnatomyInput{
100: 		PRData:   *o.prData,
101: 		Intake:   intake,
102: 		RepoPath: strp(o.input.RepoPath),
103: 	})
104: 	if err != nil {
105: 		return schemas.AnatomyResult{}, err
106: 	}
107: 	o.incInvocations(1)
108: 	o.registerCost("anatomy", resultRaw)
109: 	return mapToStruct[schemas.AnatomyResult](resultRaw), nil
110: }
111: 
112: // ---- runReviewPhases: meta-selectors → review+layer → coverage ‖ consistency → synthesize → merge-gate ----
113: 
114: func (o *Orchestrator) runReviewPhases(
115: 	ctx context.Context,
116: 	intake schemas.IntakeResult,
117: 	anatomy schemas.AnatomyResult,
118: 	reviewDepth, reviewerFeedback string,
119: ) (schemas.ReviewPlan, []schemas.ScoredFinding, error) {
120: 	plan, err := o.runMetaSelectors(ctx, intake, anatomy, reviewDepth, reviewerFeedback)
121: 	if err != nil {
122: 		return schemas.ReviewPlan{}, nil, err
123: 	}
124: 
125: 	var allFindings []schemas.ReviewFinding
126: 	var adversaryResults []schemas.AdversaryResult
127: 
128: 	if o.config.Comments.PostWorthinessGate {
129: 		// EARLY-GATE reorder: run reviewers to completion, gate, then heavy layer.
130: 		reviewerFindings, err := o.collectParallelReview(ctx, plan, reviewerFeedback)
131: 		if err != nil {
132: 			return schemas.ReviewPlan{}, nil, err
133: 		}
134: 		kept := reviewerFindings
135: 		if len(reviewerFindings) > 1 {
136: 			if sel, ok := o.postWorthinessSelect(ctx, reviewerFindings); ok && len(sel) > 0 {
137: 				kept = sel
138: 			}
139: 		}
140: 		lq := make(chan []schemas.ReviewFinding, 1)
141: 		if len(kept) > 0 {
142: 			lq <- kept
143: 		}
144: 		close(lq)
145: 		allFindings, adversaryResults, err = o.runReviewLayer(ctx, plan, lq, anatomy)
146: 		if err != nil {
147: 			return schemas.ReviewPlan{}, nil, err
148: 		}
149: 	} else {
150: 		// Gate OFF (default): review and layer stream through one channel.
151: 		var err error
152: 		allFindings, adversaryResults, err = o.streamReviewLayer(ctx, plan, anatomy, reviewerFeedback)
153: 		if err != nil {
154: 			return schemas.ReviewPlan{}, nil, err
155: 		}
156: 	}
157: 
158: 	// Phase 6 (coverage) ‖ Phase 6.7 (consistency) — independent, run concurrently.
159: 	covFindings, covAdversary, cvFindings, err := o.runCoverageAndConsistency(ctx, plan, anatomy, allFindings, adversaryResults)
160: 	if err != nil {
161: 		return schemas.ReviewPlan{}, nil, err
162: 	}
163: 	adversaryResults = covAdversary
164: 
165: 	challenged, confirmed := 0, 0
166: 	for _, r := range adversaryResults {
167: 		if r.Verdict == "challenged" {
168: 			challenged++
169: 		}
170: 		if r.Verdict == "confirmed" {
171: 			confirmed++
172: 		}
173: 	}
174: 	o.adversaryChallengedCount = challenged
175: 	o.adversaryConfirmedCount = confirmed
176: 
177: 	var consistencyNew []schemas.ReviewFinding
178: 	for _, f := range cvFindings {
179: 		if f.DimensionID == "consistency-verify" {
180: 			consistencyNew = append(consistencyNew, f)
181: 		}
182: 	}
183: 	merged := append(append([]schemas.ReviewFinding(nil), covFindings...), consistencyNew...)
184: 
185: 	scored := o.synthesize(merged, adversaryResults)
186: 
187: 	if o.config.Comments.MergeGateEnabled {
188: 		scored = gates.ClassifyFindings(ctx, o.deps.App, scored)
189: 	}
190: 	return plan, scored, nil
191: }
192: 
193: // runCoverageAndConsistency runs the coverage loop and consistency verify
194: // concurrently (asyncio.gather(coverage_loop, consistency_verify)), preserving
195: // the ordered destructuring of the two results.
196: func (o *Orchestrator) runCoverageAndConsistency(
197: 	ctx context.Context,
198: 	plan schemas.ReviewPlan,
199: 	anatomy schemas.AnatomyResult,
200: 	allFindings []schemas.ReviewFinding,
201: 	adversaryResults []schemas.AdversaryResult,
202: ) ([]schemas.ReviewFinding, []schemas.AdversaryResult, []schemas.ReviewFinding, error) {
203: 	var covFindings []schemas.ReviewFinding
204: 	var covAdversary []schemas.AdversaryResult
205: 	var cvFindings []schemas.ReviewFinding
206: 
207: 	g, gctx := errgroup.WithContext(ctx)
208: 	g.Go(func() error {
209: 		var e error
210: 		covFindings, covAdversary, e = o.runCoverageLoop(gctx, plan, anatomy,
211: 			append([]schemas.ReviewFinding(nil), allFindings...), adversaryResults)
212: 		return e
213: 	})
214: 	g.Go(func() error {
215: 		var e error
216: 		cvFindings, e = o.runConsistencyVerify(gctx, append([]schemas.ReviewFinding(nil), allFindings...))
217: 		return e
218: 	})
219: 	if err := g.Wait(); err != nil {
220: 		return nil, nil, nil, err
221: 	}
222: 	return covFindings, covAdversary, cvFindings, nil
223: }
224: 
225: // postWorthinessSelect runs the pre-layer post-worthiness gate. Returns the kept
226: // subset and ok=false when the gate errored (Python's try/except → keep all).
227: func (o *Orchestrator) postWorthinessSelect(ctx context.Context, findings []schemas.ReviewFinding) ([]schemas.ReviewFinding, bool) {
228: 	raw, err := o.rfns.postWorthiness(ctx, o.reasonerDeps(), reasoners.PostWorthinessInput{Findings: findings})
229: 	if err != nil {
230: 		return nil, false
231: 	}
232: 	// Default keep set = all indices (pw.get("keep_indices", range(len))).
233: 	keep := make(map[int]struct{})
234: 	if v, present := raw["keep_indices"]; present {
235: 		for _, i := range getIntSlice(v) {
236: 			keep[i] = struct{}{}
237: 		}
238: 	} else {
239: 		for i := range findings {
240: 			keep[i] = struct{}{}
241: 		}
242: 	}
243: 	sel := make([]schemas.ReviewFinding, 0, len(findings))
244: 	for i, f := range findings {
245: 		if _, ok := keep[i]; ok {
246: 			sel = append(sel, f)
247: 		}
248: 	}
249: 	return sel, true
250: }
251: 
252: // ---- Phase 3: META-SELECTORS ----
253: 
254: func (o *Orchestrator) runMetaSelectors(
255: 	ctx context.Context,
256: 	intake schemas.IntakeResult,
257: 	anatomy schemas.AnatomyResult,
258: 	reviewDepth, reviewerFeedback string,
259: ) (schemas.ReviewPlan, error) {
260: 	if o.budgetOrTimeoutExhausted("meta_selectors") {
261: 		return schemas.ReviewPlan{}, budgetExhaustedErr(o.budgetExhaustedMessage("meta-selectors"))
262: 	}
263: 
264: 	lensFns := map[string]func(context.Context, reasoners.Deps, reasoners.MetaInput) (map[string]any, error){
265: 		"semantic":   o.rfns.metaSemantic,
266: 		"mechanical": o.rfns.metaMechanical,
267: 		"systemic":   o.rfns.metaSystemic,
268: 	}
269: 
270: 	// Order-preserving fan-out: pre-indexed results in enabledLenses order.
271: 	type lensJob struct {
272: 		lens string
273: 		fn   func(context.Context, reasoners.Deps, reasoners.MetaInput) (map[string]any, error)
274: 	}
275: 	jobs := make([]lensJob, 0, len(enabledLenses))
276: 	for _, lens := range enabledLenses {
277: 		if fn, ok := lensFns[lens]; ok {
278: 			jobs = append(jobs, lensJob{lens: lens, fn: fn})
279: 		}
280: 	}
281: 	results := make([]schemas.MetaDimensionResult, len(jobs))
282: 	diffPatches := reasoners.OrderedPatches(o.filePatches())
283: 
284: 	g, gctx := errgroup.WithContext(ctx)
285: 	for i := range jobs {
286: 		i := i
287: 		g.Go(func() error {
288: 			raw, err := jobs[i].fn(gctx, o.reasonerDeps(), reasoners.MetaInput{
289: 				Intake:           intake,
290: 				Anatomy:          anatomy,
291: 				Depth:            reviewDepth,
292: 				RepoPath:         strp(o.input.RepoPath),
293: 				DiffPatches:      diffPatches,
294: 				ReviewerFeedback: reviewerFeedback,
295: 			})
296: 			if err != nil {
297: 				return err
298: 			}
299: 			o.incInvocations(1)
300: 			o.registerCost("meta_selectors", raw)
301: 			results[i] = mapToStruct[schemas.MetaDimensionResult](raw)
302: 			return nil
303: 		})
304: 	}
305: 	if err := g.Wait(); err != nil {
306: 		return schemas.ReviewPlan{}, err
307: 	}
308: 
309: 	o.metaSelectorResults = results
310: 	o.effectiveDepth = o.escalateDepth(reviewDepth)
311: 
312: 	// Prefix each dimension id by its lens, preserving lens then dimension order.
313: 	var allDimensions []schemas.ReviewDimension
314: 	for _, meta := range results {
315: 		for _, dim := range meta.Dimensions {
316: 			d := dim
317: 			d.ID = meta.Lens + "_" + dim.ID
318: 			allDimensions = append(allDimensions, d)
319: 		}
320: 	}
321: 	allDimensions = dedupCrossMeta(allDimensions)
322: 
323: 	if profile, ok := config.DepthProfiles[reviewDepth]; ok && len(allDimensions) > profile.MaxDimensions {
324: 		sort.SliceStable(allDimensions, func(i, j int) bool {
325: 			return allDimensions[i].Priority > allDimensions[j].Priority
326: 		})
327: 		allDimensions = allDimensions[:profile.MaxDimensions]
328: 	}
329: 
330: 	// ReviewPlan is NOT default-seeded in schemas (design §5), so seed the
331: 	// pydantic BudgetAllocation() default for total_budget explicitly — otherwise
332: 	// plan.model_dump() (into metadata) would carry an all-zero budget.
333: 	return schemas.ReviewPlan{
334: 		Dimensions:    allDimensions,
335: 		CrossRefHints: []string{},
336: 		TotalBudget:   defaultBudgetAllocation(),
337: 	}, nil
338: }
339: 
340: // defaultBudgetAllocation reproduces pydantic BudgetAllocation()'s seeded
341: // defaults (schemas defaults.go), for structs the orchestrator constructs
342: // directly (ReviewPlan.TotalBudget).
343: func defaultBudgetAllocation() schemas.BudgetAllocation {
344: 	return schemas.BudgetAllocation{
345: 		MaxCostUSD:          0.5,
346: 		MaxDurationSeconds:  60,
347: 		MaxReferenceFollows: 3,
348: 		MaxChildSpawns:      2,
349: 	}
350: }
351: 
352: // dedupCrossMeta ports _dedup_cross_meta: dedup by sorted target-files key,
353: // keeping the higher-priority dimension.
354: func dedupCrossMeta(dimensions []schemas.ReviewDimension) []schemas.ReviewDimension {
355: 	seen := map[string]schemas.ReviewDimension{}
356: 	var deduped []schemas.ReviewDimension
357: 	for _, dim := range dimensions {
358: 		key := append([]string(nil), dim.TargetFiles...)
359: 		sort.Strings(key)
360: 		keyStr := strings.Join(key, "|")
361: 		if existing, ok := seen[keyStr]; ok {
362: 			if dim.Priority > existing.Priority {
363: 				filtered := deduped[:0]
364: 				for _, d := range deduped {
365: 					if d.ID != existing.ID {
366: 						filtered = append(filtered, d)
367: 					}
368: 				}
369: 				deduped = filtered
370: 				deduped = append(deduped, dim)
371: 				seen[keyStr] = dim
372: 			}
373: 		} else {
374: 			seen[keyStr] = dim
375: 			deduped = append(deduped, dim)
376: 		}
377: 	}
378: 	return deduped
379: }
380: 
381: // ---- Phase 4: REVIEW (streaming producer) ----
382: 
383: // collectParallelReview drives runParallelReview to completion, closing the
384: // channel and collecting every batch (the gate-ON and coverage-gap pattern where
385: // Python awaits the producer fully then drains the queue).
386: func (o *Orchestrator) collectParallelReview(ctx context.Context, plan schemas.ReviewPlan, feedback string) ([]schemas.ReviewFinding, error) {
387: 	ch := make(chan []schemas.ReviewFinding)
388: 	errc := make(chan error, 1)
389: 	stats := &dimensionParseStats{}
390: 	go func() {
391: 		err := o.runParallelReview(ctx, plan, ch, 0, feedback, stats)
392: 		close(ch)
393: 		errc <- err
394: 	}()
395: 	var out []schemas.ReviewFinding
396: 	for batch := range ch {
397: 		out = append(out, batch...)
398: 	}
399: 	if err := <-errc; err != nil {
400: 		return nil, err
```
_import/usage context:_ IMPORTS: import (
IMPORTED BY: none