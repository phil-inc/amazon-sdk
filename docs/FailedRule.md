# FailedRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | [**InsuranceComplianceRuleType**](InsuranceComplianceRuleType.md) |  | 
**Message** | Pointer to **string** | Informational message about the failed rule | [optional] 

## Methods

### NewFailedRule

`func NewFailedRule(ruleType InsuranceComplianceRuleType, ) *FailedRule`

NewFailedRule instantiates a new FailedRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedRuleWithDefaults

`func NewFailedRuleWithDefaults() *FailedRule`

NewFailedRuleWithDefaults instantiates a new FailedRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *FailedRule) GetRuleType() InsuranceComplianceRuleType`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *FailedRule) GetRuleTypeOk() (*InsuranceComplianceRuleType, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *FailedRule) SetRuleType(v InsuranceComplianceRuleType)`

SetRuleType sets RuleType field to given value.


### GetMessage

`func (o *FailedRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FailedRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FailedRule) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FailedRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


