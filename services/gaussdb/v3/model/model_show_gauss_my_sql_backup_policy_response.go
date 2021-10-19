package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGaussMySqlBackupPolicyResponse struct {
	BackupPolicy   *BackupPolicy `json:"backup_policy,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowGaussMySqlBackupPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlBackupPolicyResponse", string(data)}, " ")
}