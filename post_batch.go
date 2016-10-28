package fbintegration

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
		if (end + size) > len(posts) {
			end += len(posts) - end
		} else {
			end += size
		}
	}

	return postBatches
}

// CommentParams returns a slice of BatchParams of comment requests
// for each post.
func (p PostBatch) CommentParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(p.Posts); i++ {
		params = append(params, p.Posts[i].GenerateCommentsParams())
	}

	return params
}

// TargetingParams comment pending
func (p PostBatch) TargetingParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(p.Posts); i++ {
		ad := Ad{ID: p.Posts[i].AdID, AdsetID: p.Posts[i].AdsetID}
		params = append(params, ad.CreateTargetingParams())
	}

	return params
}

// InsightParams comment pending
func (p PostBatch) InsightParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(p.Posts); i++ {
		params = append(params, p.Posts[i].GenerateInsightParams())
	}

	return params
}

// ReactionBreakdownParams comment pending
func (p PostBatch) ReactionBreakdownParams() map[string][]BatchParams {
	postParams := make(map[string][]BatchParams, len(p.Posts))

	for i := 0; i < len(p.Posts); i++ {
		postParams[p.Posts[i].ID] = p.Posts[i].GenerateReactionBreakdownParams()
	}

	return postParams
}

// TotalReactionsParams comment pending
func (p PostBatch) TotalReactionsParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(p.Posts); i++ {
		params = append(params, p.Posts[i].GenerateTotalReactionsParams())
	}

	return params
}

// TotalAdInsightsParams comment pending
func (p PostBatch) TotalAdInsightsParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(p.Posts); i++ {
		ad := Ad{ID: p.Posts[i].AdID}
		params = append(params, ad.CreateInsightParams())
	}

	return params
}

// TotalAdInsightsBreakDownParams comment pending
func (p PostBatch) TotalAdInsightsBreakDownParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(p.Posts); i++ {
		ad := Ad{ID: p.Posts[i].AdID}
		params = append(params, ad.CreateBreakdownInsightParams())
	}

	return params
}

// PostDataParams return a slice of BatchParams for post realted
// fields.
func (p PostBatch) PostDataParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(p.Posts); i++ {
		params = append(params, p.Posts[i].GeneratePostParams())
	}

	return params
}
