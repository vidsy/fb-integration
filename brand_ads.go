package fbintegration

import (
	"fmt"
	"math"
)

type (
	// BrandAds comment pending
	BrandAds struct {
		FBAdAccountID string
		Ads           []Ad
		BrandID       int64
		BrandName     string
	}
)

// NewBrandAds comment pending
func NewBrandAds(adAccountID string, brandID int64, brandName string) BrandAds {
	var ads []Ad

	return BrandAds{
		adAccountID,
		ads,
		brandID,
		brandName,
	}
}

// GenerateParams comment pending
func (ba *BrandAds) GenerateParams(adAccountID string) Params {
	return NewParams(fmt.Sprintf("/act_%s/ads", adAccountID), map[string]interface{}{
		"date_preset": "lifetime",
		"limit":       40,
		"fields":      "fields=id,creative{id, object_id},adset",
	})
}

// Add comment pending
func (ba *BrandAds) Add(ad Ad) {
	ba.Ads = append(ba.Ads, ad)
}

// GenerateSlices comment pending
func (ba *BrandAds) GenerateSlices(size int) []AdBatch {
	var adBatch []AdBatch

	batchAmount := int(math.Ceil(float64(len(ba.Ads)) / float64(size)))

	startIndex := 0
	endIndex := size

	if len(ba.Ads) < size {
		endIndex = len(ba.Ads)
	}

	for i := 0; i < batchAmount; i++ {
		batch := AdBatch{ba.Ads[startIndex:endIndex]}
		adBatch = append(adBatch, batch)

		startIndex += size
		endIndex += size

		if endIndex > len(ba.Ads) {
			endIndex = len(ba.Ads)
		}
	}

	return adBatch
}

// FindByCreativeID comment pending
func (ba *BrandAds) FindByCreativeID(creativeID string) *Ad {
	for i, ad := range ba.Ads {
		if ad.Creative.ID == creativeID {
			return &ba.Ads[i]
		}
	}

	return nil
}

// FindByPostID comment pending
func (ba *BrandAds) FindByPostID(postID string) *Ad {
	for i, ad := range ba.Ads {
		if ad.Creative.PostID == postID {
			return &ba.Ads[i]
		}
	}
	return nil
}

// FindByObjectID comment pending
func (ba *BrandAds) FindByObjectID(objectID string) *Ad {
	for i, ad := range ba.Ads {
		if ad.Post.ObjectID == objectID {
			return &ba.Ads[i]
		}
	}
	return nil
}
