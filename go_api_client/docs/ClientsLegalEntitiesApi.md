# \ClientsLegalEntitiesApi

All URIs are relative to *http://:///settings*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV4ClientsLegalEntitiesDelete**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesDelete) | **Delete** /api/v4/clients/LegalEntities | удаление списка клиентов
[**ApiV4ClientsLegalEntitiesDeleteIdPut**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesDeleteIdPut) | **Put** /api/v4/clients/LegalEntities/Delete/{id} | удаление одного клиента
[**ApiV4ClientsLegalEntitiesGroupsDictionaryGet**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesGroupsDictionaryGet) | **Get** /api/v4/clients/LegalEntities/GroupsDictionary | 
[**ApiV4ClientsLegalEntitiesImportCheckPost**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesImportCheckPost) | **Post** /api/v4/clients/LegalEntities/ImportCheck | 
[**ApiV4ClientsLegalEntitiesImportPost**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesImportPost) | **Post** /api/v4/clients/LegalEntities/Import | 
[**ApiV4ClientsLegalEntitiesManagersDictionaryGet**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesManagersDictionaryGet) | **Get** /api/v4/clients/LegalEntities/ManagersDictionary | 
[**ApiV4ClientsLegalEntitiesPatch**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesPatch) | **Patch** /api/v4/clients/LegalEntities | 
[**ApiV4ClientsLegalEntitiesPost**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesPost) | **Post** /api/v4/clients/LegalEntities | создание клиента юр. лица
[**ApiV4ClientsLegalEntitiesPut**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesPut) | **Put** /api/v4/clients/LegalEntities | Обновление клиента юр. лица
[**ApiV4ClientsLegalEntitiesRestoreIdPut**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesRestoreIdPut) | **Put** /api/v4/clients/LegalEntities/Restore/{id} | восстановление одного клиента
[**ApiV4ClientsLegalEntitiesSearchHintsGet**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesSearchHintsGet) | **Get** /api/v4/clients/LegalEntities/SearchHints | 
[**ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut**](ClientsLegalEntitiesApi.md#ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut) | **Put** /api/v4/clients/LegalEntities/ToPhysicalPerson/{id} | 



## ApiV4ClientsLegalEntitiesDelete

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesDelete(ctx).IntraService5SettingsApiModelsEntitiesDeletionModel18E3(intraService5SettingsApiModelsEntitiesDeletionModel18E3).Execute()

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
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesDelete(context.Background()).IntraService5SettingsApiModelsEntitiesDeletionModel18E3(intraService5SettingsApiModelsEntitiesDeletionModel18E3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesDelete`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesDeleteRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesDeleteIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesDeleteIdPut(ctx, id).Execute()

удаление одного клиента

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
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesDeleteIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesDeleteIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesDeleteIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesDeleteIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesDeleteIdPutRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesGroupsDictionaryGet

> []IntraService5SettingsApiModelsBaseNamedModel10D7 ApiV4ClientsLegalEntitiesGroupsDictionaryGet(ctx).Search(search).Top(top).IncludeArchived(includeArchived).Excludedids(excludedids).Execute()



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
    excludedids := []int64{int64(123)} // []int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesGroupsDictionaryGet(context.Background()).Search(search).Top(top).IncludeArchived(includeArchived).Excludedids(excludedids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesGroupsDictionaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesGroupsDictionaryGet`: []IntraService5SettingsApiModelsBaseNamedModel10D7
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesGroupsDictionaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesGroupsDictionaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **top** | **int32** |  | [default to 50]
 **includeArchived** | **bool** |  | [default to false]
 **excludedids** | **[]int64** |  | 

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


## ApiV4ClientsLegalEntitiesImportCheckPost

> IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelCheckResultB134 ApiV4ClientsLegalEntitiesImportCheckPost(ctx).IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060(intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060).Execute()



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
    intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060 := []openapiclient.IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060{*openapiclient.NewIntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060()} // []IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesImportCheckPost(context.Background()).IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060(intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesImportCheckPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesImportCheckPost`: IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelCheckResultB134
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesImportCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesImportCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060** | [**[]IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060**](IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060.md) |  | 

### Return type

[**IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelCheckResultB134**](IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelCheckResultB134.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsLegalEntitiesImportPost

> []IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelImportResult6AC0 ApiV4ClientsLegalEntitiesImportPost(ctx).IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060(intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060).Execute()



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
    intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060 := []openapiclient.IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060{*openapiclient.NewIntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060()} // []IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesImportPost(context.Background()).IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060(intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesImportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesImportPost`: []IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelImportResult6AC0
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060** | [**[]IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060**](IntraService5SettingsApiModelsUsersAndGroupsExcelImportModel0060.md) |  | 

### Return type

[**[]IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelImportResult6AC0**](IntraService5SettingsApiModelsUsersAndGroupsExcelImportModelImportResult6AC0.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsLegalEntitiesManagersDictionaryGet

> []IntraService5SettingsApiModelsBaseNamedModel10D7 ApiV4ClientsLegalEntitiesManagersDictionaryGet(ctx).Search(search).Top(top).IncludeArchived(includeArchived).Execute()



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
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesManagersDictionaryGet(context.Background()).Search(search).Top(top).IncludeArchived(includeArchived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesManagersDictionaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesManagersDictionaryGet`: []IntraService5SettingsApiModelsBaseNamedModel10D7
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesManagersDictionaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesManagersDictionaryGetRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesPatch

> ApiV4ClientsLegalEntitiesPatch(ctx).Id(id).MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F(microsoftAspNetCoreJsonPatchOperationsOperation4C6F).Execute()



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
    r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesPatch(context.Background()).Id(id).MicrosoftAspNetCoreJsonPatchOperationsOperation4C6F(microsoftAspNetCoreJsonPatchOperationsOperation4C6F).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesPatchRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesPost

> int64 ApiV4ClientsLegalEntitiesPost(ctx).IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F(intraService5SettingsApiModelsClientLegalEntityCreateModel8A9F).Execute()

создание клиента юр. лица

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
    intraService5SettingsApiModelsClientLegalEntityCreateModel8A9F := *openapiclient.NewIntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F() // IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F | группа\\компания (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesPost(context.Background()).IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F(intraService5SettingsApiModelsClientLegalEntityCreateModel8A9F).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesPost`: int64
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientLegalEntityCreateModel8A9F** | [**IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F**](IntraService5SettingsApiModelsClientLegalEntityCreateModel8A9F.md) | группа\\компания | 

### Return type

**int64**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV4ClientsLegalEntitiesPut

> ApiV4ClientsLegalEntitiesPut(ctx).IntraService5SettingsApiModelsClientLegalEntityModel62ED(intraService5SettingsApiModelsClientLegalEntityModel62ED).Execute()

Обновление клиента юр. лица

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
    intraService5SettingsApiModelsClientLegalEntityModel62ED := *openapiclient.NewIntraService5SettingsApiModelsClientLegalEntityModel62ED() // IntraService5SettingsApiModelsClientLegalEntityModel62ED | пользователь (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesPut(context.Background()).IntraService5SettingsApiModelsClientLegalEntityModel62ED(intraService5SettingsApiModelsClientLegalEntityModel62ED).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsClientLegalEntityModel62ED** | [**IntraService5SettingsApiModelsClientLegalEntityModel62ED**](IntraService5SettingsApiModelsClientLegalEntityModel62ED.md) | пользователь | 

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


## ApiV4ClientsLegalEntitiesRestoreIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesRestoreIdPut(ctx, id).Execute()

восстановление одного клиента

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
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesRestoreIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesRestoreIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesRestoreIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesRestoreIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesRestoreIdPutRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesSearchHintsGet

> []IntraService5SettingsApiModelsSearchHintsModel3054 ApiV4ClientsLegalEntitiesSearchHintsGet(ctx).SearchString(searchString).Top(top).Excludedids(excludedids).Execute()



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
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesSearchHintsGet(context.Background()).SearchString(searchString).Top(top).Excludedids(excludedids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesSearchHintsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesSearchHintsGet`: []IntraService5SettingsApiModelsSearchHintsModel3054
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesSearchHintsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesSearchHintsGetRequest struct via the builder pattern


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


## ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut

> IntraService5SettingsApiCommandsBaseCommandResult2D4E ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut(ctx, id).Execute()



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
    resp, r, err := apiClient.ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut`: IntraService5SettingsApiCommandsBaseCommandResult2D4E
    fmt.Fprintf(os.Stdout, "Response from `ClientsLegalEntitiesApi.ApiV4ClientsLegalEntitiesToPhysicalPersonIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsLegalEntitiesToPhysicalPersonIdPutRequest struct via the builder pattern


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

