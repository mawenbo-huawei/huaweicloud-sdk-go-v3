package model

import (
	"encoding/json"

	"strings"
)

// 失败原因对象。
type FailReason struct {
	// 错误码

	FailCode *string `json:"fail_code,omitempty"`
	// 边缘云失败原因列表。包含所边缘云的失败原因。

	FailMessage *string `json:"fail_message,omitempty"`
}

func (o FailReason) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FailReason struct{}"
	}

	return strings.Join([]string{"FailReason", string(data)}, " ")
}