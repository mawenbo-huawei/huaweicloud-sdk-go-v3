package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMindmapBackupByIdResponse Response Object
type DeleteMindmapBackupByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMindmapBackupByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMindmapBackupByIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteMindmapBackupByIdResponse", string(data)}, " ")
}
