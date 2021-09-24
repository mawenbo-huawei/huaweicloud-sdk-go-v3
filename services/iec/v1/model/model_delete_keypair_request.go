package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteKeypairRequest struct {
	// 密钥名称。

	KeypairName string `json:"keypair_name"`
}

func (o DeleteKeypairRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteKeypairRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeypairRequest", string(data)}, " ")
}