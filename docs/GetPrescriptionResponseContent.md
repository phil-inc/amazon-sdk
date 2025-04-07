# GetPrescriptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prescription** | [**Prescription**](Prescription.md) |  | 
**PartnerPrescriptionId** | **string** | The partner Prescription ID of the requested prescription. | 

## Methods

### NewGetPrescriptionResponseContent

`func NewGetPrescriptionResponseContent(prescription Prescription, partnerPrescriptionId string, ) *GetPrescriptionResponseContent`

NewGetPrescriptionResponseContent instantiates a new GetPrescriptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPrescriptionResponseContentWithDefaults

`func NewGetPrescriptionResponseContentWithDefaults() *GetPrescriptionResponseContent`

NewGetPrescriptionResponseContentWithDefaults instantiates a new GetPrescriptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrescription

`func (o *GetPrescriptionResponseContent) GetPrescription() Prescription`

GetPrescription returns the Prescription field if non-nil, zero value otherwise.

### GetPrescriptionOk

`func (o *GetPrescriptionResponseContent) GetPrescriptionOk() (*Prescription, bool)`

GetPrescriptionOk returns a tuple with the Prescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescription

`func (o *GetPrescriptionResponseContent) SetPrescription(v Prescription)`

SetPrescription sets Prescription field to given value.


### GetPartnerPrescriptionId

`func (o *GetPrescriptionResponseContent) GetPartnerPrescriptionId() string`

GetPartnerPrescriptionId returns the PartnerPrescriptionId field if non-nil, zero value otherwise.

### GetPartnerPrescriptionIdOk

`func (o *GetPrescriptionResponseContent) GetPartnerPrescriptionIdOk() (*string, bool)`

GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrescriptionId

`func (o *GetPrescriptionResponseContent) SetPartnerPrescriptionId(v string)`

SetPartnerPrescriptionId sets PartnerPrescriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


