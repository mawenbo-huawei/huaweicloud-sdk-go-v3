package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteResJobResponse struct {

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteResJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteResJobResponse", string(data)}, " ")
}
