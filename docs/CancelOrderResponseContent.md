# CancelOrderResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerOrderId** | **string** | The unique Order identifier provided during the creation of the order. | 

## Methods

### NewCancelOrderResponseContent

`func NewCancelOrderResponseContent(partnerOrderId string, ) *CancelOrderResponseContent`

NewCancelOrderResponseContent instantiates a new CancelOrderResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderResponseContentWithDefaults

`func NewCancelOrderResponseContentWithDefaults() *CancelOrderResponseContent`

NewCancelOrderResponseContentWithDefaults instantiates a new CancelOrderResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerOrderId

`func (o *CancelOrderResponseContent) GetPartnerOrderId() string`

GetPartnerOrderId returns the PartnerOrderId field if non-nil, zero value otherwise.

### GetPartnerOrderIdOk

`func (o *CancelOrderResponseContent) GetPartnerOrderIdOk() (*string, bool)`

GetPartnerOrderIdOk returns a tuple with the PartnerOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerOrderId

`func (o *CancelOrderResponseContent) SetPartnerOrderId(v string)`

SetPartnerOrderId sets PartnerOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


