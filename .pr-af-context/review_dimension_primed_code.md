### go/internal/schemas/defaults.go
```
1: package schemas
2: 
3: import "encoding/json"
4: 
5: // This file implements non-zero-default seeding for every struct that has at
6: // least one field whose pydantic default is not the Go zero value (design §C.1
7: // seeding table). Go's json.Unmarshal leaves an absent key at the Go zero
8: // value, whereas pydantic fills the declared default; each struct below gets a
9: // defaultXxx() constructor (also usable as the deterministic fallback a role
10: // returns on a harness parse-failure) plus an UnmarshalJSON that seeds the
11: // default before decoding — so an absent key keeps the default while a present
12: // key (even false/0/"") overrides it, matching pydantic exactly.
13: //
14: // The `type alias X` trick strips X's methods (including UnmarshalJSON) so the
15: // inner json.Unmarshal does not recurse; nested field types (e.g. Severity,
16: // BudgetAllocation) keep their own UnmarshalJSON.
17: //
18: // Where a struct in this table has a `default_factory=list` field, the
19: // constructor seeds it to a non-nil empty slice so that unmarshaling "{}" round
20: // trips to `[]` (never null), honoring the design's "empty slices marshal as []
21: // where Python default is a list" rule for the structs whose constructors this
22: // package owns.
23: 
24: // --- input.go ---
25: 
26: func defaultReviewInput() ReviewInput {
27: 	return ReviewInput{
28: 		Depth:          "auto",
29: 		Focus:          "auto",
30: 		OutputFormat:   "github",
31: 		SuggestionMode: "comment",
32: 		MaxReviewDepth: 2,
33: 		IgnorePaths:    []string{},
34: 		Hints:          []string{},
35: 	}
36: }
37: 
38: // UnmarshalJSON seeds ReviewInput's non-zero pydantic defaults. The budget-cap
39: // pointers (MaxCostUSD / MaxDurationSeconds) stay nil — they are resolved later
40: // by config.ResolveBudgetCaps / ReviewConfig.FromInput.
41: func (r *ReviewInput) UnmarshalJSON(b []byte) error {
42: 	*r = defaultReviewInput()
43: 	type alias ReviewInput
44: 	return json.Unmarshal(b, (*alias)(r))
45: }
46: 
47: // --- pipeline.go ---
48: 
49: func defaultBudgetAllocation() BudgetAllocation {
50: 	return BudgetAllocation{
51: 		MaxCostUSD:          0.5,
52: 		MaxDurationSeconds:  60,
53: 		MaxReferenceFollows: 3,
54: 		MaxChildSpawns:      2,
55: 	}
56: }
57: 
58: // UnmarshalJSON seeds BudgetAllocation defaults (0.5 / 60 / 3 / 2).
59: func (a *BudgetAllocation) UnmarshalJSON(b []byte) error {
60: 	*a = defaultBudgetAllocation()
61: 	type alias BudgetAllocation
62: 	return json.Unmarshal(b, (*alias)(a))
63: }
64: 
65: func defaultReviewDimension() ReviewDimension {
66: 	return ReviewDimension{
67: 		ContextFiles: []string{},
68: 		Priority:     1,
69: 		Budget:       defaultBudgetAllocation(),
70: 	}
71: }
72: 
73: // UnmarshalJSON seeds ReviewDimension.Priority=1 and a default Budget. When a
74: // "budget" key is present, BudgetAllocation's own UnmarshalJSON re-seeds it
75: // before applying the overrides.
76: func (d *ReviewDimension) UnmarshalJSON(b []byte) error {
77: 	*d = defaultReviewDimension()
78: 	type alias ReviewDimension
79: 	return json.Unmarshal(b, (*alias)(d))
80: }
81: 
82: func defaultSubReviewRequest() SubReviewRequest {
83: 	return SubReviewRequest{
84: 		ContextFiles: []string{},
85: 		Priority:     1,
86: 	}
87: }
88: 
89: // UnmarshalJSON seeds SubReviewRequest.Priority=1.
90: func (s *SubReviewRequest) UnmarshalJSON(b []byte) error {
91: 	*s = defaultSubReviewRequest()
92: 	type alias SubReviewRequest
93: 	return json.Unmarshal(b, (*alias)(s))
94: }
95: 
96: func defaultReviewFinding() ReviewFinding {
97: 	return ReviewFinding{
98: 		Severity:   DefaultSeverity,
99: 		Confidence: 0.5,
100: 		Tags:       []string{},
101: 	}
102: }
103: 
104: // UnmarshalJSON seeds ReviewFinding.Confidence=0.5, Severity="suggestion".
105: func (f *ReviewFinding) UnmarshalJSON(b []byte) error {
106: 	*f = defaultReviewFinding()
107: 	type alias ReviewFinding
108: 	return json.Unmarshal(b, (*alias)(f))
109: }
110: 
111: func defaultAdversaryResult() AdversaryResult {
112: 	return AdversaryResult{SeverityAdjustment: "none"}
113: }
114: 
115: // UnmarshalJSON seeds AdversaryResult.SeverityAdjustment="none".
116: func (a *AdversaryResult) UnmarshalJSON(b []byte) error {
117: 	*a = defaultAdversaryResult()
118: 	type alias AdversaryResult
119: 	return json.Unmarshal(b, (*alias)(a))
120: }
121: 
122: func defaultMetaDimensionResult() MetaDimensionResult {
123: 	return MetaDimensionResult{Confidence: 0.7}
124: }
125: 
126: // UnmarshalJSON seeds MetaDimensionResult.Confidence=0.7. Dimensions is a
127: // required field in Python (no default), so it is not seeded.
128: func (m *MetaDimensionResult) UnmarshalJSON(b []byte) error {
129: 	*m = defaultMetaDimensionResult()
130: 	type alias MetaDimensionResult
131: 	return json.Unmarshal(b, (*alias)(m))
132: }
133: 
134: // --- gates.go ---
135: 
136: func defaultCoverageGate() CoverageGate {
137: 	return CoverageGate{
138: 		GapDescriptions: []string{},
139: 		Confident:       true,
140: 	}
141: }
142: 
143: // UnmarshalJSON seeds CoverageGate.Confident=true.
144: func (c *CoverageGate) UnmarshalJSON(b []byte) error {
145: 	*c = defaultCoverageGate()
146: 	type alias CoverageGate
147: 	return json.Unmarshal(b, (*alias)(c))
148: }
149: 
150: // --- output.go ---
151: 
152: func defaultScoredFinding() ScoredFinding {
153: 	return ScoredFinding{
154: 		DiffSide:          "RIGHT",
155: 		Severity:          DefaultSeverity,
156: 		Confidence:        0.5,
157: 		Tags:              []string{},
158: 		ActiveMultipliers: []string{},
159: 	}
160: }
161: 
162: // UnmarshalJSON seeds ScoredFinding.DiffSide="RIGHT", Confidence=0.5,
163: // Severity="suggestion".
164: func (s *ScoredFinding) UnmarshalJSON(b []byte) error {
165: 	*s = defaultScoredFinding()
166: 	type alias ScoredFinding
167: 	return json.Unmarshal(b, (*alias)(s))
168: }
169: 
170: func defaultGitHubComment() GitHubComment {
171: 	return GitHubComment{Side: "RIGHT"}
172: }
173: 
174: // UnmarshalJSON seeds GitHubComment.Side="RIGHT".
175: func (c *GitHubComment) UnmarshalJSON(b []byte) error {
176: 	*c = defaultGitHubComment()
177: 	type alias GitHubComment
178: 	return json.Unmarshal(b, (*alias)(c))
179: }
```
_import/usage context:_ IMPORTS: import "encoding/json"
IMPORTED BY: none