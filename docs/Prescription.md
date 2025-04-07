# Prescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**PrescriptionStatus**](PrescriptionStatus.md) |  | 
**StatusDescription** | **string** | A more detailed explanation of the prescription&#39;s status. | 
**AmazonPharmacy** | [**Pharmacy**](Pharmacy.md) |  | 
**DrugUnits** | Pointer to **string** | Unit of the drug quantity, e.g., Caps/Tabs/ML/Dev. | [optional] 
**Ndc** | Pointer to **string** | The National Drug Code of the medication, without dashes or spaces. | [optional] 
**RefillsAuthorized** | Pointer to **float32** | Number of refills authorized on the Prescription. | [optional] 
**RefillsRemaining** | Pointer to **float32** | Number of refills remaining on the Prescription. | [optional] 
**FillsRemaining** | Pointer to **float32** | Number of total fills (first fill + refill) remaining on the Prescription. | [optional] 
**DaysSupply** | Pointer to **float32** | The number of days the authorizedQuantityPerFill will supply, mapped to Prescription&#39;s daysSupply. | [optional] 
**RxNumber** | Pointer to **string** | The pharmacy-assigned number for this prescription. | [optional] 
**TransferRequestStatus** | Pointer to [**TransferRequestStatus**](TransferRequestStatus.md) |  | [optional] 
**TotalQuantityRemaining** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**Fills** | Pointer to [**[]PrescriptionFill**](PrescriptionFill.md) | A list of fills for this prescription | [optional] 
**Prescriber** | Pointer to [**Prescriber**](Prescriber.md) |  | [optional] 

## Methods

### NewPrescription

`func NewPrescription(status PrescriptionStatus, statusDescription string, amazonPharmacy Pharmacy, ) *Prescription`

NewPrescription instantiates a new Prescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrescriptionWithDefaults

`func NewPrescriptionWithDefaults() *Prescription`

NewPrescriptionWithDefaults instantiates a new Prescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Prescription) GetStatus() PrescriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Prescription) GetStatusOk() (*PrescriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Prescription) SetStatus(v PrescriptionStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *Prescription) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *Prescription) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *Prescription) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetAmazonPharmacy

`func (o *Prescription) GetAmazonPharmacy() Pharmacy`

GetAmazonPharmacy returns the AmazonPharmacy field if non-nil, zero value otherwise.

### GetAmazonPharmacyOk

`func (o *Prescription) GetAmazonPharmacyOk() (*Pharmacy, bool)`

GetAmazonPharmacyOk returns a tuple with the AmazonPharmacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonPharmacy

`func (o *Prescription) SetAmazonPharmacy(v Pharmacy)`

SetAmazonPharmacy sets AmazonPharmacy field to given value.


### GetDrugUnits

`func (o *Prescription) GetDrugUnits() string`

GetDrugUnits returns the DrugUnits field if non-nil, zero value otherwise.

### GetDrugUnitsOk

`func (o *Prescription) GetDrugUnitsOk() (*string, bool)`

GetDrugUnitsOk returns a tuple with the DrugUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugUnits

`func (o *Prescription) SetDrugUnits(v string)`

SetDrugUnits sets DrugUnits field to given value.

### HasDrugUnits

`func (o *Prescription) HasDrugUnits() bool`

HasDrugUnits returns a boolean if a field has been set.

### GetNdc

`func (o *Prescription) GetNdc() string`

GetNdc returns the Ndc field if non-nil, zero value otherwise.

### GetNdcOk

`func (o *Prescription) GetNdcOk() (*string, bool)`

GetNdcOk returns a tuple with the Ndc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdc

`func (o *Prescription) SetNdc(v string)`

SetNdc sets Ndc field to given value.

### HasNdc

`func (o *Prescription) HasNdc() bool`

HasNdc returns a boolean if a field has been set.

### GetRefillsAuthorized

`func (o *Prescription) GetRefillsAuthorized() float32`

GetRefillsAuthorized returns the RefillsAuthorized field if non-nil, zero value otherwise.

### GetRefillsAuthorizedOk

`func (o *Prescription) GetRefillsAuthorizedOk() (*float32, bool)`

GetRefillsAuthorizedOk returns a tuple with the RefillsAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefillsAuthorized

`func (o *Prescription) SetRefillsAuthorized(v float32)`

SetRefillsAuthorized sets RefillsAuthorized field to given value.

### HasRefillsAuthorized

`func (o *Prescription) HasRefillsAuthorized() bool`

HasRefillsAuthorized returns a boolean if a field has been set.

### GetRefillsRemaining

`func (o *Prescription) GetRefillsRemaining() float32`

GetRefillsRemaining returns the RefillsRemaining field if non-nil, zero value otherwise.

### GetRefillsRemainingOk

