# \OpenapiAuthApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1AuthTokenCreate**](OpenapiAuthApi.md#OpenapiV1AuthTokenCreate) | **Post** /openapi/v1/auth-token | 


# **OpenapiV1AuthTokenCreate**
> Token OpenapiV1AuthTokenCreate(ctx, data)


企业管理员账号密码获取API-Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AuthRequest**](AuthRequest.md)|  | 

### Return type

[**Token**](Token.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

