package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListServiceSetDetailsResponse struct {

	// 服务组id
	Id *string `json:"id,omitempty"`

	// 服务组名称
	Name *string `json:"name,omitempty"`

	// 服务组描述信息
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListServiceSetDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceSetDetailsResponse", string(data)}, " ")
}