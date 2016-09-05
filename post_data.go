package fbintegration

import "time"

type (
	// PostData comment pending
	PostData struct {
		SampledPeopleReached                    float64            `json:"-"`
		PeopleReached                           float64            `json:"people_reached"`
		PeopleReachedOrganic                    float64            `json:"people_reached_organic"`
		PeopleReachedPaid                       float64            `json:"people_reached_paid"`
		EngagementRate                          float64            `json:"engagement_rate"`
		EngagedUsers                            float64            `json:"engaged_users"`
		PostConsumptions                        float64            `json:"post_consumptions"`
		PostEngagements                         float64            `json:"post_engagements"`
		Reactions                               float64            `json:"reactions"`
		ReactionsBreakdown                      map[string]float64 `json:"reactions_breakdown"`
		Comments                                float64            `json:"comments"`
		Shares                                  float64            `json:"shares"`
		Demographic                             DemographicSplit   `json:"demographic"`
		Targeting                               AdTargeting        `json:"ad_targeting"`
		PaidReach                               float64            `json:"paid_reach"`
		VideoViews                              float64            `json:"video_views"`
		VideoViewsOrganic                       float64            `json:"video_views_organic"`
		VideoViewsPaid                          float64            `json:"video_views_paid"`
		ViewRate                                float64            `json:"view_rate"`
		UniqueViewers                           float64            `json:"unique_viewers"`
		ViewsToNinetyFivePercentComplete        float64            `json:"views_to_ninety_five_percent_complete"`
		ViewsToNinetyFivePercentCompleteOrganic float64            `json:"views_to_ninety_five_percent_complete_organic"`
		ViewsToNinetyFivePercentCompletePaid    float64            `json:"views_to_ninety_five_percent_complete_paid"`
		PercentViewsCompleted                   float64            `json:"percent_views_completed"`
		AverageDurationWatched                  float64            `json:"average_duration_watched"`
		OverallMinutesViewed                    float64            `json:"overall_minutes_viewed"`
		Spend                                   float64            `json:"spend"`
		VideoViewCost                           float64            `json:"video_view_cost"`
		CreatedTimestamp                        time.Time          `json:"created_timestamp"`
	}
)
