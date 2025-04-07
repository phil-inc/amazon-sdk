# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**PartnerPrescriptionId** | **string** | The partner ID of the prescription this order is associated with. | 
**StatusReasonCode** | Pointer to [**OrderStatusReasonCode**](OrderStatusReasonCode.md) |  | [optional] 
**StatusDescription** | Pointer to **string** | A more detailed explanation of the order&#39;s status. | [optional] 
**AdditionalOrderDetails** | Pointer to [**AdditionalOrderDetails**](AdditionalOrderDetails.md) |  | [optional] 
**FilledDate** | Pointer to **string** | The date the fill was created, in the format &#39;YYYY-MM-DD&#39;. | [optional] 
**ShipDate** | Pointer to **string** | The date the fill was shipped, in the format &#39;YYYY-MM-DD&#39;. | [optional] 
**FillNumber** | Pointer to **float32** | The number of times this prescription has been filled. | [optional] 
**QuantityDispensed** | Pointer to **float32** | The number of units in dispense. | [optional] 

## Methods

### NewOrder

`func NewOrder(status OrderStatus, partnerPrescriptionId string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Order) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetPartnerPrescriptionId

`func (o *Order) GetPartnerPrescriptionId() string`

GetPartnerPrescriptionId returns the PartnerPrescriptionId field if non-nil, zero value otherwise.

### GetPartnerPrescriptionIdOk

`func (o *Order) GetPartnerPrescriptionIdOk() (*string, bool)`

GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrescriptionId

`func (o *Order) SetPartnerPrescriptionId(v string)`

SetPartnerPrescriptionId sets PartnerPrescriptionId field to given value.


### GetStatusReasonCode

`func (o *Order) GetStatusReasonCode() OrderStatusReasonCode`

GetStatusReasonCode returns the StatusReasonCode field if non-nil, zero value otherwise.

### GetStatusReasonCodeOk

`func (o *Order) GetStatusReasonCodeOk() (*OrderStatusReasonCode, bool)`

GetStatusReasonCodeOk returns a tuple with the StatusReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReasonCode

`func (o *Order) SetStatusReasonCode(v OrderStatusReasonCode)`

SetStatusReasonCode sets StatusReasonCode field to given value.

### HasStatusReasonCode

`func (o *Order) HasStatusReasonCode() bool`

HasStatusReasonCode returns a boolean if a field has been set.

### GetStatusDescription

`func (o *Order) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *Order) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *Order) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *Order) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetAdditionalOrderDetails

`func (o *Order) GetAdditionalOrderDetails() AdditionalOrderDetails`

GetAdditionalOrderDetails returns the AdditionalOrderDetails field if non-nil, zero value otherwise.

### GetAdditionalOrderDetailsOk

`func (o *Order) GetAdditionalOrderDetailsOk() (*AdditionalOrderDetails, bool)`

GetAdditionalOrderDetailsOk returns a tuple with the AdditionalOrderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalOrderDetails

`func (o *Order) SetAdditionalOrderDetails(v AdditionalOrderDetails)`

SetAdditionalOrderDetails sets AdditionalOrderDetails field to given value.

### HasAdditionalOrderDetails

`func (o *Order) HasAdditionalOrderDetails() bool`

HasAdditionalOrderDetails returns a boolean if a field has been set.

### GetFilledDate

`func (o *Order) GetFilledDate() string`

GetFilledDate returns the FilledDate field if non-nil, zero value otherwise.

### GetFilledDateOk

`func (o *Order) GetFilledDateOk() (*string, bool)`

GetFilledDateOk returns a tuple with the FilledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledDate

`func (o *Order) SetFilledDate(v string)`

SetFilledDate sets FilledDate field to given value.

### HasFilledDate

`func (o *Order) HasFilledDate() bool`

HasFilledDate returns a boolean if a field has been set.

### GetShipDate

`func (o *Order) GetShipDate() string`

GetShipDate returns the ShipDate field if non-nil, zero value otherwise.

### GetShipDateOk

`func (o *Order) GetShipDateOk() (*string, bool)`

GetShipDateOk returns a tuple with the ShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipDate

`func (o *Order) SetShipDate(v string)`

SetShipDate sets ShipDate field to given value.

### HasShipDate

`func (o *Order) HasShipDate() bool`

HasShipDate returns a boolean if a field has been set.

### GetFillNumber

`func (o *Order) GetFillNumber() float32`

GetFillNumber returns the FillNumber field if non-nil, zero value otherwise.

### GetFillNumberOk

`func (o *Order) GetFillNumberOk() (*float32, bool)`

GetFillNumberOk returns a tuple with the FillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillNumber

`func (o *Order) SetFillNumber(v float32)`

SetFillNumber sets FillNumber field to given value.

### HasFillNumber

`func (o *Order) HasFillNumber() bool`

HasFillNumber returns a boolean if a field has been set.

### GetQuantityDispensed

`func (o *Order) GetQuantityDispensed() float32`

GetQuantityDispensed returns the QuantityDispensed field if non-nil, zero value otherwise.

### GetQuantityDispensedOk

`func (o *Order) GetQuantityDispensedOk() (*float32, bool)`

GetQuantityDispensedOk returns a tuple with the QuantityDispensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDispensed

`func (o *Order) SetQuantityDispensed(v float32)`

SetQuantityDispensed sets QuantityDispensed field to given value.

### HasQuantityDispensed

`func (o *Order) HasQuantityDispensed() bool`

HasQuantityDispensed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


