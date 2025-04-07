# PatientDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**Name**](Name.md) |  | 
**DateOfBirth** | **string** | Patient&#39;s date of birth in the format &#39;yyyy-mm-dd&#39;. | 
**SexAssignedAtBirth** | [**SexAssignedAtBirth**](SexAssignedAtBirth.md) |  | 
**ContactInfo** | [**[]ContactInfo**](ContactInfo.md) | List of contact information of the patient. | 
**Gender** | Pointer to **string** | Gender of the patient. | [optional] 
**Email** | Pointer to **string** | Email address of the patient. | [optional] 
**SmsConsent** | **bool** | Patient&#39;s consent to receive SMS notifications for signup and other communication. | 
**AllergyDetails** | [**AllergyDetails**](AllergyDetails.md) |  | 
**ExistingMedicalConditionDetails** | [**ExistingMedicalConditionsDetails**](ExistingMedicalConditionsDetails.md) |  | 
**PregnancyStatus** | [**PregnancyStatus**](PregnancyStatus.md) |  | 
**ExistingMedicationDetails** | Pointer to [**ExistingMedicationDetails**](ExistingMedicationDetails.md) |  | [optional] 

## Methods

### NewPatientDetails

`func NewPatientDetails(name Name, dateOfBirth string, sexAssignedAtBirth SexAssignedAtBirth, contactInfo []ContactInfo, smsConsent bool, allergyDetails AllergyDetails, existingMedicalConditionDetails ExistingMedicalConditionsDetails, pregnancyStatus PregnancyStatus, ) *PatientDetails`

NewPatientDetails instantiates a new PatientDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatientDetailsWithDefaults

`func NewPatientDetailsWithDefaults() *PatientDetails`

NewPatientDetailsWithDefaults instantiates a new PatientDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatientDetails) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatientDetails) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatientDetails) SetName(v Name)`

SetName sets Name field to given value.


### GetDateOfBirth

`func (o *PatientDetails) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PatientDetails) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PatientDetails) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetSexAssignedAtBirth

`func (o *PatientDetails) GetSexAssignedAtBirth() SexAssignedAtBirth`

GetSexAssignedAtBirth returns the SexAssignedAtBirth field if non-nil, zero value otherwise.

### GetSexAssignedAtBirthOk

`func (o *PatientDetails) GetSexAssignedAtBirthOk() (*SexAssignedAtBirth, bool)`

GetSexAssignedAtBirthOk returns a tuple with the SexAssignedAtBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexAssignedAtBirth

`func (o *PatientDetails) SetSexAssignedAtBirth(v SexAssignedAtBirth)`

SetSexAssignedAtBirth sets SexAssignedAtBirth field to given value.


### GetContactInfo

`func (o *PatientDetails) GetContactInfo() []ContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *PatientDetails) GetContactInfoOk() (*[]ContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *PatientDetails) SetContactInfo(v []ContactInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetGender

`func (o *PatientDetails) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PatientDetails) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PatientDetails) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PatientDetails) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetEmail

`func (o *PatientDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatientDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatientDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatientDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSmsConsent

`func (o *PatientDetails) GetSmsConsent() bool`

GetSmsConsent returns the SmsConsent field if non-nil, zero value otherwise.

### GetSmsConsentOk

`func (o *PatientDetails) GetSmsConsentOk() (*bool, bool)`

GetSmsConsentOk returns a tuple with the SmsConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsConsent

`func (o *PatientDetails) SetSmsConsent(v bool)`

SetSmsConsent sets SmsConsent field to given value.


### GetAllergyDetails

`func (o *PatientDetails) GetAllergyDetails() AllergyDetails`

GetAllergyDetails returns the AllergyDetails field if non-nil, zero value otherwise.

### GetAllergyDetailsOk

`func (o *PatientDetails) GetAllergyDetailsOk() (*AllergyDetails, bool)`

GetAllergyDetailsOk returns a tuple with the AllergyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllergyDetails

`func (o *PatientDetails) SetAllergyDetails(v AllergyDetails)`

SetAllergyDetails sets AllergyDetails field to given value.


### GetExistingMedicalConditionDetails

`func (o *PatientDetails) GetExistingMedicalConditionDetails() ExistingMedicalConditionsDetails`

GetExistingMedicalConditionDetails returns the ExistingMedicalConditionDetails field if non-nil, zero value otherwise.

### GetExistingMedicalConditionDetailsOk

`func (o *PatientDetails) GetExistingMedicalConditionDetailsOk() (*ExistingMedicalConditionsDetails, bool)`

GetExistingMedicalConditionDetailsOk returns a tuple with the ExistingMedicalConditionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMedicalConditionDetails

`func (o *PatientDetails) SetExistingMedicalConditionDetails(v ExistingMedicalConditionsDetails)`

SetExistingMedicalConditionDetails sets ExistingMedicalConditionDetails field to given value.


### GetPregnancyStatus

`func (o *PatientDetails) GetPregnancyStatus() PregnancyStatus`

GetPregnancyStatus returns the PregnancyStatus field if non-nil, zero value otherwise.

### GetPregnancyStatusOk

`func (o *PatientDetails) GetPregnancyStatusOk() (*PregnancyStatus, bool)`

GetPregnancyStatusOk returns a tuple with the PregnancyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPregnancyStatus

`func (o *PatientDetails) SetPregnancyStatus(v PregnancyStatus)`

SetPregnancyStatus sets PregnancyStatus field to given value.


### GetExistingMedicationDetails

`func (o *PatientDetails) GetExistingMedicationDetails() ExistingMedicationDetails`

GetExistingMedicationDetails returns the ExistingMedicationDetails field if non-nil, zero value otherwise.

### GetExistingMedicationDetailsOk

`func (o *PatientDetails) GetExistingMedicationDetailsOk() (*ExistingMedicationDetails, bool)`

GetExistingMedicationDetailsOk returns a tuple with the ExistingMedicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMedicationDetails

`func (o *PatientDetails) SetExistingMedicationDetails(v ExistingMedicationDetails)`

SetExistingMedicationDetails sets ExistingMedicationDetails field to given value.

### HasExistingMedicationDetails

`func (o *PatientDetails) HasExistingMedicationDetails() bool`

HasExistingMedicationDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


