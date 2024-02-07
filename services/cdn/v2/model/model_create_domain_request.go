package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainRequest Request Object
type CreateDomainRequest struct {
	Body *CreateDomainRequestBody `json:"body,omitempty"`
}

func (o CreateDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainRequest", string(data)}, " ")
}
