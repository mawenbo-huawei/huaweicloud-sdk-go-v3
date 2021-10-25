package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateProtectedInstanceRequest struct {
	Body *CreateProtectedInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateProtectedInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProtectedInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateProtectedInstanceRequest", string(data)}, " ")
}