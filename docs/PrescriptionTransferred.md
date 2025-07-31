# PrescriptionTransferred

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrescriptionFillTransferInType** | [**PrescriptionFillTransferInType**](PrescriptionFillTransferInType.md) |  | 
**FirstFillDate** | Pointer to **string** | The date on which the prescription was first filled by the pharmacy, in the format &#39;YYYY-MM-DD&#39;. This date helps track the initial dispensing of the medication to the patient. This field is required if prescription is previously dispensed (PrescriptionFillTransferInType.DISPENSED_BEFORE). | [optional] 
**LastFillDate** | Pointer to **string** | The date on which the prescription was most recently filled by the pharmacy, in the format &#39;YYYY-MM-DD&#39;. This date helps track the latest dispensing of the medication and is useful for monitoring adherence and refill patterns. This field is required if prescription is previously dispensed (PrescriptionFillTransferInType.DISPENSED_BEFORE). | [optional] 
**RefillsRemaining** | Pointer to **float64** | The number of refills left on the prescription including the first fill. This field will be marked as required for both NewRx and previously dispensed Rx starting from 2025-03-12. | [optional] 
**RefillsTransferred** | Pointer to **float64** | The number of refills being transferred to Amazon Pharmacy including the first fill. | [optional] 
**QuantityTransferred** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**QuantityRemaining** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**PreviousDispensedMedications** | Pointer to [**[]PreviousDispensedMedication**](PreviousDispensedMedication.md) | Contains information about previous medication dispenses for the prescription. Each entry represents a medication dispensed by the sending pharmacy. This provides details about actual dispensing when the filled medication differs from what was prescribed. When PrescriptionFillTransferInType.DISPENSED_BEFORE is true, this object must contain either at least one aggregate dispense detail or a list of unique previous dispenses. This field will be marked as required for previously dispensed Rx starting from 2025-03-12. | [optional] 
**PharmacyRxNumber** | Pointer to **string** | Prescription number assigned by the sending pharmacy. This field will be marked as required starting from 2025-03-12. | [optional] 

## Methods

### NewPrescriptionTransferred

`func NewPrescriptionTransferred(prescriptionFillTransferInType PrescriptionFillTransferInType, ) *PrescriptionTransferred`

NewPrescriptionTransferred instantiates a new PrescriptionTransferred object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrescriptionTransferredWithDefaults

`func NewPrescriptionTransferredWithDefaults() *PrescriptionTransferred`

NewPrescriptionTransferredWithDefaults instantiates a new PrescriptionTransferred object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrescriptionFillTransferInType

`func (o *PrescriptionTransferred) GetPrescriptionFillTransferInType() PrescriptionFillTransferInType`

GetPrescriptionFillTransferInType returns the PrescriptionFillTransferInType field if non-nil, zero value otherwise.

### GetPrescriptionFillTransferInTypeOk

`func (o *PrescriptionTransferred) GetPrescriptionFillTransferInTypeOk() (*PrescriptionFillTransferInType, bool)`

GetPrescriptionFillTransferInTypeOk returns a tuple with the PrescriptionFillTransferInType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriptionFillTransferInType

`func (o *PrescriptionTransferred) SetPrescriptionFillTransferInType(v PrescriptionFillTransferInType)`

SetPrescriptionFillTransferInType sets PrescriptionFillTransferInType field to given value.


### GetFirstFillDate

`func (o *PrescriptionTransferred) GetFirstFillDate() string`

GetFirstFillDate returns the FirstFillDate field if non-nil, zero value otherwise.

### GetFirstFillDateOk

`func (o *PrescriptionTransferred) GetFirstFillDateOk() (*string, bool)`

GetFirstFillDateOk returns a tuple with the FirstFillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFillDate

`func (o *PrescriptionTransferred) SetFirstFillDate(v string)`

SetFirstFillDate sets FirstFillDate field to given value.

### HasFirstFillDate

`func (o *PrescriptionTransferred) HasFirstFillDate() bool`

HasFirstFillDate returns a boolean if a field has been set.

### GetLastFillDate

`func (o *PrescriptionTransferred) GetLastFillDate() string`

GetLastFillDate returns the LastFillDate field if non-nil, zero value otherwise.

### GetLastFillDateOk

`func (o *PrescriptionTransferred) GetLastFillDateOk() (*string, bool)`

GetLastFillDateOk returns a tuple with the LastFillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFillDate

`func (o *PrescriptionTransferred) SetLastFillDate(v string)`

SetLastFillDate sets LastFillDate field to given value.

### HasLastFillDate

`func (o *PrescriptionTransferred) HasLastFillDate() bool`

HasLastFillDate returns a boolean if a field has been set.

