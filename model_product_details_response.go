/*
 * Product Info.
 *
 * The product info application uses the cloud provider APIs to asynchronously fetch and parse instance type attributes and prices, while storing the results in an in memory cache and making it available as structured data through a REST API.
 *
 * API version: 0.0.1
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProductDetailsResponse ProductDetailsResponse Api object to be mapped to product info response
type ProductDetailsResponse struct {
	// Products represents a slice of products for a given provider (VMs with attributes and process)
	Products *[]ProductDetails `json:"products,omitempty"`
	// ScrapingTime represents scraping time for a given provider in milliseconds
	ScrapingTime *string `json:"scrapingTime,omitempty"`
}

// NewProductDetailsResponse instantiates a new ProductDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetailsResponse() *ProductDetailsResponse {
	this := ProductDetailsResponse{}
	return &this
}

// NewProductDetailsResponseWithDefaults instantiates a new ProductDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailsResponseWithDefaults() *ProductDetailsResponse {
	this := ProductDetailsResponse{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *ProductDetailsResponse) GetProducts() []ProductDetails {
	if o == nil || o.Products == nil {
		var ret []ProductDetails
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailsResponse) GetProductsOk() (*[]ProductDetails, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *ProductDetailsResponse) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ProductDetails and assigns it to the Products field.
func (o *ProductDetailsResponse) SetProducts(v []ProductDetails) {
	o.Products = &v
}

// GetScrapingTime returns the ScrapingTime field value if set, zero value otherwise.
func (o *ProductDetailsResponse) GetScrapingTime() string {
	if o == nil || o.ScrapingTime == nil {
		var ret string
		return ret
	}
	return *o.ScrapingTime
}

// GetScrapingTimeOk returns a tuple with the ScrapingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailsResponse) GetScrapingTimeOk() (*string, bool) {
	if o == nil || o.ScrapingTime == nil {
		return nil, false
	}
	return o.ScrapingTime, true
}

// HasScrapingTime returns a boolean if a field has been set.
func (o *ProductDetailsResponse) HasScrapingTime() bool {
	if o != nil && o.ScrapingTime != nil {
		return true
	}

	return false
}

// SetScrapingTime gets a reference to the given string and assigns it to the ScrapingTime field.
func (o *ProductDetailsResponse) SetScrapingTime(v string) {
	o.ScrapingTime = &v
}

func (o ProductDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	if o.ScrapingTime != nil {
		toSerialize["scrapingTime"] = o.ScrapingTime
	}
	return json.Marshal(toSerialize)
}

type NullableProductDetailsResponse struct {
	value *ProductDetailsResponse
	isSet bool
}

func (v NullableProductDetailsResponse) Get() *ProductDetailsResponse {
	return v.value
}

func (v *NullableProductDetailsResponse) Set(val *ProductDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetailsResponse(val *ProductDetailsResponse) *NullableProductDetailsResponse {
	return &NullableProductDetailsResponse{value: val, isSet: true}
}

func (v NullableProductDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

