package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteRuleAclDto struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	ObjectId string `json:"object_id"`

	// 规则id列表，批量删除规则时传入的id列表。id值可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。
	RuleIds []string `json:"rule_ids"`
}

func (o DeleteRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleAclDto struct{}"
	}

	return strings.Join([]string{"DeleteRuleAclDto", string(data)}, " ")
}
