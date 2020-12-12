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
// TeamInfo struct for TeamInfo
type TeamInfo struct {
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 团队拥有者用户
	Creater string `json:"creater,omitempty"`
	// 企业ID
	EnterpriseId string `json:"enterprise_id"`
	// 是否激活
	IsActive bool `json:"is_active,omitempty"`
	// 团队开通的数据中心数量
	RegionNum int32 `json:"region_num,omitempty"`
	// 用户在团队中拥有的角色
	RoleInfos []RoleInfo `json:"role_infos,omitempty"`
	// 团队的组件数量
	ServiceNum int32 `json:"service_num,omitempty"`
	// 团队别名
	TenantAlias string `json:"tenant_alias"`
	// 团队ID
	TenantId string `json:"tenant_id"`
	// 团队名称
	TenantName string `json:"tenant_name"`
}
