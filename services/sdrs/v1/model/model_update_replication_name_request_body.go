package model

import (
	"encoding/json"

	"strings"
)

// 更新复制对名称请求体
type UpdateReplicationNameRequestBody struct {
	Replication *UpdateReplicationNameRequestParams `json:"replication"`
}

func (o UpdateReplicationNameRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateReplicationNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateReplicationNameRequestBody", string(data)}, " ")
}