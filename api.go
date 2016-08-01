package fbIntegration

import (
	facebookLib "github.com/huandu/facebook"
)

const facebookAPIVersion = "v2.7"

type (
	// API comment pending
	API struct {
		Session     *facebookLib.Session
		AccessToken string
	}
)

// NewAPI comment pending
func NewAPI() API {
	api := API{&facebookLib.Session{}, ""}
	api.Session.Version = facebookAPIVersion
	return api
}

// Batch comment pending
func (f *API) Batch(params ...facebookLib.Params) ([]*APIResponse, *Error) {
	var apiResponses []*APIResponse

	results, err := f.Session.BatchApi(params...)

	if err != nil {
		return []*APIResponse{}, NewError(err)
	}

	for _, result := range results {
		batch, err := result.Batch()

		if err != nil {
			return []*APIResponse{}, NewError(err)
		}

		apiResponse := NewAPIResponse(f, &batch.Result)
		apiResponses = append(apiResponses, apiResponse)
	}

	return apiResponses, nil
}

//Get comment pending
func (f *API) Get(path string, params facebookLib.Params) (*APIResponse, *Error) {
	results, err := f.Session.Get(path, params)

	if err != nil {
		wrappedError := NewError(err)
		return nil, wrappedError
	}

	return NewAPIResponse(f, &results), nil
}

// SetAccessToken comment pending
func (f *API) SetAccessToken(accessToken string) {
	f.Session.SetAccessToken(accessToken)
	f.AccessToken = accessToken
}
