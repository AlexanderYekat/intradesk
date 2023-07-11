# \ClientsApi

All URIs are relative to *http://:///settings*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV4ClientsInitUsersPut**](ClientsApi.md#ApiV4ClientsInitUsersPut) | **Put** /api/v4/clients/InitUsers | 



## ApiV4ClientsInitUsersPut

> IntraService5SettingsApiModelsInitClientUserOutputModel5E23 ApiV4ClientsInitUsersPut(ctx).IntraService5SettingsApiModelsInitClientUserInputModelE478(intraService5SettingsApiModelsInitClientUserInputModelE478).Execute()



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
    intraService5SettingsApiModelsInitClientUserInputModelE478 := *openapiclient.NewIntraService5SettingsApiModelsInitClientUserInputModelE478() // IntraService5SettingsApiModelsInitClientUserInputModelE478 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientsApi.ApiV4ClientsInitUsersPut(context.Background()).IntraService5SettingsApiModelsInitClientUserInputModelE478(intraService5SettingsApiModelsInitClientUserInputModelE478).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientsApi.ApiV4ClientsInitUsersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV4ClientsInitUsersPut`: IntraService5SettingsApiModelsInitClientUserOutputModel5E23
    fmt.Fprintf(os.Stdout, "Response from `ClientsApi.ApiV4ClientsInitUsersPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV4ClientsInitUsersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intraService5SettingsApiModelsInitClientUserInputModelE478** | [**IntraService5SettingsApiModelsInitClientUserInputModelE478**](IntraService5SettingsApiModelsInitClientUserInputModelE478.md) |  | 

### Return type

[**IntraService5SettingsApiModelsInitClientUserOutputModel5E23**](IntraService5SettingsApiModelsInitClientUserOutputModel5E23.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, application/json-patch+json, text/json, application/*+json
- **Accept**: application/json;odata.metadata=minimal;odata.streaming=true, application/json;odata.metadata=minimal;odata.streaming=false, application/json;odata.metadata=minimal, application/json;odata.metadata=full;odata.streaming=true, application/json;odata.metadata=full;odata.streaming=false, application/json;odata.metadata=full, application/json;odata.metadata=none;odata.streaming=true, application/json;odata.metadata=none;odata.streaming=false, application/json;odata.metadata=none, application/json;odata.streaming=true, application/json;odata.streaming=false, application/json, application/xml, application/prs.odatatestxx-odata, text/plain, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

