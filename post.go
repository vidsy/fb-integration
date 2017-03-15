package fbintegration

import (
	"fmt"

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
	return NewBatchParams(fmt.Sprintf("%s/comments?summary=true&limit=0&period=lifetime", p.ID))
}

// GenerateInsightParams comment pending
func (p Post) GenerateInsightParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/insights/post_video_views_unique,post_engaged_users,post_video_complete_views_organic,post_video_complete_views_paid,post_consumptions,post_impressions,post_impressions_paid,post_impressions_unique,post_impressions_paid_unique,post_video_views,post_video_views_paid,post_video_views_organic,post_video_view_time,post_video_avg_time_watched,post_stories_by_action_type?period=lifetime&limit=20", p.ID))
}

// GenerateParams comment pending
func (p *Post) GenerateParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s?fields=%s", p.ID, "object_id,message"))
}

// GeneratePostParams creates params for getting back data on a
// particular post.
func (p *Post) GeneratePostParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s?fields=%s", p.ID, "created_time,shares"))
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

// ReactionTypes slice of currently supported reaction type for facebook posts.
func (p Post) ReactionTypes() []string {
	return []string{
		"LIKE", "LOVE", "WOW", "HAHA", "SAD", "ANGRY", "THANKFUL",
	}
}
