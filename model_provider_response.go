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

// ProviderResponse ProviderResponse is the response used for the requested provider
type ProviderResponse struct {
	Provider *Provider `json:"provider,omitempty"`
}

// NewProviderResponse instantiates a new ProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderResponse() *ProviderResponse {
	this := ProviderResponse{}
	return &this
}

// NewProviderResponseWithDefaults instantiates a new ProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderResponseWithDefaults() *ProviderResponse {
	this := ProviderResponse{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ProviderResponse) GetProvider() Provider {
	if o == nil || o.Provider == nil {
		var ret Provider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderResponse) GetProviderOk() (*Provider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ProviderResponse) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given Provider and assigns it to the Provider field.
func (o *ProviderResponse) SetProvider(v Provider) {
	o.Provider = &v
}

func (o ProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableProviderResponse struct {
	value *ProviderResponse
	isSet bool
}

func (v NullableProviderResponse) Get() *ProviderResponse {
	return v.value
}

func (v *NullableProviderResponse) Set(val *ProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderResponse(val *ProviderResponse) *NullableProviderResponse {
	return &NullableProviderResponse{value: val, isSet: true}
}

func (v NullableProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


