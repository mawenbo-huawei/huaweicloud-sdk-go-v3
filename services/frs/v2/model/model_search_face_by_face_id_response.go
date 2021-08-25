package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SearchFaceByFaceIdResponse struct {
	// 查找的人脸集合，详见[SearchFace](zh-cn_topic_0106912071.xml)。 调用失败时无此字段。

	Faces          *[]SearchFace `json:"faces,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o SearchFaceByFaceIdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SearchFaceByFaceIdResponse struct{}"
	}

	return strings.Join([]string{"SearchFaceByFaceIdResponse", string(data)}, " ")
}