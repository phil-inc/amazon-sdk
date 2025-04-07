# PrescriptionFill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FillNumber** | **float32** | A positive integer representing the position of this fill in a                     sequence of fills for a given prescription, starts with 0 and incremented from 1. | 
**DateCreated** | Pointer to **string** | The date-time that this record was created. | [optional] 
**DatePurchased** | Pointer to **string** | The date-time that this fill was purchased. | [optional] 
**DateComplete** | Pointer to **string** | The date-time that this order was shipped and status changed to COMPLETE. | [optional] 
**Status** | [**PrescriptionFillStatus**](PrescriptionFillStatus.md) |  | 
**DaysSupply** | Pointer to **float32** | The number of days supply for this fill. | [optional] 
**Quantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 

## Methods

### NewPrescriptionFill

`func NewPrescriptionFill(fillNumber float32, status PrescriptionFillStatus, ) *PrescriptionFill`

NewPrescriptionFill instantiates a new PrescriptionFill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrescriptionFillWithDefaults

`func NewPrescriptionFillWithDefaults() *PrescriptionFill`

NewPrescriptionFillWithDefaults instantiates a new PrescriptionFill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFillNumber

`func (o *PrescriptionFill) GetFillNumber() float32`

GetFillNumber returns the FillNumber field if non-nil, zero value otherwise.

### GetFillNumberOk

`func (o *PrescriptionFill) GetFillNumberOk() (*float32, bool)`

GetFillNumberOk returns a tuple with the FillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillNumber

`func (o *PrescriptionFill) SetFillNumber(v float32)`

SetFillNumber sets FillNumber field to given value.


### GetDateCreated

`func (o *PrescriptionFill) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PrescriptionFill) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PrescriptionFill) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PrescriptionFill) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDatePurchased

`func (o *PrescriptionFill) GetDatePurchased() string`

GetDatePurchased returns the DatePurchased field if non-nil, zero value otherwise.

### GetDatePurchasedOk

`func (o *PrescriptionFill) GetDatePurchasedOk() (*string, bool)`

GetDatePurchasedOk returns a tuple with the DatePurchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePurchased

`func (o *PrescriptionFill) SetDatePurchased(v string)`

SetDatePurchased sets DatePurchased field to given value.

### HasDatePurchased

`func (o *PrescriptionFill) HasDatePurchased() bool`

HasDatePurchased returns a boolean if a field has been set.

### GetDateComplete

`func (o *PrescriptionFill) GetDateComplete() string`

GetDateComplete returns the DateComplete field if non-nil, zero value otherwise.

### GetDateCompleteOk

`func (o *PrescriptionFill) GetDateCompleteOk() (*string, bool)`

GetDateCompleteOk returns a tuple with the DateComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateComplete

`func (o *PrescriptionFill) SetDateComplete(v string)`

SetDateComplete sets DateComplete field to given value.

### HasDateComplete

`func (o *PrescriptionFill) HasDateComplete() bool`

HasDateComplete returns a boolean if a field has been set.

### GetStatus

`func (o *PrescriptionFill) GetStatus() PrescriptionFillStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrescriptionFill) GetStatusOk() (*PrescriptionFillStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrescriptionFill) SetStatus(v PrescriptionFillStatus)`

SetStatus sets Status field to given value.


### GetDaysSupply

`func (o *PrescriptionFill) GetDaysSupply() float32`

GetDaysSupply returns the DaysSupply field if non-nil, zero value otherwise.

### GetDaysSupplyOk

`func (o *PrescriptionFill) GetDaysSupplyOk() (*float32, bool)`

GetDaysSupplyOk returns a tuple with the DaysSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSupply

`func (o *PrescriptionFill) SetDaysSupply(v float32)`

SetDaysSupply sets DaysSupply field to given value.

### HasDaysSupply

`func (o *PrescriptionFill) HasDaysSupply() bool`

HasDaysSupply returns a boolean if a field has been set.

### GetQuantity

`func (o *PrescriptionFill) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PrescriptionFill) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PrescriptionFill) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PrescriptionFill) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


