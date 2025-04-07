# Diagnosis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icd10Code** | Pointer to **string** | International Classification of Diseases (ICD 10) Code of the diagnosis. | [optional] 
**Description** | Pointer to **string** | A human-readable description of the diagnosis. | [optional] 

## Methods

### NewDiagnosis

`func NewDiagnosis() *Diagnosis`

NewDiagnosis instantiates a new Diagnosis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosisWithDefaults

`func NewDiagnosisWithDefaults() *Diagnosis`

NewDiagnosisWithDefaults instantiates a new Diagnosis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcd10Code

`func (o *Diagnosis) GetIcd10Code() string`

GetIcd10Code returns the Icd10Code field if non-nil, zero value otherwise.

### GetIcd10CodeOk

`func (o *Diagnosis) GetIcd10CodeOk() (*string, bool)`

GetIcd10CodeOk returns a tuple with the Icd10Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcd10Code

`func (o *Diagnosis) SetIcd10Code(v string)`

SetIcd10Code sets Icd10Code field to given value.

### HasIcd10Code

`func (o *Diagnosis) HasIcd10Code() bool`

HasIcd10Code returns a boolean if a field has been set.

### GetDescription

`func (o *Diagnosis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Diagnosis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Diagnosis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Diagnosis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


