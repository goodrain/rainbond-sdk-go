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
// TcpGatewayRule struct for TcpGatewayRule
type TcpGatewayRule struct {
	ID int32 `json:"ID,omitempty"`
	// 容器端口
	ContainerPort int32 `json:"container_port,omitempty"`
	// ip+port
	EndPoint string `json:"end_point"`
	// 是否已开启对外端口
	IsOuterService bool `json:"is_outer_service,omitempty"`
	// 服务协议：tcp,udp
	Protocol string `json:"protocol,omitempty"`
	// region id
	RegionId string `json:"region_id"`
	// 扩展功能
	RuleExtensions *string `json:"rule_extensions,omitempty"`
	// 组件别名
	ServiceAlias string `json:"service_alias,omitempty"`
	// 组件id
	ServiceId string `json:"service_id"`
	// 组件名
	ServiceName string `json:"service_name"`
	// tcp_rule_id
	TcpRuleId string `json:"tcp_rule_id"`
	// 租户id
	TenantId string `json:"tenant_id"`
	// 类型（默认：0， 自定义：1）
	Type int32 `json:"type,omitempty"`
}
