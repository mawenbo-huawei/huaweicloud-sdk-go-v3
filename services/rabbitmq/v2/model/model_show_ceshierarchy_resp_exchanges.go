package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespExchanges struct {

	// exchange名称。
	Name *string `json:"name,omitempty"`

	// 对应的vhost。
	Vhost *string `json:"vhost,omitempty"`
}

func (o ShowCeshierarchyRespExchanges) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespExchanges struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespExchanges", string(data)}, " ")
}
