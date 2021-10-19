package model

import (
	"encoding/json"

	"strings"
)

// 查询资源返回的异常信息
type QueryError struct {
	// 错误编码。

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息。

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o QueryError) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryError struct{}"
	}

	return strings.Join([]string{"QueryError", string(data)}, " ")
}