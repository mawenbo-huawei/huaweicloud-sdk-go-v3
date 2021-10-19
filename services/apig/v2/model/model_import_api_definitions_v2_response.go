package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ImportApiDefinitionsV2Response struct {
	// API分组编号

	GroupId *string `json:"group_id,omitempty"`
	// 导入失败信息

	Failure *[]SwaggerFailureResp `json:"failure,omitempty"`
	// 导入成功信息

	Success *[]SwaggerSuccessResp `json:"success,omitempty"`

	Swagger        *SwaggerInfoResp `json:"swagger,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ImportApiDefinitionsV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsV2Response struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsV2Response", string(data)}, " ")
}