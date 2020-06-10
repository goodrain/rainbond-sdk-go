# \OpenapiTeamApi

All URIs are relative to *http://0.0.0.0:7070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1TeamsCertificatesCreate**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesCreate) | **Post** /openapi/v1/teams/{team_id}/certificates | 
[**OpenapiV1TeamsCertificatesDelete**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesDelete) | **Delete** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**OpenapiV1TeamsCertificatesList**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesList) | **Get** /openapi/v1/teams/{team_id}/certificates | 
[**OpenapiV1TeamsCertificatesRead**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesRead) | **Get** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**OpenapiV1TeamsCertificatesUpdate**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesUpdate) | **Put** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**OpenapiV1TeamsCreate**](OpenapiTeamApi.md#OpenapiV1TeamsCreate) | **Post** /openapi/v1/teams | 
[**OpenapiV1TeamsDelete**](OpenapiTeamApi.md#OpenapiV1TeamsDelete) | **Delete** /openapi/v1/teams/{team_id} | 
[**OpenapiV1TeamsList**](OpenapiTeamApi.md#OpenapiV1TeamsList) | **Get** /openapi/v1/teams | 
[**OpenapiV1TeamsRead**](OpenapiTeamApi.md#OpenapiV1TeamsRead) | **Get** /openapi/v1/teams/{team_id} | 
[**OpenapiV1TeamsRegionsServicesList**](OpenapiTeamApi.md#OpenapiV1TeamsRegionsServicesList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/services | 
[**OpenapiV1TeamsUpdate**](OpenapiTeamApi.md#OpenapiV1TeamsUpdate) | **Put** /openapi/v1/teams/{team_id} | 
[**OpenapiV1TeamsUsersCreate**](OpenapiTeamApi.md#OpenapiV1TeamsUsersCreate) | **Post** /openapi/v1/teams/{team_id}/users/{user_id} | 
[**OpenapiV1TeamsUsersDelete**](OpenapiTeamApi.md#OpenapiV1TeamsUsersDelete) | **Delete** /openapi/v1/teams/{team_id}/users/{user_id} | 
[**OpenapiV1TeamsUsersList**](OpenapiTeamApi.md#OpenapiV1TeamsUsersList) | **Get** /openapi/v1/teams/{team_id}/users | 
[**OpenapiV1TeamsUsersUpdate**](OpenapiTeamApi.md#OpenapiV1TeamsUsersUpdate) | **Put** /openapi/v1/teams/{team_id}/users/{user_id} | 


# **OpenapiV1TeamsCertificatesCreate**
> TeamCertificatesR OpenapiV1TeamsCertificatesCreate(ctx, teamId, data)


添加证书

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
  **data** | [**TeamCertificatesC**](TeamCertificatesC.md)|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesDelete**
> OpenapiV1TeamsCertificatesDelete(ctx, certificateId, teamId)


删除证书

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesList**
> TeamCertificatesL OpenapiV1TeamsCertificatesList(ctx, teamId, optional)


获取团队下证书列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
 **optional** | ***OpenapiV1TeamsCertificatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1TeamsCertificatesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| 页码 | 
 **pageSize** | **optional.Float32**| 每页数量 | 

### Return type

[**TeamCertificatesL**](TeamCertificatesL.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesRead**
> TeamCertificatesR OpenapiV1TeamsCertificatesRead(ctx, certificateId, teamId)


获取团队下证书列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesUpdate**
> TeamCertificatesR OpenapiV1TeamsCertificatesUpdate(ctx, certificateId, teamId, data)


更新证书

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 
  **teamId** | **string**|  | 
  **data** | [**TeamCertificatesC**](TeamCertificatesC.md)|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCreate**
> TeamBaseInfo OpenapiV1TeamsCreate(ctx, data)


add team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreateTeamReq**](CreateTeamReq.md)|  | 

### Return type

[**TeamBaseInfo**](TeamBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsDelete**
> OpenapiV1TeamsDelete(ctx, teamId)


删除团队

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsList**
> ListTeamResp OpenapiV1TeamsList(ctx, optional)


获取用户所在团队列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiV1TeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1TeamsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

# **OpenapiV1TeamsRead**
> TeamInfo OpenapiV1TeamsRead(ctx, teamId)


获取团队

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 

### Return type

[**TeamInfo**](TeamInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsServicesList**
> ListRegionTeamServices OpenapiV1TeamsRegionsServicesList(ctx, regionName, teamId, optional)


获取团队下指定数据中心组件信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 
 **optional** | ***OpenapiV1TeamsRegionsServicesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1TeamsRegionsServicesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eid** | **optional.String**| 根据数据中心名称搜索 | 
 **teamName** | **optional.String**| 根据数据中心名称搜索 | 

### Return type

[**ListRegionTeamServices**](ListRegionTeamServices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsUpdate**
> OpenapiV1TeamsUpdate(ctx, teamId, data)


更新团队信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
  **data** | [**UpdateTeamInfoReq**](UpdateTeamInfoReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsUsersCreate**
> OpenapiV1TeamsUsersCreate(ctx, teamId, userId, data)


add team user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
  **userId** | **string**|  | 
  **data** | [**CreateTeamUserReq**](CreateTeamUserReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsUsersDelete**
> OpenapiV1TeamsUsersDelete(ctx, teamId, userId)


将用户从团队中移除

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsUsersList**
> ListTeamUsersResp OpenapiV1TeamsUsersList(ctx, teamId, optional)


获取团队用户列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
 **optional** | ***OpenapiV1TeamsUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiV1TeamsUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| 用户名、邮箱、手机号搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListTeamUsersResp**](ListTeamUsersResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsUsersUpdate**
> OpenapiV1TeamsUsersUpdate(ctx, teamId, userId, data)


update team user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
  **userId** | **string**|  | 
  **data** | [**CreateTeamUserReq**](CreateTeamUserReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

