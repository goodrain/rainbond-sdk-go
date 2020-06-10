/*
 * Rainbond Open API
 *
 * Rainbond open api
 *
 * API version: v1
 * Contact: barnett@goodrain.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CreateTeamReq struct {
	// 团队所属人，未提供时默认使用登录用户作为所属人
	Creater int32 `json:"creater,omitempty"`
	// 团队所属企业ID,未提供时默认使用请求用户企业ID
	EnterpriseId string `json:"enterprise_id"`
	// 默认开通的数据中心，未指定则不开通
	Region string `json:"region,omitempty"`
	// 团队名称
	TenantName string `json:"tenant_name"`
}
