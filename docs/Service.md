# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsStatic** | Pointer to **bool** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsStatic

`func (o *Service) GetIsStatic() bool`

GetIsStatic returns the IsStatic field if non-nil, zero value otherwise.

### GetIsStaticOk

`func (o *Service) GetIsStaticOk() (*bool, bool)`

GetIsStaticOk returns a tuple with the IsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatic

`func (o *Service) SetIsStatic(v bool)`

SetIsStatic sets IsStatic field to given value.

### HasIsStatic

`func (o *Service) HasIsStatic() bool`

HasIsStatic returns a boolean if a field has been set.

### GetService

`func (o *Service) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Service) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Service) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *Service) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


