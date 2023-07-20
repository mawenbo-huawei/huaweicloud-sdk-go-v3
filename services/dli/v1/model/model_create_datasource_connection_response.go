package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatasourceConnectionResponse Response Object
type CreateDatasourceConnectionResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId   *string `json:"connection_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatasourceConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateDatasourceConnectionResponse", string(data)}, " ")
}