`func (o *Prescription) GetRefillsRemainingOk() (*float32, bool)`

GetRefillsRemainingOk returns a tuple with the RefillsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefillsRemaining

`func (o *Prescription) SetRefillsRemaining(v float32)`

SetRefillsRemaining sets RefillsRemaining field to given value.

### HasRefillsRemaining

`func (o *Prescription) HasRefillsRemaining() bool`

HasRefillsRemaining returns a boolean if a field has been set.

### GetFillsRemaining

`func (o *Prescription) GetFillsRemaining() float32`

GetFillsRemaining returns the FillsRemaining field if non-nil, zero value otherwise.

### GetFillsRemainingOk

`func (o *Prescription) GetFillsRemainingOk() (*float32, bool)`

GetFillsRemainingOk returns a tuple with the FillsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillsRemaining

`func (o *Prescription) SetFillsRemaining(v float32)`

SetFillsRemaining sets FillsRemaining field to given value.

### HasFillsRemaining

`func (o *Prescription) HasFillsRemaining() bool`

HasFillsRemaining returns a boolean if a field has been set.

### GetDaysSupply

`func (o *Prescription) GetDaysSupply() float32`

GetDaysSupply returns the DaysSupply field if non-nil, zero value otherwise.

### GetDaysSupplyOk

`func (o *Prescription) GetDaysSupplyOk() (*float32, bool)`

GetDaysSupplyOk returns a tuple with the DaysSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSupply

`func (o *Prescription) SetDaysSupply(v float32)`

SetDaysSupply sets DaysSupply field to given value.

### HasDaysSupply

`func (o *Prescription) HasDaysSupply() bool`

HasDaysSupply returns a boolean if a field has been set.

### GetRxNumber

`func (o *Prescription) GetRxNumber() string`

GetRxNumber returns the RxNumber field if non-nil, zero value otherwise.

### GetRxNumberOk

`func (o *Prescription) GetRxNumberOk() (*string, bool)`

GetRxNumberOk returns a tuple with the RxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxNumber

`func (o *Prescription) SetRxNumber(v string)`

SetRxNumber sets RxNumber field to given value.

### HasRxNumber

`func (o *Prescription) HasRxNumber() bool`

HasRxNumber returns a boolean if a field has been set.

### GetTransferRequestStatus

`func (o *Prescription) GetTransferRequestStatus() TransferRequestStatus`

GetTransferRequestStatus returns the TransferRequestStatus field if non-nil, zero value otherwise.

### GetTransferRequestStatusOk

`func (o *Prescription) GetTransferRequestStatusOk() (*TransferRequestStatus, bool)`

GetTransferRequestStatusOk returns a tuple with the TransferRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferRequestStatus

`func (o *Prescription) SetTransferRequestStatus(v TransferRequestStatus)`

SetTransferRequestStatus sets TransferRequestStatus field to given value.

### HasTransferRequestStatus

`func (o *Prescription) HasTransferRequestStatus() bool`

HasTransferRequestStatus returns a boolean if a field has been set.

### GetTotalQuantityRemaining

`func (o *Prescription) GetTotalQuantityRemaining() Quantity`

GetTotalQuantityRemaining returns the TotalQuantityRemaining field if non-nil, zero value otherwise.

### GetTotalQuantityRemainingOk

`func (o *Prescription) GetTotalQuantityRemainingOk() (*Quantity, bool)`

GetTotalQuantityRemainingOk returns a tuple with the TotalQuantityRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantityRemaining

`func (o *Prescription) SetTotalQuantityRemaining(v Quantity)`

SetTotalQuantityRemaining sets TotalQuantityRemaining field to given value.

### HasTotalQuantityRemaining

`func (o *Prescription) HasTotalQuantityRemaining() bool`

HasTotalQuantityRemaining returns a boolean if a field has been set.

### GetFills

`func (o *Prescription) GetFills() []PrescriptionFill`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *Prescription) GetFillsOk() (*[]PrescriptionFill, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *Prescription) SetFills(v []PrescriptionFill)`

SetFills sets Fills field to given value.

### HasFills

`func (o *Prescription) HasFills() bool`

HasFills returns a boolean if a field has been set.

### GetPrescriber

`func (o *Prescription) GetPrescriber() Prescriber`

GetPrescriber returns the Prescriber field if non-nil, zero value otherwise.

### GetPrescriberOk

`func (o *Prescription) GetPrescriberOk() (*Prescriber, bool)`

GetPrescriberOk returns a tuple with the Prescriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriber

`func (o *Prescription) SetPrescriber(v Prescriber)`

SetPrescriber sets Prescriber field to given value.

### HasPrescriber

`func (o *Prescription) HasPrescriber() bool`

HasPrescriber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


