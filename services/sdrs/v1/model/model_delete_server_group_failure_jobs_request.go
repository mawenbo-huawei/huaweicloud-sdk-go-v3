package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteServerGroupFailureJobsRequest struct {
	// 保护组ID。

	ServerGroupId string `json:"server_group_id"`
}

func (o DeleteServerGroupFailureJobsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServerGroupFailureJobsRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupFailureJobsRequest", string(data)}, " ")
}