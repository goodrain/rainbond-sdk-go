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

type UpdateAncmReqSerilizer struct {
	// A标签文字
	ATag string `json:"a_tag,omitempty"`
	// a标签跳转地址
	ATagUrl string `json:"a_tag_url,omitempty"`
	// 通知是否启用
	Active bool `json:"active,omitempty"`
	// 通知内容
	Content string `json:"content,omitempty"`
	// 通知的等级
	Level string `json:"level,omitempty"`
	// 通知标题
	Title string `json:"title,omitempty"`
	// 通知类型
	Type_ string `json:"type,omitempty"`
}
