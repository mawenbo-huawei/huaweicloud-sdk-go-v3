package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalEipsResponse Response Object
type ListGlobalEipsResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 全局弹性公网IP
	GlobalEips *[]ListBindingGeip `json:"global_eips,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGlobalEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipsResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalEipsResponse", string(data)}, " ")
}
