# InsuranceComplianceExceptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedRules** | [**[]FailedRule**](FailedRule.md) | List of specific rules that were violated | 
**Message** | **string** | Informational message about the insurance compliance exception | 

## Methods

### NewInsuranceComplianceExceptionResponseContent

`func NewInsuranceComplianceExceptionResponseContent(failedRules []FailedRule, message string, ) *InsuranceComplianceExceptionResponseContent`

NewInsuranceComplianceExceptionResponseContent instantiates a new InsuranceComplianceExceptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsuranceComplianceExceptionResponseContentWithDefaults

`func NewInsuranceComplianceExceptionResponseContentWithDefaults() *InsuranceComplianceExceptionResponseContent`

NewInsuranceComplianceExceptionResponseContentWithDefaults instantiates a new InsuranceComplianceExceptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedRules

`func (o *InsuranceComplianceExceptionResponseContent) GetFailedRules() []FailedRule`

GetFailedRules returns the FailedRules field if non-nil, zero value otherwise.

### GetFailedRulesOk

`func (o *InsuranceComplianceExceptionResponseContent) GetFailedRulesOk() (*[]FailedRule, bool)`

GetFailedRulesOk returns a tuple with the FailedRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRules

`func (o *InsuranceComplianceExceptionResponseContent) SetFailedRules(v []FailedRule)`

SetFailedRules sets FailedRules field to given value.


### GetMessage

`func (o *InsuranceComplianceExceptionResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InsuranceComplianceExceptionResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InsuranceComplianceExceptionResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


