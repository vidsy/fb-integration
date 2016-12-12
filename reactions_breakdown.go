package fbintegration

type (
	// ReactionsBreakdown holds details of a reaction breakdown for a post
	// or group of posts.
	ReactionsBreakdown struct {
		Type  string
		Value float64
	}
)

// Increment increments the value by the given amount.
func (r *ReactionsBreakdown) Increment(amount float64) {
	r.Value += amount
}
