package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTestCaseResultResponse struct {
	// 接口调用失败错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 接口调用失败错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorDetail    *ErrorDetailInfo `json:"error_detail,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateTestCaseResultResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTestCaseResultResponse struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseResultResponse", string(data)}, " ")
}