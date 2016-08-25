package fbintegration

import (
	"fmt"
	facebookLib "github.com/huandu/facebook"
	"github.com/pariz/gountries"
	"reflect"
)

type (
	// AdTargeting comment pending
	AdTargeting struct {
		Locations []string `json:"locations"`
		AgeMin    float64  `json:"age_min"`
		AgeMax    float64  `json:"age_max"`
		Interests []string `json:"interests"`
	}
)

// NewAdTargetingFromResult comment pending
func NewAdTargetingFromResult(results *facebookLib.Result) AdTargeting {
	targeting := results.Get("targeting")
	query := gountries.New()
	adTargeting := AdTargeting{}

	if targeting != nil {
		countries := results.Get("targeting.geo_locations.countries")
		if countries != nil {
			countryList := reflect.ValueOf(countries)
			for i := 0; i < countryList.Len(); i++ {
				country := results.Get(fmt.Sprintf("targeting.geo_locations.countries.%d", i)).(string)
				countryName, err := query.FindCountryByAlpha(country)
				if err != nil {
					adTargeting.Locations = append(adTargeting.Locations, country)
				} else {
					adTargeting.Locations = append(adTargeting.Locations, countryName.Name.Common)
				}
			}
		}

		ageMin := results.Get("targeting.age_min")
		if ageMin != nil {
			adTargeting.AgeMin = ageMin.(float64)
		}

		ageMax := results.Get("targeting.age_max")
		if ageMax != nil {
			adTargeting.AgeMax = ageMax.(float64)
		}

		interests := results.Get("targeting.interests")
		if interests != nil {
			adTargeting.Interests = addInterests(interests, results)
		}

		flexibleInterests := results.Get("targeting.flexible_spec.interests")
		if flexibleInterests != nil {
			adTargeting.Interests = addInterests(flexibleInterests, results)
		}
	}

	return adTargeting
}

func addInterests(interests interface{}, results *facebookLib.Result) []string {
	var filteredInterests []string
	interestsList := reflect.ValueOf(interests)
	for i := 0; i < interestsList.Len(); i++ {
		interest := results.Get(fmt.Sprintf("targeting.interests.%d.name", i)).(string)
		filteredInterests = append(filteredInterests, interest)
	}

	return filteredInterests
}
