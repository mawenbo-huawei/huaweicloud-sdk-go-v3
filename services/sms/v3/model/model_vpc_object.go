package model

import (
	"encoding/json"

	"strings"
)

// vpc对象
type VpcObject struct {
	// 虚拟私有云ID，如果是自动创建，填“autoCreate”

	Id string `json:"id"`
	// 虚拟私有云名称

	Name string `json:"name"`
	// VPC的网段，默认192.168.0.0/16

	Cidr *string `json:"cidr,omitempty"`
}

func (o VpcObject) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VpcObject struct{}"
	}

	return strings.Join([]string{"VpcObject", string(data)}, " ")
}