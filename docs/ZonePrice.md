# ZonePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float64** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewZonePrice

`func NewZonePrice() *ZonePrice`

NewZonePrice instantiates a new ZonePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonePriceWithDefaults

`func NewZonePriceWithDefaults() *ZonePrice`

NewZonePriceWithDefaults instantiates a new ZonePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *ZonePrice) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ZonePrice) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ZonePrice) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ZonePrice) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetZone

`func (o *ZonePrice) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ZonePrice) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ZonePrice) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ZonePrice) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


