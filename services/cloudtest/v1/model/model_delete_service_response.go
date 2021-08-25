package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteServiceResponse struct {
	// 接口调用失败错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 接口调用失败错误信息

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteServiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServiceResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceResponse", string(data)}, " ")
}