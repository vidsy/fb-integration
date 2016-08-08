package fbintegration

import (
	facebookLib "github.com/huandu/facebook"
)

type (
	// PostResults comment pending
	PostResults struct {
		Engagement        []*facebookLib.Result
		Insights          *facebookLib.Result
		ReactionBreakdown []*facebookLib.Result
		TotalReactions    *facebookLib.Result
	}
)
