package fbintegration

import (
	facebookLib "github.com/huandu/facebook"
)

type (
	// PostResults comment pending
	PostResults struct {
		Insights          *facebookLib.Result
		ReactionBreakdown []*facebookLib.Result
		TotalReactions    *facebookLib.Result
	}
)
