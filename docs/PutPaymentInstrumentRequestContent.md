# PutPaymentInstrumentRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerPatientId** | **string** | Identifier of the Patient who owns the Payment Instrument. | 
**PartnerPaymentInstrumentId** | **string** | A unique identifier provided by the partner that identifies the payment instrument, this can be used later for placing orders. | 
**PaymentMethod** | [**PaymentMethod**](PaymentMethod.md) |  | 
**CardDetails** | [**CardDetails**](CardDetails.md) |  | 
**BillingAddress** | [**Address**](Address.md) |  | 

## Methods

### NewPutPaymentInstrumentRequestContent

`func NewPutPaymentInstrumentRequestContent(partnerPatientId string, partnerPaymentInstrumentId string, paymentMethod PaymentMethod, cardDetails CardDetails, billingAddress Address, ) *PutPaymentInstrumentRequestContent`

NewPutPaymentInstrumentRequestContent instantiates a new PutPaymentInstrumentRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPaymentInstrumentRequestContentWithDefaults

`func NewPutPaymentInstrumentRequestContentWithDefaults() *PutPaymentInstrumentRequestContent`

NewPutPaymentInstrumentRequestContentWithDefaults instantiates a new PutPaymentInstrumentRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerPatientId

`func (o *PutPaymentInstrumentRequestContent) GetPartnerPatientId() string`

GetPartnerPatientId returns the PartnerPatientId field if non-nil, zero value otherwise.

### GetPartnerPatientIdOk

`func (o *PutPaymentInstrumentRequestContent) GetPartnerPatientIdOk() (*string, bool)`

GetPartnerPatientIdOk returns a tuple with the PartnerPatientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPatientId

`func (o *PutPaymentInstrumentRequestContent) SetPartnerPatientId(v string)`

SetPartnerPatientId sets PartnerPatientId field to given value.


### GetPartnerPaymentInstrumentId

`func (o *PutPaymentInstrumentRequestContent) GetPartnerPaymentInstrumentId() string`

GetPartnerPaymentInstrumentId returns the PartnerPaymentInstrumentId field if non-nil, zero value otherwise.

### GetPartnerPaymentInstrumentIdOk

`func (o *PutPaymentInstrumentRequestContent) GetPartnerPaymentInstrumentIdOk() (*string, bool)`

GetPartnerPaymentInstrumentIdOk returns a tuple with the PartnerPaymentInstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPaymentInstrumentId

`func (o *PutPaymentInstrumentRequestContent) SetPartnerPaymentInstrumentId(v string)`

SetPartnerPaymentInstrumentId sets PartnerPaymentInstrumentId field to given value.


### GetPaymentMethod

`func (o *PutPaymentInstrumentRequestContent) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PutPaymentInstrumentRequestContent) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PutPaymentInstrumentRequestContent) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetCardDetails

`func (o *PutPaymentInstrumentRequestContent) GetCardDetails() CardDetails`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *PutPaymentInstrumentRequestContent) GetCardDetailsOk() (*CardDetails, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *PutPaymentInstrumentRequestContent) SetCardDetails(v CardDetails)`

SetCardDetails sets CardDetails field to given value.


### GetBillingAddress

`func (o *PutPaymentInstrumentRequestContent) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PutPaymentInstrumentRequestContent) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PutPaymentInstrumentRequestContent) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


