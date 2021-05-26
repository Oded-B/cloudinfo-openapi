# \ProductsApi

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProducts**](ProductsApi.md#GetProducts) | **Get** /providers/{provider}/services/{service}/regions/{region}/products | Provides a list of available machine types on a given provider in a specific region.



## GetProducts

> ProductDetailsResponse GetProducts(ctx, provider, service, region).Execute()

Provides a list of available machine types on a given provider in a specific region.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    provider := "provider_example" // string | 
    service := "service_example" // string | 
    region := "region_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetProducts(context.Background(), provider, service, region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProducts`: ProductDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 
**service** | **string** |  | 
**region** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProductDetailsResponse**](ProductDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

