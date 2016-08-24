package fbintegration

import "encoding/json"

type (
	// PostsSummary comment pending
	PostsSummary struct {
		VideosPosted      int     `json:"videos_posted"`
		CampaignReach     float64 `json:"campaign_reach"`
		CampaignViews     float64 `json:"campaign_views"`
		MinutesViewed     float64 `json:"minutes_viewed"`
		UnqiueViewers     float64 `json:"unqiue_viewers"`
		OverallViewRate   float64 `json:"overall_view_rate"`
		TotalEngagements  float64 `json:"total_engagements"`
		EngagementRate    float64 `json:"engagement_rate"`
		TopViewedVideoID  string  `json:"top_viewed_video_id"`
		TopEngagedVideoID string  `json:"top_engaged_video_id"`
	}
)

// NewPostsSummary comment pending
func NewPostsSummary(posts []*Post) PostsSummary {
	var ps PostsSummary

	ps.Reactions = make(map[string]float64)

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
		ps.TotalActions += post.Data.Actions

		for _, reactionType := range post.ReactionTypes() {
			ps.Reactions[reactionType] += post.Data.Reactions[reactionType]
		}
	}

	ps.EngagementRate = calculateEngagementRate(ps.TotalActions, ps.TotalReach)
	ps.TotalViewRate = calculateViewRate(ps.TotalReach, ps.TotalUniqueVideoViews)

	ps.TotalVideosUsed = calculateVideosUsed(posts)
	ps.TopVideoID = findTopVideoFromEngagementRate(posts)

	return ps
}

// calculateViewRate comment pending
func calculateViewRate(totalReach float64, totalUniqueVideoViews float64) float64 {
	return (totalUniqueVideoViews / totalReach) * 100
}

// ToJSON comment pending
func (p *PostsSummary) ToJSON() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func calculateAverage(total float64, divisor int) float64 {
	return total / float64(divisor)
}

func calculateVideosUsed(posts []*Post) int {
	videosUsed := make(map[string]*interface{})

	for _, post := range posts {
		if _, exists := videosUsed[post.ObjectID]; !exists {
			videosUsed[post.ObjectID] = nil
		}
	}

	return len(videosUsed)
}

func findTopVideoFromEngagementRate(posts []*Post) string {
	topPost := posts[0]

	for _, post := range posts {
		if post.Data.EngagementRate > topPost.Data.EngagementRate {
			topPost = post
		}
	}

	return topPost.ObjectID

}
