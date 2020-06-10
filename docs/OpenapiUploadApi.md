# \OpenapiUploadApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1UploadFileCreate**](OpenapiUploadApi.md#OpenapiV1UploadFileCreate) | **Post** /openapi/v1/upload-file | 


# **OpenapiV1UploadFileCreate**
> UploadFileResp OpenapiV1UploadFileCreate(ctx, optional)


上传文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1UploadFileCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1UploadFileCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File**| 文件 | 

### Return type

[**UploadFileResp**](UploadFileResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

