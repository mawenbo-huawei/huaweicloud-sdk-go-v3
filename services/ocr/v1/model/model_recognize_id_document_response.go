package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeIdDocumentResponse struct {
	Result         *IdDocumentItem `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RecognizeIdDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeIdDocumentResponse struct{}"
	}

	return strings.Join([]string{"RecognizeIdDocumentResponse", string(data)}, " ")
}