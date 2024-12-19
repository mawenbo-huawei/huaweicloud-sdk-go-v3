package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupStrategyRequest 备份策略请求信息
type BackupStrategyRequest struct {

	// 策略ID。
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称。添加备份策略时为必选字段。策略名称在4位到92位之间，必须以字母开头，不区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他特殊字符，并且名称唯一。
	PolicyName *string `json:"policy_name,omitempty"`

	// 执行策略。添加备份策略时为必选字段。符合Cron表达式格式。
	BackupStrategy *string `json:"backup_strategy,omitempty"`

	// 备份类型： - full：全量。 - increment：增量。
	BackupType *string `json:"backup_type,omitempty"`

	// 备份级别： - cluster：集群级。
	BackupLevel *string `json:"backup_level,omitempty"`
}

func (o BackupStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyRequest struct{}"
	}

	return strings.Join([]string{"BackupStrategyRequest", string(data)}, " ")
}
