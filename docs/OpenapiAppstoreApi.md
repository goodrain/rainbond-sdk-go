# \OpenapiAppstoreApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1AppstoresList**](OpenapiAppstoreApi.md#OpenapiV1AppstoresList) | **Get** /openapi/v1/appstores | 
[**OpenapiV1AppstoresUpdate**](OpenapiAppstoreApi.md#OpenapiV1AppstoresUpdate) | **Put** /openapi/v1/appstores/{eid} | 


# **OpenapiV1AppstoresList**
> ListAppStoreInfosResp OpenapiV1AppstoresList(ctx, optional)


获取应用市场信息列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1AppstoresListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1AppstoresListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| 按企业名称, 企业别名搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListAppStoreInfosResp**](ListAppStoreInfosResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1AppstoresUpdate**
> AppStoreInfo OpenapiV1AppstoresUpdate(ctx, eid, data)


修改应用市场信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eid** | **string**|  | 
  **data** | [**UpdAppStoreInfoReq**](UpdAppStoreInfoReq.md)|  | 

### Return type

[**AppStoreInfo**](AppStoreInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

