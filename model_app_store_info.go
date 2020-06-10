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

type AppStoreInfo struct {
	// 应用市场API地址
	AccessUrl string `json:"access_url"`
	// 应用市场名称
	AppstoreName string `json:"appstore_name"`
	// 企业别名
	EnterpriseAlias string `json:"enterprise_alias"`
	// 企业ID(联合云ID)
	EnterpriseId string `json:"enterprise_id"`
	// 企业名称
	EnterpriseName string `json:"enterprise_name"`
}
