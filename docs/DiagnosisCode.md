# DiagnosisCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icd10Code** | Pointer to **string** | International Classification of Diseases (ICD 10) Code of the diagnosis. Required for Hubs partners. Not required if codeWithQualifier is passed in. | [optional] 
**CodeWithQualifier** | Pointer to [**DiagnosisCodeWithQualifier**](DiagnosisCodeWithQualifier.md) |  | [optional] 
**Description** | **string** | A string that must contain at least one non-whitespace character, potentially preceded by whitespace.  Here&#39;s how it&#39;s checked: - &#39;^&#39; asserts the beginning of the string. - &#39;\\s*&#39; allows any number of whitespace characters at the start of the string, including none. - &#39;\\S&#39; ensures there is at least one non-whitespace character in the string. - &#39;.*$&#39; allows any characters to follow the non-whitespace character, extending to the end of the string.  This ensures that the string is not entirely whitespace, although it can start with whitespace and can contain any characters after the first non-whitespace character.  Note: This naturally enforces a minimum length of 1 due to 1 non-whitespace character requirement. | 

## Methods

### NewDiagnosisCode

`func NewDiagnosisCode(description string, ) *DiagnosisCode`

NewDiagnosisCode instantiates a new DiagnosisCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosisCodeWithDefaults

`func NewDiagnosisCodeWithDefaults() *DiagnosisCode`

NewDiagnosisCodeWithDefaults instantiates a new DiagnosisCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcd10Code

`func (o *DiagnosisCode) GetIcd10Code() string`

GetIcd10Code returns the Icd10Code field if non-nil, zero value otherwise.

### GetIcd10CodeOk

`func (o *DiagnosisCode) GetIcd10CodeOk() (*string, bool)`

GetIcd10CodeOk returns a tuple with the Icd10Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcd10Code

`func (o *DiagnosisCode) SetIcd10Code(v string)`

SetIcd10Code sets Icd10Code field to given value.

### HasIcd10Code

`func (o *DiagnosisCode) HasIcd10Code() bool`

HasIcd10Code returns a boolean if a field has been set.

### GetCodeWithQualifier

`func (o *DiagnosisCode) GetCodeWithQualifier() DiagnosisCodeWithQualifier`

GetCodeWithQualifier returns the CodeWithQualifier field if non-nil, zero value otherwise.

### GetCodeWithQualifierOk

`func (o *DiagnosisCode) GetCodeWithQualifierOk() (*DiagnosisCodeWithQualifier, bool)`

GetCodeWithQualifierOk returns a tuple with the CodeWithQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeWithQualifier

`func (o *DiagnosisCode) SetCodeWithQualifier(v DiagnosisCodeWithQualifier)`

SetCodeWithQualifier sets CodeWithQualifier field to given value.

### HasCodeWithQualifier

`func (o *DiagnosisCode) HasCodeWithQualifier() bool`

HasCodeWithQualifier returns a boolean if a field has been set.

### GetDescription

`func (o *DiagnosisCode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiagnosisCode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiagnosisCode) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


