package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadbalancerAutoscalingOption 参数解释：弹性扩缩容配置信息。负载均衡器开启弹性扩缩容后，可根据负载情况自动调整负载均衡器的规格。  约束限制：该字段已经废弃，建议不要使用该字段。  仅当租户白名单放开后该字段才有效，开启弹性扩缩容后，l4_flavor_id和l7_flavor_id表示该负载均衡器实例弹性规格的上限。  [如果传入该字段，创建的弹性实例将会有保底规格并产生对应保底规格的费用。](tag:cmcc,ctc,dc2,dt,dt_test,fcs_arm,fcs_dt,fm,h3dc,hcs,hcso_dt,HEC,hk_sbc,hk_tm,hws,hws_hk,hws_hn,hws_ocb,hws_test,kvm,nohcs,nornal_iec,ocb,sbc,sfsturb,tlf,tlf_test,tm)  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
type CreateLoadbalancerAutoscalingOption struct {

	// 参数解释：负载均衡器实例弹性扩缩容开关。  取值范围：  - true：开启。  - false：关闭。
	Enable bool `json:"enable"`

	// 参数解释：实例弹性扩缩容的最小七层规格ID。  约束限制： 该字段已经废弃，建议不要使用该字段。
	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o CreateLoadbalancerAutoscalingOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerAutoscalingOption struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerAutoscalingOption", string(data)}, " ")
}
