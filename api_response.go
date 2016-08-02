package fbintegration

import (
	facebookLib "github.com/huandu/facebook"
)

type (
	// APIResponse comment pending
	APIResponse struct {
		*API
		Result      *facebookLib.Result
		PagedResult *facebookLib.PagingResult
		Page        int
	}
)

// NewAPIResponse comment pending
func NewAPIResponse(facebookAPI *API, results *facebookLib.Result) *APIResponse {
	return &APIResponse{
		facebookAPI,
		results,
		nil,
		1,
	}
}

// Data comment pending
func (a *APIResponse) Data(paginate bool) ([]facebookLib.Result, error) {
	if paginate {
		if a.PagedResult == nil {
			paging, err := a.Result.Paging(a.Session)
			if err != nil {
				return []facebookLib.Result{}, err
			}

			a.PagedResult = paging
		} else {
			a.PagedResult.Next()
		}

		a.Page++
		return a.PagedResult.Data(), nil
	}

	var results []facebookLib.Result
	results = append(results, *a.Result)

	return results, nil
}

// HasResults comment pending
func (a *APIResponse) HasResults() bool {
	if a.PagedResult == nil {
		return true
	}
	return a.PagedResult.HasNext()
}
