package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogConfigRequest Request Object
type UpdateLogConfigRequest struct {

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取(ListFirewallDetail.xml)。
	FwInstanceId string `json:"fw_instance_id"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *LogConfigDto `json:"body,omitempty"`
}

func (o UpdateLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogConfigRequest", string(data)}, " ")
}
