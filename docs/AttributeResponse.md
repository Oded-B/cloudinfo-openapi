# AttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **string** |  | [optional] 
**AttributeValues** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewAttributeResponse

`func NewAttributeResponse() *AttributeResponse`

NewAttributeResponse instantiates a new AttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeResponseWithDefaults

`func NewAttributeResponseWithDefaults() *AttributeResponse`

NewAttributeResponseWithDefaults instantiates a new AttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *AttributeResponse) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AttributeResponse) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AttributeResponse) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *AttributeResponse) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeValues

`func (o *AttributeResponse) GetAttributeValues() []float64`

GetAttributeValues returns the AttributeValues field if non-nil, zero value otherwise.

### GetAttributeValuesOk

`func (o *AttributeResponse) GetAttributeValuesOk() (*[]float64, bool)`

GetAttributeValuesOk returns a tuple with the AttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValues

`func (o *AttributeResponse) SetAttributeValues(v []float64)`

SetAttributeValues sets AttributeValues field to given value.

### HasAttributeValues

`func (o *AttributeResponse) HasAttributeValues() bool`

HasAttributeValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


