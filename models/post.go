package fbintegration

type (
	// Post comment pending
	Post struct {
		ID       string `facebook:"id"        json:"post_id"`
		ObjectID string `facebook:"object_id" json:"object_id"`
	}
)
