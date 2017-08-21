package fbintegration

import (
	"fmt"

	"strings"

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

	return Ad{
		ID:       id,
		AdsetID:  adsetID,
		Creative: &Creative{ID: creativeID},
		Post:     &Post{},
	}
}

// CreateBatchParams comment pending
func (a *Ad) CreateBatchParams() BatchParams {
	fields := []string{
		"effective_object_story_id",
		"object_id",
		"object_type",
	}

	uri := fmt.Sprintf(
		"%s?fields=%s",
		a.Creative.ID,
		strings.Join(fields, ","),
	)

	return NewBatchParams(uri)
}

// CreateInsightParams comment pending
func (a *Ad) CreateInsightParams() BatchParams {
	fields := []string{
		"objective",
		"estimated_ad_recallers",
		"inline_link_clicks",
		"inline_post_engagement",
		"actions",
		"ctr",
		"clicks",
		"impressions",
		"inline_link_click_ctr",
		"reach",
		"spend",
		"total_actions",
		"total_unique_actions",
		"unique_actions",
		"video_p95_watched_actions",
	}

	uri := fmt.Sprintf(
		"%s/insights?fields=%s&date_preset=lifetime",
		a.ID,
		strings.Join(fields, ","),
	)

	return NewBatchParams(uri)
}

// CreateBreakdownInsightParams comment pending
func (a *Ad) CreateBreakdownInsightParams() BatchParams {
	fields := []string{
		"reach",
	}

	breakdowns := []string{
		"age",
		"gender",
	}

	uri := fmt.Sprintf(
		"%s/insights?fields=%s&date_preset=lifetime&breakdowns=%s",
		a.ID,
		strings.Join(fields, ","),
		strings.Join(breakdowns, ","),
	)

	return NewBatchParams(uri)
}

// CreateTargetingParams comment pending
func (a *Ad) CreateTargetingParams() BatchParams {
	fields := []string{
		"targeting",
	}

	uri := fmt.Sprintf(
		"%s?fields=%s&date_preset=lifetime",
		a.AdsetID,
		strings.Join(fields, ","),
	)

	return NewBatchParams(uri)
}
