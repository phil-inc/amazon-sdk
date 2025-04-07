# MedicalCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Patient&#39;s condition code. Condition code should follow the ConditionCodeType input&#39;s standard code format. | 
**CodeType** | [**ConditionCodeType**](ConditionCodeType.md) |  | 
**TextValue** | Pointer to **string** | A human-readable text value of the medical condition. | [optional] 

## Methods

### NewMedicalCondition

`func NewMedicalCondition(code string, codeType ConditionCodeType, ) *MedicalCondition`

NewMedicalCondition instantiates a new MedicalCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMedicalConditionWithDefaults

`func NewMedicalConditionWithDefaults() *MedicalCondition`

NewMedicalConditionWithDefaults instantiates a new MedicalCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MedicalCondition) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MedicalCondition) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MedicalCondition) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeType

`func (o *MedicalCondition) GetCodeType() ConditionCodeType`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *MedicalCondition) GetCodeTypeOk() (*ConditionCodeType, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *MedicalCondition) SetCodeType(v ConditionCodeType)`

SetCodeType sets CodeType field to given value.


### GetTextValue

`func (o *MedicalCondition) GetTextValue() string`

GetTextValue returns the TextValue field if non-nil, zero value otherwise.

### GetTextValueOk

`func (o *MedicalCondition) GetTextValueOk() (*string, bool)`

GetTextValueOk returns a tuple with the TextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextValue

`func (o *MedicalCondition) SetTextValue(v string)`

SetTextValue sets TextValue field to given value.

### HasTextValue

`func (o *MedicalCondition) HasTextValue() bool`

HasTextValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


