/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAlarmTemplateRequest struct {
	ContentType string                          `json:"Content-Type"`
	Body        *CreateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateAlarmTemplateRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateAlarmTemplateRequest", string(data)}, " ")
}