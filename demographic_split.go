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
func (d *DemographicSplit) Push(demographicSplitItem DemographicSplitItem) {
	d.Data = append(d.Data, demographicSplitItem)
}

// BuildMap builds a map representation of the breakdown.
func (d DemographicSplit) BuildMap() (map[string]map[string]float64, error) {
	demographicSplit := map[string]map[string]float64{}

	for _, demographicSplitItem := range d.Data {
		_, exists := demographicSplit[demographicSplitItem.AgeRange]

		if !exists {
			demographicSplit[demographicSplitItem.AgeRange] = map[string]float64{"male": 0, "female": 0}
		}

		value, err := strconv.ParseFloat(demographicSplitItem.Value, 64)

		if err != nil {
			return nil, err
		}

		demographicSplit[demographicSplitItem.AgeRange][demographicSplitItem.Gender] = (value / d.Total) * 100
	}

	return demographicSplit, nil
}

// MarshalJSON comment pending
func (d DemographicSplit) MarshalJSON() ([]byte, error) {
	demographicSplitMap, err := d.BuildMap()
	if err != nil {
		return []byte{}, err
	}

	json, err := json.Marshal(demographicSplitMap)

	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func getResultValue(results *facebookLib.Result, recordIndex int, key string) string {
	query := fmt.Sprintf("data.%d.%s", recordIndex, key)
	return results.Get(query).(string)
}
