# ExistingMedication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ndc** | **string** | The National Drug Code of the medication, without dashes or spaces. | 
**Rxcui** | Pointer to **string** | RxNorm concept unique identifier | [optional] 
**Description** | Pointer to **string** | Name or description of the drug. | [optional] 
**Strength** | Pointer to [**Strength**](Strength.md) |  | [optional] 
**Quantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 

## Methods

### NewExistingMedication

`func NewExistingMedication(ndc string, ) *ExistingMedication`

NewExistingMedication instantiates a new ExistingMedication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExistingMedicationWithDefaults

`func NewExistingMedicationWithDefaults() *ExistingMedication`

NewExistingMedicationWithDefaults instantiates a new ExistingMedication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNdc

`func (o *ExistingMedication) GetNdc() string`

GetNdc returns the Ndc field if non-nil, zero value otherwise.

### GetNdcOk

`func (o *ExistingMedication) GetNdcOk() (*string, bool)`

GetNdcOk returns a tuple with the Ndc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdc

`func (o *ExistingMedication) SetNdc(v string)`

SetNdc sets Ndc field to given value.


### GetRxcui

`func (o *ExistingMedication) GetRxcui() string`

GetRxcui returns the Rxcui field if non-nil, zero value otherwise.

### GetRxcuiOk

`func (o *ExistingMedication) GetRxcuiOk() (*string, bool)`

GetRxcuiOk returns a tuple with the Rxcui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxcui

`func (o *ExistingMedication) SetRxcui(v string)`

SetRxcui sets Rxcui field to given value.

### HasRxcui

`func (o *ExistingMedication) HasRxcui() bool`

HasRxcui returns a boolean if a field has been set.

### GetDescription

`func (o *ExistingMedication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExistingMedication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExistingMedication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExistingMedication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStrength

`func (o *ExistingMedication) GetStrength() Strength`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ExistingMedication) GetStrengthOk() (*Strength, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ExistingMedication) SetStrength(v Strength)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *ExistingMedication) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetQuantity

`func (o *ExistingMedication) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ExistingMedication) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ExistingMedication) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ExistingMedication) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


