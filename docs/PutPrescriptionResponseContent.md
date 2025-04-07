# PutPrescriptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceivingPharmacy** | [**Pharmacy**](Pharmacy.md) |  | 
**PartnerPrescriptionId** | **string** | The partner Prescription ID of the added prescription. | 

## Methods

### NewPutPrescriptionResponseContent

`func NewPutPrescriptionResponseContent(receivingPharmacy Pharmacy, partnerPrescriptionId string, ) *PutPrescriptionResponseContent`

NewPutPrescriptionResponseContent instantiates a new PutPrescriptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPrescriptionResponseContentWithDefaults

`func NewPutPrescriptionResponseContentWithDefaults() *PutPrescriptionResponseContent`

NewPutPrescriptionResponseContentWithDefaults instantiates a new PutPrescriptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceivingPharmacy

`func (o *PutPrescriptionResponseContent) GetReceivingPharmacy() Pharmacy`

GetReceivingPharmacy returns the ReceivingPharmacy field if non-nil, zero value otherwise.

### GetReceivingPharmacyOk

`func (o *PutPrescriptionResponseContent) GetReceivingPharmacyOk() (*Pharmacy, bool)`

GetReceivingPharmacyOk returns a tuple with the ReceivingPharmacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingPharmacy

`func (o *PutPrescriptionResponseContent) SetReceivingPharmacy(v Pharmacy)`

SetReceivingPharmacy sets ReceivingPharmacy field to given value.


### GetPartnerPrescriptionId

`func (o *PutPrescriptionResponseContent) GetPartnerPrescriptionId() string`

GetPartnerPrescriptionId returns the PartnerPrescriptionId field if non-nil, zero value otherwise.

### GetPartnerPrescriptionIdOk

`func (o *PutPrescriptionResponseContent) GetPartnerPrescriptionIdOk() (*string, bool)`

GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrescriptionId

`func (o *PutPrescriptionResponseContent) SetPartnerPrescriptionId(v string)`

SetPartnerPrescriptionId sets PartnerPrescriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


