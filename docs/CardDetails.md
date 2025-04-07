# CardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedCardNumber** | **string** | The encrypted card number. | 
**CardHolderName** | **string** | The Name of Card Holder owner as it appears on the Card. | 
**ExpirationMonth** | **float32** | The expiration month of the card. | 
**ExpirationYear** | **float32** | The expiration year of the card. | 

## Methods

### NewCardDetails

`func NewCardDetails(encryptedCardNumber string, cardHolderName string, expirationMonth float32, expirationYear float32, ) *CardDetails`

NewCardDetails instantiates a new CardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDetailsWithDefaults

`func NewCardDetailsWithDefaults() *CardDetails`

NewCardDetailsWithDefaults instantiates a new CardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedCardNumber

`func (o *CardDetails) GetEncryptedCardNumber() string`

GetEncryptedCardNumber returns the EncryptedCardNumber field if non-nil, zero value otherwise.

### GetEncryptedCardNumberOk

`func (o *CardDetails) GetEncryptedCardNumberOk() (*string, bool)`

GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedCardNumber

`func (o *CardDetails) SetEncryptedCardNumber(v string)`

SetEncryptedCardNumber sets EncryptedCardNumber field to given value.


### GetCardHolderName

`func (o *CardDetails) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *CardDetails) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *CardDetails) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.


### GetExpirationMonth

`func (o *CardDetails) GetExpirationMonth() float32`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *CardDetails) GetExpirationMonthOk() (*float32, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *CardDetails) SetExpirationMonth(v float32)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationYear

`func (o *CardDetails) GetExpirationYear() float32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *CardDetails) GetExpirationYearOk() (*float32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *CardDetails) SetExpirationYear(v float32)`

SetExpirationYear sets ExpirationYear field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


