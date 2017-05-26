package fbintegration

import (
	"encoding/json"
	"fmt"

	"strings"

	facebookLib "github.com/huandu/facebook"
)

type (
	// Post contains details for a given post including the name, id and insights.
	Post struct {
		AdType   string       `json:"ad_type"`
		Name     string       `facebook:"message"   json:"name"`
		ID       string       `facebook:"id"        json:"post_id"`
		AdID     string       `json:"ad_id"`
		AdsetID  string       `json:"adset_id"`
		ObjectID string       `json:"object_id"`
		Results  *PostResults `json:"-"`
		Data     *PostData    `json:"data,omitempty"`
	}
)

// NewPostFromResult creates a new Post struct from a facebookLib.Result.
func NewPostFromResult(result facebookLib.Result) Post {
	var post Post
	post.Results = &PostResults{}
	result.DecodeField("", &post)
	return post
}

// GenerateCommentsParams create the params for getting comments for a post.
func (p Post) GenerateCommentsParams() BatchParams {
	uri := fmt.Sprintf("%s/comments?summary=true&limit=0&period=lifetime", p.ID)

	return NewBatchParams(uri)
}

// GenerateInsightParams create the params for getting insights for a post from the Graph API.
func (p Post) GenerateInsightParams() BatchParams {
	fields := []string{
		"post_consumptions",
		"post_engaged_users",
		"post_impressions",
		"post_impressions_paid",
		"post_impressions_paid_unique",
		"post_impressions_unique",
		"post_stories_by_action_type",
		"post_video_avg_time_watched",
		"post_video_complete_views_organic",
		"post_video_complete_views_paid",
		"post_video_view_time",
		"post_video_views",
		"post_video_views_organic",
		"post_video_views_paid",
		"post_video_views_unique",
	}

	uri := fmt.Sprintf(
		"%s/insights/%s?period=lifetime&limit=20",
		p.ID,
		strings.Join(fields, ","),
	)

	return NewBatchParams(uri)
}

// GenerateParams creates the params for getting information on a post.
func (p *Post) GenerateParams() BatchParams {
	fields := []string{
		"object_id",
		"message",
	}

	uri := fmt.Sprintf(
		"%s?fields=%s",
		p.ID,
		strings.Join(fields, ","),
	)

	return NewBatchParams(uri)
}

// GeneratePostParams creates params for getting back data on a
// particular post.
func (p *Post) GeneratePostParams() BatchParams {
	fields := []string{
		"created_time",
		"shares",
	}

	uri := fmt.Sprintf(
		"%s?fields=%s",
		p.ID,
		strings.Join(fields, ","),
	)

	return NewBatchParams(uri)
}

// GenerateReactionBreakdownParams creates params for getting the reaction breakdown for a post.
func (p Post) GenerateReactionBreakdownParams() []BatchParams {
	var params []BatchParams

	for _, reaction := range p.ReactionTypes() {
		uri := fmt.Sprintf(
			"%s/reactions?limit=0&summary=total_count&type=%s",
			p.ID,
			reaction,
		)

		params = append(params, NewBatchParams(uri))
	}

	return params
}

// GenerateSharesParams creates params for getting the shared count for a post.
func (p Post) GenerateSharesParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/sharedposts", p.ID))
}

// GenerateTotalReactionsParams creates params for gettting the total reaction count.
func (p Post) GenerateTotalReactionsParams() BatchParams {
	return NewBatchParams(fmt.Sprintf("%s/reactions?limit=0&summary=total_count", p.ID))
}

// ReactionTypes slice of currently supported reaction type for facebook posts.
func (p Post) ReactionTypes() []string {
	return []string{
		"ANGRY",
		"HAHA",
		"LIKE",
		"LOVE",
		"SAD",
		"THANKFUL",
		"WOW",
	}
}

// ToJSON marshal the current struct into a byte array then
// return a string representation of that.
func (p *Post) ToJSON() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
