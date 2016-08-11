package fbintegration

import (
	"fmt"
	facebookLib "github.com/huandu/facebook"
)

type (
	// AudienceSplit commend pending
	AudienceSplit struct {
		Gender   string `json:"gender"`
		AgeRange string `json:"age_range"`
		Value    string `json:"value"`
	}
)

// NewAudienceSplitFromResult comment pending
func NewAudienceSplitFromResult(results *facebookLib.Result, recordIndex int) *AudienceSplit {
	return &AudienceSplit{
		getResultValue(results, recordIndex, "gender"),
		getResultValue(results, recordIndex, "age"),
		getResultValue(results, recordIndex, "reach"),
	}
}

func getResultValue(results *facebookLib.Result, recordIndex int, key string) string {
	query := fmt.Sprintf("data.%d.%s", recordIndex, key)
	return results.Get(query).(string)
}
