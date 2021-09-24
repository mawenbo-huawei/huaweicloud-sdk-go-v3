package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddNodeRequest struct {
	// 集群 ID，获取方式请参见[[如何获取接口URI中参数](https://support.huaweicloud.com/api-cce/cce_02_0271.html)](tag:hws)[[如何获取接口URI中参数](https://support.huaweicloud.com/intl/zh-cn/api-cce/cce_02_0271.html)](tag:hws_hk)

	ClusterId string `json:"cluster_id"`

	Body *AddNodeList `json:"body,omitempty"`
}

func (o AddNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddNodeRequest struct{}"
	}

	return strings.Join([]string{"AddNodeRequest", string(data)}, " ")
}