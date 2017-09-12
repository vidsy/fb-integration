package fbintegration

import (
	"fmt"

	facebookLib "github.com/huandu/facebook"
)

type (
	// Params comment pending
	Params struct {
		Endpoint string
		Params   map[string]interface{}
	}

	// BatchParams comment pending
	BatchParams struct {
		Method      string
		RelativeURL string
	}
)

// NewParams comment pending
func NewParams(endpoint string, params map[string]interface{}) Params {
	return Params{endpoint, params}
}

// NewBatchParams comment pending
func NewBatchParams(endpoint string) BatchParams {
	return BatchParams{
		RelativeURL: endpoint,
	}
}

// ToFbParams comment pending
func (p BatchParams) ToFbParams() facebookLib.Params {
	if p.Method == "" {
		p.Method = "GET"
	}

	p.RelativeURL = fmt.Sprintf("%s/%s", facebookAPIVersion, p.RelativeURL)

	return facebookLib.Params{
		"method":       p.Method,
		"relative_url": p.RelativeURL,
	}
}

// ToFbParams comment pending
func (p Params) ToFbParams() facebookLib.Params {
	return facebookLib.MakeParams(p.Params)
}
