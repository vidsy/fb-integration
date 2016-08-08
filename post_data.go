package fbintegration

type (
	// PostData comment pending
	PostData struct {
		Impressions        int            `json:"impressions"`
		PaidImpressions    int            `json:"paid_impressions"`
		OrganicImpressions int            `json:"organic_impressions"`
		Reach              int            `json:"reach"`
		PaidReach          int            `json:"paid_reach"`
		OrganicReach       int            `json:"organic_reach"`
		VideoViews         int            `json:"video_views"`
		PaidVideoViews     int            `json:"paid_video_views"`
		OrganicVideoViews  int            `json:"organic_video_views"`
		UniqueVideoViews   int            `json:"unique_video_views"`
		MinutesViewed      int            `json:"minutes_viewed"`
		AverageDuration    int            `json:"average_duration"`
		ReactionsTotal     int            `json:"reactions_total"`
		Reactions          map[string]int `json:"reactions"`
		Comments           int            `json:"comments"`
		Shares             int            `json:"shares"`
	}
)
