# FieldErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DemographicErrors** | Pointer to [**[]PossibleDemographicFieldsInError**](PossibleDemographicFieldsInError.md) | List of PossibleDemographicFieldsInError enums | [optional] 
**InsuranceErrors** | Pointer to [**[]PossibleInsuranceFieldsInError**](PossibleInsuranceFieldsInError.md) | List of PossibleInsuranceFieldsInError enums | [optional] 

## Methods

### NewFieldErrors

`func NewFieldErrors() *FieldErrors`

NewFieldErrors instantiates a new FieldErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldErrorsWithDefaults

`func NewFieldErrorsWithDefaults() *FieldErrors`

NewFieldErrorsWithDefaults instantiates a new FieldErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDemographicErrors

`func (o *FieldErrors) GetDemographicErrors() []PossibleDemographicFieldsInError`

GetDemographicErrors returns the DemographicErrors field if non-nil, zero value otherwise.

### GetDemographicErrorsOk

`func (o *FieldErrors) GetDemographicErrorsOk() (*[]PossibleDemographicFieldsInError, bool)`

GetDemographicErrorsOk returns a tuple with the DemographicErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographicErrors

`func (o *FieldErrors) SetDemographicErrors(v []PossibleDemographicFieldsInError)`

SetDemographicErrors sets DemographicErrors field to given value.

### HasDemographicErrors

`func (o *FieldErrors) HasDemographicErrors() bool`

HasDemographicErrors returns a boolean if a field has been set.

### GetInsuranceErrors

`func (o *FieldErrors) GetInsuranceErrors() []PossibleInsuranceFieldsInError`

GetInsuranceErrors returns the InsuranceErrors field if non-nil, zero value otherwise.

### GetInsuranceErrorsOk

`func (o *FieldErrors) GetInsuranceErrorsOk() (*[]PossibleInsuranceFieldsInError, bool)`

GetInsuranceErrorsOk returns a tuple with the InsuranceErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceErrors

`func (o *FieldErrors) SetInsuranceErrors(v []PossibleInsuranceFieldsInError)`

SetInsuranceErrors sets InsuranceErrors field to given value.

### HasInsuranceErrors

`func (o *FieldErrors) HasInsuranceErrors() bool`

HasInsuranceErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


