# ShippingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingId** | Pointer to **string** | The tracking ID associated with the shipment. | [optional] 
**Carrier** | Pointer to **string** | The name of the carrier fulfilling the delivery of the shipment. | [optional] 
**EstimatedArrival** | Pointer to **string** | The estimated arrival date in the format &#39;YYYY-MM-DDDD&#39; in UTC, e.g., 2023-04-17T03:00:00.000Z. | [optional] 

## Methods

### NewShippingDetails

`func NewShippingDetails() *ShippingDetails`

NewShippingDetails instantiates a new ShippingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingDetailsWithDefaults

`func NewShippingDetailsWithDefaults() *ShippingDetails`

NewShippingDetailsWithDefaults instantiates a new ShippingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingId

`func (o *ShippingDetails) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *ShippingDetails) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *ShippingDetails) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *ShippingDetails) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetCarrier

`func (o *ShippingDetails) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ShippingDetails) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ShippingDetails) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ShippingDetails) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetEstimatedArrival

`func (o *ShippingDetails) GetEstimatedArrival() string`

GetEstimatedArrival returns the EstimatedArrival field if non-nil, zero value otherwise.

### GetEstimatedArrivalOk

`func (o *ShippingDetails) GetEstimatedArrivalOk() (*string, bool)`

GetEstimatedArrivalOk returns a tuple with the EstimatedArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrival

`func (o *ShippingDetails) SetEstimatedArrival(v string)`

SetEstimatedArrival sets EstimatedArrival field to given value.

### HasEstimatedArrival

`func (o *ShippingDetails) HasEstimatedArrival() bool`

HasEstimatedArrival returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


