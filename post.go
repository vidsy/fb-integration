package fbintegration

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	facebookLib "github.com/huandu/facebook"
)

type (
	// Post comment pending
	Post struct {
		Name     string       `facebook:"message" json:"name"`
		ID       string       `facebook:"id"        json:"post_id"`
		AdID     string       `json:"ad_id"`
		ObjectID string       `facebook:"object_id" json:"object_id"`
		Results  *PostResults `json:"-"`
		Data     *PostData    `json:"data,omitempty"`
	}
)

// NewPostFromResult comment pending
func NewPostFromResult(result facebookLib.Result) Post {
	var post Post
	post.Results = &PostResults{}
	result.DecodeField("", &post)
	return post
}

// GenerateCommentsParams comment pending
func (p Post) GenerateCommentsParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/comments?summary=true&filter=stream", p.ID),
	}
}

// GenerateInsightParams comment pending
func (p Post) GenerateInsightParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/insights/post_video_views_unique,post_engaged_users,post_video_complete_views_organic,post_video_complete_views_paid,post_consumptions,post_impressions,post_impressions_paid,post_impressions_unique,post_impressions_paid_unique,post_video_views,post_video_views_paid,post_video_views_organic,post_video_view_time,post_video_avg_time_watched?period=lifetime&limit=20", p.ID),
	}
}

// GenerateParams comments pending
func (p *Post) GenerateParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s?fields=%s", p.ID, "object_id,message"),
	}
}

// GenerateReactionBreakdownParams comment pending
func (p Post) GenerateReactionBreakdownParams() []facebookLib.Params {
	var params []facebookLib.Params

	for _, reaction := range p.ReactionTypes() {
		params = append(params,
			facebookLib.Params{
				"method":       facebookLib.GET,
				"relative_url": fmt.Sprintf("%s/reactions?limit=0&summary=total_count&type=%s", p.ID, reaction),
			},
		)
	}

	return params
}

