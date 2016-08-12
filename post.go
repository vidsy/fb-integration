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
		"relative_url": fmt.Sprintf("%s/insights/post_impressions,post_impressions_paid,post_impressions_unique,post_impressions_paid_unique,post_video_views_paid,post_video_views_organic,post_video_views_organic_unique,post_video_views_paid_unique,post_video_view_time,post_video_avg_time_watched?period=lifetime&limit=20", p.ID),
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
		total := int(minutesViewed["value"].(float64))
		p.Data.MinutesViewed = ((total / 1000) / 60)
	}

	averageDuration := p.getInsightsValue("post_video_avg_time_watched")
	if averageDuration != nil {
		total := int(averageDuration["value"].(float64))
		p.Data.AverageDuration = (total / 1000)
	}

	p.Data.ReactionsTotal = p.getReactionsTotal(p.Results.TotalReactions)

	p.Data.Reactions = make(map[string]int)
	for i, reactionType := range p.ReactionTypes() {
		p.Data.Reactions[reactionType] = p.getReactionsTotal(p.Results.ReactionBreakdown[i])
	}

	p.Data.Comments = p.getComments()

	p.Data.Shares = p.getShares()

	totalClicks := p.getAdInsightsValue("clicks")
	if totalClicks != nil {
		int64Value, err := strconv.ParseInt(totalClicks.(string), 10, 32)
		if err == nil {
			p.Data.Clicks = int64Value
		}
	}

	totalSpend := p.getAdInsightsValue("spend")
	if totalSpend != nil {
		p.Data.Spend = totalSpend.(float64)
	}

	uniquePeopleEngaged := p.getAdInsightsValue("total_unique_actions")
	if uniquePeopleEngaged != nil {
		int64Value, err := strconv.ParseInt(uniquePeopleEngaged.(string), 10, 32)
		if err == nil {
			p.Data.UniquePeopleEngaged = int64Value
		}
	}

	p.Data.Actions = int64(p.Data.ReactionsTotal) + int64(p.Data.Comments) + int64(p.Data.Shares) + p.Data.Clicks
	if p.Data.Actions > 0 && p.Data.UniqueVideoViews > 0 {
		p.Data.EngagementPercentPeopleViewed = p.round((float64(p.Data.Actions)/float64(p.Data.UniqueVideoViews))*100, 2)
	}

	if p.Data.VideoViews > 0 && p.Data.Impressions > 0 {
		p.Data.ViewRate = p.round((float64(p.Data.VideoViews)/float64(p.Data.Impressions))*100, 2)
	}

	if p.Data.OrganicVideoViews > 0 || p.Data.PaidVideoViews > 0 {
		p.Data.VideoViewCost = p.round(p.Data.Spend/(float64(p.Data.OrganicVideoViews)+float64(p.Data.PaidVideoViews)), 4)
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

func (p Post) generateAudienceSplit() []*AudienceSplit {
	data := p.Results.AdBreakdownInsights.Get("data")
	slice := reflect.ValueOf(data)
	audienceSplits := make([]*AudienceSplit, slice.Len())

	for i := 0; i < slice.Len(); i++ {
		audienceSplits[i] = NewAudienceSplitFromResult(p.Results.AdBreakdownInsights, i)
	}

	return audienceSplits
}

func (p Post) getComments() int {
	return int(p.Results.Engagement[0].Get("summary.total_count").(float64))
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

func (p Post) getShares() int {
	data := p.Results.Engagement[1].Get("data")
	slice := reflect.ValueOf(data)

	return slice.Len()
}

func (p Post) getAdInsightsValue(key string) interface{} {
	query := fmt.Sprintf("data.0.%s", key)
	value := p.Results.AdInsights.Get(query)

	if value != nil {
		return value
	}

	return nil
}

func (p Post) getReactionsTotal(result *facebookLib.Result) int {
	return int(result.Get("summary.total_count").(float64))
}

func (p Post) round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}
