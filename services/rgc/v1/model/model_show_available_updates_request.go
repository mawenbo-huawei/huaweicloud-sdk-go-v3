package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableUpdatesRequest Request Object
type ShowAvailableUpdatesRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ShowAvailableUpdatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableUpdatesRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableUpdatesRequest", string(data)}, " ")
}