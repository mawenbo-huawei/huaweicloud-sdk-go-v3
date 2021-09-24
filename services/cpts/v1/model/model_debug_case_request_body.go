package model

import (
	"encoding/json"

	"strings"
)

// DebugCaseRequestBody
type DebugCaseRequestBody struct {
	// status

	Status int32 `json:"status"`
	// cluster_id

	ClusterId int32 `json:"cluster_id"`
	// cluster_type

	ClusterType string `json:"cluster_type"`
	// without_package

	WithoutPackage int32 `json:"without_package"`
}

func (o DebugCaseRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DebugCaseRequestBody struct{}"
	}

	return strings.Join([]string{"DebugCaseRequestBody", string(data)}, " ")
}