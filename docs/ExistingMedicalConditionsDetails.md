# ExistingMedicalConditionsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoKnownExistingMedicalConditions** | Pointer to **bool** | Currently supported but on path to deprecation as we will plan to switch to existingMedicalConditionsDescriptor This shape is deprecated since 2024-05-15: Use existingMedicalConditionsDescriptor instead | [optional] 
**ExistingMedicalConditions** | Pointer to [**[]MedicalCondition**](MedicalCondition.md) | If noKnownExistingConditions is false, conditions is required | [optional] 
**ExistingMedicalConditionsDescriptor** | Pointer to [**ExistingMedicalConditionsDescriptor**](ExistingMedicalConditionsDescriptor.md) |  | [optional] 

## Methods

### NewExistingMedicalConditionsDetails

`func NewExistingMedicalConditionsDetails() *ExistingMedicalConditionsDetails`

NewExistingMedicalConditionsDetails instantiates a new ExistingMedicalConditionsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExistingMedicalConditionsDetailsWithDefaults

`func NewExistingMedicalConditionsDetailsWithDefaults() *ExistingMedicalConditionsDetails`

NewExistingMedicalConditionsDetailsWithDefaults instantiates a new ExistingMedicalConditionsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoKnownExistingMedicalConditions

`func (o *ExistingMedicalConditionsDetails) GetNoKnownExistingMedicalConditions() bool`

GetNoKnownExistingMedicalConditions returns the NoKnownExistingMedicalConditions field if non-nil, zero value otherwise.

### GetNoKnownExistingMedicalConditionsOk

`func (o *ExistingMedicalConditionsDetails) GetNoKnownExistingMedicalConditionsOk() (*bool, bool)`

GetNoKnownExistingMedicalConditionsOk returns a tuple with the NoKnownExistingMedicalConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoKnownExistingMedicalConditions

`func (o *ExistingMedicalConditionsDetails) SetNoKnownExistingMedicalConditions(v bool)`

SetNoKnownExistingMedicalConditions sets NoKnownExistingMedicalConditions field to given value.

### HasNoKnownExistingMedicalConditions

`func (o *ExistingMedicalConditionsDetails) HasNoKnownExistingMedicalConditions() bool`

HasNoKnownExistingMedicalConditions returns a boolean if a field has been set.

### GetExistingMedicalConditions

`func (o *ExistingMedicalConditionsDetails) GetExistingMedicalConditions() []MedicalCondition`

GetExistingMedicalConditions returns the ExistingMedicalConditions field if non-nil, zero value otherwise.

### GetExistingMedicalConditionsOk

`func (o *ExistingMedicalConditionsDetails) GetExistingMedicalConditionsOk() (*[]MedicalCondition, bool)`

GetExistingMedicalConditionsOk returns a tuple with the ExistingMedicalConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMedicalConditions

`func (o *ExistingMedicalConditionsDetails) SetExistingMedicalConditions(v []MedicalCondition)`

SetExistingMedicalConditions sets ExistingMedicalConditions field to given value.

### HasExistingMedicalConditions

`func (o *ExistingMedicalConditionsDetails) HasExistingMedicalConditions() bool`

HasExistingMedicalConditions returns a boolean if a field has been set.

### GetExistingMedicalConditionsDescriptor

`func (o *ExistingMedicalConditionsDetails) GetExistingMedicalConditionsDescriptor() ExistingMedicalConditionsDescriptor`

GetExistingMedicalConditionsDescriptor returns the ExistingMedicalConditionsDescriptor field if non-nil, zero value otherwise.

### GetExistingMedicalConditionsDescriptorOk

`func (o *ExistingMedicalConditionsDetails) GetExistingMedicalConditionsDescriptorOk() (*ExistingMedicalConditionsDescriptor, bool)`

GetExistingMedicalConditionsDescriptorOk returns a tuple with the ExistingMedicalConditionsDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMedicalConditionsDescriptor

`func (o *ExistingMedicalConditionsDetails) SetExistingMedicalConditionsDescriptor(v ExistingMedicalConditionsDescriptor)`

SetExistingMedicalConditionsDescriptor sets ExistingMedicalConditionsDescriptor field to given value.

### HasExistingMedicalConditionsDescriptor

`func (o *ExistingMedicalConditionsDetails) HasExistingMedicalConditionsDescriptor() bool`

HasExistingMedicalConditionsDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


