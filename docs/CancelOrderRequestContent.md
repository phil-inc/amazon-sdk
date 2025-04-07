# CancelOrderRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancellationReason** | **string** | The reason for cancelling this order. | 

## Methods

### NewCancelOrderRequestContent

`func NewCancelOrderRequestContent(cancellationReason string, ) *CancelOrderRequestContent`

NewCancelOrderRequestContent instantiates a new CancelOrderRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderRequestContentWithDefaults

`func NewCancelOrderRequestContentWithDefaults() *CancelOrderRequestContent`

NewCancelOrderRequestContentWithDefaults instantiates a new CancelOrderRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancellationReason

`func (o *CancelOrderRequestContent) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *CancelOrderRequestContent) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *CancelOrderRequestContent) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


