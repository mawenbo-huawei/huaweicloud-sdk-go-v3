package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlCreateReadonlyNodeRequest 只读节点信息
type MysqlCreateReadonlyNodeRequest struct {

	// 指定创建的只读节点故障倒换优先级。  故障倒换优先级的取值范围为1~16，数字越小，优先级越大，即故障倒换时，主节点会优先倒换到优先级高的只读节点上，优先级相同的只读节点选为主节点的概率相同。最多支持9个只读节点设置故障倒换优先级，超过9个的只读节点优先级默认为-1，表示不会参与倒换。可通过修改节点的故障倒换优先级来进行调整。
	Priorities []int32 `json:"priorities"`

	// 创建包周期时可指定，表示是否自动从客户的账户中支付，此字段不影响自动续订的支付方式。  - true，为自动支付，默认该方式。 - false，为手动支付。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`

	// 可用区。可指定可用区创建只读节点，不传该参数时默认为自动选择可用区。  调用[查询数据库规格](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlFlavors.html)获取，其中az_status中的key为availability_zone。  注：指定可用区创建只读节点，可能由于资源不足创建失败。
	AvailabilityZones *[]string `json:"availability_zones,omitempty"`
}

func (o MysqlCreateReadonlyNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlCreateReadonlyNodeRequest struct{}"
	}

	return strings.Join([]string{"MysqlCreateReadonlyNodeRequest", string(data)}, " ")
}
