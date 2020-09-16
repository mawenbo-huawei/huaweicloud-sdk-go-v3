/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowHealthMonitorResponse struct {
	// 请求ID。 注：自动生成 。
	RequestId     *string        `json:"request_id,omitempty"`
	Healthmonitor *HealthMonitor `json:"healthmonitor,omitempty"`
}

func (o ShowHealthMonitorResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowHealthMonitorResponse", string(data)}, " ")
}