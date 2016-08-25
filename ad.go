package fbintegration

import (
	"fmt"
	facebookLib "github.com/huandu/facebook"
)

type (
	// Ad comment pending
	Ad struct {
		ID       string
		AdsetID  string
		Creative *Creative
		Post     *Post
	}
)

// NewAd comment pending
func NewAd(result *facebookLib.Result) Ad {
	var id string
	var creativeID string
	var adsetID string

	result.DecodeField("id", &id)
	result.DecodeField("creative.id", &creativeID)
	result.DecodeField("adset.id", &adsetID)

	ad := Ad{
		id,
		adsetID,
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

// CreateBreakdownInsightParams comment pending
func (a *Ad) CreateBreakdownInsightParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/insights?fields=reach&date_preset=lifetime&breakdowns=age,gender", a.ID),
	}
}

// CreateTargetingParams comment pending
func (a *Ad) CreateTargetingParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s?fields=targeting&date_preset=lifetime", a.AdsetID),
	}
}