### GetRefillsRemaining

`func (o *PrescriptionTransferred) GetRefillsRemaining() float64`

GetRefillsRemaining returns the RefillsRemaining field if non-nil, zero value otherwise.

### GetRefillsRemainingOk

`func (o *PrescriptionTransferred) GetRefillsRemainingOk() (*float64, bool)`

GetRefillsRemainingOk returns a tuple with the RefillsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefillsRemaining

`func (o *PrescriptionTransferred) SetRefillsRemaining(v float64)`

SetRefillsRemaining sets RefillsRemaining field to given value.

### HasRefillsRemaining

`func (o *PrescriptionTransferred) HasRefillsRemaining() bool`

HasRefillsRemaining returns a boolean if a field has been set.

### GetRefillsTransferred

`func (o *PrescriptionTransferred) GetRefillsTransferred() float64`

GetRefillsTransferred returns the RefillsTransferred field if non-nil, zero value otherwise.

### GetRefillsTransferredOk

`func (o *PrescriptionTransferred) GetRefillsTransferredOk() (*float64, bool)`

GetRefillsTransferredOk returns a tuple with the RefillsTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefillsTransferred

`func (o *PrescriptionTransferred) SetRefillsTransferred(v float64)`

SetRefillsTransferred sets RefillsTransferred field to given value.

### HasRefillsTransferred

`func (o *PrescriptionTransferred) HasRefillsTransferred() bool`

HasRefillsTransferred returns a boolean if a field has been set.

### GetQuantityTransferred

`func (o *PrescriptionTransferred) GetQuantityTransferred() Quantity`

GetQuantityTransferred returns the QuantityTransferred field if non-nil, zero value otherwise.

### GetQuantityTransferredOk

`func (o *PrescriptionTransferred) GetQuantityTransferredOk() (*Quantity, bool)`

GetQuantityTransferredOk returns a tuple with the QuantityTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityTransferred

`func (o *PrescriptionTransferred) SetQuantityTransferred(v Quantity)`

SetQuantityTransferred sets QuantityTransferred field to given value.

### HasQuantityTransferred

`func (o *PrescriptionTransferred) HasQuantityTransferred() bool`

HasQuantityTransferred returns a boolean if a field has been set.

### GetQuantityRemaining

`func (o *PrescriptionTransferred) GetQuantityRemaining() Quantity`

GetQuantityRemaining returns the QuantityRemaining field if non-nil, zero value otherwise.

### GetQuantityRemainingOk

`func (o *PrescriptionTransferred) GetQuantityRemainingOk() (*Quantity, bool)`

GetQuantityRemainingOk returns a tuple with the QuantityRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRemaining

`func (o *PrescriptionTransferred) SetQuantityRemaining(v Quantity)`

SetQuantityRemaining sets QuantityRemaining field to given value.

### HasQuantityRemaining

`func (o *PrescriptionTransferred) HasQuantityRemaining() bool`

HasQuantityRemaining returns a boolean if a field has been set.

### GetPreviousDispensedMedications

`func (o *PrescriptionTransferred) GetPreviousDispensedMedications() []PreviousDispensedMedication`

GetPreviousDispensedMedications returns the PreviousDispensedMedications field if non-nil, zero value otherwise.

### GetPreviousDispensedMedicationsOk

`func (o *PrescriptionTransferred) GetPreviousDispensedMedicationsOk() (*[]PreviousDispensedMedication, bool)`

GetPreviousDispensedMedicationsOk returns a tuple with the PreviousDispensedMedications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousDispensedMedications

`func (o *PrescriptionTransferred) SetPreviousDispensedMedications(v []PreviousDispensedMedication)`

SetPreviousDispensedMedications sets PreviousDispensedMedications field to given value.

### HasPreviousDispensedMedications

`func (o *PrescriptionTransferred) HasPreviousDispensedMedications() bool`

HasPreviousDispensedMedications returns a boolean if a field has been set.

### GetPharmacyRxNumber

`func (o *PrescriptionTransferred) GetPharmacyRxNumber() string`

GetPharmacyRxNumber returns the PharmacyRxNumber field if non-nil, zero value otherwise.

### GetPharmacyRxNumberOk

`func (o *PrescriptionTransferred) GetPharmacyRxNumberOk() (*string, bool)`

GetPharmacyRxNumberOk returns a tuple with the PharmacyRxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmacyRxNumber

`func (o *PrescriptionTransferred) SetPharmacyRxNumber(v string)`

SetPharmacyRxNumber sets PharmacyRxNumber field to given value.

### HasPharmacyRxNumber

`func (o *PrescriptionTransferred) HasPharmacyRxNumber() bool`

HasPharmacyRxNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


