package fbintegration

import (
	"fmt"
	facebookLib "github.com/huandu/facebook"
)

type (
	Params struct {
		Endpoint string
		Params   map[string]interface{}
	}

	// BatchParams comment pending
	BatchParams struct {
		Method      string
		RelativeUrl string
	}
)

func NewParams(endpoint string, params map[string]interface{}) Params {
	return Params{endpoint, params}
}

func NewBatchParams(endpoint string) BatchParams {
	return BatchParams{RelativeUrl: endpoint}
}

func (p BatchParams) ToFBParams() facebookLib.Params {
	if p.Method == "" {
		p.Method = "GET"
	}

	p.RelativeUrl = fmt.Sprintf("%s/%s", facebookAPIVersion, p.RelativeUrl)

	return facebookLib.Params{
		"method":       p.Method,
		"relative_url": p.RelativeUrl,
	}
}

func (p Params) ToFBParams() facebookLib.Params {
	return facebookLib.MakeParams(p.Params)
}
