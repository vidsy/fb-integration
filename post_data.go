package fbintegration

type (
	// PostData comment pending
	PostData struct {
		Impressions                   float64            `json:"impressions"`
		PaidImpressions               float64            `json:"paid_impressions"`
		OrganicImpressions            float64            `json:"organic_impressions"`
		Reach                         float64            `json:"reach"`
		PaidReach                     float64            `json:"paid_reach"`
		OrganicReach                  float64            `json:"organic_reach"`
		VideoViews                    float64            `json:"video_views"`
		PaidVideoViews                float64            `json:"paid_video_views"`
		OrganicVideoViews             float64            `json:"organic_video_views"`
		UniqueVideoViews              float64            `json:"unique_video_views"`
		MinutesViewed                 float64            `json:"minutes_viewed"`
		AverageDuration               float64            `json:"average_duration"`
		ReactionsTotal                float64            `json:"reactions_total"`
		Reactions                     map[string]float64 `json:"reactions"`
		Comments                      float64            `json:"comments"`
		Shares                        float64            `json:"shares"`
		Clicks                        float64            `json:"clicks"`
		UniquePeopleEngaged           float64            `json:"unique_people_engaged"`
		Spend                         float64            `json:"spend"`
		VideoViewCost                 float64            `json:"video_view_cost"`
		Actions                       float64            `json:"actions"`
		EngagementPercentPeopleViewed float64            `json:"engagement_percent_people_viewed"`
		ViewRate                      float64            `json:"view_rate"`
		AudienceSplit                 AudienceSplit      `json:"audience_split"`
	}
)
