# ProductDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Burst** | Pointer to **bool** | Burst this is derived for now | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CpusPerVm** | Pointer to **float64** |  | [optional] 
**CurrentGen** | Pointer to **bool** | CurrentGen signals whether the instance type generation is the current one. Only applies for amazon | [optional] 
**GpusPerVm** | Pointer to **float64** |  | [optional] 
**MemPerVm** | Pointer to **float64** |  | [optional] 
**NtwPerf** | Pointer to **string** |  | [optional] 
**NtwPerfCategory** | Pointer to **string** |  | [optional] 
**OnDemandPrice** | Pointer to **float64** |  | [optional] 
**SpotPrice** | Pointer to [**[]ZonePrice**](ZonePrice.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProductDetails

`func NewProductDetails() *ProductDetails`

NewProductDetails instantiates a new ProductDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDetailsWithDefaults

`func NewProductDetailsWithDefaults() *ProductDetails`

NewProductDetailsWithDefaults instantiates a new ProductDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ProductDetails) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProductDetails) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProductDetails) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProductDetails) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBurst

`func (o *ProductDetails) GetBurst() bool`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *ProductDetails) GetBurstOk() (*bool, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *ProductDetails) SetBurst(v bool)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *ProductDetails) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetCategory

`func (o *ProductDetails) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProductDetails) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProductDetails) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ProductDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCpusPerVm

`func (o *ProductDetails) GetCpusPerVm() float64`

GetCpusPerVm returns the CpusPerVm field if non-nil, zero value otherwise.

### GetCpusPerVmOk

`func (o *ProductDetails) GetCpusPerVmOk() (*float64, bool)`

GetCpusPerVmOk returns a tuple with the CpusPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerVm

`func (o *ProductDetails) SetCpusPerVm(v float64)`

SetCpusPerVm sets CpusPerVm field to given value.

### HasCpusPerVm

`func (o *ProductDetails) HasCpusPerVm() bool`

HasCpusPerVm returns a boolean if a field has been set.

### GetCurrentGen

`func (o *ProductDetails) GetCurrentGen() bool`

GetCurrentGen returns the CurrentGen field if non-nil, zero value otherwise.

### GetCurrentGenOk

`func (o *ProductDetails) GetCurrentGenOk() (*bool, bool)`

GetCurrentGenOk returns a tuple with the CurrentGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentGen

`func (o *ProductDetails) SetCurrentGen(v bool)`

SetCurrentGen sets CurrentGen field to given value.

### HasCurrentGen

`func (o *ProductDetails) HasCurrentGen() bool`

HasCurrentGen returns a boolean if a field has been set.

### GetGpusPerVm

`func (o *ProductDetails) GetGpusPerVm() float64`

GetGpusPerVm returns the GpusPerVm field if non-nil, zero value otherwise.

### GetGpusPerVmOk

`func (o *ProductDetails) GetGpusPerVmOk() (*float64, bool)`

GetGpusPerVmOk returns a tuple with the GpusPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpusPerVm

`func (o *ProductDetails) SetGpusPerVm(v float64)`

SetGpusPerVm sets GpusPerVm field to given value.

### HasGpusPerVm

`func (o *ProductDetails) HasGpusPerVm() bool`

HasGpusPerVm returns a boolean if a field has been set.

### GetMemPerVm

`func (o *ProductDetails) GetMemPerVm() float64`

GetMemPerVm returns the MemPerVm field if non-nil, zero value otherwise.

### GetMemPerVmOk

`func (o *ProductDetails) GetMemPerVmOk() (*float64, bool)`

GetMemPerVmOk returns a tuple with the MemPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemPerVm

`func (o *ProductDetails) SetMemPerVm(v float64)`

SetMemPerVm sets MemPerVm field to given value.

### HasMemPerVm

`func (o *ProductDetails) HasMemPerVm() bool`

HasMemPerVm returns a boolean if a field has been set.

### GetNtwPerf

`func (o *ProductDetails) GetNtwPerf() string`

GetNtwPerf returns the NtwPerf field if non-nil, zero value otherwise.

### GetNtwPerfOk

`func (o *ProductDetails) GetNtwPerfOk() (*string, bool)`

GetNtwPerfOk returns a tuple with the NtwPerf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtwPerf

`func (o *ProductDetails) SetNtwPerf(v string)`

SetNtwPerf sets NtwPerf field to given value.

### HasNtwPerf

`func (o *ProductDetails) HasNtwPerf() bool`

HasNtwPerf returns a boolean if a field has been set.

### GetNtwPerfCategory

`func (o *ProductDetails) GetNtwPerfCategory() string`

GetNtwPerfCategory returns the NtwPerfCategory field if non-nil, zero value otherwise.

### GetNtwPerfCategoryOk

`func (o *ProductDetails) GetNtwPerfCategoryOk() (*string, bool)`

GetNtwPerfCategoryOk returns a tuple with the NtwPerfCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtwPerfCategory

`func (o *ProductDetails) SetNtwPerfCategory(v string)`

SetNtwPerfCategory sets NtwPerfCategory field to given value.

### HasNtwPerfCategory

`func (o *ProductDetails) HasNtwPerfCategory() bool`

HasNtwPerfCategory returns a boolean if a field has been set.

### GetOnDemandPrice

`func (o *ProductDetails) GetOnDemandPrice() float64`

GetOnDemandPrice returns the OnDemandPrice field if non-nil, zero value otherwise.

### GetOnDemandPriceOk

`func (o *ProductDetails) GetOnDemandPriceOk() (*float64, bool)`

GetOnDemandPriceOk returns a tuple with the OnDemandPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandPrice

`func (o *ProductDetails) SetOnDemandPrice(v float64)`

SetOnDemandPrice sets OnDemandPrice field to given value.

### HasOnDemandPrice

`func (o *ProductDetails) HasOnDemandPrice() bool`

HasOnDemandPrice returns a boolean if a field has been set.

### GetSpotPrice

`func (o *ProductDetails) GetSpotPrice() []ZonePrice`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *ProductDetails) GetSpotPriceOk() (*[]ZonePrice, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *ProductDetails) SetSpotPrice(v []ZonePrice)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *ProductDetails) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.

### GetType

`func (o *ProductDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZones

`func (o *ProductDetails) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ProductDetails) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ProductDetails) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ProductDetails) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


