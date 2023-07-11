# \ClientsPhysicalPersonsApi

All URIs are relative to *http://:///settings*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV4ClientsPhysicalPersonsBlockEmailUseridPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsBlockEmailUseridPut) | **Put** /api/v4/clients/PhysicalPersons/BlockEmail/{userid} | блокировка емейла пользователя клиента физ лица
[**ApiV4ClientsPhysicalPersonsDelete**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsDelete) | **Delete** /api/v4/clients/PhysicalPersons | удаление списка клиентов
[**ApiV4ClientsPhysicalPersonsDeleteIdPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsDeleteIdPut) | **Put** /api/v4/clients/PhysicalPersons/Delete/{id} | удаление одного клиента физ лица
[**ApiV4ClientsPhysicalPersonsDeleteUserUseridPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsDeleteUserUseridPut) | **Put** /api/v4/clients/PhysicalPersons/DeleteUser/{userid} | удаление одного пользователя клиента физ лица
[**ApiV4ClientsPhysicalPersonsManagersDictionaryGet**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsManagersDictionaryGet) | **Get** /api/v4/clients/PhysicalPersons/ManagersDictionary | 
[**ApiV4ClientsPhysicalPersonsPatch**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsPatch) | **Patch** /api/v4/clients/PhysicalPersons | Патч физ лица (можно найти как по id, так и по contactPersonId)
[**ApiV4ClientsPhysicalPersonsPost**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsPost) | **Post** /api/v4/clients/PhysicalPersons | создание клиента физ. лица
[**ApiV4ClientsPhysicalPersonsPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsPut) | **Put** /api/v4/clients/PhysicalPersons | Обновление клиента физ. лица
[**ApiV4ClientsPhysicalPersonsRestoreIdPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsRestoreIdPut) | **Put** /api/v4/clients/PhysicalPersons/Restore/{id} | восстановление одного клиента физ лица
[**ApiV4ClientsPhysicalPersonsRestoreUserUseridPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsRestoreUserUseridPut) | **Put** /api/v4/clients/PhysicalPersons/RestoreUser/{userid} | восстановление одного пользователя клиента физ лица
[**ApiV4ClientsPhysicalPersonsToLegalEntityIdPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsToLegalEntityIdPut) | **Put** /api/v4/clients/PhysicalPersons/ToLegalEntity/{id} | 
[**ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut**](ClientsPhysicalPersonsApi.md#ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut) | **Put** /api/v4/clients/PhysicalPersons/UnblockEmail/{userid} | разблокировка емейла пользователя клиента физ лица



## ApiV4ClientsPhysicalPersonsBlockEmailUseridPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsBlockEmailUseridPut(ctx, userid).Execute()

блокировка емейла пользователя клиента физ лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userid := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsBlockEmailUseridPut(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsBlockEmailUseridPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsBlockEmailUseridPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsBlockEmailUseridPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsBlockEmailUseridPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsDelete

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsDelete(ctx).IntraService5SettingsApiModelsEntitiesDeletionModel18E3(intraService5SettingsApiModelsEntitiesDeletionModel18E3).Execute()

удаление списка клиентов

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    intraService5SettingsApiModelsEntitiesDeletionModel18E3 := *openapiclient.NewIntraService5SettingsApiModelsEntitiesDeletionModel18E3() // IntraService5SettingsApiModelsEntitiesDeletionModel18E3 | модель удаления пользователей (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDelete(context.Background()).IntraService5SettingsApiModelsEntitiesDeletionModel18E3(intraService5SettingsApiModelsEntitiesDeletionModel18E3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsDelete`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsEntitiesDeletionModel18E3** | [**IntraService5SettingsApiModelsEntitiesDeletionModel18E3**](IntraService5SettingsApiModelsEntitiesDeletionModel18E3.md) | модель удаления пользователей | 

### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsDeleteIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsDeleteIdPut(ctx, id).Execute()

удаление одного клиента физ лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDeleteIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDeleteIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsDeleteIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDeleteIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsDeleteIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsDeleteUserUseridPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsDeleteUserUseridPut(ctx, userid).Execute()

удаление одного пользователя клиента физ лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userid := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDeleteUserUseridPut(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDeleteUserUseridPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsDeleteUserUseridPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsDeleteUserUseridPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsDeleteUserUseridPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsManagersDictionaryGet

> []IntraService5SettingsApiModelsBaseNamedModel10D7 ApiV4ClientsPhysicalPersonsManagersDictionaryGet(ctx).Search(search).Top(top).IncludeArchived(includeArchived).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    search := "search_example" // string |  (optional)
    top := int32(56) // int32 |  (optional) (default to 50)
    includeArchived := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsManagersDictionaryGet(context.Background()).Search(search).Top(top).IncludeArchived(includeArchived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsManagersDictionaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsManagersDictionaryGet`: []IntraService5SettingsApiModelsBaseNamedModel10D7
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsManagersDictionaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsManagersDictionaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **top** | **int32** |  | [default to 50]
 **includeArchived** | **bool** |  | [default to false]

### Return type

[**[]IntraService5SettingsApiModelsBaseNamedModel10D7**](IntraService5SettingsApiModelsBaseNamedModel10D7.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsPatch

> ApiV4ClientsPhysicalPersonsPatch(ctx).Id(id).MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F(microsoftAspNetCoreJsonPatchOperationsOperation4C6F).Execute()

Патч физ лица (можно найти как по id, так и по contactPersonId)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int64(789) // int64 |  (optional)
    microsoftAspNetCoreJsonPatchOperationsOperation4C6F := []openapiclient.MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F{*openapiclient.NewMicrosoftAspNetCoreJsonPatchOperationsOperation4C6F()} // []MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsPatch(context.Background()).Id(id).MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F(microsoftAspNetCoreJsonPatchOperationsOperation4C6F).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** |  | 
 **microsoftAspNetCoreJsonPatchOperationsOperation4C6F** | [**[]MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F**](MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsPost

> ApiV4ClientsPhysicalPersonsPost(ctx).IntraService5SettingsApiModelsClientPhysicalPersonModelE777(intraService5SettingsApiModelsClientPhysicalPersonModelE777).Execute()

создание клиента физ. лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    intraService5SettingsApiModelsClientPhysicalPersonModelE777 := *openapiclient.NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777() // IntraService5SettingsApiModelsClientPhysicalPersonModelE777 | группа\\компания (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsPost(context.Background()).IntraService5SettingsApiModelsClientPhysicalPersonModelE777(intraService5SettingsApiModelsClientPhysicalPersonModelE777).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientPhysicalPersonModelE777** | [**IntraService5SettingsApiModelsClientPhysicalPersonModelE777**](IntraService5SettingsApiModelsClientPhysicalPersonModelE777.md) | группа\\компания | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsPut

> ApiV4ClientsPhysicalPersonsPut(ctx).IntraService5SettingsApiModelsClientPhysicalPersonModelE777(intraService5SettingsApiModelsClientPhysicalPersonModelE777).Execute()

Обновление клиента физ. лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    intraService5SettingsApiModelsClientPhysicalPersonModelE777 := *openapiclient.NewIntraService5SettingsApiModelsClientPhysicalPersonModelE777() // IntraService5SettingsApiModelsClientPhysicalPersonModelE777 | пользователь (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsPut(context.Background()).IntraService5SettingsApiModelsClientPhysicalPersonModelE777(intraService5SettingsApiModelsClientPhysicalPersonModelE777).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientPhysicalPersonModelE777** | [**IntraService5SettingsApiModelsClientPhysicalPersonModelE777**](IntraService5SettingsApiModelsClientPhysicalPersonModelE777.md) | пользователь | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsRestoreIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsRestoreIdPut(ctx, id).Execute()

восстановление одного клиента физ лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsRestoreIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsRestoreIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsRestoreIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsRestoreIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsRestoreIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsRestoreUserUseridPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsRestoreUserUseridPut(ctx, userid).Execute()

восстановление одного пользователя клиента физ лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userid := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsRestoreUserUseridPut(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsRestoreUserUseridPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsRestoreUserUseridPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsRestoreUserUseridPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsRestoreUserUseridPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsToLegalEntityIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsToLegalEntityIdPut(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsToLegalEntityIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsToLegalEntityIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsToLegalEntityIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsToLegalEntityIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsToLegalEntityIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut(ctx, userid).Execute()

разблокировка емейла пользователя клиента физ лица

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userid := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut(context.Background(), userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsPhysicalPersonsApi.ApiV4ClientsPhysicalPersonsUnblockEmailUseridPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsPhysicalPersonsUnblockEmailUseridPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiCommandsBaseCommandResult2D4E**](IntraService5SettingsApiCommandsBaseCommandResult2D4E.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

