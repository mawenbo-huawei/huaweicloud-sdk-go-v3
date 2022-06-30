package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeExitEntryPermitResponse struct {

	// 调用成功时表示调用结果。  调用失败时无此字段。
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RecognizeExitEntryPermitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeExitEntryPermitResponse struct{}"
	}

	return strings.Join([]string{"RecognizeExitEntryPermitResponse", string(data)}, " ")
}
