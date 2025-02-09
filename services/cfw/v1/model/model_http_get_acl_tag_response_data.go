package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpGetAclTagResponseData struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 规则标签总数
	Total *int32 `json:"total,omitempty"`

	// 规则标签列表
	Records *[]TagsVo `json:"records,omitempty"`
}

func (o HttpGetAclTagResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpGetAclTagResponseData struct{}"
	}

	return strings.Join([]string{"HttpGetAclTagResponseData", string(data)}, " ")
}
