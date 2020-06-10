# \OpenapiEntrepriseApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1ConfigsList**](OpenapiEntrepriseApi.md#OpenapiV1ConfigsList) | **Get** /openapi/v1/configs | 
[**OpenapiV1EnterprisesList**](OpenapiEntrepriseApi.md#OpenapiV1EnterprisesList) | **Get** /openapi/v1/enterprises | 
[**OpenapiV1EnterprisesRead**](OpenapiEntrepriseApi.md#OpenapiV1EnterprisesRead) | **Get** /openapi/v1/enterprises/{eid} | 
[**OpenapiV1EnterprisesResourceList**](OpenapiEntrepriseApi.md#OpenapiV1EnterprisesResourceList) | **Get** /openapi/v1/enterprises/{eid}/resource | 
[**OpenapiV1EnterprisesUpdate**](OpenapiEntrepriseApi.md#OpenapiV1EnterprisesUpdate) | **Put** /openapi/v1/enterprises/{eid} | 


# **OpenapiV1ConfigsList**
> EnterpriseConfigSeralizer OpenapiV1ConfigsList(ctx, )


获取企业配置信息

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EnterpriseConfigSeralizer**](EnterpriseConfigSeralizer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1EnterprisesList**
> ListEntsResp OpenapiV1EnterprisesList(ctx, optional)


获取企业列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1EnterprisesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1EnterprisesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| 按企业名称, 企业别名搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListEntsResp**](ListEntsResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1EnterprisesRead**
> EnterpriseInfo OpenapiV1EnterprisesRead(ctx, eid)


获取企业信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eid** | **string**|  | 

### Return type

[**EnterpriseInfo**](EnterpriseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1EnterprisesResourceList**
> EnterpriseSource OpenapiV1EnterprisesResourceList(ctx, eid)


获取企业使用资源信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eid** | **string**|  | 

### Return type

[**EnterpriseSource**](EnterpriseSource.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1EnterprisesUpdate**
> OpenapiV1EnterprisesUpdate(ctx, eid, eid2, name, alias)


更新企业信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eid** | **string**|  | 
  **eid2** | **string**|  | 
  **name** | **string**|  | 
  **alias** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

