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

import (
	"time"
)

type CertificatesR struct {
	// 证书名称
	Alias string `json:"alias"`
	// 证书类型
	CertificateType string `json:"certificate_type"`
	// 过期时间
	EndData time.Time `json:"end_data"`
	// 是否过期
	HasExpired bool `json:"has_expired"`
	// id
	Id int32 `json:"id"`
	// 证书来源
	IssuedBy string `json:"issued_by"`
	// 域名列表
	IssuedTo []string `json:"issued_to"`
}
