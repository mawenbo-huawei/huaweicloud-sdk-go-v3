package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilterV2 struct {

	// |参数名称：运算符| |参数的约束及描述：0：仅包含1：仅排除 此参数不携带或携带值为null时，不作为筛选条件。|
	Operator *int32 `json:"operator,omitempty"`

	FilterFactor *FilterFactor `json:"filter_factor,omitempty"`
}

func (o FilterV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterV2 struct{}"
	}

	return strings.Join([]string{"FilterV2", string(data)}, " ")
}
