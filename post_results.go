package fbintegration

import (
	facebookLib "github.com/huandu/facebook"
)

type (
	// PostResults comment pending
	PostResults struct {
		Targeting           *facebookLib.Result
		Engagement          []*facebookLib.Result
		Insights            *facebookLib.Result
		AdInsights          *facebookLib.Result
		AdBreakdownInsights *facebookLib.Result
		ReactionBreakdown   []*facebookLib.Result
		TotalReactions      *facebookLib.Result
	}
)
