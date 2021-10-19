package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchListEdgeAppsResponse struct {
	// 总记录数

	Count *int32 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`
	// 每页记录数

	EdgeApps       *[]QueryApplicationBriefResponseDto `json:"edge_apps,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o BatchListEdgeAppsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchListEdgeAppsResponse struct{}"
	}

	return strings.Join([]string{"BatchListEdgeAppsResponse", string(data)}, " ")
}