// GenerateSharesParams comment pending
func (p Post) GenerateSharesParams() facebookLib.Params {
	return facebookLib.Params{
		"method":       facebookLib.GET,
		"relative_url": fmt.Sprintf("%s/sharedposts", p.ID),
	}
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
	p.Data = &PostData{}

	peopleReached := p.getInsightsValue("post_impressions_unique")
	if peopleReached != nil {
		p.Data.PeopleReached = peopleReached["value"].(float64)
	}

	peopleReachedPaid := p.getInsightsValue("post_impressions_paid_unique")
	if peopleReachedPaid != nil {
		p.Data.PeopleReachedPaid = peopleReachedPaid["value"].(float64)
	}

	p.Data.PeopleReachedPaid = (p.Data.PeopleReached - p.Data.PeopleReachedPaid)

	impressions := p.getInsightsValue("post_impressions")
	if impressions != nil {
		p.Data.Impressions = impressions["value"].(float64)
	}

	paidImpressions := p.getInsightsValue("post_impressions_paid")
	if impressions != nil {
		p.Data.PaidImpressions = paidImpressions["value"].(float64)
	}

	organicVideoCompletion := p.getInsightsValue("post_video_complete_views_organic")
	if organicVideoCompletion != nil {
		p.Data.OrganicVideoCompletion = organicVideoCompletion["value"].(float64)
	}

	paidVideoCompletion := p.getInsightsValue("post_video_complete_views_paid")
	if paidVideoCompletion != nil {
		p.Data.PaidVideoCompletion = paidVideoCompletion["value"].(float64)
	}

	p.Data.VideoCompletion = p.Data.PaidVideoCompletion + p.Data.OrganicVideoCompletion

	lifetimeEngagedUsers := p.getInsightsValue("post_engaged_users")
	if lifetimeEngagedUsers != nil {
		p.Data.LifetimeEngagedUsers = lifetimeEngagedUsers["value"].(float64)
	}

	p.Data.OrganicImpressions = (p.Data.Impressions - p.Data.PaidImpressions)

	reach := p.getInsightsValue("post_impressions_unique")
	if impressions != nil {
		p.Data.Reach = reach["value"].(float64)
	}

	sampledReach := p.getAdInsightsValue("reach")
	if sampledReach != nil {
		float64Value, err := strconv.ParseFloat(sampledReach.(string), 64)
		if err == nil {
			p.Data.SampledReach = float64Value
		}
	}

	paidReach := p.getInsightsValue("post_impressions_paid_unique")
	if impressions != nil {
		p.Data.PaidReach = paidReach["value"].(float64)
	}

	p.Data.OrganicReach = (p.Data.Reach - p.Data.PaidReach)

	paidViews := p.getInsightsValue("post_video_views_paid")
	if paidViews != nil {
		p.Data.PaidVideoViews = paidViews["value"].(float64)
	}

	organicViews := p.getInsightsValue("post_video_views_organic")
	if organicViews != nil {
		p.Data.OrganicVideoViews = organicViews["value"].(float64)
	}

	videoViews := p.getInsightsValue("post_video_views")
	if videoViews != nil {
		p.Data.VideoViews = videoViews["value"].(float64)
	}

	p.Data.UniqueVideoViews = p.getInsightsValue("post_video_views_unique")["value"].(float64)

	minutesViewed := p.getInsightsValue("post_video_view_time")
	if organicViews != nil {
		total := minutesViewed["value"].(float64)
		p.Data.MinutesViewed = ((total / 1000) / 60)
	}

	averageDuration := p.getInsightsValue("post_video_avg_time_watched")
	if averageDuration != nil {
		total := averageDuration["value"].(float64)
		p.Data.AverageDuration = (total / 1000)
	}

	p.Data.VideoCompletionPercent = (p.Data.VideoCompletion / p.Data.VideoViews) * 100
	p.Data.ReactionsTotal = p.getReactionsTotal(p.Results.TotalReactions)

	p.Data.Reactions = make(map[string]float64)
	for i, reactionType := range p.ReactionTypes() {
		p.Data.Reactions[reactionType] = p.getReactionsTotal(p.Results.ReactionBreakdown[i])
	}

	p.Data.Comments = p.getComments()

	p.Data.Shares = p.getShares()

	totalClicks := p.getInsightsValue("post_consumptions")
	if totalClicks != nil {
		p.Data.Clicks = totalClicks["value"].(float64)
	}

	totalSpend := p.getAdInsightsValue("spend")
	if totalSpend != nil {
		p.Data.Spend = totalSpend.(float64)
	}

	p.Data.Actions = p.Data.ReactionsTotal + p.Data.Comments + p.Data.Shares + p.Data.Clicks
	if p.Data.Actions > 0 && p.Data.Reach > 0 {
		p.Data.EngagementRate = (p.Data.Actions / p.Data.Reach) * 100
	}

	if p.Data.UniqueVideoViews > 0 && p.Data.Reach > 0 {
		p.Data.ViewRate = (p.Data.UniqueVideoViews / p.Data.Reach) * 100
	}

	if p.Data.OrganicVideoViews > 0 || p.Data.PaidVideoViews > 0 {
		p.Data.VideoViewCost = p.Data.Spend / (p.Data.OrganicVideoViews + p.Data.PaidVideoViews)
	}
	p.Data.AudienceSplit = p.generateAudienceSplit()
}

// ReactionTypes comment pending
func (p Post) ReactionTypes() []string {
	return []string{
		"LIKE", "LOVE", "WOW", "HAHA", "SAD", "ANGRY", "THANKFUL",
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

func (p Post) generateAudienceSplit() AudienceSplit {
	return NewAudienceSplitFromResult(p.Results.AdBreakdownInsights, p.Data.SampledReach)
}

func (p Post) getComments() float64 {
	return p.Results.Engagement[0].Get("summary.total_count").(float64)
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

func (p Post) getShares() float64 {
	queryBase := "data.0.unique_actions"
	data := p.Results.AdInsights.Get(queryBase)
	slice := reflect.ValueOf(data)

	for i := 0; i < slice.Len(); i++ {
		query := fmt.Sprintf("%s.%d.action_type", queryBase, i)
		actionType := p.Results.AdInsights.Get(query)

		if actionType == "post" {
			query := fmt.Sprintf("%s.%d.value", queryBase, i)
			return p.Results.AdInsights.Get(query).(float64)
		}
	}
	return 0
}

func (p Post) getAdInsightsValue(key string) interface{} {
	query := fmt.Sprintf("data.0.%s", key)
	value := p.Results.AdInsights.Get(query)

	if value != nil {
		return value
	}

	return nil
}

func (p Post) getReactionsTotal(result *facebookLib.Result) float64 {
	return result.Get("summary.total_count").(float64)
}
