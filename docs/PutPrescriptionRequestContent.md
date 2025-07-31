# PutPrescriptionRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerPrescriptionId** | **string** | A unique Prescription Identifier provided by the partner. | 
**PartnerPatientId** | **string** | A unique Patient Identifier provided by the partner. | 
**Prescriber** | [**Prescriber**](Prescriber.md) |  | 
**SupervisingPhysician** | Pointer to [**Prescriber**](Prescriber.md) |  | [optional] 
**MedicationPrescribed** | [**MedicationPrescribed**](MedicationPrescribed.md) |  | 
**SendingPharmacy** | Pointer to [**Pharmacy**](Pharmacy.md) |  | [optional] 
**PrescriptionTransferInDetails** | Pointer to [**PrescriptionTransferInDetails**](PrescriptionTransferInDetails.md) |  | [optional] 
**ReceivingPharmacy** | [**Pharmacy**](Pharmacy.md) |  | 
**Observation** | Pointer to [**Observation**](Observation.md) |  | [optional] 

## Methods

### NewPutPrescriptionRequestContent

`func NewPutPrescriptionRequestContent(partnerPrescriptionId string, partnerPatientId string, prescriber Prescriber, medicationPrescribed MedicationPrescribed, receivingPharmacy Pharmacy, ) *PutPrescriptionRequestContent`

NewPutPrescriptionRequestContent instantiates a new PutPrescriptionRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPrescriptionRequestContentWithDefaults

`func NewPutPrescriptionRequestContentWithDefaults() *PutPrescriptionRequestContent`

NewPutPrescriptionRequestContentWithDefaults instantiates a new PutPrescriptionRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerPrescriptionId

`func (o *PutPrescriptionRequestContent) GetPartnerPrescriptionId() string`

GetPartnerPrescriptionId returns the PartnerPrescriptionId field if non-nil, zero value otherwise.

### GetPartnerPrescriptionIdOk

`func (o *PutPrescriptionRequestContent) GetPartnerPrescriptionIdOk() (*string, bool)`

GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrescriptionId

`func (o *PutPrescriptionRequestContent) SetPartnerPrescriptionId(v string)`

SetPartnerPrescriptionId sets PartnerPrescriptionId field to given value.


### GetPartnerPatientId

`func (o *PutPrescriptionRequestContent) GetPartnerPatientId() string`

GetPartnerPatientId returns the PartnerPatientId field if non-nil, zero value otherwise.

### GetPartnerPatientIdOk

`func (o *PutPrescriptionRequestContent) GetPartnerPatientIdOk() (*string, bool)`

GetPartnerPatientIdOk returns a tuple with the PartnerPatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPatientId

`func (o *PutPrescriptionRequestContent) SetPartnerPatientId(v string)`

SetPartnerPatientId sets PartnerPatientId field to given value.


### GetPrescriber

`func (o *PutPrescriptionRequestContent) GetPrescriber() Prescriber`

GetPrescriber returns the Prescriber field if non-nil, zero value otherwise.

### GetPrescriberOk

`func (o *PutPrescriptionRequestContent) GetPrescriberOk() (*Prescriber, bool)`

GetPrescriberOk returns a tuple with the Prescriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriber

`func (o *PutPrescriptionRequestContent) SetPrescriber(v Prescriber)`

SetPrescriber sets Prescriber field to given value.


### GetSupervisingPhysician

`func (o *PutPrescriptionRequestContent) GetSupervisingPhysician() Prescriber`

GetSupervisingPhysician returns the SupervisingPhysician field if non-nil, zero value otherwise.

### GetSupervisingPhysicianOk

`func (o *PutPrescriptionRequestContent) GetSupervisingPhysicianOk() (*Prescriber, bool)`

GetSupervisingPhysicianOk returns a tuple with the SupervisingPhysician field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisingPhysician

`func (o *PutPrescriptionRequestContent) SetSupervisingPhysician(v Prescriber)`

SetSupervisingPhysician sets SupervisingPhysician field to given value.

