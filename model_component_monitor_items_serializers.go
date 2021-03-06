/*
 * Rainbond Open API
 *
 * Rainbond open api
 *
 * API version: v1
 * Contact: barnett@goodrain.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ComponentMonitorItemsSerializers struct for ComponentMonitorItemsSerializers
type ComponentMonitorItemsSerializers struct {
	Data ComponentMonitorBaseSerializers `json:"data,omitempty"`
	// 监控项
	MonitorItem string `json:"monitor_item"`
	// 监控状态
	Status string `json:"status,omitempty"`
}
