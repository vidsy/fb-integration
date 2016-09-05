package fbintegration

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	facebookLib "github.com/huandu/facebook"
)

type (
	// DemographicSplit comment pending
	DemographicSplit struct {
		Total float64
		Data  []DemographicSplitItem
	}

	// DemographicSplitItem comment pending
	DemographicSplitItem struct {
		Gender   string `json:"gender"`
		AgeRange string `json:"age_range"`
		Value    string `json:"value"`
	}
)

// NewDemographicSplitFromResult comment pending
func NewDemographicSplitFromResult(results *facebookLib.Result, total float64) DemographicSplit {
	demographicSplit := DemographicSplit{Total: total}
	data := results.Get("data")
	slice := reflect.ValueOf(data)

	for i := 0; i < slice.Len(); i++ {
		demographicSplitItem := DemographicSplitItem{
			getResultValue(results, i, "gender"),
			getResultValue(results, i, "age"),
			getResultValue(results, i, "reach"),
		}
		demographicSplit.Push(demographicSplitItem)
	}

	return demographicSplit
}

// Push comment pending
func (as *DemographicSplit) Push(demographicSplitItem DemographicSplitItem) {
	as.Data = append(as.Data, demographicSplitItem)
}

// MarshalJSON comment pending
func (as DemographicSplit) MarshalJSON() ([]byte, error) {
	demographicSplit := map[string]map[string]float64{}

	for _, demographicSplitItem := range as.Data {
		_, exists := demographicSplit[demographicSplitItem.AgeRange]

		if !exists {
			demographicSplit[demographicSplitItem.AgeRange] = map[string]float64{"male": 0, "female": 0}
		}

		value, err := strconv.ParseFloat(demographicSplitItem.Value, 64)

		if err != nil {
			return []byte{}, err
		}

		demographicSplit[demographicSplitItem.AgeRange][demographicSplitItem.Gender] = (value / as.Total) * 100
	}

	json, err := json.Marshal(demographicSplit)

	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func getResultValue(results *facebookLib.Result, recordIndex int, key string) string {
	query := fmt.Sprintf("data.%d.%s", recordIndex, key)
	return results.Get(query).(string)
}
