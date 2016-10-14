package fbintegration

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	facebookLib "github.com/huandu/facebook"
)

type (
	// Post comment pending
	Post struct {
		Name     string       `facebook:"message"   json:"name"`
		ID       string       `facebook:"id"        json:"post_id"`
		AdID     string       `json:"ad_id"`
		AdsetID  string       `json:"adset_id"`
		ObjectID string       `json:"object_id"`
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
func (p Post) GenerateCommentsParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/comments?summary=true&filter=stream", p.ID))
}

// GenerateInsightParams comment pending
func (p Post) GenerateInsightParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/insights/post_video_views_unique,post_engaged_users,post_video_complete_views_organic,post_video_complete_views_paid,post_consumptions,post_impressions,post_impressions_paid,post_impressions_unique,post_impressions_paid_unique,post_video_views,post_video_views_paid,post_video_views_organic,post_video_view_time,post_video_avg_time_watched,post_stories_by_action_type?period=lifetime&limit=20", p.ID))
}

// GenerateParams comment pending
func (p *Post) GenerateParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s?fields=%s", p.ID, "object_id,message"))
}

// GeneratePostCreatedTimestampParams comment pending
func (p *Post) GeneratePostCreatedTimestampParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s?fields=%s", p.ID, "created_time"))
}

// GenerateReactionBreakdownParams comment pending
func (p Post) GenerateReactionBreakdownParams() []BatchParams {
	var params []BatchParams

	for _, reaction := range p.ReactionTypes() {
		params = append(params, NewBatchParams(fmt.Sprintf("%s/reactions?limit=0&summary=total_count&type=%s", p.ID, reaction)))
	}

	return params
}

// GenerateSharesParams comment pending
func (p Post) GenerateSharesParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/sharedposts", p.ID))
}

// GenerateTotalReactionsParams comment pending
func (p Post) GenerateTotalReactionsParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/reactions?limit=0&summary=total_count", p.ID))
}

