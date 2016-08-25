package fbintegration

type (
	// AdBatch comment pending
	AdBatch struct {
		Ads []Ad
	}
)

// CreativeParams comment pending
func (a *AdBatch) CreativeParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(a.Ads); i++ {
		params = append(params, a.Ads[i].Creative.GenerateParams())
	}

	return params
}

// PostParams comment pending
func (a *AdBatch) PostParams() []BatchParams {
	var params []BatchParams

	for i := 0; i < len(a.Ads); i++ {
		params = append(params, a.Ads[i].Post.GenerateParams())
	}

	return params
}
