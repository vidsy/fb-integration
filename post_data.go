package fbintegration

type (
	// PostData comment pending
	PostData struct {
		PeopleReached        float64            `json:"people_reached"`
		PeopleReachedOrganic float64            `json:"people_reached_organic"`
		PeopleReachedPaid    float64            `json:"people_reached_paid"`
		EngagementRate       float64            `json:"engagement_rate"`
		EngagedUsers         float64            `json:"engaged_users"`
		PostConsumptions     float64            `json:"post_consumptions"`
		PostEngagements      float64            `json:"post_engagements"`
		Reactions            float64            `json:"reactions"`
		Comments             float64            `json:"comments"`
		Shares               float64            `json:"shares"`
		ReactionsBreakdown   map[string]float64 `json:"reactions"`
		Demographic          DemographicSplit   `json:"demographic"`
		PaidReach            float64            `json:"paid_reach"`
		//Targeting            Targeting          `json:"targeting"`
		VideoViews                              float64 `json:"video_views"`
		VideoViewsOrganic                       float64 `json:"video_views_organic"`
		VideoViewsPaid                          float64 `json:"video_views_paid"`
		ViewRate                                float64 `json:"view_rate"`
		UniqueViewers                           float64 `json:"unique_viewers"`
		ViewsToNinetyFivePercentComplete        float64 `json:"views_to_ninety_five_percent_complete"`
		ViewsToNinetyFivePercentCompleteOrganic float64 `json:"views_to_ninety_five_percent_complete_organic"`
		ViewsToNinetyFivePercentCompletePaid    float64 `json:"views_to_ninety_five_percent_complete_paid"`
		PercentViewsCompleted                   float64 `json:"percent_views_completed"`
		AverageDurationWatched                  float64 `json:"average_duration_watched"`
		OverMinutesViewed                       float64 `json:"overall_minutes_viewed"`

		Impressions            float64            `json:"impressions"`
		PaidImpressions        float64            `json:"paid_impressions"`
		OrganicImpressions     float64            `json:"organic_impressions"`
		Reach                  float64            `json:"reach"`
		SampledReach           float64            `json:"sampled_reach"`
		PaidReach              float64            `json:"paid_reach"`
		OrganicReach           float64            `json:"organic_reach"`
		VideoViews             float64            `json:"video_views"`
		VideoCompletion        float64            `json:"video_completion"`
		VideoCompletionPercent float64            `json:"video_completion_percent"`
		OrganicVideoCompletion float64            `json:"organic_video_completion"`
		PaidVideoCompletion    float64            `json:"paid_video_completion"`
		PaidVideoViews         float64            `json:"paid_video_views"`
		OrganicVideoViews      float64            `json:"organic_video_views"`
		UniqueVideoViews       float64            `json:"unique_video_views"`
		MinutesViewed          float64            `json:"minutes_viewed"`
		AverageDuration        float64            `json:"average_duration"`
		ReactionsTotal         float64            `json:"reactions_total"`
		Reactions              map[string]float64 `json:"reactions"`
		Comments               float64            `json:"comments"`
		Shares                 float64            `json:"shares"`
		Clicks                 float64            `json:"clicks"`
		Spend                  float64            `json:"spend"`
		VideoViewCost          float64            `json:"video_view_cost"`
		Actions                float64            `json:"actions"`
		EngagementRate         float64            `json:"engagement_rate"`
		ViewRate               float64            `json:"view_rate"`
		AudienceSplit          AudienceSplit      `json:"audience_split"`
		LifetimeEngagedUsers   float64            `json:"lifetime_engaged_users"`
	}
)
