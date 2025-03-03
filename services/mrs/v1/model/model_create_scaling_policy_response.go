package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingPolicyResponse Response Object
type CreateScalingPolicyResponse struct {

	// 操作结果。 - succeeded：操作成功 - 操作失败时返回的错误码信息如[错误码](https://support.huaweicloud.com/api-mrs/ErrorCode.html)所示。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyResponse", string(data)}, " ")
}
