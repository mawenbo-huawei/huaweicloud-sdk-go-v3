package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterConfigResponseInfo struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群开启防护节点数量
	ProtectNodeNum *int32 `json:"protect_node_num,omitempty"`

	// 集群防护中断节点数量
	ProtectInterruptNodeNum *int32 `json:"protect_interrupt_node_num,omitempty"`

	// 集群防护中断节点数量
	UnprotectNodeNum *int32 `json:"unprotect_node_num,omitempty"`

	// 集群节点总数
	NodeTotalNum *int32 `json:"node_total_num,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 付费模式 | on_demand 按需 free 免费
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 开启agent自动升级
	AutoUpgrade *bool `json:"auto_upgrade,omitempty"`

	// 优先使用包周期配额；默认0
	PreferPacketCycle *int32 `json:"prefer_packet_cycle,omitempty"`

	// vpc id
	VpcId *string `json:"vpc_id,omitempty"`

	// cce protection type
	ProtectType *string `json:"protect_type,omitempty"`

	// protection status
	ProtectStatus *string `json:"protect_status,omitempty"`

	// cluster type
	ClusterType *string `json:"cluster_type,omitempty"`

	// fail reason
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o ClusterConfigResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterConfigResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterConfigResponseInfo", string(data)}, " ")
}