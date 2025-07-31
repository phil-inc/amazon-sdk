# Quantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float64** | Total quantity of the prescription. | 
**UnitOfMeasureCode** | Pointer to **string** | A NCI Code, NCPDP Drug QuantityUnitOfMeasure Terminology, e.g., C48542. This shape is deprecated: This field will be replaced by &#x60;quantityUnitOfMeasureCode&#x60;. This field will no longer be supported starting from 2025-03-12. | [optional] 
**QuantityUnitOfMeasureCode** | Pointer to [**QuantityUnitOfMeasureCode**](QuantityUnitOfMeasureCode.md) |  | [optional] 
**CodeListQualifier** | Pointer to [**QuantityCodeListQualifier**](QuantityCodeListQualifier.md) |  | [optional] 

## Methods

### NewQuantity

`func NewQuantity(value float64, ) *Quantity`

NewQuantity instantiates a new Quantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantityWithDefaults

`func NewQuantityWithDefaults() *Quantity`

NewQuantityWithDefaults instantiates a new Quantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Quantity) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Quantity) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Quantity) SetValue(v float64)`

SetValue sets Value field to given value.


### GetUnitOfMeasureCode

`func (o *Quantity) GetUnitOfMeasureCode() string`

GetUnitOfMeasureCode returns the UnitOfMeasureCode field if non-nil, zero value otherwise.

### GetUnitOfMeasureCodeOk

`func (o *Quantity) GetUnitOfMeasureCodeOk() (*string, bool)`

GetUnitOfMeasureCodeOk returns a tuple with the UnitOfMeasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureCode

`func (o *Quantity) SetUnitOfMeasureCode(v string)`

SetUnitOfMeasureCode sets UnitOfMeasureCode field to given value.

### HasUnitOfMeasureCode

`func (o *Quantity) HasUnitOfMeasureCode() bool`

HasUnitOfMeasureCode returns a boolean if a field has been set.

### GetQuantityUnitOfMeasureCode

`func (o *Quantity) GetQuantityUnitOfMeasureCode() QuantityUnitOfMeasureCode`

GetQuantityUnitOfMeasureCode returns the QuantityUnitOfMeasureCode field if non-nil, zero value otherwise.

### GetQuantityUnitOfMeasureCodeOk

`func (o *Quantity) GetQuantityUnitOfMeasureCodeOk() (*QuantityUnitOfMeasureCode, bool)`

GetQuantityUnitOfMeasureCodeOk returns a tuple with the QuantityUnitOfMeasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUnitOfMeasureCode

`func (o *Quantity) SetQuantityUnitOfMeasureCode(v QuantityUnitOfMeasureCode)`

SetQuantityUnitOfMeasureCode sets QuantityUnitOfMeasureCode field to given value.

### HasQuantityUnitOfMeasureCode

`func (o *Quantity) HasQuantityUnitOfMeasureCode() bool`

HasQuantityUnitOfMeasureCode returns a boolean if a field has been set.

### GetCodeListQualifier

`func (o *Quantity) GetCodeListQualifier() QuantityCodeListQualifier`

GetCodeListQualifier returns the CodeListQualifier field if non-nil, zero value otherwise.

### GetCodeListQualifierOk

`func (o *Quantity) GetCodeListQualifierOk() (*QuantityCodeListQualifier, bool)`

GetCodeListQualifierOk returns a tuple with the CodeListQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeListQualifier

`func (o *Quantity) SetCodeListQualifier(v QuantityCodeListQualifier)`

SetCodeListQualifier sets CodeListQualifier field to given value.

### HasCodeListQualifier

`func (o *Quantity) HasCodeListQualifier() bool`

HasCodeListQualifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


