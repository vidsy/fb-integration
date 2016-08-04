package fbintegration

import (
	"encoding/json"
	"fmt"
	"reflect"

	facebookLib "github.com/huandu/facebook"
)

type (
	// Post comment pending
	Post struct {
		ID       string `facebook:"id"        json:"post_id"`
		ObjectID string `facebook:"object_id" json:"object_id"`
		Name     string `json:"name"`
		Results  struct {
			Insights          *facebookLib.Result
			ReactionBreakdown []*facebookLib.Result
			TotalReactions    *facebookLib.Result
		}
		Data struct {
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
		} `json:"data"`
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

	for _, reaction := range p.reactionTypes() {
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
func (p *Post) ParseResults() {
	impressions := p.getInsightsValue("post_impressions")
	if impressions != nil {
		p.Data.Impressions = int(impressions["value"].(float64))
	}

	paidImpressions := p.getInsightsValue("post_impressions_paid")
	if impressions != nil {
		p.Data.PaidImpressions = int(paidImpressions["value"].(float64))
	}

	p.Data.OrganicImpressions = (p.Data.Impressions - p.Data.PaidImpressions)

	reach := p.getInsightsValue("post_impressions_unique")
	if impressions != nil {
		p.Data.Reach = int(reach["value"].(float64))
	}

	paidReach := p.getInsightsValue("post_impressions_paid_unique")
	if impressions != nil {
		p.Data.PaidReach = int(paidReach["value"].(float64))
	}

	p.Data.OrganicReach = (p.Data.Reach - p.Data.PaidReach)

	paidViews := p.getInsightsValue("post_video_views_paid")
	if paidViews != nil {
		p.Data.PaidVideoViews = int(paidViews["value"].(float64))
	}

	organicViews := p.getInsightsValue("post_video_views_organic")
	if organicViews != nil {
		p.Data.OrganicVideoViews = int(organicViews["value"].(float64))
	}

	p.Data.VideoViews = (p.Data.PaidVideoViews + p.Data.OrganicVideoViews)

	uniquePaidViews := int(p.getInsightsValue("post_video_views_paid_unique")["value"].(float64))
	uniqueOrganicViews := int(p.getInsightsValue("post_video_views_organic_unique")["value"].(float64))

	p.Data.UniqueVideoViews = (uniquePaidViews + uniqueOrganicViews)

	minutesViewed := p.getInsightsValue("post_video_view_time")
	if organicViews != nil {
		p.Data.MinutesViewed = int(minutesViewed["value"].(float64))
	}

	averageDuration := p.getInsightsValue("post_video_avg_time_watched")
	if organicViews != nil {
		p.Data.AverageDuration = int(averageDuration["value"].(float64))
	}

	p.Data.ReactionsTotal = p.getReactionsTotal(p.Results.TotalReactions)

	p.Data.Reactions = make(map[string]int)
	for i, reactionType := range p.reactionTypes() {
		p.Data.Reactions[reactionType] = p.getReactionsTotal(p.Results.ReactionBreakdown[i])
	}
}

// ToJSON comment pending
func (p *Post) ToJSON() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (p Post) getInsightsValue(key string) map[string]interface{} {
	data := p.Results.Insights.Get("data")
	slice := reflect.ValueOf(data)

	for i := 0; i < slice.Len(); i++ {
		query := fmt.Sprintf("data.%d.name", i)
		name := p.Results.Insights.Get(query)

		if name == key {
			query = fmt.Sprintf("data.%d.values", i)
			values := p.Results.Insights.Get(query).([]interface{})

			if len(values) > 0 {
				query = fmt.Sprintf("data.%d.values.0", i)
				return p.Results.Insights.Get(query).(map[string]interface{})
			}

			return nil
		}
	}

	return nil
}

func (p Post) getReactionsTotal(result *facebookLib.Result) int {
	return int(result.Get("summary.total_count").(float64))
}

func (p Post) reactionTypes() []string {
	return []string{
		"LIKE", "LOVE", "WOW", "HAHA", "SAD", "ANGRY", "THANKFUL",
	}
}
