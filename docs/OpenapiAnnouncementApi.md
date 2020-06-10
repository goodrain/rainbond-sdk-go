# \OpenapiAnnouncementApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1AnnouncementsCreate**](OpenapiAnnouncementApi.md#OpenapiV1AnnouncementsCreate) | **Post** /openapi/v1/announcements | 
[**OpenapiV1AnnouncementsDelete**](OpenapiAnnouncementApi.md#OpenapiV1AnnouncementsDelete) | **Delete** /openapi/v1/announcements/{aid} | 
[**OpenapiV1AnnouncementsList**](OpenapiAnnouncementApi.md#OpenapiV1AnnouncementsList) | **Get** /openapi/v1/announcements | 
[**OpenapiV1AnnouncementsUpdate**](OpenapiAnnouncementApi.md#OpenapiV1AnnouncementsUpdate) | **Put** /openapi/v1/announcements/{aid} | 


# **OpenapiV1AnnouncementsCreate**
> OpenapiV1AnnouncementsCreate(ctx, data)


添加站内信

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreateAncmReqSerilizer**](CreateAncmReqSerilizer.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1AnnouncementsDelete**
> OpenapiV1AnnouncementsDelete(ctx, aid)


删除站内信

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1AnnouncementsList**
> ListAnnouncementResp OpenapiV1AnnouncementsList(ctx, optional)


获取站内信列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1AnnouncementsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1AnnouncementsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListAnnouncementResp**](ListAnnouncementResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1AnnouncementsUpdate**
> OpenapiV1AnnouncementsUpdate(ctx, aid, data)


更新站内信

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **aid** | **string**|  | 
  **data** | [**UpdateAncmReqSerilizer**](UpdateAncmReqSerilizer.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

