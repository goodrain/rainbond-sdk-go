# \OpenapiGatewayApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1HttpdomainsList**](OpenapiGatewayApi.md#OpenapiV1HttpdomainsList) | **Get** /openapi/v1/httpdomains | 
[**OpenapiV1TeamsRegionsAppsHttpdomainsDelete**](OpenapiGatewayApi.md#OpenapiV1TeamsRegionsAppsHttpdomainsDelete) | **Delete** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 
[**OpenapiV1TeamsRegionsAppsHttpdomainsList**](OpenapiGatewayApi.md#OpenapiV1TeamsRegionsAppsHttpdomainsList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
[**OpenapiV1TeamsRegionsAppsHttpdomainsUpdate**](OpenapiGatewayApi.md#OpenapiV1TeamsRegionsAppsHttpdomainsUpdate) | **Put** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains/{rule_id} | 


# **OpenapiV1HttpdomainsList**
> []EnterpriseHttpGatewayRule OpenapiV1HttpdomainsList(ctx, optional)


获取企业应用http访问策略列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1HttpdomainsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1HttpdomainsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoSsl** | **optional.Bool**| 查询条件，是否为需要自动匹配证书的策略 | 

### Return type

[**[]EnterpriseHttpGatewayRule**](EnterpriseHTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsHttpdomainsDelete**
> HttpGatewayRule OpenapiV1TeamsRegionsAppsHttpdomainsDelete(ctx, appId, regionName, ruleId, teamId)


删除HTTP访问策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **ruleId** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsHttpdomainsList**
> []HttpGatewayRule OpenapiV1TeamsRegionsAppsHttpdomainsList(ctx, appId, regionName, teamId)


获取应用http访问策略列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**[]HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsHttpdomainsUpdate**
> HttpGatewayRule OpenapiV1TeamsRegionsAppsHttpdomainsUpdate(ctx, appId, regionName, ruleId, teamId, data)


更新HTTP访问策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **ruleId** | **string**|  | 
  **teamId** | **string**|  | 
  **data** | [**UpdatePostHttpGatewayRule**](UpdatePostHttpGatewayRule.md)|  | 

### Return type

[**HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

