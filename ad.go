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
func (a *Ad) CreateBatchParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s?fields=object_id,object_type,effective_object_story_id", a.Creative.ID))
}

// CreateInsightParams comment pending
func (a *Ad) CreateInsightParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/insights?fields=unique_actions,reach,spend&date_preset=lifetime", a.ID))
}

// CreateBreakdownInsightParams comment pending
func (a *Ad) CreateBreakdownInsightParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/insights?fields=reach&date_preset=lifetime&breakdowns=age,gender", a.ID))
}

// CreateTargetingParams comment pending
func (a *Ad) CreateTargetingParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s?fields=targeting&date_preset=lifetime", a.AdsetID))
}
