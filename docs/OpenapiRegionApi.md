# \OpenapiRegionApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1RegionsCreate**](OpenapiRegionApi.md#OpenapiV1RegionsCreate) | **Post** /openapi/v1/regions | 
[**OpenapiV1RegionsDelete**](OpenapiRegionApi.md#OpenapiV1RegionsDelete) | **Delete** /openapi/v1/regions/{region_id} | 
[**OpenapiV1RegionsList**](OpenapiRegionApi.md#OpenapiV1RegionsList) | **Get** /openapi/v1/regions | 
[**OpenapiV1RegionsRead**](OpenapiRegionApi.md#OpenapiV1RegionsRead) | **Get** /openapi/v1/regions/{region_id} | 
[**OpenapiV1RegionsStatusUpdate**](OpenapiRegionApi.md#OpenapiV1RegionsStatusUpdate) | **Put** /openapi/v1/regions/{region_id}/status | 
[**OpenapiV1RegionsUpdate**](OpenapiRegionApi.md#OpenapiV1RegionsUpdate) | **Put** /openapi/v1/regions/{region_id} | 


# **OpenapiV1RegionsCreate**
> RegionInfo OpenapiV1RegionsCreate(ctx, data)


添加数据中心

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AddRegionRequest**](AddRegionRequest.md)|  | 

### Return type

[**RegionInfo**](RegionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1RegionsDelete**
> RegionInfo OpenapiV1RegionsDelete(ctx, regionId)


删除指定数据中心元数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **string**|  | 

### Return type

[**RegionInfo**](RegionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1RegionsList**
> []RegionInfoResp OpenapiV1RegionsList(ctx, )


获取全部数据中心列表

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RegionInfoResp**](RegionInfoResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1RegionsRead**
> RegionInfo OpenapiV1RegionsRead(ctx, regionId)


获取指定数据中心数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **string**|  | 

### Return type

[**RegionInfo**](RegionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1RegionsStatusUpdate**
> RegionInfo OpenapiV1RegionsStatusUpdate(ctx, regionId, data)


修改数据中心的状态(上线, 下线, 设为维护)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **string**|  | 
  **data** | [**UpdateRegionStatusReq**](UpdateRegionStatusReq.md)|  | 

### Return type

[**RegionInfo**](RegionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1RegionsUpdate**
> RegionInfo OpenapiV1RegionsUpdate(ctx, regionId, data)


更新指定数据中心元数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **string**|  | 
  **data** | [**UpdateRegionReq**](UpdateRegionReq.md)|  | 

### Return type

[**RegionInfo**](RegionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

