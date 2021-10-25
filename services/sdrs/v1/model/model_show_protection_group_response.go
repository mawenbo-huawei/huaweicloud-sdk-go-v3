package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProtectionGroupResponse struct {
	ServerGroup    *ShowProtectionGroupParams `json:"server_group,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowProtectionGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectionGroupResponse", string(data)}, " ")
}