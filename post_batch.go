package fbintegration

import facebookLib "github.com/huandu/facebook"

type (
	// PostBatch comment pending
	PostBatch struct {
		Posts []*Post
	}
)

// GeneratePostBatchSlices comment pending
func GeneratePostBatchSlices(posts []*Post, size int) []PostBatch {
	var postBatches []PostBatch

	amount := (len(posts) / size) + 1
	start := 0

	var end int
	if len(posts) < size {
		end = len(posts)
	} else {
		end = size
	}

	for i := 0; i < amount; i++ {
		batch := PostBatch{
			posts[start:end],
		}

		postBatches = append(postBatches, batch)
		start += size
		end += size
	}

	return postBatches
}

// EngagementParams comment pending
func (p PostBatch) EngagementParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(p.Posts); i++ {
		params = append(params, p.Posts[i].GenerateCommentsParams())
		params = append(params, p.Posts[i].GenerateSharesParams())
	}

	return params
}

// InsightParams comment pending
func (p PostBatch) InsightParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(p.Posts); i++ {
		params = append(params, p.Posts[i].GenerateInsightParams())
	}

	return params
}

// ReactionBreakdownParams comment pending
func (p PostBatch) ReactionBreakdownParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(p.Posts); i++ {
		for _, p := range p.Posts[i].GenerateReactionBreakdownParams() {
			params = append(params, p)
		}
	}

	return params
}

// TotalReactionsParams comment pending
func (p PostBatch) TotalReactionsParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(p.Posts); i++ {
		params = append(params, p.Posts[i].GenerateTotalReactionsParams())
	}

	return params
}

// TotalAdInsightsParams comment pending
func (p PostBatch) TotalAdInsightsParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(p.Posts); i++ {
		ad := Ad{ID: p.Posts[i].AdID}
		params = append(params, ad.CreateInsightParams())
	}

	return params
}

// TotalAdInsightsBreakDownsParams comment pending
func (p PostBatch) TotalAdInsightsBreakDownParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(p.Posts); i++ {
		ad := Ad{ID: p.Posts[i].AdID}
		params = append(params, ad.CreateBreakdownInsightParams())
	}

	return params
}
