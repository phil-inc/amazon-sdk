# PutPatientRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerPatientId** | **string** | A unique Patient Identifier provided while creating the patient profile. | 
**PatientDetails** | [**PatientDetails**](PatientDetails.md) |  | 
**ResidentialAddress** | [**Address**](Address.md) |  | 
**Insurances** | Pointer to [**[]Insurance**](Insurance.md) | List of insurances for this patient. | [optional] 

## Methods

### NewPutPatientRequestContent

`func NewPutPatientRequestContent(partnerPatientId string, patientDetails PatientDetails, residentialAddress Address, ) *PutPatientRequestContent`

NewPutPatientRequestContent instantiates a new PutPatientRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPatientRequestContentWithDefaults

`func NewPutPatientRequestContentWithDefaults() *PutPatientRequestContent`

NewPutPatientRequestContentWithDefaults instantiates a new PutPatientRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerPatientId

`func (o *PutPatientRequestContent) GetPartnerPatientId() string`

GetPartnerPatientId returns the PartnerPatientId field if non-nil, zero value otherwise.

### GetPartnerPatientIdOk

`func (o *PutPatientRequestContent) GetPartnerPatientIdOk() (*string, bool)`

GetPartnerPatientIdOk returns a tuple with the PartnerPatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPatientId

`func (o *PutPatientRequestContent) SetPartnerPatientId(v string)`

SetPartnerPatientId sets PartnerPatientId field to given value.


### GetPatientDetails

`func (o *PutPatientRequestContent) GetPatientDetails() PatientDetails`

GetPatientDetails returns the PatientDetails field if non-nil, zero value otherwise.

### GetPatientDetailsOk

`func (o *PutPatientRequestContent) GetPatientDetailsOk() (*PatientDetails, bool)`

GetPatientDetailsOk returns a tuple with the PatientDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatientDetails

`func (o *PutPatientRequestContent) SetPatientDetails(v PatientDetails)`

SetPatientDetails sets PatientDetails field to given value.


### GetResidentialAddress

`func (o *PutPatientRequestContent) GetResidentialAddress() Address`

GetResidentialAddress returns the ResidentialAddress field if non-nil, zero value otherwise.

### GetResidentialAddressOk

`func (o *PutPatientRequestContent) GetResidentialAddressOk() (*Address, bool)`

GetResidentialAddressOk returns a tuple with the ResidentialAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialAddress

`func (o *PutPatientRequestContent) SetResidentialAddress(v Address)`

SetResidentialAddress sets ResidentialAddress field to given value.


### GetInsurances

`func (o *PutPatientRequestContent) GetInsurances() []Insurance`

GetInsurances returns the Insurances field if non-nil, zero value otherwise.

### GetInsurancesOk

`func (o *PutPatientRequestContent) GetInsurancesOk() (*[]Insurance, bool)`

GetInsurancesOk returns a tuple with the Insurances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsurances

`func (o *PutPatientRequestContent) SetInsurances(v []Insurance)`

SetInsurances sets Insurances field to given value.

### HasInsurances

`func (o *PutPatientRequestContent) HasInsurances() bool`

HasInsurances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


