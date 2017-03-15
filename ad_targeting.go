package fbintegration

import (
	facebookLib "github.com/huandu/facebook"

	"github.com/pariz/gountries"
)

type (
	// AdTargeting comment pending
	AdTargeting struct {
		AgeMin             int      `json:"age_min"`
		AgeMax             int      `json:"age_max"`
		DevicePlatforms    []string `json:"device_platforms"`
		PublisherPlatforms []string `json:"publisher_platforms"`
		InstagramPositions []string `json:"instagram_positions"`
		FacebookPositions  []string `json:"facebook_postitions"`
		Genders            struct {
			Male   bool `json:"male"`
			Female bool `json:"female"`
		} `json:"genders"`
		Interests []string `json:"interests"`
		Locations []string `json:"locations"`
	}

	// AdTargetingPayload comment pending
	AdTargetingPayload struct {
		Targeting struct {
			AgeMin             int      `facebook:"age_min"`
			AgeMax             int      `facebook:"age_max"`
			DevicePlatforms    []string `facebook:"device_platforms"`
			PublisherPlatforms []string `facebook:"publisher_platforms"`
			InstagramPositions []string `facebook:"instagram_positions"`
			FacebookPositions  []string `facebook:"facebook_postitions"`
			FlexibleSpec       []struct {
				Interests []struct {
					Name string `facebook:"name"`
				} `facebook:"interests"`
			} `facebook:"flexible_spec"`

			Genders      []int `facebook:"genders"`
			GeoLocations struct {
				Countries     []string `facebook:"countries"`
				CountryGroups []string `facebook:"country_groups"`
				Cities        []struct {
					Name string `facebook:"name"`
				} `facebook:"cities"`
			} `facebook:"geo_locations"`
			Interests []struct {
				Name string `facebook:"name"`
			} `facebook:"interests"`
		} `facebook:"targeting"`
	}
)

// AdType returns the ad type based on publisher platform and position.
func (at AdTargeting) AdType() string {
	if len(at.InstagramPositions) > 0 {
		for _, position := range at.InstagramPositions {
			if position == "story" {
				return "instagram_story"
			}
		}
	}

	return "facebook"
}

// HasInterests comment pending
func (atp AdTargetingPayload) HasInterests() bool {
	return len(atp.Targeting.Interests) > 0
}

// HasFlexibleInterests comment pending
func (atp AdTargetingPayload) HasFlexibleInterests() bool {
	if len(atp.Targeting.FlexibleSpec) > 0 {
		for _, flexibleSpec := range atp.Targeting.FlexibleSpec {
			if len(flexibleSpec.Interests) > 0 {
				return true
			}
		}
	}
	return false
}

// FlexibleInterests comment pending
func (atp AdTargetingPayload) FlexibleInterests() []string {
	var interests []string
	for index, flexibleSpec := range atp.Targeting.FlexibleSpec {
		if len(flexibleSpec.Interests) > 0 {
			for _, interest := range atp.Targeting.FlexibleSpec[index].Interests {
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

	adTargeting.DevicePlatforms = adTargetingPayload.Targeting.DevicePlatforms
	adTargeting.PublisherPlatforms = adTargetingPayload.Targeting.PublisherPlatforms
	adTargeting.FacebookPositions = adTargetingPayload.Targeting.FacebookPositions
	adTargeting.InstagramPositions = adTargetingPayload.Targeting.InstagramPositions

	return adTargeting
}
