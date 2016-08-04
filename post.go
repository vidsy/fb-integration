package fbintegration

import (
	"fmt"
	"log"
	"reflect"

	facebookLib "github.com/huandu/facebook"
)

type (
	// Post comment pending
	Post struct {
		ID                 string `facebook:"id"        json:"post_id"`
		ObjectID           string `facebook:"object_id" json:"object_id"`
		Insights           *facebookLib.Result
		ReactionBreakdown  []*facebookLib.Result
		TotalReactions     *facebookLib.Result
		Name               string         `json:"name"`
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
	}
)

// GenerateInsightParams comment pending
func (p Post) GenerateInsightParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/insights/post_impressions,post_impressions_paid,post_impressions_unique,post_impressions_paid_unique,post_video_views_paid,post_video_views_organic,post_video_views_organic_unique,post_video_views_paid_unique,post_video_view_time,post_video_avg_time_watched?period=lifetime&limit=20", p.ID),
	}
}

// GenerateReactionBreakdownParams comment pending
func (p Post) GenerateReactionBreakdownParams() []facebookLib.Params {
	var params []facebookLib.Params

	reactions := []string{
		"LIKE", "LOVE", "WOW", "HAHA", "SAD", "ANGRY", "THANKFUL",
	}

	for _, reaction := range reactions {
		params = append(params,
			facebookLib.Params{
				"method":       facebookLib.GET,
				"relative_url": fmt.Sprintf("%s/reactions?limit=0&summary=total_count&type=%s", p.ID, reaction),
			},
		)
	}

	return params
}

// GenerateTotalReactionsParams comment pending
func (p Post) GenerateTotalReactionsParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/reactions?limit=0&summary=total_count", p.ID),
	}
}

// ParseResults comment pending
func (p Post) ParseResults() {
	impressions := p.getInsightsValue("post_impressions")
	log.Println("post_impressions")
	if impressions != nil {
		batman := int(impressions["value"].(float64))
		log.Println(batman)
		p.Impressions = batman
	}

	paidImpressions := p.getInsightsValue("post_impressions_paid")
	if impressions != nil {
		p.PaidImpressions = int(paidImpressions["value"].(float64))
	}

	p.OrganicImpressions = (p.Impressions - p.PaidImpressions)

	reach := p.getInsightsValue("post_impressions_unique")
	if impressions != nil {
		p.Reach = int(reach["value"].(float64))
	}

	paidReach := p.getInsightsValue("post_impressions_paid_unique")
	if impressions != nil {
		p.PaidReach = int(paidReach["value"].(float64))
	}

	p.OrganicReach = (p.Reach - p.PaidReach)

	paidViews := p.getInsightsValue("post_video_views_paid")
	if paidViews != nil {
		p.PaidVideoViews = int(paidViews["value"].(float64))
	}

	organicViews := p.getInsightsValue("post_video_views_organic")
	if organicViews != nil {
		p.OrganicVideoViews = int(organicViews["value"].(float64))
	}

	p.VideoViews = (p.PaidVideoViews + p.OrganicVideoViews)

	uniquePaidViews := int(p.getInsightsValue("post_video_views_paid_unique")["value"].(float64))
	uniqueOrganicViews := int(p.getInsightsValue("post_video_views_organic_unique")["value"].(float64))

	p.UniqueVideoViews = (uniquePaidViews + uniqueOrganicViews)

	minutesViewed := p.getInsightsValue("post_video_view_time")
	if organicViews != nil {
		p.MinutesViewed = int(minutesViewed["value"].(float64))
	}

	averageDuration := p.getInsightsValue("post_video_avg_time_watched")
	if organicViews != nil {
		p.AverageDuration = int(averageDuration["value"].(float64))
	}
}

func (p Post) getInsightsValue(key string) map[string]interface{} {
	data := p.Insights.Get("data")
	slice := reflect.ValueOf(data)

	for i := 0; i < slice.Len(); i++ {
		query := fmt.Sprintf("data.%d.name", i)
		name := p.Insights.Get(query)

		if name == key {
			query = fmt.Sprintf("data.%d.values", i)
			values := p.Insights.Get(query).([]interface{})

			if len(values) > 0 {
				query = fmt.Sprintf("data.%d.values.0", i)
				return p.Insights.Get(query).(map[string]interface{})
			}

			return nil
		}
	}

	return nil
}
