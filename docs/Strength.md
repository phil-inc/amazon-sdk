# Strength

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Dosage Strength, e.g., 120. | [optional] 
**FormCode** | Pointer to **string** | NCI Code that specifies the strength form code, e.g., C42998. This shape is deprecated: This field will be replaced by &#x60;strengthFormCode&#x60;. This field will no longer be supported starting from 2025-03-12. | [optional] 
**StrengthFormCode** | Pointer to [**StrengthFormCode**](StrengthFormCode.md) |  | [optional] 
**UnitOfMeasureCode** | Pointer to **string** | A NCI Code that specifies the strength unit, e.g., C48152. This shape is deprecated: This field will be replaced by &#x60;strengthUnitOfMeasureCode&#x60;. This field will no longer be supported starting from 2025-03-12. | [optional] 
**StrengthUnitOfMeasureCode** | Pointer to [**StrengthUnitOfMeasureCode**](StrengthUnitOfMeasureCode.md) |  | [optional] 

## Methods

### NewStrength

`func NewStrength() *Strength`

NewStrength instantiates a new Strength object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrengthWithDefaults

`func NewStrengthWithDefaults() *Strength`

NewStrengthWithDefaults instantiates a new Strength object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Strength) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Strength) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Strength) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Strength) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFormCode

`func (o *Strength) GetFormCode() string`

GetFormCode returns the FormCode field if non-nil, zero value otherwise.

### GetFormCodeOk

`func (o *Strength) GetFormCodeOk() (*string, bool)`

GetFormCodeOk returns a tuple with the FormCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormCode

`func (o *Strength) SetFormCode(v string)`

SetFormCode sets FormCode field to given value.

### HasFormCode

`func (o *Strength) HasFormCode() bool`

HasFormCode returns a boolean if a field has been set.

### GetStrengthFormCode

`func (o *Strength) GetStrengthFormCode() StrengthFormCode`

GetStrengthFormCode returns the StrengthFormCode field if non-nil, zero value otherwise.

### GetStrengthFormCodeOk

`func (o *Strength) GetStrengthFormCodeOk() (*StrengthFormCode, bool)`

GetStrengthFormCodeOk returns a tuple with the StrengthFormCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrengthFormCode

`func (o *Strength) SetStrengthFormCode(v StrengthFormCode)`

SetStrengthFormCode sets StrengthFormCode field to given value.

### HasStrengthFormCode

`func (o *Strength) HasStrengthFormCode() bool`

HasStrengthFormCode returns a boolean if a field has been set.

### GetUnitOfMeasureCode

`func (o *Strength) GetUnitOfMeasureCode() string`

GetUnitOfMeasureCode returns the UnitOfMeasureCode field if non-nil, zero value otherwise.

### GetUnitOfMeasureCodeOk

`func (o *Strength) GetUnitOfMeasureCodeOk() (*string, bool)`

GetUnitOfMeasureCodeOk returns a tuple with the UnitOfMeasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureCode

`func (o *Strength) SetUnitOfMeasureCode(v string)`

SetUnitOfMeasureCode sets UnitOfMeasureCode field to given value.

### HasUnitOfMeasureCode

`func (o *Strength) HasUnitOfMeasureCode() bool`

HasUnitOfMeasureCode returns a boolean if a field has been set.

### GetStrengthUnitOfMeasureCode

`func (o *Strength) GetStrengthUnitOfMeasureCode() StrengthUnitOfMeasureCode`

GetStrengthUnitOfMeasureCode returns the StrengthUnitOfMeasureCode field if non-nil, zero value otherwise.

### GetStrengthUnitOfMeasureCodeOk

`func (o *Strength) GetStrengthUnitOfMeasureCodeOk() (*StrengthUnitOfMeasureCode, bool)`

GetStrengthUnitOfMeasureCodeOk returns a tuple with the StrengthUnitOfMeasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrengthUnitOfMeasureCode

`func (o *Strength) SetStrengthUnitOfMeasureCode(v StrengthUnitOfMeasureCode)`

SetStrengthUnitOfMeasureCode sets StrengthUnitOfMeasureCode field to given value.

### HasStrengthUnitOfMeasureCode

`func (o *Strength) HasStrengthUnitOfMeasureCode() bool`

HasStrengthUnitOfMeasureCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


