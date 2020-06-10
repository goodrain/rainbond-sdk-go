# HttpGatewayRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **int32** |  | [optional] [default to null]
**AutoSsl** | **bool** | 是否自动匹配证书，升级为https，如果开启，由外部服务完成升级 | [optional] [default to null]
**AutoSslConfig** | **string** | 自动分发证书配置 | [optional] [default to null]
**CertificateId** | **int32** | 证书ID | [optional] [default to null]
**ContainerPort** | **int32** | 容器端口 | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | 创建时间 | [optional] [default to null]
**DomainCookie** | **string** | 域名cookie | [optional] [default to null]
**DomainHeander** | **string** | 域名heander | [optional] [default to null]
**DomainName** | **string** | 域名 | [default to null]
**DomainPath** | **string** | 域名path | [optional] [default to null]
**DomainType** | **string** | 组件域名类型 | [optional] [default to null]
**HttpRuleId** | **string** | http_rule_id | [default to null]
**IsOuterService** | **bool** | 是否已开启对外端口 | [optional] [default to null]
**IsSenior** | **bool** | 是否有高级路由 | [optional] [default to null]
**Protocol** | **string** | 域名类型 http https httptphttps httpandhttps | [optional] [default to null]
**RegionId** | **string** | region id | [default to null]
**RuleExtensions** | **string** | 扩展功能 | [optional] [default to null]
**ServiceAlias** | **string** | 组件别名 | [optional] [default to null]
**ServiceId** | **string** | 组件id | [default to null]
**ServiceName** | **string** | 组件名 | [default to null]
**TenantId** | **string** | 租户id | [default to null]
**TheWeight** | **int32** | 权重 | [optional] [default to null]
**Type_** | **int32** | 类型（默认：0， 自定义：1） | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


