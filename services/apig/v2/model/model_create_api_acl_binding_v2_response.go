package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateApiAclBindingV2Response struct {
	// API与ACL的绑定关系列表

	AclBindings    *[]AclApiBindingInfo `json:"acl_bindings,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateApiAclBindingV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateApiAclBindingV2Response struct{}"
	}

	return strings.Join([]string{"CreateApiAclBindingV2Response", string(data)}, " ")
}