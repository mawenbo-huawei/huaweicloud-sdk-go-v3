package model

import (
	"encoding/json"

	"strings"
)

// 保护实例挂载信息结构
type ProtectedInstanceAttachment struct {
	// 复制对ID。

	Replication string `json:"replication"`
	// 挂载点。

	Device string `json:"device"`
}

func (o ProtectedInstanceAttachment) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProtectedInstanceAttachment struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceAttachment", string(data)}, " ")
}