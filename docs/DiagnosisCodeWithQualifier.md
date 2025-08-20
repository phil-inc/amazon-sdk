# DiagnosisCodeWithQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | String representation of the diagnosis code | 
**CodeQualifier** | [**DiagnosisCodeQualifier**](DiagnosisCodeQualifier.md) |  | 

## Methods

### NewDiagnosisCodeWithQualifier

`func NewDiagnosisCodeWithQualifier(code string, codeQualifier DiagnosisCodeQualifier, ) *DiagnosisCodeWithQualifier`

NewDiagnosisCodeWithQualifier instantiates a new DiagnosisCodeWithQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosisCodeWithQualifierWithDefaults

`func NewDiagnosisCodeWithQualifierWithDefaults() *DiagnosisCodeWithQualifier`

NewDiagnosisCodeWithQualifierWithDefaults instantiates a new DiagnosisCodeWithQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DiagnosisCodeWithQualifier) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DiagnosisCodeWithQualifier) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DiagnosisCodeWithQualifier) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeQualifier

`func (o *DiagnosisCodeWithQualifier) GetCodeQualifier() DiagnosisCodeQualifier`

GetCodeQualifier returns the CodeQualifier field if non-nil, zero value otherwise.

### GetCodeQualifierOk

`func (o *DiagnosisCodeWithQualifier) GetCodeQualifierOk() (*DiagnosisCodeQualifier, bool)`

GetCodeQualifierOk returns a tuple with the CodeQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeQualifier

`func (o *DiagnosisCodeWithQualifier) SetCodeQualifier(v DiagnosisCodeQualifier)`

SetCodeQualifier sets CodeQualifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


