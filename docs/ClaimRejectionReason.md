# ClaimRejectionReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimRejectionType** | Pointer to [**ClaimRejectionType**](ClaimRejectionType.md) |  | [optional] 
**RejectionTypeDescription** | Pointer to **string** | This field contains the description explaining the claimRejectionType | [optional] 
**FieldErrors** | Pointer to [**FieldErrors**](FieldErrors.md) |  | [optional] 
**RejectedInsuranceDetails** | Pointer to [**RejectedInsuranceDetails**](RejectedInsuranceDetails.md) |  | [optional] 
**DaysSupplyDetails** | Pointer to [**DaysSupplyDetails**](DaysSupplyDetails.md) |  | [optional] 

## Methods

### NewClaimRejectionReason

`func NewClaimRejectionReason() *ClaimRejectionReason`

NewClaimRejectionReason instantiates a new ClaimRejectionReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimRejectionReasonWithDefaults

`func NewClaimRejectionReasonWithDefaults() *ClaimRejectionReason`

NewClaimRejectionReasonWithDefaults instantiates a new ClaimRejectionReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimRejectionType

`func (o *ClaimRejectionReason) GetClaimRejectionType() ClaimRejectionType`

GetClaimRejectionType returns the ClaimRejectionType field if non-nil, zero value otherwise.

### GetClaimRejectionTypeOk

`func (o *ClaimRejectionReason) GetClaimRejectionTypeOk() (*ClaimRejectionType, bool)`

GetClaimRejectionTypeOk returns a tuple with the ClaimRejectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRejectionType

`func (o *ClaimRejectionReason) SetClaimRejectionType(v ClaimRejectionType)`

SetClaimRejectionType sets ClaimRejectionType field to given value.

### HasClaimRejectionType

`func (o *ClaimRejectionReason) HasClaimRejectionType() bool`

HasClaimRejectionType returns a boolean if a field has been set.

### GetRejectionTypeDescription

`func (o *ClaimRejectionReason) GetRejectionTypeDescription() string`

GetRejectionTypeDescription returns the RejectionTypeDescription field if non-nil, zero value otherwise.

### GetRejectionTypeDescriptionOk

`func (o *ClaimRejectionReason) GetRejectionTypeDescriptionOk() (*string, bool)`

GetRejectionTypeDescriptionOk returns a tuple with the RejectionTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionTypeDescription

`func (o *ClaimRejectionReason) SetRejectionTypeDescription(v string)`

SetRejectionTypeDescription sets RejectionTypeDescription field to given value.

### HasRejectionTypeDescription

`func (o *ClaimRejectionReason) HasRejectionTypeDescription() bool`

HasRejectionTypeDescription returns a boolean if a field has been set.

### GetFieldErrors

`func (o *ClaimRejectionReason) GetFieldErrors() FieldErrors`

GetFieldErrors returns the FieldErrors field if non-nil, zero value otherwise.

### GetFieldErrorsOk

`func (o *ClaimRejectionReason) GetFieldErrorsOk() (*FieldErrors, bool)`

GetFieldErrorsOk returns a tuple with the FieldErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldErrors

`func (o *ClaimRejectionReason) SetFieldErrors(v FieldErrors)`

SetFieldErrors sets FieldErrors field to given value.

### HasFieldErrors

`func (o *ClaimRejectionReason) HasFieldErrors() bool`

HasFieldErrors returns a boolean if a field has been set.

### GetRejectedInsuranceDetails

`func (o *ClaimRejectionReason) GetRejectedInsuranceDetails() RejectedInsuranceDetails`

GetRejectedInsuranceDetails returns the RejectedInsuranceDetails field if non-nil, zero value otherwise.

### GetRejectedInsuranceDetailsOk

`func (o *ClaimRejectionReason) GetRejectedInsuranceDetailsOk() (*RejectedInsuranceDetails, bool)`

GetRejectedInsuranceDetailsOk returns a tuple with the RejectedInsuranceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedInsuranceDetails

`func (o *ClaimRejectionReason) SetRejectedInsuranceDetails(v RejectedInsuranceDetails)`

SetRejectedInsuranceDetails sets RejectedInsuranceDetails field to given value.

### HasRejectedInsuranceDetails

`func (o *ClaimRejectionReason) HasRejectedInsuranceDetails() bool`

HasRejectedInsuranceDetails returns a boolean if a field has been set.

### GetDaysSupplyDetails

`func (o *ClaimRejectionReason) GetDaysSupplyDetails() DaysSupplyDetails`

GetDaysSupplyDetails returns the DaysSupplyDetails field if non-nil, zero value otherwise.

### GetDaysSupplyDetailsOk

`func (o *ClaimRejectionReason) GetDaysSupplyDetailsOk() (*DaysSupplyDetails, bool)`

GetDaysSupplyDetailsOk returns a tuple with the DaysSupplyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSupplyDetails

`func (o *ClaimRejectionReason) SetDaysSupplyDetails(v DaysSupplyDetails)`

SetDaysSupplyDetails sets DaysSupplyDetails field to given value.

### HasDaysSupplyDetails

`func (o *ClaimRejectionReason) HasDaysSupplyDetails() bool`

HasDaysSupplyDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


