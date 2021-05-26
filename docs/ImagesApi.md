# \ImagesApi

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImages**](ImagesApi.md#GetImages) | **Get** /providers/{provider}/services/{service}/regions/{region}/images | Provides a list of available images on a given provider in a specific region for a service.



## GetImages

> []Image GetImages(ctx, provider, service, region).Gpu(gpu).Version(version).Os(os).PkeVersion(pkeVersion).LatestOnly(latestOnly).Execute()

Provides a list of available images on a given provider in a specific region for a service.

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
    gpu := "gpu_example" // string |  (optional)
    version := "version_example" // string |  (optional)
    os := "os_example" // string |  (optional)
    pkeVersion := "pkeVersion_example" // string |  (optional)
    latestOnly := "latestOnly_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImagesApi.GetImages(context.Background(), provider, service, region).Gpu(gpu).Version(version).Os(os).PkeVersion(pkeVersion).LatestOnly(latestOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.GetImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImages`: []Image
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.GetImages`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gpu** | **string** |  | 
 **version** | **string** |  | 
 **os** | **string** |  | 
 **pkeVersion** | **string** |  | 
 **latestOnly** | **string** |  | 

### Return type

[**[]Image**](Image.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

