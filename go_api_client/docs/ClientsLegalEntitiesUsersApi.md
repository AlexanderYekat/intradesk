# \ClientsLegalEntitiesUsersApi

All URIs are relative to *http://:///settings*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut) | **Put** /api/v4/clients/LegalEntities/Users/BlockEmail/{id} | удаление одного пользователя
[**ApiV4ClientsLegalEntitiesUsersDelete**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersDelete) | **Delete** /api/v4/clients/LegalEntities/Users | удаление списка пользователей
[**ApiV4ClientsLegalEntitiesUsersDeleteIdPut**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersDeleteIdPut) | **Put** /api/v4/clients/LegalEntities/Users/Delete/{id} | удаление одного пользователя
[**ApiV4ClientsLegalEntitiesUsersMergePut**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersMergePut) | **Put** /api/v4/clients/LegalEntities/Users/Merge | объединение пользователей юр лица
[**ApiV4ClientsLegalEntitiesUsersPatch**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersPatch) | **Patch** /api/v4/clients/LegalEntities/Users | 
[**ApiV4ClientsLegalEntitiesUsersPost**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersPost) | **Post** /api/v4/clients/LegalEntities/Users | Создание пользователя
[**ApiV4ClientsLegalEntitiesUsersPut**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersPut) | **Put** /api/v4/clients/LegalEntities/Users | Обновление пользователя клиента
[**ApiV4ClientsLegalEntitiesUsersRestoreIdPut**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersRestoreIdPut) | **Put** /api/v4/clients/LegalEntities/Users/Restore/{id} | восстановление одного пользователя
[**ApiV4ClientsLegalEntitiesUsersRoleHintsGet**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersRoleHintsGet) | **Get** /api/v4/clients/LegalEntities/Users/RoleHints | Поиск ролей с определенным типом. Используется на форме фильтра для пользователей клиентов и сотрудников. Не используется контроллер ролей, так как не хватает прав доступа для общих разделов клиента.
[**ApiV4ClientsLegalEntitiesUsersSearchHintsGet**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersSearchHintsGet) | **Get** /api/v4/clients/LegalEntities/Users/SearchHints | поиск пользователей клиентов
[**ApiV4ClientsLegalEntitiesUsersToEmployeePut**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersToEmployeePut) | **Put** /api/v4/clients/LegalEntities/Users/ToEmployee | перевод пользователя клиента в сотрудника
[**ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut) | **Put** /api/v4/clients/LegalEntities/Users/UnblockEmail/{id} | восстановление одного пользователя
[**ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet) | **Get** /api/v4/clients/LegalEntities/Users/UserDevices/{userId} | получить устройства пользователя (тип и токен)
[**ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet**](ClientsLegalEntitiesUsersApi.md#ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet) | **Get** /api/v4/clients/LegalEntities/Users/ViewUserPhoto/{id} | получение фото пользователя



## ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut(ctx, id).Execute()

удаление одного пользователя

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
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersBlockEmailIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersBlockEmailIdPutRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesUsersDelete

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesUsersDelete(ctx).IntraService5SettingsApiModelsEntitiesDeletionModel18E3(intraService5SettingsApiModelsEntitiesDeletionModel18E3).Execute()

удаление списка пользователей

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
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersDelete(context.Background()).IntraService5SettingsApiModelsEntitiesDeletionModel18E3(intraService5SettingsApiModelsEntitiesDeletionModel18E3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersDelete`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersDeleteRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesUsersDeleteIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesUsersDeleteIdPut(ctx, id).Execute()

удаление одного пользователя

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
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersDeleteIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersDeleteIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersDeleteIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersDeleteIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersDeleteIdPutRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesUsersMergePut

> ApiV4ClientsLegalEntitiesUsersMergePut(ctx).IntraService5SettingsApiModelsClientsMergeClientUsersModelF151(intraService5SettingsApiModelsClientsMergeClientUsersModelF151).Execute()

объединение пользователей юр лица

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
    intraService5SettingsApiModelsClientsMergeClientUsersModelF151 := *openapiclient.NewIntraService5SettingsApiModelsClientsMergeClientUsersModelF151() // IntraService5SettingsApiModelsClientsMergeClientUsersModelF151 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersMergePut(context.Background()).IntraService5SettingsApiModelsClientsMergeClientUsersModelF151(intraService5SettingsApiModelsClientsMergeClientUsersModelF151).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersMergePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersMergePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientsMergeClientUsersModelF151** | [**IntraService5SettingsApiModelsClientsMergeClientUsersModelF151**](IntraService5SettingsApiModelsClientsMergeClientUsersModelF151.md) |  | 

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


## ApiV4ClientsLegalEntitiesUsersPatch

> ApiV4ClientsLegalEntitiesUsersPatch(ctx).Id(id).MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F(microsoftAspNetCoreJsonPatchOperationsOperation4C6F).Execute()



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
    r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersPatch(context.Background()).Id(id).MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F(microsoftAspNetCoreJsonPatchOperationsOperation4C6F).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersPatchRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesUsersPost

> ApiV4ClientsLegalEntitiesUsersPost(ctx).IntraService5SettingsApiModelsClientUserModel38EA(intraService5SettingsApiModelsClientUserModel38EA).Execute()

Создание пользователя

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
    intraService5SettingsApiModelsClientUserModel38EA := *openapiclient.NewIntraService5SettingsApiModelsClientUserModel38EA() // IntraService5SettingsApiModelsClientUserModel38EA | пользователь (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersPost(context.Background()).IntraService5SettingsApiModelsClientUserModel38EA(intraService5SettingsApiModelsClientUserModel38EA).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientUserModel38EA** | [**IntraService5SettingsApiModelsClientUserModel38EA**](IntraService5SettingsApiModelsClientUserModel38EA.md) | пользователь | 

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


## ApiV4ClientsLegalEntitiesUsersPut

> ApiV4ClientsLegalEntitiesUsersPut(ctx).IntraService5SettingsApiModelsClientUserModel38EA(intraService5SettingsApiModelsClientUserModel38EA).Execute()

Обновление пользователя клиента

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
    intraService5SettingsApiModelsClientUserModel38EA := *openapiclient.NewIntraService5SettingsApiModelsClientUserModel38EA() // IntraService5SettingsApiModelsClientUserModel38EA | пользователь (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersPut(context.Background()).IntraService5SettingsApiModelsClientUserModel38EA(intraService5SettingsApiModelsClientUserModel38EA).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientUserModel38EA** | [**IntraService5SettingsApiModelsClientUserModel38EA**](IntraService5SettingsApiModelsClientUserModel38EA.md) | пользователь | 

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


## ApiV4ClientsLegalEntitiesUsersRestoreIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesUsersRestoreIdPut(ctx, id).Execute()

восстановление одного пользователя

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
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersRestoreIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersRestoreIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersRestoreIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersRestoreIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersRestoreIdPutRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesUsersRoleHintsGet

> []IntraService5SettingsApiModelsFilterHintModelECE3 ApiV4ClientsLegalEntitiesUsersRoleHintsGet(ctx).SearchString(searchString).Execute()

Поиск ролей с определенным типом. Используется на форме фильтра для пользователей клиентов и сотрудников. Не используется контроллер ролей, так как не хватает прав доступа для общих разделов клиента.

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
    searchString := "searchString_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersRoleHintsGet(context.Background()).SearchString(searchString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersRoleHintsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersRoleHintsGet`: []IntraService5SettingsApiModelsFilterHintModelECE3
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersRoleHintsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersRoleHintsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **string** |  | 

### Return type

[**[]IntraService5SettingsApiModelsFilterHintModelECE3**](IntraService5SettingsApiModelsFilterHintModelECE3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsLegalEntitiesUsersSearchHintsGet

> []IntraService5SettingsApiModelsSearchHintsModel3054 ApiV4ClientsLegalEntitiesUsersSearchHintsGet(ctx).SearchString(searchString).Top(top).Excludedids(excludedids).Execute()

поиск пользователей клиентов

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
    searchString := "searchString_example" // string |  (optional)
    top := int32(56) // int32 |  (optional)
    excludedids := []int64{int64(123)} // []int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersSearchHintsGet(context.Background()).SearchString(searchString).Top(top).Excludedids(excludedids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersSearchHintsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersSearchHintsGet`: []IntraService5SettingsApiModelsSearchHintsModel3054
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersSearchHintsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersSearchHintsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **string** |  | 
 **top** | **int32** |  | 
 **excludedids** | **[]int64** |  | 

### Return type

[**[]IntraService5SettingsApiModelsSearchHintsModel3054**](IntraService5SettingsApiModelsSearchHintsModel3054.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsLegalEntitiesUsersToEmployeePut

> ApiV4ClientsLegalEntitiesUsersToEmployeePut(ctx).IntraService5SettingsApiModelsClientUserToEmployeeModel1D42(intraService5SettingsApiModelsClientUserToEmployeeModel1D42).Execute()

перевод пользователя клиента в сотрудника

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
    intraService5SettingsApiModelsClientUserToEmployeeModel1D42 := *openapiclient.NewIntraService5SettingsApiModelsClientUserToEmployeeModel1D42() // IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersToEmployeePut(context.Background()).IntraService5SettingsApiModelsClientUserToEmployeeModel1D42(intraService5SettingsApiModelsClientUserToEmployeeModel1D42).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersToEmployeePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersToEmployeePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientUserToEmployeeModel1D42** | [**IntraService5SettingsApiModelsClientUserToEmployeeModel1D42**](IntraService5SettingsApiModelsClientUserToEmployeeModel1D42.md) |  | 

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


## ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut(ctx, id).Execute()

восстановление одного пользователя

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
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersUnblockEmailIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersUnblockEmailIdPutRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet

> []IntraService5SettingsApiModelsUserDeviceDto4EDC ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet(ctx, userId).Execute()

получить устройства пользователя (тип и токен)

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
    userId := int64(789) // int64 | идентификатор пользователя

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet`: []IntraService5SettingsApiModelsUserDeviceDto4EDC
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | идентификатор пользователя | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersUserDevicesUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IntraService5SettingsApiModelsUserDeviceDto4EDC**](IntraService5SettingsApiModelsUserDeviceDto4EDC.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet

> IntraService5SettingsApiModelsUserPhotosB2B1 ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet(ctx, id).Execute()

получение фото пользователя

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
    id := int64(789) // int64 | идентификатор пользователя

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet`: IntraService5SettingsApiModelsUserPhotosB2B1
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesUsersApi.ApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | идентификатор пользователя | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesUsersViewUserPhotoIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntraService5SettingsApiModelsUserPhotosB2B1**](IntraService5SettingsApiModelsUserPhotosB2B1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

