package model

import (
	"encoding/json"

	"strings"
)

// 更新终端节点路由表请求结构体
type UpdateEndpointRoutetableRequestBody struct {
	// 路由表ID列表。

	Routetables []string `json:"routetables"`
}

func (o UpdateEndpointRoutetableRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEndpointRoutetableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRoutetableRequestBody", string(data)}, " ")
}