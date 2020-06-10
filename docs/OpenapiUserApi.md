# \OpenapiUserApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1AdministratorsCreate**](OpenapiUserApi.md#OpenapiV1AdministratorsCreate) | **Post** /openapi/v1/administrators | 
[**OpenapiV1AdministratorsList**](OpenapiUserApi.md#OpenapiV1AdministratorsList) | **Get** /openapi/v1/administrators | 
[**OpenapiV1UserChangepwdUpdate**](OpenapiUserApi.md#OpenapiV1UserChangepwdUpdate) | **Put** /openapi/v1/user/changepwd | 
[**OpenapiV1UsersAdministratorDelete**](OpenapiUserApi.md#OpenapiV1UsersAdministratorDelete) | **Delete** /openapi/v1/users/{user_id}/administrator | 
[**OpenapiV1UsersCreate**](OpenapiUserApi.md#OpenapiV1UsersCreate) | **Post** /openapi/v1/users | 
[**OpenapiV1UsersDelete**](OpenapiUserApi.md#OpenapiV1UsersDelete) | **Delete** /openapi/v1/users/{user_id} | 
[**OpenapiV1UsersList**](OpenapiUserApi.md#OpenapiV1UsersList) | **Get** /openapi/v1/users | 
[**OpenapiV1UsersRead**](OpenapiUserApi.md#OpenapiV1UsersRead) | **Get** /openapi/v1/users/{user_id} | 
[**OpenapiV1UsersTeamsList**](OpenapiUserApi.md#OpenapiV1UsersTeamsList) | **Get** /openapi/v1/users/{user_id}/teams | 
[**OpenapiV1UsersUpdate**](OpenapiUserApi.md#OpenapiV1UsersUpdate) | **Put** /openapi/v1/users/{user_id} | 


# **OpenapiV1AdministratorsCreate**
> OpenapiV1AdministratorsCreate(ctx, data)


添加企业用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreateAdminUserReq**](CreateAdminUserReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1AdministratorsList**
> ListUsersRespView OpenapiV1AdministratorsList(ctx, optional)


获取企业管理员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1AdministratorsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1AdministratorsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **optional.String**| 企业ID | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListUsersRespView**](ListUsersRespView.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UserChangepwdUpdate**
> OpenapiV1UserChangepwdUpdate(ctx, data)


修改用户密码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ChangePassWdUser**](ChangePassWdUser.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UsersAdministratorDelete**
> OpenapiV1UsersAdministratorDelete(ctx, userId)


删除企业管理员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UsersCreate**
> OpenapiV1UsersCreate(ctx, data)


添加普通用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreateUser**](CreateUser.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UsersDelete**
> OpenapiV1UsersDelete(ctx, userId)


删除用户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UsersList**
> ListUsersRespView OpenapiV1UsersList(ctx, optional)


获取用户列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1UsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1UsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| 用户名、邮箱、手机号搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListUsersRespView**](ListUsersRespView.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UsersRead**
> UserInfo OpenapiV1UsersRead(ctx, userId)


根据用户ID获取用户信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UsersTeamsList**
> ListTeamResp OpenapiV1UsersTeamsList(ctx, userId, optional)


获取用户的团队列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***OpenapiV1UsersTeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1UsersTeamsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eid** | **optional.String**| 企业ID | 
 **query** | **optional.String**| 团队名称搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListTeamResp**](ListTeamResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1UsersUpdate**
> OpenapiV1UsersUpdate(ctx, userId, data)


更新用户信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **data** | [**UpdateUser**](UpdateUser.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

