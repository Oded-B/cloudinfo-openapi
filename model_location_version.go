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

// LocationVersion LocationVersion struct for displaying version information per location
type LocationVersion struct {
	Default *string `json:"default,omitempty"`
	Location *string `json:"location,omitempty"`
	Versions *[]string `json:"versions,omitempty"`
}

// NewLocationVersion instantiates a new LocationVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationVersion() *LocationVersion {
	this := LocationVersion{}
	return &this
}

// NewLocationVersionWithDefaults instantiates a new LocationVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationVersionWithDefaults() *LocationVersion {
	this := LocationVersion{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *LocationVersion) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationVersion) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *LocationVersion) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *LocationVersion) SetDefault(v string) {
	o.Default = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LocationVersion) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationVersion) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LocationVersion) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *LocationVersion) SetLocation(v string) {
	o.Location = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *LocationVersion) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationVersion) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *LocationVersion) HasVersions() bool {
	if o != nil && o.Versions != nil {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *LocationVersion) SetVersions(v []string) {
	o.Versions = &v
}

func (o LocationVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return json.Marshal(toSerialize)
}

type NullableLocationVersion struct {
	value *LocationVersion
	isSet bool
}

func (v NullableLocationVersion) Get() *LocationVersion {
	return v.value
}

func (v *NullableLocationVersion) Set(val *LocationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationVersion(val *LocationVersion) *NullableLocationVersion {
	return &NullableLocationVersion{value: val, isSet: true}
}

func (v NullableLocationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


