package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTasksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTasksResponse struct{}"
	}

	return strings.Join([]string{"DeleteTasksResponse", string(data)}, " ")
}