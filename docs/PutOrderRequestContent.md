# PutOrderRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerOrderId** | **string** | A unique Order identifier provided by the partner. | 
**PartnerPrescriptionId** | **string** | A unique Prescription Identifier provided by the partner. | 
**ShippingAddress** | [**Address**](Address.md) |  | 
**PartnerPaymentInstrumentId** | Pointer to **string** | Customer&#39;s payment instrument ID, previously created using the payment-instrument API. This field is required when the quoted copay amount is greater than zero and placeOrderOnDate is null. | [optional] 
**QuotedCopay** | Pointer to **string** | An estimated copay for the customer. If the actual copay is higher than the estimated copay by a certain amount($XX)     or percentage (XX%), the shipment will be put on hold. | [optional] 
**PlaceOrderOnDate** | Pointer to **string** | Order will be held until this date and sent for order placement only when the date is passed.     It is quoted in the format &#39;YYYY-MM-DD&#39;. | [optional] 
**OrderInsuranceDetails** | Pointer to [**OrderInsuranceDetails**](OrderInsuranceDetails.md) |  | [optional] 

## Methods

### NewPutOrderRequestContent

`func NewPutOrderRequestContent(partnerOrderId string, partnerPrescriptionId string, shippingAddress Address, ) *PutOrderRequestContent`

NewPutOrderRequestContent instantiates a new PutOrderRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutOrderRequestContentWithDefaults

`func NewPutOrderRequestContentWithDefaults() *PutOrderRequestContent`

NewPutOrderRequestContentWithDefaults instantiates a new PutOrderRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerOrderId

`func (o *PutOrderRequestContent) GetPartnerOrderId() string`

GetPartnerOrderId returns the PartnerOrderId field if non-nil, zero value otherwise.

### GetPartnerOrderIdOk

`func (o *PutOrderRequestContent) GetPartnerOrderIdOk() (*string, bool)`

GetPartnerOrderIdOk returns a tuple with the PartnerOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerOrderId

`func (o *PutOrderRequestContent) SetPartnerOrderId(v string)`

SetPartnerOrderId sets PartnerOrderId field to given value.


### GetPartnerPrescriptionId

`func (o *PutOrderRequestContent) GetPartnerPrescriptionId() string`

GetPartnerPrescriptionId returns the PartnerPrescriptionId field if non-nil, zero value otherwise.

### GetPartnerPrescriptionIdOk

`func (o *PutOrderRequestContent) GetPartnerPrescriptionIdOk() (*string, bool)`

GetPartnerPrescriptionIdOk returns a tuple with the PartnerPrescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrescriptionId

`func (o *PutOrderRequestContent) SetPartnerPrescriptionId(v string)`

SetPartnerPrescriptionId sets PartnerPrescriptionId field to given value.


### GetShippingAddress

`func (o *PutOrderRequestContent) GetShippingAddress() Address`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *PutOrderRequestContent) GetShippingAddressOk() (*Address, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *PutOrderRequestContent) SetShippingAddress(v Address)`

SetShippingAddress sets ShippingAddress field to given value.


### GetPartnerPaymentInstrumentId

`func (o *PutOrderRequestContent) GetPartnerPaymentInstrumentId() string`

GetPartnerPaymentInstrumentId returns the PartnerPaymentInstrumentId field if non-nil, zero value otherwise.

### GetPartnerPaymentInstrumentIdOk

`func (o *PutOrderRequestContent) GetPartnerPaymentInstrumentIdOk() (*string, bool)`

GetPartnerPaymentInstrumentIdOk returns a tuple with the PartnerPaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPaymentInstrumentId

`func (o *PutOrderRequestContent) SetPartnerPaymentInstrumentId(v string)`

SetPartnerPaymentInstrumentId sets PartnerPaymentInstrumentId field to given value.

### HasPartnerPaymentInstrumentId

`func (o *PutOrderRequestContent) HasPartnerPaymentInstrumentId() bool`

HasPartnerPaymentInstrumentId returns a boolean if a field has been set.

### GetQuotedCopay

`func (o *PutOrderRequestContent) GetQuotedCopay() string`

GetQuotedCopay returns the QuotedCopay field if non-nil, zero value otherwise.

### GetQuotedCopayOk

`func (o *PutOrderRequestContent) GetQuotedCopayOk() (*string, bool)`

GetQuotedCopayOk returns a tuple with the QuotedCopay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedCopay

`func (o *PutOrderRequestContent) SetQuotedCopay(v string)`

SetQuotedCopay sets QuotedCopay field to given value.

### HasQuotedCopay

`func (o *PutOrderRequestContent) HasQuotedCopay() bool`

HasQuotedCopay returns a boolean if a field has been set.

### GetPlaceOrderOnDate

`func (o *PutOrderRequestContent) GetPlaceOrderOnDate() string`

GetPlaceOrderOnDate returns the PlaceOrderOnDate field if non-nil, zero value otherwise.

### GetPlaceOrderOnDateOk

`func (o *PutOrderRequestContent) GetPlaceOrderOnDateOk() (*string, bool)`

GetPlaceOrderOnDateOk returns a tuple with the PlaceOrderOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOrderOnDate

`func (o *PutOrderRequestContent) SetPlaceOrderOnDate(v string)`

SetPlaceOrderOnDate sets PlaceOrderOnDate field to given value.

### HasPlaceOrderOnDate

`func (o *PutOrderRequestContent) HasPlaceOrderOnDate() bool`

HasPlaceOrderOnDate returns a boolean if a field has been set.

### GetOrderInsuranceDetails

`func (o *PutOrderRequestContent) GetOrderInsuranceDetails() OrderInsuranceDetails`

GetOrderInsuranceDetails returns the OrderInsuranceDetails field if non-nil, zero value otherwise.

### GetOrderInsuranceDetailsOk

`func (o *PutOrderRequestContent) GetOrderInsuranceDetailsOk() (*OrderInsuranceDetails, bool)`

GetOrderInsuranceDetailsOk returns a tuple with the OrderInsuranceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderInsuranceDetails

`func (o *PutOrderRequestContent) SetOrderInsuranceDetails(v OrderInsuranceDetails)`

SetOrderInsuranceDetails sets OrderInsuranceDetails field to given value.

### HasOrderInsuranceDetails

`func (o *PutOrderRequestContent) HasOrderInsuranceDetails() bool`

HasOrderInsuranceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


