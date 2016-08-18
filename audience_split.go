package fbintegration

import (
	"encoding/json"
	"fmt"
	facebookLib "github.com/huandu/facebook"
	"reflect"
	"strconv"
)

type (
	// AudienceSplit comment pending
	AudienceSplit struct {
		Total float64
		Data  []AudienceSplitItem
	}

	// AudienceSplitItem comment pending
	AudienceSplitItem struct {
		Gender   string `json:"gender"`
		AgeRange string `json:"age_range"`
		Value    string `json:"value"`
	}
)

// NewAudienceSplitFromResult comment pending
func NewAudienceSplitFromResult(results *facebookLib.Result, total float64) AudienceSplit {
	audienceSplit := AudienceSplit{Total: total}
	data := results.Get("data")
	slice := reflect.ValueOf(data)

	for i := 0; i < slice.Len(); i++ {
		audienceSplitItem := AudienceSplitItem{
			getResultValue(results, i, "gender"),
			getResultValue(results, i, "age"),
			getResultValue(results, i, "reach"),
		}
		audienceSplit.Push(audienceSplitItem)
	}

	return audienceSplit
}

func (as *AudienceSplit) Push(audienceSplitItem AudienceSplitItem) {
	as.Data = append(as.Data, audienceSplitItem)
}

func (as AudienceSplit) MarshalJSON() ([]byte, error) {
	audienceSplit := map[string]map[string]float64{}

	for _, audienceSplitItem := range as.Data {
		_, exists := audienceSplit[audienceSplitItem.AgeRange]

		if !exists {
			audienceSplit[audienceSplitItem.AgeRange] = map[string]float64{"male": 0, "female": 0}
		}

		value, err := strconv.ParseFloat(audienceSplitItem.Value, 64)

		if err != nil {
			return []byte{}, err
		}

		audienceSplit[audienceSplitItem.AgeRange][audienceSplitItem.Gender] = (value / as.Total) * 100
	}

	json, err := json.Marshal(audienceSplit)

	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func getResultValue(results *facebookLib.Result, recordIndex int, key string) string {
	query := fmt.Sprintf("data.%d.%s", recordIndex, key)
	return results.Get(query).(string)
}
