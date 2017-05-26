package fbintegration

import (
	"fmt"
	"strings"

	facebookLib "github.com/huandu/facebook"
)

const videoType = "VIDEO"

type (
	// Creative comment pending
	Creative struct {
		ID              string                  `facebook:"id"`
		ObjectStorySpec CreativeObjectStorySpec `facebook:"object_story_spec"`
		ObjectType      string                  `facebook:"object_type"`
		PostID          string                  `facebook:"effective_object_story_id"`
	}

	// CreativeObjectStorySpec comment pending
	CreativeObjectStorySpec struct {
		VideoData CreativeObjectVideoData `facebook:"video_data"`
	}

	// CreativeObjectVideoData comment pending
	CreativeObjectVideoData struct {
		VideoID     string `facebook:"video_id"`
		Description string `facebook:"description"`
	}
)

// NewCreativeFromResult comment pending
func NewCreativeFromResult(result facebookLib.Result) Creative {
	var creative Creative
	result.DecodeField("", &creative)
	return creative
}

// GenerateParams comments pending
func (c *Creative) GenerateParams() BatchParams {
	fields := []string{
		"object_id",
		"object_type",
		"effective_object_story_id",
		"object_story_spec",
	}

	uri := fmt.Sprintf("%s?fields=%s", c.ID, strings.Join(fields, ","))

	return NewBatchParams(uri)
}

// IsVideo comment pending
func (c *Creative) IsVideo() bool {
	if c.ObjectType == videoType {
		return true
	}
	return false
}

// HasObjectID comment pending
func (c Creative) HasObjectID() bool {
	if c.ObjectStorySpec.VideoData.VideoID != "" {
		return true
	}
	return false
}

// ObjectID comment pending
func (c Creative) ObjectID() string {
	if c.HasObjectID() {
		return c.ObjectStorySpec.VideoData.VideoID
	}

	return ""
}
