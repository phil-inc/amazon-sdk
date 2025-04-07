# ClaimRejectionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimRejectionReasons** | Pointer to [**[]ClaimRejectionReason**](ClaimRejectionReason.md) | This field represents reasons for claim rejections | [optional] 

## Methods

### NewClaimRejectionDetails

`func NewClaimRejectionDetails() *ClaimRejectionDetails`

NewClaimRejectionDetails instantiates a new ClaimRejectionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimRejectionDetailsWithDefaults

`func NewClaimRejectionDetailsWithDefaults() *ClaimRejectionDetails`

NewClaimRejectionDetailsWithDefaults instantiates a new ClaimRejectionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimRejectionReasons

`func (o *ClaimRejectionDetails) GetClaimRejectionReasons() []ClaimRejectionReason`

GetClaimRejectionReasons returns the ClaimRejectionReasons field if non-nil, zero value otherwise.

### GetClaimRejectionReasonsOk

`func (o *ClaimRejectionDetails) GetClaimRejectionReasonsOk() (*[]ClaimRejectionReason, bool)`

GetClaimRejectionReasonsOk returns a tuple with the ClaimRejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRejectionReasons

`func (o *ClaimRejectionDetails) SetClaimRejectionReasons(v []ClaimRejectionReason)`

SetClaimRejectionReasons sets ClaimRejectionReasons field to given value.

### HasClaimRejectionReasons

`func (o *ClaimRejectionDetails) HasClaimRejectionReasons() bool`

HasClaimRejectionReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


