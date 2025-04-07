# DrugDbCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A string that must contain at least one non-whitespace character, potentially preceded by whitespace.  Here&#39;s how it&#39;s checked: - &#39;^&#39; asserts the beginning of the string. - &#39;\\s*&#39; allows any number of whitespace characters at the start of the string, including none. - &#39;\\S&#39; ensures there is at least one non-whitespace character in the string. - &#39;.*$&#39; allows any characters to follow the non-whitespace character, extending to the end of the string.  This ensures that the string is not entirely whitespace, although it can start with whitespace and can contain any characters after the first non-whitespace character.  Note: This naturally enforces a minimum length of 1 due to 1 non-whitespace character requirement. | 
**Qualifier** | [**DrugDbCodeQualifier**](DrugDbCodeQualifier.md) |  | 

## Methods

### NewDrugDbCode

`func NewDrugDbCode(code string, qualifier DrugDbCodeQualifier, ) *DrugDbCode`

NewDrugDbCode instantiates a new DrugDbCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugDbCodeWithDefaults

`func NewDrugDbCodeWithDefaults() *DrugDbCode`

NewDrugDbCodeWithDefaults instantiates a new DrugDbCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DrugDbCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DrugDbCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DrugDbCode) SetCode(v string)`

SetCode sets Code field to given value.


### GetQualifier

`func (o *DrugDbCode) GetQualifier() DrugDbCodeQualifier`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *DrugDbCode) GetQualifierOk() (*DrugDbCodeQualifier, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *DrugDbCode) SetQualifier(v DrugDbCodeQualifier)`

SetQualifier sets Qualifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


