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
		TotalEngagement   float64 `json:"total_engagement"`
		EngagementRate    float64 `json:"engagement_rate"`
		TopViewedVideoID  string  `json:"top_viewed_video_id"`
		TopEngagedVideoID string  `json:"top_engaged_video_id"`
	}
)

// NewPostsSummary comment pending
func NewPostsSummary(posts []*Post) PostsSummary {
	var ps PostsSummary
	var totalUniqueViewers float64
	var totalPeopleReached float64
	var totalPostConsumptions float64
	var totalPostEngagements float64

	topViewedVideoPost := posts[0]
	topEngagedVideoPost := posts[0]

	videosUsed := make(map[string]*interface{})

	for _, post := range posts {
		ps.CampaignReach += post.Data.PeopleReached
		ps.CampaignViews += post.Data.VideoViews
		ps.MinutesViewed += post.Data.OverallMinutesViewed
		ps.UnqiueViewers += post.Data.UniqueViewers
		ps.TotalEngagement += post.Data.PostConsumptions + post.Data.PostEngagements

		totalUniqueViewers += post.Data.UniqueViewers
		totalPeopleReached += post.Data.PeopleReached
		totalPostConsumptions += post.Data.PostConsumptions
		totalPostEngagements += post.Data.PostEngagements

		if post.Data.VideoViews > topViewedVideoPost.Data.VideoViews {
			topViewedVideoPost = post
		}

		if post.Data.EngagementRate > topEngagedVideoPost.Data.EngagementRate {
			topEngagedVideoPost = post
		}

		if _, exists := videosUsed[post.ObjectID]; !exists {
			videosUsed[post.ObjectID] = nil
		}
	}

	ps.OverallViewRate = calculateViewRate(totalUniqueViewers, totalPeopleReached)
	ps.EngagementRate = calculateEngagementRate(totalPostConsumptions, totalPostEngagements, totalPeopleReached)
	ps.TopEngagedVideoID = topEngagedVideoPost.ObjectID
	ps.TopViewedVideoID = topViewedVideoPost.ObjectID
	ps.VideosPosted = len(videosUsed)

	return ps
}

// calculateViewRate comment pending
func calculateViewRate(totalUniqueVideoViews float64, totalPeopleReached float64) float64 {
	if totalUniqueVideoViews > 0 && totalPeopleReached > 0 {
		return (totalUniqueVideoViews / totalPeopleReached) * 100
	} else {
		return 0
	}
}

// calculateEngagementRate comment pending
func calculateEngagementRate(totalPostConsumptions float64, totalPostEngagements float64, totalPeopleReached float64) float64 {
	if totalPostConsumptions > 0 && totalPostEngagements > 0 && totalPeopleReached > 0 {
		return ((totalPostConsumptions + totalPostEngagements) / totalPeopleReached) * 100
	} else {
		return 0
	}
}

// ToJSON comment pending
func (p *PostsSummary) ToJSON() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func calculateVideosPosted(posts []*Post) int {
	videosUsed := make(map[string]*interface{})

	for _, post := range posts {
		if _, exists := videosUsed[post.ObjectID]; !exists {
			videosUsed[post.ObjectID] = nil
		}
	}

	return len(videosUsed)
}
