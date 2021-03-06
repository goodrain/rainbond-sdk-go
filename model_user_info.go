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
// UserInfo struct for UserInfo
type UserInfo struct {
	// 注册ip
	ClientIp *string `json:"client_ip,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 邮件地址
	Email *string `json:"email,omitempty"`
	// enterprise_id
	EnterpriseId string `json:"enterprise_id,omitempty"`
	// 激活状态
	IsActive bool `json:"is_active,omitempty"`
	// 用户昵称
	NickName string `json:"nick_name,omitempty"`
	// 用户来源
	Origion *string `json:"origion,omitempty"`
	// 手机号码
	Phone *string `json:"phone,omitempty"`
	UserId int32 `json:"user_id"`
}
