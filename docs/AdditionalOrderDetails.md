# AdditionalOrderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingDetails** | Pointer to [**ShippingDetails**](ShippingDetails.md) |  | [optional] 
**PlaceOrderOnDate** | Pointer to **string** | API attribute where orders are held till this date and sent for order placement only when the date is passed. | [optional] 
**ClaimRejectedDetails** | Pointer to [**ClaimRejectionDetails**](ClaimRejectionDetails.md) |  | [optional] 

## Methods

### NewAdditionalOrderDetails

`func NewAdditionalOrderDetails() *AdditionalOrderDetails`

NewAdditionalOrderDetails instantiates a new AdditionalOrderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalOrderDetailsWithDefaults

`func NewAdditionalOrderDetailsWithDefaults() *AdditionalOrderDetails`

NewAdditionalOrderDetailsWithDefaults instantiates a new AdditionalOrderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingDetails

`func (o *AdditionalOrderDetails) GetShippingDetails() ShippingDetails`

GetShippingDetails returns the ShippingDetails field if non-nil, zero value otherwise.

### GetShippingDetailsOk

`func (o *AdditionalOrderDetails) GetShippingDetailsOk() (*ShippingDetails, bool)`

GetShippingDetailsOk returns a tuple with the ShippingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDetails

`func (o *AdditionalOrderDetails) SetShippingDetails(v ShippingDetails)`

SetShippingDetails sets ShippingDetails field to given value.

### HasShippingDetails

`func (o *AdditionalOrderDetails) HasShippingDetails() bool`

HasShippingDetails returns a boolean if a field has been set.

### GetPlaceOrderOnDate

`func (o *AdditionalOrderDetails) GetPlaceOrderOnDate() string`

GetPlaceOrderOnDate returns the PlaceOrderOnDate field if non-nil, zero value otherwise.

### GetPlaceOrderOnDateOk

`func (o *AdditionalOrderDetails) GetPlaceOrderOnDateOk() (*string, bool)`

GetPlaceOrderOnDateOk returns a tuple with the PlaceOrderOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOrderOnDate

`func (o *AdditionalOrderDetails) SetPlaceOrderOnDate(v string)`

SetPlaceOrderOnDate sets PlaceOrderOnDate field to given value.

### HasPlaceOrderOnDate

`func (o *AdditionalOrderDetails) HasPlaceOrderOnDate() bool`

HasPlaceOrderOnDate returns a boolean if a field has been set.

### GetClaimRejectedDetails

`func (o *AdditionalOrderDetails) GetClaimRejectedDetails() ClaimRejectionDetails`

GetClaimRejectedDetails returns the ClaimRejectedDetails field if non-nil, zero value otherwise.

### GetClaimRejectedDetailsOk

`func (o *AdditionalOrderDetails) GetClaimRejectedDetailsOk() (*ClaimRejectionDetails, bool)`

GetClaimRejectedDetailsOk returns a tuple with the ClaimRejectedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRejectedDetails

`func (o *AdditionalOrderDetails) SetClaimRejectedDetails(v ClaimRejectionDetails)`

SetClaimRejectedDetails sets ClaimRejectedDetails field to given value.

### HasClaimRejectedDetails

`func (o *AdditionalOrderDetails) HasClaimRejectedDetails() bool`

HasClaimRejectedDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