// ParseResults comment pending
func (p *Post) ParseResults() {
	p.Data = &PostData{}

	sampledPeopleReached := p.getAdInsightsValue("reach")
	if sampledPeopleReached != nil {
		float64Value, err := strconv.ParseFloat(sampledPeopleReached.(string), 64)
		if err == nil {
			p.Data.SampledPeopleReached = float64Value
		}
	}

	peopleReached := p.getInsightsValue("post_impressions_unique")
	if peopleReached != nil {
		p.Data.PeopleReached = peopleReached["value"].(float64)
	}

	peopleReachedPaid := p.getInsightsValue("post_impressions_paid_unique")
	if peopleReachedPaid != nil {
		p.Data.PeopleReachedPaid = peopleReachedPaid["value"].(float64)
	}

	p.Data.PeopleReachedOrganic = (p.Data.PeopleReached - p.Data.PeopleReachedPaid)

	p.Data.Reactions = p.getReactionsTotal(p.Results.TotalReactions)

	p.Data.ReactionsBreakdown = make(map[string]float64)
	for i, reactionType := range p.ReactionTypes() {
		p.Data.ReactionsBreakdown[reactionType] = p.getReactionsTotal(p.Results.ReactionBreakdown[i])
	}

	p.Data.Comments = p.getActionTypeTotal("comment")
	p.Data.Shares = p.getActionTypeTotal("share")

	postConsumptions := p.getInsightsValue("post_consumptions")
	if postConsumptions != nil {
		p.Data.PostConsumptions = postConsumptions["value"].(float64)
	}

	p.Data.PostEngagements = p.Data.Reactions + p.Data.Comments + p.Data.Shares

	engagedUsers := p.getInsightsValue("post_engaged_users")
	if engagedUsers != nil {
		p.Data.EngagedUsers = engagedUsers["value"].(float64)
	}

	if p.Data.PeopleReached > 0 {
		p.Data.EngagementRate = (((p.Data.PostConsumptions + p.Data.PostEngagements) / p.Data.PeopleReached) * 100)
	}

	p.Data.Demographic = p.generateDemographic()
	p.Data.Targeting = p.generateTargeting()

	paidReach := p.getInsightsValue("post_impressions_paid_unique")
	if paidReach != nil {
		p.Data.PaidReach = paidReach["value"].(float64)
	}

	videoViews := p.getInsightsValue("post_video_views")
	if videoViews != nil {
		p.Data.VideoViews = videoViews["value"].(float64)
	}

	videoViewsOrganic := p.getInsightsValue("post_video_views_organic")
	if videoViewsOrganic != nil {
		p.Data.VideoViewsOrganic = videoViewsOrganic["value"].(float64)
	}

	videoViewsPaid := p.getInsightsValue("post_video_views_paid")
	if videoViewsPaid != nil {
		p.Data.VideoViewsPaid = videoViewsPaid["value"].(float64)
	}

	videoViewsUnique := p.getInsightsValue("post_video_views_unique")
	if videoViewsUnique != nil {
		p.Data.UniqueViewers = videoViewsUnique["value"].(float64)
	}

	if p.Data.PeopleReached > 0 {
		p.Data.ViewRate = (p.Data.UniqueViewers / p.Data.PeopleReached) * 100
	}

	viewsToNinetyFivePercentCompleteOrganic := p.getInsightsValue("post_video_complete_views_organic")
	if viewsToNinetyFivePercentCompleteOrganic != nil {
		p.Data.ViewsToNinetyFivePercentCompleteOrganic = viewsToNinetyFivePercentCompleteOrganic["value"].(float64)
	}

	viewsToNinetyFivePercentCompletePaid := p.getInsightsValue("post_video_complete_views_paid")
	if viewsToNinetyFivePercentCompletePaid != nil {
		p.Data.ViewsToNinetyFivePercentCompletePaid = viewsToNinetyFivePercentCompletePaid["value"].(float64)
	}

	p.Data.ViewsToNinetyFivePercentComplete = p.Data.ViewsToNinetyFivePercentCompletePaid + p.Data.ViewsToNinetyFivePercentCompleteOrganic

	if p.Data.VideoViews > 0 {
		p.Data.PercentViewsCompleted = (p.Data.ViewsToNinetyFivePercentComplete / p.Data.VideoViews) * 100
	}

	averageDurationWatched := p.getInsightsValue("post_video_avg_time_watched")
	if averageDurationWatched != nil {
		total := averageDurationWatched["value"].(float64)
		p.Data.AverageDurationWatched = (total / 1000)
	}

	overallMinutesViewed := p.getInsightsValue("post_video_view_time")
	if overallMinutesViewed != nil {
		total := overallMinutesViewed["value"].(float64)
		p.Data.OverallMinutesViewed = ((total / 1000) / 60)
	}

	totalSpend := p.getAdInsightsValue("spend")
	if totalSpend != nil {
		p.Data.Spend = totalSpend.(float64)
	}

	if p.Data.VideoViewsOrganic > 0 || p.Data.VideoViewsPaid > 0 {
		p.Data.VideoViewCost = p.Data.Spend / (p.Data.VideoViewsOrganic + p.Data.VideoViewsPaid)
	}

	createdTimestamp, err := p.getCreatedTimestamp()
	if err == nil {
		p.Data.CreatedTimestamp = createdTimestamp
	}
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

func (p Post) generateDemographic() DemographicSplit {
	return NewDemographicSplitFromResult(p.Results.AdBreakdownInsights, p.Data.SampledPeopleReached)
}

func (p Post) generateTargeting() AdTargeting {
	return NewAdTargetingFromResult(p.Results.Targeting)
}

func (p Post) getComments() float64 {
	return p.Results.Engagement[0].Get("summary.total_count").(float64)
}

func (p Post) getCreatedTimestamp() (time.Time, error) {
	layout := "2006-01-02T15:04:05-0700"
	t := p.Results.CreatedTimestamp.Get("created_time")

	return time.Parse(layout, t.(string))
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

func (p Post) getActionTypeTotal(actionType string) float64 {
	postStories := p.getInsightsValue("post_stories_by_action_type")
	if postStories["value"] == nil {
		return 0
	}

	reflectedMap := reflect.ValueOf(postStories["value"])
	valuesMap := reflectedMap.Interface().(map[string]interface{})
	value := valuesMap[actionType]

	if value != nil {
		return valuesMap[actionType].(float64)
	}
	return 0
}

func (p Post) getReactionsTotal(result *facebookLib.Result) float64 {
	return result.Get("summary.total_count").(float64)
}
