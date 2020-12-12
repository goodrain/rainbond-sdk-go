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
// RegionInfoResp struct for RegionInfoResp
type RegionInfoResp struct {
	// api cert file
	CertFile *string `json:"cert_file,omitempty"`
	// 集群描述
	Desc string `json:"desc"`
	// 集群http应用访问根域名
	Httpdomain string `json:"httpdomain,omitempty"`
	// api cert key file
	KeyFile *string `json:"key_file,omitempty"`
	// 集群别名
	RegionAlias string `json:"region_alias"`
	// region id
	RegionId string `json:"region_id,omitempty"`
	// 集群名称
	RegionName string `json:"region_name"`
	// 数据中心范围 private|public
	Scope string `json:"scope,omitempty"`
	// api ca file
	SslCaCert *string `json:"ssl_ca_cert,omitempty"`
	// 集群状态 0：编辑中 1:启用 2：停用 3:维护中
	Status string `json:"status"`
	// 集群tcp应用访问根域名
	Tcpdomain string `json:"tcpdomain,omitempty"`
	// 集群API url
	Url string `json:"url"`
	// 集群Websocket url
	Wsurl string `json:"wsurl,omitempty"`
}
