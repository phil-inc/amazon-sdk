# PreviousDispensedMedication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrugDescription** | **string** | Name or description of the medication being prescribed. | 
**Ndc** | **string** | The National Drug Code of the medication, without dashes or spaces. | 
**DrugDbCode** | Pointer to [**DrugDbCode**](DrugDbCode.md) |  | [optional] 
**Quantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**DaysSupply** | Pointer to **float32** | Estimated number of days a prescription would last. | [optional] 
**Directions** | Pointer to **string** | Clinical directions meaningful to the patient as written by the prescriber, e.g., &#39;Take 1 tablet by mouth once weekly.&#39; | [optional] 
**LastFillDate** | **string** | A string representing a date in the format &#39;YYYY-MM-DD&#39;. | 
**DispenseAsWritten** | Pointer to **string** | An indicator that the prescriber has instructed to dispense as written (DAW) code. This is typically used when a brand is chosen over a generic alternative based on the direction of the prescriber. 0 means no DAW is specified 1 means by prescriber 2 means a patient/user request 9 means other. | [optional] 

## Methods

### NewPreviousDispensedMedication

`func NewPreviousDispensedMedication(drugDescription string, ndc string, lastFillDate string, ) *PreviousDispensedMedication`

NewPreviousDispensedMedication instantiates a new PreviousDispensedMedication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviousDispensedMedicationWithDefaults

`func NewPreviousDispensedMedicationWithDefaults() *PreviousDispensedMedication`

NewPreviousDispensedMedicationWithDefaults instantiates a new PreviousDispensedMedication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrugDescription

`func (o *PreviousDispensedMedication) GetDrugDescription() string`

GetDrugDescription returns the DrugDescription field if non-nil, zero value otherwise.

### GetDrugDescriptionOk

`func (o *PreviousDispensedMedication) GetDrugDescriptionOk() (*string, bool)`

GetDrugDescriptionOk returns a tuple with the DrugDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugDescription

`func (o *PreviousDispensedMedication) SetDrugDescription(v string)`

SetDrugDescription sets DrugDescription field to given value.


### GetNdc

`func (o *PreviousDispensedMedication) GetNdc() string`

GetNdc returns the Ndc field if non-nil, zero value otherwise.

### GetNdcOk

`func (o *PreviousDispensedMedication) GetNdcOk() (*string, bool)`

GetNdcOk returns a tuple with the Ndc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdc

`func (o *PreviousDispensedMedication) SetNdc(v string)`

SetNdc sets Ndc field to given value.


### GetDrugDbCode

`func (o *PreviousDispensedMedication) GetDrugDbCode() DrugDbCode`

GetDrugDbCode returns the DrugDbCode field if non-nil, zero value otherwise.

### GetDrugDbCodeOk

`func (o *PreviousDispensedMedication) GetDrugDbCodeOk() (*DrugDbCode, bool)`

GetDrugDbCodeOk returns a tuple with the DrugDbCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugDbCode

`func (o *PreviousDispensedMedication) SetDrugDbCode(v DrugDbCode)`

SetDrugDbCode sets DrugDbCode field to given value.

### HasDrugDbCode

`func (o *PreviousDispensedMedication) HasDrugDbCode() bool`

HasDrugDbCode returns a boolean if a field has been set.

### GetQuantity

`func (o *PreviousDispensedMedication) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PreviousDispensedMedication) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PreviousDispensedMedication) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PreviousDispensedMedication) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDaysSupply

`func (o *PreviousDispensedMedication) GetDaysSupply() float32`

GetDaysSupply returns the DaysSupply field if non-nil, zero value otherwise.

### GetDaysSupplyOk

`func (o *PreviousDispensedMedication) GetDaysSupplyOk() (*float32, bool)`

GetDaysSupplyOk returns a tuple with the DaysSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSupply

`func (o *PreviousDispensedMedication) SetDaysSupply(v float32)`

SetDaysSupply sets DaysSupply field to given value.

### HasDaysSupply

`func (o *PreviousDispensedMedication) HasDaysSupply() bool`

HasDaysSupply returns a boolean if a field has been set.

### GetDirections

`func (o *PreviousDispensedMedication) GetDirections() string`

GetDirections returns the Directions field if non-nil, zero value otherwise.

### GetDirectionsOk

`func (o *PreviousDispensedMedication) GetDirectionsOk() (*string, bool)`

GetDirectionsOk returns a tuple with the Directions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirections

`func (o *PreviousDispensedMedication) SetDirections(v string)`

SetDirections sets Directions field to given value.

### HasDirections

`func (o *PreviousDispensedMedication) HasDirections() bool`

HasDirections returns a boolean if a field has been set.

### GetLastFillDate

`func (o *PreviousDispensedMedication) GetLastFillDate() string`

GetLastFillDate returns the LastFillDate field if non-nil, zero value otherwise.

### GetLastFillDateOk

`func (o *PreviousDispensedMedication) GetLastFillDateOk() (*string, bool)`

GetLastFillDateOk returns a tuple with the LastFillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFillDate

`func (o *PreviousDispensedMedication) SetLastFillDate(v string)`

SetLastFillDate sets LastFillDate field to given value.


### GetDispenseAsWritten

`func (o *PreviousDispensedMedication) GetDispenseAsWritten() string`

GetDispenseAsWritten returns the DispenseAsWritten field if non-nil, zero value otherwise.

### GetDispenseAsWrittenOk

`func (o *PreviousDispensedMedication) GetDispenseAsWrittenOk() (*string, bool)`

GetDispenseAsWrittenOk returns a tuple with the DispenseAsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispenseAsWritten

`func (o *PreviousDispensedMedication) SetDispenseAsWritten(v string)`

SetDispenseAsWritten sets DispenseAsWritten field to given value.

### HasDispenseAsWritten

`func (o *PreviousDispensedMedication) HasDispenseAsWritten() bool`

HasDispenseAsWritten returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


