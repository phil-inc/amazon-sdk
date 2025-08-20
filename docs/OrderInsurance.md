# OrderInsurance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pcn** | Pointer to **string** | Processor Control Number assigned by the processor. PBMs use it to distinguish between plans. It&#39;s a secondary identifier used in addition to the binNumber. | [optional] 
**GroupNumber** | Pointer to **string** | The insurance group number found on the insurance card that signifies which group the cardholder belongs to. | [optional] 
**BinNumber** | **string** | A 6-digit number that identifies the Pharmacy Benefit Manager (PBM) that this insurance uses for pharmacy claims. | 
**PayerType** | Pointer to [**PayerType**](PayerType.md) |  | [optional] 
**PayerName** | Pointer to **string** | The name of the payer associated with the insurance plan. | [optional] 
**StateCode** | Pointer to **string** | Optional state code related to the insurance plan, used for state-specific insurance information. | [optional] 
**PersonCode** | **string** | Insurance coverages are shared between family members. Each person code identifies a covered individual family member. The primary insurance holder is often 01 or 001. | 
**RelationshipCode** | Pointer to [**PatientRelationship**](PatientRelationship.md) |  | [optional] 
**StartDate** | Pointer to **string** | Start date for the insurance coverage. This indicates when the insurance coverage begins. | [optional] 
**ExpiryDate** | Pointer to **string** | Expiration date for the insurance coverage. This indicates when the insurance will no longer be valid. | [optional] 
**CardholderId** | Pointer to **string** | The memberId of this person&#39;s insurance information. It includes their suffix and is the primary identifier for their insurance coverage. Must be alphanumeric and between 1 to 20 characters. | [optional] 

## Methods

### NewOrderInsurance

`func NewOrderInsurance(binNumber string, personCode string, ) *OrderInsurance`

NewOrderInsurance instantiates a new OrderInsurance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInsuranceWithDefaults

`func NewOrderInsuranceWithDefaults() *OrderInsurance`

NewOrderInsuranceWithDefaults instantiates a new OrderInsurance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcn

`func (o *OrderInsurance) GetPcn() string`

GetPcn returns the Pcn field if non-nil, zero value otherwise.

### GetPcnOk

`func (o *OrderInsurance) GetPcnOk() (*string, bool)`

GetPcnOk returns a tuple with the Pcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcn

`func (o *OrderInsurance) SetPcn(v string)`

SetPcn sets Pcn field to given value.

### HasPcn

`func (o *OrderInsurance) HasPcn() bool`

HasPcn returns a boolean if a field has been set.

### GetGroupNumber

`func (o *OrderInsurance) GetGroupNumber() string`

GetGroupNumber returns the GroupNumber field if non-nil, zero value otherwise.

### GetGroupNumberOk

`func (o *OrderInsurance) GetGroupNumberOk() (*string, bool)`

GetGroupNumberOk returns a tuple with the GroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumber

`func (o *OrderInsurance) SetGroupNumber(v string)`

SetGroupNumber sets GroupNumber field to given value.

### HasGroupNumber

`func (o *OrderInsurance) HasGroupNumber() bool`

HasGroupNumber returns a boolean if a field has been set.

### GetBinNumber

`func (o *OrderInsurance) GetBinNumber() string`

GetBinNumber returns the BinNumber field if non-nil, zero value otherwise.

### GetBinNumberOk

`func (o *OrderInsurance) GetBinNumberOk() (*string, bool)`

GetBinNumberOk returns a tuple with the BinNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinNumber

`func (o *OrderInsurance) SetBinNumber(v string)`

SetBinNumber sets BinNumber field to given value.


### GetPayerType

`func (o *OrderInsurance) GetPayerType() PayerType`

GetPayerType returns the PayerType field if non-nil, zero value otherwise.

### GetPayerTypeOk

`func (o *OrderInsurance) GetPayerTypeOk() (*PayerType, bool)`

GetPayerTypeOk returns a tuple with the PayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerType

`func (o *OrderInsurance) SetPayerType(v PayerType)`

SetPayerType sets PayerType field to given value.

### HasPayerType

`func (o *OrderInsurance) HasPayerType() bool`

HasPayerType returns a boolean if a field has been set.

### GetPayerName

`func (o *OrderInsurance) GetPayerName() string`

GetPayerName returns the PayerName field if non-nil, zero value otherwise.

### GetPayerNameOk

`func (o *OrderInsurance) GetPayerNameOk() (*string, bool)`

GetPayerNameOk returns a tuple with the PayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerName

`func (o *OrderInsurance) SetPayerName(v string)`

SetPayerName sets PayerName field to given value.

### HasPayerName

`func (o *OrderInsurance) HasPayerName() bool`

HasPayerName returns a boolean if a field has been set.

### GetStateCode

`func (o *OrderInsurance) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *OrderInsurance) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *OrderInsurance) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *OrderInsurance) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetPersonCode

`func (o *OrderInsurance) GetPersonCode() string`

GetPersonCode returns the PersonCode field if non-nil, zero value otherwise.

### GetPersonCodeOk

`func (o *OrderInsurance) GetPersonCodeOk() (*string, bool)`

GetPersonCodeOk returns a tuple with the PersonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonCode

`func (o *OrderInsurance) SetPersonCode(v string)`

SetPersonCode sets PersonCode field to given value.


### GetRelationshipCode

`func (o *OrderInsurance) GetRelationshipCode() PatientRelationship`

GetRelationshipCode returns the RelationshipCode field if non-nil, zero value otherwise.

### GetRelationshipCodeOk

`func (o *OrderInsurance) GetRelationshipCodeOk() (*PatientRelationship, bool)`

GetRelationshipCodeOk returns a tuple with the RelationshipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipCode

`func (o *OrderInsurance) SetRelationshipCode(v PatientRelationship)`

SetRelationshipCode sets RelationshipCode field to given value.

### HasRelationshipCode

`func (o *OrderInsurance) HasRelationshipCode() bool`

HasRelationshipCode returns a boolean if a field has been set.

### GetStartDate

`func (o *OrderInsurance) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OrderInsurance) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OrderInsurance) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *OrderInsurance) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *OrderInsurance) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *OrderInsurance) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *OrderInsurance) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *OrderInsurance) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetCardholderId

`func (o *OrderInsurance) GetCardholderId() string`

GetCardholderId returns the CardholderId field if non-nil, zero value otherwise.

### GetCardholderIdOk

`func (o *OrderInsurance) GetCardholderIdOk() (*string, bool)`

GetCardholderIdOk returns a tuple with the CardholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderId

`func (o *OrderInsurance) SetCardholderId(v string)`

SetCardholderId sets CardholderId field to given value.

### HasCardholderId

`func (o *OrderInsurance) HasCardholderId() bool`

HasCardholderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


