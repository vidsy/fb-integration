package fbintegration

import (
	facebookLib "github.com/huandu/facebook"
)

type (
	//AdBatch comment pending
	AdBatch struct {
		Ads []Ad
	}
)

//CreativeParams comment pending
func (a *AdBatch) CreativeParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(a.Ads); i++ {
		params = append(params, a.Ads[i].Creative.GenerateParams())
	}

	return params
}

//PostParams comment pending
func (a *AdBatch) PostParams() []facebookLib.Params {
	var params []facebookLib.Params

	for i := 0; i < len(a.Ads); i++ {
		params = append(params, a.Ads[i].Post.GenerateParams())
	}

	return params
}
