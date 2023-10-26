package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var (
	CN_SOUTH_4 = region.NewRegion("cn-south-4",
		"https://cgs.cn-south-4.myhuaweicloud.com")
	AE_AD_1 = region.NewRegion("ae-ad-1",
		"https://hss.ae-ad-1.myhuaweicloud.com")
)

var staticFields = map[string]*region.Region{
	"cn-south-4": CN_SOUTH_4,
	"ae-ad-1":    AE_AD_1,
}

var provider = region.DefaultProviderChain("CGS")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
