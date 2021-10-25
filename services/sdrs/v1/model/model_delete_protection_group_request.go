package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteProtectionGroupRequest struct {
	// 保护组ID。

	ServerGroupId string `json:"server_group_id"`
}

func (o DeleteProtectionGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectionGroupRequest", string(data)}, " ")
}