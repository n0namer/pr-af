package node

import "github.com/Agent-Field/pr-af/go/internal/schemas"

func formatReviewResult(format string, result schemas.ReviewResult) any {
	switch format {
	case "markdown":
		return result.Review.Body
	case "sarif":
		return reviewResultSARIF(result)
	default:
		return result
	}
}

func reviewResultSARIF(result schemas.ReviewResult) map[string]any {
	results := make([]any, 0, len(result.Findings))
	for _, f := range result.Findings {
		end := f.LineEnd
		if end <= 0 {
			end = f.LineStart
		}
		results = append(results, map[string]any{
			"ruleId": "pr-af." + string(f.Severity),
			"level": sarifLevel(f.Severity),
			"message": map[string]any{"text": f.Title + ": " + f.Body},
			"locations": []any{map[string]any{
				"physicalLocation": map[string]any{
					"artifactLocation": map[string]any{"uri": f.FilePath},
					"region": map[string]any{"startLine": f.LineStart, "endLine": end},
				},
			}},
		})
	}
	return map[string]any{
		"version": "2.1.0",
		"$schema": "https://json.schemastore.org/sarif-2.1.0.json",
		"runs": []any{map[string]any{
			"tool": map[string]any{"driver": map[string]any{"name": "PR-AF"}},
			"results": results,
		}},
	}
}

func sarifLevel(severity schemas.Severity) string {
	switch severity {
	case "critical", "important":
		return "error"
	case "suggestion":
		return "warning"
	default:
		return "note"
	}
}
