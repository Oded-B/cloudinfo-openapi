# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**Gpu** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *Image) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Image) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Image) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Image) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetGpu

`func (o *Image) GetGpu() bool`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *Image) GetGpuOk() (*bool, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *Image) SetGpu(v bool)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *Image) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Image) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *Image) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Image) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Image) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Image) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *Image) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Image) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Image) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Image) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


