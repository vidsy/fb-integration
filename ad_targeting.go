package fbintegration

import (
	facebookLib "github.com/huandu/facebook"
	"github.com/pariz/gountries"
)

type (
	// AdTargeting comment pending
	AdTargeting struct {
		Locations []string `json:"locations"`
		AgeMin    int      `json:"age_min"`
		AgeMax    int      `json:"age_max"`
		Interests []string `json:"interests"`
		Genders   struct {
			Male   bool `json:"male"`
			Female bool `json:"female"`
		} `json:"genders"`
	}

	// AdTargetingPayload comment pending
	AdTargetingPayload struct {
		Targeting struct {
			GeoLocations struct {
				Countries     []string `facebook:"countries"`
				CountryGroups []string `facebook"country_groups"`
				Cities        []struct {
					Name string `facebook:"name"`
				} `facebook:"cities"`
			} `facebook:"geo_locations"`
			AgeMin    int   `facebook:"age_min"`
			AgeMax    int   `facebook:"age_max"`
			Genders   []int `facebook:"genders"`
			Interests []struct {
				Name string `facebook:"name"`
			} `facebook:"interests"`
			FlexibleSpecs []struct {
				Interests []struct {
					Name string `facebook:"name"`
				} `facebook:"interests"`
			} `facebook:"flexible_spec"`
		} `facebook:"targeting"`
	}
)

func (atp AdTargetingPayload) HasInterests() bool {
	return len(atp.Targeting.Interests) > 0
}

func (atp AdTargetingPayload) HasFlexibleInterests() bool {
	if len(atp.Targeting.FlexibleSpecs) > 0 {
		for _, flexibleSpec := range atp.Targeting.FlexibleSpecs {
			if len(flexibleSpec.Interests) > 0 {
				return true
			}
		}
	}
	return false
}

func (atp AdTargetingPayload) FlexibleInterests() []string {
	var interests []string
	for index, flexibleSpec := range atp.Targeting.FlexibleSpecs {
		if len(flexibleSpec.Interests) > 0 {
			for _, interest := range atp.Targeting.FlexibleSpecs[index].Interests {
				interests = append(interests, interest.Name)
			}
		}
	}

	return interests
}

// NewAdTargetingFromResult comment pending
func NewAdTargetingFromResult(results *facebookLib.Result) AdTargeting {
	query := gountries.New()
	adTargeting := AdTargeting{}

	var adTargetingPayload AdTargetingPayload
	results.DecodeField("", &adTargetingPayload)
	targeting := adTargetingPayload.Targeting

	if adTargetingPayload.HasInterests() {
		for _, interest := range targeting.Interests {
			adTargeting.Interests = append(adTargeting.Interests, interest.Name)
		}
	}

	if adTargetingPayload.HasFlexibleInterests() {
		adTargeting.Interests = append(adTargeting.Interests, adTargetingPayload.FlexibleInterests()...)
	}

	if len(targeting.GeoLocations.Countries) > 0 {
		for _, country := range targeting.GeoLocations.Countries {
			countryName, err := query.FindCountryByAlpha(country)
			if err != nil {
				adTargeting.Locations = append(adTargeting.Locations, country)
			} else {
				adTargeting.Locations = append(adTargeting.Locations, countryName.Name.Common)
			}
		}
	}

	if len(targeting.GeoLocations.Cities) > 0 {
		for _, city := range targeting.GeoLocations.Cities {
			adTargeting.Locations = append(adTargeting.Locations, city.Name)
		}
	}

	adTargeting.AgeMin = targeting.AgeMin
	adTargeting.AgeMax = targeting.AgeMax

	if len(targeting.Genders) > 0 {
		for _, gender := range targeting.Genders {
			switch gender {
			case 1:
				adTargeting.Genders.Male = true
			case 2:
				adTargeting.Genders.Female = true
			}
		}
	} else {
		adTargeting.Genders.Male = true
		adTargeting.Genders.Female = true
	}

	return adTargeting
}
