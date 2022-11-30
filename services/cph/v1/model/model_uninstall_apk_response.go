package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UninstallApkResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表。
	Jobs           *[]PhoneJob `json:"jobs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UninstallApkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallApkResponse struct{}"
	}

	return strings.Join([]string{"UninstallApkResponse", string(data)}, " ")
}