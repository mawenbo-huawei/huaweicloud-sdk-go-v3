package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EmbeddedDatabaseWatermark struct {

	// 添加水印的内容
	WatermarkContent string `json:"watermark_content"`

	// 水印密钥
	WatermarkKey string `json:"watermark_key"`

	// 字段类型列表，最大长度100。使用时，至少包含两个字段，一个“primary_key”为true表示主键，一个为false用来嵌入水印
	Columns []Columns `json:"columns"`

	// 嵌入可提取水印的行数限制1500-50000行。内容建议：数据具有随机性，内容重复率低
	Data []map[string]interface{} `json:"data"`
}

func (o EmbeddedDatabaseWatermark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmbeddedDatabaseWatermark struct{}"
	}

	return strings.Join([]string{"EmbeddedDatabaseWatermark", string(data)}, " ")
}
