# PostHttpGatewayRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSsl** | **bool** | 是否自动匹配证书，升级为https，如果开启，由外部服务完成升级 | [optional] [default to null]
**AutoSslConfig** | **bool** | 自动分发证书配置 | [optional] [default to null]
**CertificateId** | **int32** | 证书id | [optional] [default to null]
**ContainerPort** | **int32** | 绑定端口 | [default to null]
**DomainCookie** | **string** | 域名cookie | [optional] [default to null]
**DomainHeader** | **string** | 域名header | [optional] [default to null]
**DomainName** | **string** | 域名 | [default to null]
**DomainPath** | **string** | 域名路径 | [optional] [default to null]
**RuleExtensions** | **[]string** | 规则扩展 | [optional] [default to null]
**ServiceId** | **string** | 应用组件id | [default to null]
**TheWeight** | **int32** |  | [optional] [default to null]
**WhetherOpen** | **bool** | 是否开放 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


