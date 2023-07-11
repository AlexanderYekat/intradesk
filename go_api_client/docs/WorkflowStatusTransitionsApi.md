# \WorkflowStatusTransitionsApi

All URIs are relative to *http://:///settings*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OdataV4WorkflowStatusTransitionsGet**](WorkflowStatusTransitionsApi.md#OdataV4WorkflowStatusTransitionsGet) | **Get** /odata/v4/WorkflowStatusTransitions | 



## OdataV4WorkflowStatusTransitionsGet

> map[string]interface{} OdataV4WorkflowStatusTransitionsGet(ctx).Select_(select_).Expand(expand).Filter(filter).Orderby(orderby).Count(count).Execute()



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
    select_ := "select__example" // string |  (optional)
    expand := "expand_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    orderby := "orderby_example" // string |  (optional)
    count := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowStatusTransitionsApi.OdataV4WorkflowStatusTransitionsGet(context.Background()).Select_(select_).Expand(expand).Filter(filter).Orderby(orderby).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowStatusTransitionsApi.OdataV4WorkflowStatusTransitionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OdataV4WorkflowStatusTransitionsGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowStatusTransitionsApi.OdataV4WorkflowStatusTransitionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOdataV4WorkflowStatusTransitionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** |  | 
 **expand** | **string** |  | 
 **filter** | **string** |  | 
 **orderby** | **string** |  | 
 **count** | **bool** |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