### HasSupervisingPhysician

`func (o *PutPrescriptionRequestContent) HasSupervisingPhysician() bool`

HasSupervisingPhysician returns a boolean if a field has been set.

### GetMedicationPrescribed

`func (o *PutPrescriptionRequestContent) GetMedicationPrescribed() MedicationPrescribed`

GetMedicationPrescribed returns the MedicationPrescribed field if non-nil, zero value otherwise.

### GetMedicationPrescribedOk

`func (o *PutPrescriptionRequestContent) GetMedicationPrescribedOk() (*MedicationPrescribed, bool)`

GetMedicationPrescribedOk returns a tuple with the MedicationPrescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicationPrescribed

`func (o *PutPrescriptionRequestContent) SetMedicationPrescribed(v MedicationPrescribed)`

SetMedicationPrescribed sets MedicationPrescribed field to given value.


### GetSendingPharmacy

`func (o *PutPrescriptionRequestContent) GetSendingPharmacy() Pharmacy`

GetSendingPharmacy returns the SendingPharmacy field if non-nil, zero value otherwise.

### GetSendingPharmacyOk

`func (o *PutPrescriptionRequestContent) GetSendingPharmacyOk() (*Pharmacy, bool)`

GetSendingPharmacyOk returns a tuple with the SendingPharmacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingPharmacy

`func (o *PutPrescriptionRequestContent) SetSendingPharmacy(v Pharmacy)`

SetSendingPharmacy sets SendingPharmacy field to given value.

### HasSendingPharmacy

`func (o *PutPrescriptionRequestContent) HasSendingPharmacy() bool`

HasSendingPharmacy returns a boolean if a field has been set.

### GetPrescriptionTransferInDetails

`func (o *PutPrescriptionRequestContent) GetPrescriptionTransferInDetails() PrescriptionTransferInDetails`

GetPrescriptionTransferInDetails returns the PrescriptionTransferInDetails field if non-nil, zero value otherwise.

### GetPrescriptionTransferInDetailsOk

`func (o *PutPrescriptionRequestContent) GetPrescriptionTransferInDetailsOk() (*PrescriptionTransferInDetails, bool)`

GetPrescriptionTransferInDetailsOk returns a tuple with the PrescriptionTransferInDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriptionTransferInDetails

`func (o *PutPrescriptionRequestContent) SetPrescriptionTransferInDetails(v PrescriptionTransferInDetails)`

SetPrescriptionTransferInDetails sets PrescriptionTransferInDetails field to given value.

### HasPrescriptionTransferInDetails

`func (o *PutPrescriptionRequestContent) HasPrescriptionTransferInDetails() bool`

HasPrescriptionTransferInDetails returns a boolean if a field has been set.

### GetReceivingPharmacy

`func (o *PutPrescriptionRequestContent) GetReceivingPharmacy() Pharmacy`

GetReceivingPharmacy returns the ReceivingPharmacy field if non-nil, zero value otherwise.

### GetReceivingPharmacyOk

`func (o *PutPrescriptionRequestContent) GetReceivingPharmacyOk() (*Pharmacy, bool)`

GetReceivingPharmacyOk returns a tuple with the ReceivingPharmacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingPharmacy

`func (o *PutPrescriptionRequestContent) SetReceivingPharmacy(v Pharmacy)`

SetReceivingPharmacy sets ReceivingPharmacy field to given value.


### GetObservation

`func (o *PutPrescriptionRequestContent) GetObservation() Observation`

GetObservation returns the Observation field if non-nil, zero value otherwise.

### GetObservationOk

`func (o *PutPrescriptionRequestContent) GetObservationOk() (*Observation, bool)`

GetObservationOk returns a tuple with the Observation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservation

`func (o *PutPrescriptionRequestContent) SetObservation(v Observation)`

SetObservation sets Observation field to given value.

### HasObservation

`func (o *PutPrescriptionRequestContent) HasObservation() bool`

HasObservation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


