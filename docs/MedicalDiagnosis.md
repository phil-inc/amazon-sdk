# MedicalDiagnosis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClinicalInformationQualifier** | [**ClinicalInformationQualifier**](ClinicalInformationQualifier.md) |  | 
**Primary** | [**DiagnosisCode**](DiagnosisCode.md) |  | 
**Secondary** | Pointer to [**DiagnosisCode**](DiagnosisCode.md) |  | [optional] 

## Methods

### NewMedicalDiagnosis

`func NewMedicalDiagnosis(clinicalInformationQualifier ClinicalInformationQualifier, primary DiagnosisCode, ) *MedicalDiagnosis`

NewMedicalDiagnosis instantiates a new MedicalDiagnosis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMedicalDiagnosisWithDefaults

`func NewMedicalDiagnosisWithDefaults() *MedicalDiagnosis`

NewMedicalDiagnosisWithDefaults instantiates a new MedicalDiagnosis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClinicalInformationQualifier

`func (o *MedicalDiagnosis) GetClinicalInformationQualifier() ClinicalInformationQualifier`

GetClinicalInformationQualifier returns the ClinicalInformationQualifier field if non-nil, zero value otherwise.

### GetClinicalInformationQualifierOk

`func (o *MedicalDiagnosis) GetClinicalInformationQualifierOk() (*ClinicalInformationQualifier, bool)`

GetClinicalInformationQualifierOk returns a tuple with the ClinicalInformationQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClinicalInformationQualifier

`func (o *MedicalDiagnosis) SetClinicalInformationQualifier(v ClinicalInformationQualifier)`

SetClinicalInformationQualifier sets ClinicalInformationQualifier field to given value.


### GetPrimary

`func (o *MedicalDiagnosis) GetPrimary() DiagnosisCode`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *MedicalDiagnosis) GetPrimaryOk() (*DiagnosisCode, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *MedicalDiagnosis) SetPrimary(v DiagnosisCode)`

SetPrimary sets Primary field to given value.


### GetSecondary

`func (o *MedicalDiagnosis) GetSecondary() DiagnosisCode`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *MedicalDiagnosis) GetSecondaryOk() (*DiagnosisCode, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *MedicalDiagnosis) SetSecondary(v DiagnosisCode)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *MedicalDiagnosis) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


