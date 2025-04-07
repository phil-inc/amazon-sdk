# PostPrescriptionTransferRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerPrescriptionId** | **string** | A unique Prescription Identifier provided by the partner. | 
**TransferToPharmacy** | [**TransferToPharmacy**](TransferToPharmacy.md) |  | 
**ReasonForTransfer** | Pointer to **string** | The reason for transferring out this prescription. | [optional] 

## Methods

### NewPostPrescriptionTransferRequestContent

`func NewPostPrescriptionTransferRequestContent(partnerPrescriptionId string, transferToPharmacy TransferToPharmacy, ) *PostPrescriptionTransferRequestContent`

NewPostPrescriptionTransferRequestContent instantiates a new PostPrescriptionTransferRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPrescriptionTransferRequestContentWithDefaults

`func NewPostPrescriptionTransferRequestContentWithDefaults() *PostPrescriptionTransferRequestContent`

NewPostPrescriptionTransferRequestContentWithDefaults instantiates a new PostPrescriptionTransferRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerPrescriptionId

`func (o *PostPrescriptionTransferRequestContent) GetPartnerPrescriptionId() string`

GetPartnerPrescriptionId returns the PartnerPrescriptionId field if non-nil, zero value otherwise.

### GetPartnerPrescriptionIdOk

`func (o *PostPrescriptionTransferRequestContent) GetPartnerPrescriptionIdOk() (*string, bool)`

GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrescriptionId

`func (o *PostPrescriptionTransferRequestContent) SetPartnerPrescriptionId(v string)`

SetPartnerPrescriptionId sets PartnerPrescriptionId field to given value.


### GetTransferToPharmacy

`func (o *PostPrescriptionTransferRequestContent) GetTransferToPharmacy() TransferToPharmacy`

GetTransferToPharmacy returns the TransferToPharmacy field if non-nil, zero value otherwise.

### GetTransferToPharmacyOk

`func (o *PostPrescriptionTransferRequestContent) GetTransferToPharmacyOk() (*TransferToPharmacy, bool)`

GetTransferToPharmacyOk returns a tuple with the TransferToPharmacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferToPharmacy

`func (o *PostPrescriptionTransferRequestContent) SetTransferToPharmacy(v TransferToPharmacy)`

SetTransferToPharmacy sets TransferToPharmacy field to given value.


### GetReasonForTransfer

`func (o *PostPrescriptionTransferRequestContent) GetReasonForTransfer() string`

GetReasonForTransfer returns the ReasonForTransfer field if non-nil, zero value otherwise.

### GetReasonForTransferOk

`func (o *PostPrescriptionTransferRequestContent) GetReasonForTransferOk() (*string, bool)`

GetReasonForTransferOk returns a tuple with the ReasonForTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForTransfer

`func (o *PostPrescriptionTransferRequestContent) SetReasonForTransfer(v string)`

SetReasonForTransfer sets ReasonForTransfer field to given value.

### HasReasonForTransfer

`func (o *PostPrescriptionTransferRequestContent) HasReasonForTransfer() bool`

HasReasonForTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


