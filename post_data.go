package fbintegration

import "time"

type (
	// PostData comment pending
	PostData struct {
		AverageDurationWatched                  float64            `json:"average_duration_watched"`
		Comments                                float64            `json:"comments"`
		CreatedTimestamp                        time.Time          `json:"created_timestamp"`
		CTR                                     float64            `json:"ctr"`
		CTRRate                                 float64            `json:"ctr_rate"`
		Demographic                             DemographicSplit   `json:"demographic"`
		DeepEngagementRate                      float64            `json:"deep_engagement_rate"`
		EngagementRate                          float64            `json:"engagement_rate"`
		EngagedUsers                            float64            `json:"engaged_users"`
		Impressions                             float64            `json:"impressions"`
		InlineLinkClickCTR                      float64            `json:"inline_link_click_ctr"`
		InlinePostEngagement                    float64            `json:"inline_post_engagement"`
		OverallMinutesViewed                    float64            `json:"overall_minutes_viewed"`
		PaidReach                               float64            `json:"paid_reach"`
		PeopleReached                           float64            `json:"people_reached"`
		PeopleReachedOrganic                    float64            `json:"people_reached_organic"`
		PeopleReachedPaid                       float64            `json:"people_reached_paid"`
		PercentViewsCompleted                   float64            `json:"percent_views_completed"`
		PostConsumptions                        float64            `json:"post_consumptions"`
		PostEngagements                         float64            `json:"post_engagements"`
		Reactions                               float64            `json:"reactions"`
		ReactionsBreakdown                      map[string]float64 `json:"reactions_breakdown"`
		Result                                  float64            `json:"result"`
		ResultRate                              float64            `json:"result_rate"`
		ResultName                              string             `json:"result_name"`
		SampledPeopleReached                    float64            `json:"-"`
		Shares                                  float64            `json:"shares"`
		Spend                                   float64            `json:"spend"`
		Targeting                               AdTargeting        `json:"ad_targeting"`
		UniqueViewers                           float64            `json:"unique_viewers"`
		ViewRate                                float64            `json:"view_rate"`
		ViewsToNinetyFivePercentComplete        float64            `json:"views_to_ninety_five_percent_complete"`
		ViewsToNinetyFivePercentCompleteOrganic float64            `json:"views_to_ninety_five_percent_complete_organic"`
		ViewsToNinetyFivePercentCompletePaid    float64            `json:"views_to_ninety_five_percent_complete_paid"`
		VideoViewCost                           float64            `json:"video_view_cost"`
		VideoViews                              float64            `json:"video_views"`
		VideoViewsOrganic                       float64            `json:"video_views_organic"`
		VideoViewsPaid                          float64            `json:"video_views_paid"`
	}
)
