# Patient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**PatientStatus**](PatientStatus.md) |  | 
**StatusReasonCode** | Pointer to [**PatientStatusReasonCode**](PatientStatusReasonCode.md) |  | [optional] 
**StatusDescription** | Pointer to **string** | A human-readable description of the patient&#39;s status. | [optional] 
**Name** | [**Name**](Name.md) |  | 
**DateOfBirth** | **string** | Patient&#39;s date of birth in the format &#39;yyyy-mm-dd&#39;. | 
**SexAssignedAtBirth** | [**SexAssignedAtBirth**](SexAssignedAtBirth.md) |  | 
**ContactInfo** | [**[]ContactInfo**](ContactInfo.md) | List of contact information of the patient. | 
**ResidentialAddress** | [**Address**](Address.md) |  | 
**Email** | Pointer to **string** | Email address of the patient. | [optional] 

## Methods

### NewPatient

`func NewPatient(status PatientStatus, name Name, dateOfBirth string, sexAssignedAtBirth SexAssignedAtBirth, contactInfo []ContactInfo, residentialAddress Address, ) *Patient`

NewPatient instantiates a new Patient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatientWithDefaults

`func NewPatientWithDefaults() *Patient`

NewPatientWithDefaults instantiates a new Patient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Patient) GetStatus() PatientStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Patient) GetStatusOk() (*PatientStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Patient) SetStatus(v PatientStatus)`

SetStatus sets Status field to given value.


### GetStatusReasonCode

`func (o *Patient) GetStatusReasonCode() PatientStatusReasonCode`

GetStatusReasonCode returns the StatusReasonCode field if non-nil, zero value otherwise.

### GetStatusReasonCodeOk

`func (o *Patient) GetStatusReasonCodeOk() (*PatientStatusReasonCode, bool)`

GetStatusReasonCodeOk returns a tuple with the StatusReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReasonCode

`func (o *Patient) SetStatusReasonCode(v PatientStatusReasonCode)`

SetStatusReasonCode sets StatusReasonCode field to given value.

### HasStatusReasonCode

`func (o *Patient) HasStatusReasonCode() bool`

HasStatusReasonCode returns a boolean if a field has been set.

### GetStatusDescription

`func (o *Patient) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *Patient) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *Patient) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *Patient) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetName

`func (o *Patient) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Patient) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Patient) SetName(v Name)`

SetName sets Name field to given value.


### GetDateOfBirth

`func (o *Patient) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Patient) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Patient) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetSexAssignedAtBirth

`func (o *Patient) GetSexAssignedAtBirth() SexAssignedAtBirth`

GetSexAssignedAtBirth returns the SexAssignedAtBirth field if non-nil, zero value otherwise.

### GetSexAssignedAtBirthOk

`func (o *Patient) GetSexAssignedAtBirthOk() (*SexAssignedAtBirth, bool)`

GetSexAssignedAtBirthOk returns a tuple with the SexAssignedAtBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexAssignedAtBirth

`func (o *Patient) SetSexAssignedAtBirth(v SexAssignedAtBirth)`

SetSexAssignedAtBirth sets SexAssignedAtBirth field to given value.


### GetContactInfo

`func (o *Patient) GetContactInfo() []ContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *Patient) GetContactInfoOk() (*[]ContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *Patient) SetContactInfo(v []ContactInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetResidentialAddress

`func (o *Patient) GetResidentialAddress() Address`

GetResidentialAddress returns the ResidentialAddress field if non-nil, zero value otherwise.

### GetResidentialAddressOk

`func (o *Patient) GetResidentialAddressOk() (*Address, bool)`

GetResidentialAddressOk returns a tuple with the ResidentialAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialAddress

`func (o *Patient) SetResidentialAddress(v Address)`

SetResidentialAddress sets ResidentialAddress field to given value.


### GetEmail

`func (o *Patient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Patient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Patient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Patient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


