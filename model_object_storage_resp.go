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

type ObjectStorageResp struct {
	AccessKey string `json:"access_key"`
	BucketName string `json:"bucket_name"`
	Endpoint string `json:"endpoint"`
	Provider string `json:"provider"`
	SecretKey string `json:"secret_key"`
}
