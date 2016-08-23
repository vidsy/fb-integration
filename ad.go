package fbintegration

import (
	"fmt"
	facebookLib "github.com/huandu/facebook"
)

type (
	// Ad comment pending
	Ad struct {
		ID       string
		Creative *Creative
		Post     *Post
	}
)

// NewAd comment pending
func NewAd(result *facebookLib.Result) Ad {
	var id string
	var creativeID string

	result.DecodeField("id", &id)
	result.DecodeField("creative.id", &creativeID)

	ad := Ad{
		id,
		&Creative{creativeID, "", ""},
		&Post{},
	}

	return ad
}

// CreateBatchParams comment pending
func (a *Ad) CreateBatchParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": a.Creative.ID,
		"fields":       "object_id,object_type,effective_object_story_id",
	}
}

// CreateInsightParams comment pending
func (a *Ad) CreateInsightParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/insights?fields=unique_actions,reach,spend&date_preset=lifetime", a.ID),
	}
}

//CreateBreakdownInsightParams comment pending
func (a *Ad) CreateBreakdownInsightParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/insights?fields=reach&date_preset=lifetime&breakdowns=age,gender", a.ID),
	}
}
