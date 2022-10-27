package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://sdg.cn-north-4.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://sdg.cn-north-1.myhuaweicloud.com")
var CN_NORTH_9 = region.NewRegion("cn-north-9", "https://sdg.cn-north-9.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://sdg.cn-east-2.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://sdg.cn-east-3.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://sdg.cn-south-1.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://sdg.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://sdg.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://sdg.ap-southeast-3.myhuaweicloud.com")
var SA_BRAZIL_1 = region.NewRegion("sa-brazil-1", "https://sdg.sa-brazil-1.myhuaweicloud.com")
var LA_SOUTH_2 = region.NewRegion("la-south-2", "https://sdg.la-south-2.myhuaweicloud.com")
var AF_SOUTH_1 = region.NewRegion("af-south-1", "https://sdg.af-south-1.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"cn-north-9":     CN_NORTH_9,
	"cn-east-2":      CN_EAST_2,
	"cn-east-3":      CN_EAST_3,
	"cn-south-1":     CN_SOUTH_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-south-2":     LA_SOUTH_2,
	"af-south-1":     AF_SOUTH_1,
}

var provider = region.DefaultProviderChain("DSC")

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
