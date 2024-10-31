package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpPunishmentRuleRequest Request Object
type DeleteHttpPunishmentRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteHttpPunishmentRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpPunishmentRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpPunishmentRuleRequest", string(data)}, " ")
}
