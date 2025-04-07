# RejectedInsuranceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pcn** | Pointer to **string** | Processor Control Number assigned by the processor. PBMs use it to distinguish between plans. It&#39;s a secondary identifier used in addition to the binNumber. | [optional] 
**GroupNumber** | Pointer to **string** | The insurance group number found on the insurance card that signifies which group the cardholder belongs to. | [optional] 
**BinNumber** | **string** | A 6-digit number that identifies the Pharmacy Benefit Manager (PBM) that this insurance uses for pharmacy claims. | 
**PayerType** | Pointer to [**PayerType**](PayerType.md) |  | [optional] 
**PayerName** | Pointer to **string** | The name of the payer associated with the insurance plan. | [optional] 
**StateCode** | Pointer to **string** | Optional state code related to the insurance plan, used for state-specific insurance information. | [optional] 

## Methods

### NewRejectedInsuranceDetails

`func NewRejectedInsuranceDetails(binNumber string, ) *RejectedInsuranceDetails`

NewRejectedInsuranceDetails instantiates a new RejectedInsuranceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectedInsuranceDetailsWithDefaults

`func NewRejectedInsuranceDetailsWithDefaults() *RejectedInsuranceDetails`

NewRejectedInsuranceDetailsWithDefaults instantiates a new RejectedInsuranceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcn

`func (o *RejectedInsuranceDetails) GetPcn() string`

GetPcn returns the Pcn field if non-nil, zero value otherwise.

### GetPcnOk

`func (o *RejectedInsuranceDetails) GetPcnOk() (*string, bool)`

GetPcnOk returns a tuple with the Pcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcn

`func (o *RejectedInsuranceDetails) SetPcn(v string)`

SetPcn sets Pcn field to given value.

### HasPcn

`func (o *RejectedInsuranceDetails) HasPcn() bool`

HasPcn returns a boolean if a field has been set.

### GetGroupNumber

`func (o *RejectedInsuranceDetails) GetGroupNumber() string`

GetGroupNumber returns the GroupNumber field if non-nil, zero value otherwise.

### GetGroupNumberOk

`func (o *RejectedInsuranceDetails) GetGroupNumberOk() (*string, bool)`

GetGroupNumberOk returns a tuple with the GroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumber

`func (o *RejectedInsuranceDetails) SetGroupNumber(v string)`

SetGroupNumber sets GroupNumber field to given value.

### HasGroupNumber

`func (o *RejectedInsuranceDetails) HasGroupNumber() bool`

HasGroupNumber returns a boolean if a field has been set.

### GetBinNumber

`func (o *RejectedInsuranceDetails) GetBinNumber() string`

GetBinNumber returns the BinNumber field if non-nil, zero value otherwise.

### GetBinNumberOk

`func (o *RejectedInsuranceDetails) GetBinNumberOk() (*string, bool)`

GetBinNumberOk returns a tuple with the BinNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinNumber

`func (o *RejectedInsuranceDetails) SetBinNumber(v string)`

SetBinNumber sets BinNumber field to given value.


### GetPayerType

`func (o *RejectedInsuranceDetails) GetPayerType() PayerType`

GetPayerType returns the PayerType field if non-nil, zero value otherwise.

### GetPayerTypeOk

`func (o *RejectedInsuranceDetails) GetPayerTypeOk() (*PayerType, bool)`

GetPayerTypeOk returns a tuple with the PayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerType

`func (o *RejectedInsuranceDetails) SetPayerType(v PayerType)`

SetPayerType sets PayerType field to given value.

### HasPayerType

`func (o *RejectedInsuranceDetails) HasPayerType() bool`

HasPayerType returns a boolean if a field has been set.

### GetPayerName

`func (o *RejectedInsuranceDetails) GetPayerName() string`

GetPayerName returns the PayerName field if non-nil, zero value otherwise.

### GetPayerNameOk

`func (o *RejectedInsuranceDetails) GetPayerNameOk() (*string, bool)`

GetPayerNameOk returns a tuple with the PayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerName

`func (o *RejectedInsuranceDetails) SetPayerName(v string)`

SetPayerName sets PayerName field to given value.

### HasPayerName

`func (o *RejectedInsuranceDetails) HasPayerName() bool`

HasPayerName returns a boolean if a field has been set.

### GetStateCode

`func (o *RejectedInsuranceDetails) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *RejectedInsuranceDetails) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *RejectedInsuranceDetails) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *RejectedInsuranceDetails) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


