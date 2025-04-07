# PrescriptionTransferInDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendingPharmacy** | [**Pharmacy**](Pharmacy.md) |  | 
**PrescriptionTransferred** | [**PrescriptionTransferred**](PrescriptionTransferred.md) |  | 

## Methods

### NewPrescriptionTransferInDetails

`func NewPrescriptionTransferInDetails(sendingPharmacy Pharmacy, prescriptionTransferred PrescriptionTransferred, ) *PrescriptionTransferInDetails`

NewPrescriptionTransferInDetails instantiates a new PrescriptionTransferInDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrescriptionTransferInDetailsWithDefaults

`func NewPrescriptionTransferInDetailsWithDefaults() *PrescriptionTransferInDetails`

NewPrescriptionTransferInDetailsWithDefaults instantiates a new PrescriptionTransferInDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendingPharmacy

`func (o *PrescriptionTransferInDetails) GetSendingPharmacy() Pharmacy`

GetSendingPharmacy returns the SendingPharmacy field if non-nil, zero value otherwise.

### GetSendingPharmacyOk

`func (o *PrescriptionTransferInDetails) GetSendingPharmacyOk() (*Pharmacy, bool)`

GetSendingPharmacyOk returns a tuple with the SendingPharmacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingPharmacy

`func (o *PrescriptionTransferInDetails) SetSendingPharmacy(v Pharmacy)`

SetSendingPharmacy sets SendingPharmacy field to given value.


### GetPrescriptionTransferred

`func (o *PrescriptionTransferInDetails) GetPrescriptionTransferred() PrescriptionTransferred`

GetPrescriptionTransferred returns the PrescriptionTransferred field if non-nil, zero value otherwise.

### GetPrescriptionTransferredOk

`func (o *PrescriptionTransferInDetails) GetPrescriptionTransferredOk() (*PrescriptionTransferred, bool)`

GetPrescriptionTransferredOk returns a tuple with the PrescriptionTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriptionTransferred

`func (o *PrescriptionTransferInDetails) SetPrescriptionTransferred(v PrescriptionTransferred)`

SetPrescriptionTransferred sets PrescriptionTransferred field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


