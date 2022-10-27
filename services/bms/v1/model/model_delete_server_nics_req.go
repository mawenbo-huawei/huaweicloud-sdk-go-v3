package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 需要解绑的网卡列表信息
type DeleteServerNicsReq struct {

	//
	Nics *[]ServerNics `json:"nics,omitempty"`
}

func (o DeleteServerNicsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerNicsReq struct{}"
	}

	return strings.Join([]string{"DeleteServerNicsReq", string(data)}, " ")
}
