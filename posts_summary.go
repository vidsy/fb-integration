package fbintegration

import "encoding/json"

type (
	// PostsSummary comment pending
	PostsSummary struct {
		TotalImpressions                   float64            `json:"total_impressions"`
		TotalPaidImpressions               float64            `json:"total_paid_impressions"`
		TotalOrganicImpressions            float64            `json:"total_organic_impressions"`
		TotalReach                         float64            `json:"total_reach"`
		TotalPaidReach                     float64            `json:"total_paid_reach"`
		TotalOrganicReach                  float64            `json:"total_organic_reach"`
		TotalVideoViews                    float64            `json:"total_video_views"`
		TotalPaidVideoViews                float64            `json:"total_paid_video_views"`
		TotalOrganicVideoViews             float64            `json:"total_organic_video_views"`
		TotalUniqueVideoViews              float64            `json:"total_unique_video_views"`
		TotalMinutesViewed                 float64            `json:"total_minutes_viewed"`
		TotalVideosUsed                    int                `json:"total_videos_used"`
		ReactionsTotal                     float64            `json:"reactions_total"`
		Reactions                          map[string]float64 `json:"reactions"`
		TotalClicks                        float64            `json:"total_clicks"`
		TotalUniquePeopleEngaged           float64            `json:"total_unique_people_engaged"`
		TotalActions                       float64            `json:"total_actions"`
		TotalEngagementPercentPeopleViewed float64            `json:"total_engagement_percent_people_viewed"`
		TotalViewRate                      float64            `json:"total_view_rate"`
		TopVideoID                         string             `json:"top_video_id"`
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
		ps.TotalEngagementPercentPeopleViewed += post.Data.EngagementPercentPeopleViewed
		ps.TotalViewRate += post.Data.ViewRate

		for _, reactionType := range post.ReactionTypes() {
			ps.Reactions[reactionType] += post.Data.Reactions[reactionType]
		}
	}

	ps.TotalEngagementPercentPeopleViewed = calculateAverage(ps.TotalEngagementPercentPeopleViewed, len(posts))
	ps.TotalViewRate = calculateAverage(ps.TotalViewRate, len(posts))

	ps.TotalVideosUsed = calculateVideosUsed(posts)
	ps.TopVideoID = findTopVideoFromEngagementPercentPeopleViewedVideo(posts)

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

func calculateAverage(total float64, divisor int) float64 {
	return total / float64(divisor)
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

func findTopVideoFromEngagementPercentPeopleViewedVideo(posts []*Post) string {
	topPost := posts[0]

	for _, post := range posts {
		if post.Data.EngagementPercentPeopleViewed > topPost.Data.EngagementPercentPeopleViewed {
			topPost = post
		}
	}

	return topPost.ObjectID

}
