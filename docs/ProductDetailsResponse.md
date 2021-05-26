# ProductDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]ProductDetails**](ProductDetails.md) | Products represents a slice of products for a given provider (VMs with attributes and process) | [optional] 
**ScrapingTime** | Pointer to **string** | ScrapingTime represents scraping time for a given provider in milliseconds | [optional] 

## Methods

### NewProductDetailsResponse

`func NewProductDetailsResponse() *ProductDetailsResponse`

NewProductDetailsResponse instantiates a new ProductDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDetailsResponseWithDefaults

`func NewProductDetailsResponseWithDefaults() *ProductDetailsResponse`

NewProductDetailsResponseWithDefaults instantiates a new ProductDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *ProductDetailsResponse) GetProducts() []ProductDetails`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductDetailsResponse) GetProductsOk() (*[]ProductDetails, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductDetailsResponse) SetProducts(v []ProductDetails)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ProductDetailsResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetScrapingTime

`func (o *ProductDetailsResponse) GetScrapingTime() string`

GetScrapingTime returns the ScrapingTime field if non-nil, zero value otherwise.

### GetScrapingTimeOk

`func (o *ProductDetailsResponse) GetScrapingTimeOk() (*string, bool)`

GetScrapingTimeOk returns a tuple with the ScrapingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapingTime

`func (o *ProductDetailsResponse) SetScrapingTime(v string)`

SetScrapingTime sets ScrapingTime field to given value.

### HasScrapingTime

`func (o *ProductDetailsResponse) HasScrapingTime() bool`

HasScrapingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


