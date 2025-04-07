# Allergy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Patient&#39;s allergy code, Allergy code should follow the AllergyCodeType input&#39;s standard code format. | 
**CodeType** | [**AllergyCodeType**](AllergyCodeType.md) |  | 
**TextValue** | **string** | A human-readable text value of the allergy. Eg, Peanut Allergy | 

## Methods

### NewAllergy

`func NewAllergy(code string, codeType AllergyCodeType, textValue string, ) *Allergy`

NewAllergy instantiates a new Allergy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllergyWithDefaults

`func NewAllergyWithDefaults() *Allergy`

NewAllergyWithDefaults instantiates a new Allergy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Allergy) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Allergy) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Allergy) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeType

`func (o *Allergy) GetCodeType() AllergyCodeType`

GetCodeType returns the CodeType field if non-nil, zero value otherwise.

### GetCodeTypeOk

`func (o *Allergy) GetCodeTypeOk() (*AllergyCodeType, bool)`

GetCodeTypeOk returns a tuple with the CodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeType

`func (o *Allergy) SetCodeType(v AllergyCodeType)`

SetCodeType sets CodeType field to given value.


### GetTextValue

`func (o *Allergy) GetTextValue() string`

GetTextValue returns the TextValue field if non-nil, zero value otherwise.

### GetTextValueOk

`func (o *Allergy) GetTextValueOk() (*string, bool)`

GetTextValueOk returns a tuple with the TextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextValue

`func (o *Allergy) SetTextValue(v string)`

SetTextValue sets TextValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


