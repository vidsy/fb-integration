package fbintegration

import (
	"fmt"

	facebookLib "github.com/huandu/facebook"
)

type (
	// Post comment pending
	Post struct {
		ID       string `facebook:"id"        json:"post_id"`
		ObjectID string `facebook:"object_id" json:"object_id"`
	}
)

// GenerateInsightParams comment pending
func (p Post) GenerateInsightParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/insights/post_impressions,post_impressions_paid,post_impressions_unique,post_impressions_paid_unique,post_video_views_paid,post_video_views_organic,post_video_views_organic_unique,post_video_views_paid_unique,post_video_view_time,post_video_avg_time_watched?period=lifetime&limit=20", p.ID),
	}
}
