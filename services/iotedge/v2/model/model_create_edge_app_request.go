package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEdgeAppRequest struct {
	Body *CreateEdgeApplicationRequestDto `json:"body,omitempty"`
}

func (o CreateEdgeAppRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeAppRequest", string(data)}, " ")
}