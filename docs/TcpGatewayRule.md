# TcpGatewayRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **int32** |  | [optional] [readonly] 
**ContainerPort** | **int32** | 容器端口 | [optional] 
**EndPoint** | **string** | ip+port | 
**IsOuterService** | **bool** | 是否已开启对外端口 | [optional] 
**Protocol** | **string** | 服务协议：tcp,udp | [optional] 
**RegionId** | **string** | region id | 
**RuleExtensions** | Pointer to **string** | 扩展功能 | [optional] 
**ServiceAlias** | **string** | 组件别名 | [optional] 
**ServiceId** | **string** | 组件id | 
**ServiceName** | **string** | 组件名 | 
**TcpRuleId** | **string** | tcp_rule_id | 
**TenantId** | **string** | 租户id | 
**Type** | **int32** | 类型（默认：0， 自定义：1） | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


