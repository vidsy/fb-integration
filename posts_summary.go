package fbintegration

import "encoding/json"

type (
	// PostsSummary comment pending
	PostsSummary struct {
		TotalImpressions         int            `json:"total_impressions"`
		TotalPaidImpressions     int            `json:"total_paid_impressions"`
		TotalOrganicImpressions  int            `json:"total_organic_impressions"`
		TotalReach               int            `json:"total_reach"`
		TotalPaidReach           int            `json:"total_paid_reach"`
		TotalOrganicReach        int            `json:"total_organic_reach"`
		TotalVideoViews          int            `json:"total_video_views"`
		TotalPaidVideoViews      int            `json:"total_paid_video_views"`
		TotalOrganicVideoViews   int            `json:"total_organic_video_views"`
		TotalUniqueVideoViews    int            `json:"total_unique_video_views"`
		TotalMinutesViewed       int            `json:"total_minutes_viewed"`
		TotalVideosUsed          int            `json:"total_videos_used"`
		ReactionsTotal           int            `json:"reactions_total"`
		Reactions                map[string]int `json:"reactions"`
		TotalClicks              int            `json:"total_clicks"`
		TotalUniquePeopleEngaged int            `json:"total_unique_people_engaged"`
		TotalSpend               int            `json:"total_spend"`
	}
)

// NewPostsSummary comment pending
func NewPostsSummary(posts []*Post) PostsSummary {
	var ps PostsSummary

	ps.Reactions = make(map[string]int)

	for _, post := range posts {
		ps.TotalImpressions += post.Data.Impressions
		ps.TotalPaidImpressions += post.Data.PaidImpressions
		ps.TotalOrganicImpressions += post.Data.OrganicImpressions
		ps.TotalReach += post.Data.Reach
		ps.TotalPaidReach += post.Data.PaidReach
		ps.TotalOrganicReach += post.Data.OrganicReach
		ps.TotalVideoViews += post.Data.VideoViews
		ps.TotalPaidVideoViews += post.Data.PaidVideoViews
		ps.TotalOrganicVideoViews += post.Data.OrganicVideoViews
		ps.TotalUniqueVideoViews += post.Data.UniqueVideoViews
		ps.TotalMinutesViewed += post.Data.MinutesViewed
		ps.ReactionsTotal += post.Data.ReactionsTotal

		for _, reactionType := range post.ReactionTypes() {
			ps.Reactions[reactionType] += post.Data.Reactions[reactionType]
		}
	}

	ps.TotalVideosUsed = calculateVideosUsed(posts)

	return ps
}

// ToJSON comment pending
func (p *PostsSummary) ToJSON() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func calculateVideosUsed(posts []*Post) int {
	videosUsed := make(map[string]*Post)

	for _, post := range posts {
		if _, exists := videosUsed[post.ObjectID]; !exists {
			videosUsed[post.ObjectID] = post
		}
	}

	return len(videosUsed)
}
