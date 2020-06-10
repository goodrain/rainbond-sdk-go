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

type TeamCertificatesC struct {
	// 证书名称
	Alias string `json:"alias"`
	// 证书key
	Certificate string `json:"certificate"`
	// 证书类型
	CertificateType string `json:"certificate_type"`
	// 证书
	PrivateKey string `json:"private_key"`
}
