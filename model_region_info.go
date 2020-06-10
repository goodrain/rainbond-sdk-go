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

type RegionInfo struct {
	// 验证文件
	CertFile string `json:"cert_file,omitempty"`
	// 数据中心描述
	Desc string `json:"desc,omitempty"`
	// 数据中心http应用访问根域名
	Httpdomain string `json:"httpdomain"`
	// 验证的key
	KeyFile string `json:"key_file,omitempty"`
	// 数据中心别名
	RegionAlias string `json:"region_alias"`
	// region id
	RegionId string `json:"region_id"`
	// 数据中心名称,不可修改
	RegionName string `json:"region_name"`
	// 数据中心范围 private|public
	Scope string `json:"scope,omitempty"`
	// 数据中心访问ca证书地址
	SslCaCert string `json:"ssl_ca_cert,omitempty"`
	// 数据中心状态 0：编辑中 1:启用 2：停用 3:维护中
	Status string `json:"status"`
	// 数据中心tcp应用访问根域名
	Tcpdomain string `json:"tcpdomain"`
	// 数据中心token
	Token string `json:"token,omitempty"`
	// 数据中心API url
	Url string `json:"url"`
	// 数据中心Websocket url
	Wsurl string `json:"wsurl"`
}
