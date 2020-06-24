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
// ComponentMonitorBaseSerializers struct for ComponentMonitorBaseSerializers
type ComponentMonitorBaseSerializers struct {
	// 返回类型
	ResultType string `json:"resultType"`
	Result []MonitorDataSerializers `json:"result,omitempty"`
	Getlist map[string]string `json:"getlist,omitempty"`
}