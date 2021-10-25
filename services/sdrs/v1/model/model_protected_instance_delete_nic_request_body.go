package model

import (
	"encoding/json"

	"strings"
)

// 保护实例删除网卡请求体
type ProtectedInstanceDeleteNicRequestBody struct {
	// 网卡Port ID。

	NicId string `json:"nic_id"`
}

func (o ProtectedInstanceDeleteNicRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProtectedInstanceDeleteNicRequestBody struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceDeleteNicRequestBody", string(data)}, " ")
}