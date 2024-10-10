package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"strings"
)

type UpdateKvRequestBody struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	// 用户自定义的主键名及值。 > 内容字段：主键字段名和值，组合索引多个元素。
	PrimaryKey *bson.D `bson:"primary_key"`

	ConditionExpression *ConditionExpression `bson:"condition_expression,omitempty"`

	UpdateFields *UpdateFields `bson:"update_fields,omitempty"`
}

func (o UpdateKvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKvRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKvRequestBody", string(data)}, " ")
}