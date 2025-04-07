# Insurance

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
**PartnerInsuranceId** | **string** | A unique identifier to identify the partner-provided insurance. | 
**EncryptedCardholderId** | **string** | Insurance card number encrypted using the patient-information-encryption-key. The unencrypted value cannot exceed 20 characters. | 

## Methods

### NewInsurance

`func NewInsurance(binNumber string, personCode string, partnerInsuranceId string, encryptedCardholderId string, ) *Insurance`

NewInsurance instantiates a new Insurance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsuranceWithDefaults

`func NewInsuranceWithDefaults() *Insurance`

NewInsuranceWithDefaults instantiates a new Insurance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcn

`func (o *Insurance) GetPcn() string`

GetPcn returns the Pcn field if non-nil, zero value otherwise.

### GetPcnOk

`func (o *Insurance) GetPcnOk() (*string, bool)`

GetPcnOk returns a tuple with the Pcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcn

`func (o *Insurance) SetPcn(v string)`

SetPcn sets Pcn field to given value.

### HasPcn

`func (o *Insurance) HasPcn() bool`

HasPcn returns a boolean if a field has been set.

### GetGroupNumber

`func (o *Insurance) GetGroupNumber() string`

GetGroupNumber returns the GroupNumber field if non-nil, zero value otherwise.

### GetGroupNumberOk

`func (o *Insurance) GetGroupNumberOk() (*string, bool)`

GetGroupNumberOk returns a tuple with the GroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumber

`func (o *Insurance) SetGroupNumber(v string)`

SetGroupNumber sets GroupNumber field to given value.

### HasGroupNumber

`func (o *Insurance) HasGroupNumber() bool`

HasGroupNumber returns a boolean if a field has been set.

### GetBinNumber

`func (o *Insurance) GetBinNumber() string`

GetBinNumber returns the BinNumber field if non-nil, zero value otherwise.

### GetBinNumberOk

`func (o *Insurance) GetBinNumberOk() (*string, bool)`

GetBinNumberOk returns a tuple with the BinNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinNumber

`func (o *Insurance) SetBinNumber(v string)`

SetBinNumber sets BinNumber field to given value.


### GetPayerType

`func (o *Insurance) GetPayerType() PayerType`

GetPayerType returns the PayerType field if non-nil, zero value otherwise.

### GetPayerTypeOk

`func (o *Insurance) GetPayerTypeOk() (*PayerType, bool)`

GetPayerTypeOk returns a tuple with the PayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerType

`func (o *Insurance) SetPayerType(v PayerType)`

SetPayerType sets PayerType field to given value.

### HasPayerType

`func (o *Insurance) HasPayerType() bool`

HasPayerType returns a boolean if a field has been set.

### GetPayerName

`func (o *Insurance) GetPayerName() string`

GetPayerName returns the PayerName field if non-nil, zero value otherwise.

### GetPayerNameOk

`func (o *Insurance) GetPayerNameOk() (*string, bool)`

GetPayerNameOk returns a tuple with the PayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerName

`func (o *Insurance) SetPayerName(v string)`

SetPayerName sets PayerName field to given value.

### HasPayerName

`func (o *Insurance) HasPayerName() bool`

HasPayerName returns a boolean if a field has been set.

### GetStateCode

`func (o *Insurance) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *Insurance) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *Insurance) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *Insurance) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetPersonCode

`func (o *Insurance) GetPersonCode() string`

GetPersonCode returns the PersonCode field if non-nil, zero value otherwise.

### GetPersonCodeOk

`func (o *Insurance) GetPersonCodeOk() (*string, bool)`

GetPersonCodeOk returns a tuple with the PersonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonCode

`func (o *Insurance) SetPersonCode(v string)`

SetPersonCode sets PersonCode field to given value.


### GetRelationshipCode

`func (o *Insurance) GetRelationshipCode() PatientRelationship`

GetRelationshipCode returns the RelationshipCode field if non-nil, zero value otherwise.

### GetRelationshipCodeOk

`func (o *Insurance) GetRelationshipCodeOk() (*PatientRelationship, bool)`

GetRelationshipCodeOk returns a tuple with the RelationshipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipCode

`func (o *Insurance) SetRelationshipCode(v PatientRelationship)`

SetRelationshipCode sets RelationshipCode field to given value.

### HasRelationshipCode

`func (o *Insurance) HasRelationshipCode() bool`

HasRelationshipCode returns a boolean if a field has been set.

### GetStartDate

`func (o *Insurance) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Insurance) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Insurance) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Insurance) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Insurance) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Insurance) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Insurance) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Insurance) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetPartnerInsuranceId

`func (o *Insurance) GetPartnerInsuranceId() string`

GetPartnerInsuranceId returns the PartnerInsuranceId field if non-nil, zero value otherwise.

### GetPartnerInsuranceIdOk

`func (o *Insurance) GetPartnerInsuranceIdOk() (*string, bool)`

GetPartnerInsuranceIdOk returns a tuple with the PartnerInsuranceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerInsuranceId

`func (o *Insurance) SetPartnerInsuranceId(v string)`

SetPartnerInsuranceId sets PartnerInsuranceId field to given value.


### GetEncryptedCardholderId

`func (o *Insurance) GetEncryptedCardholderId() string`

GetEncryptedCardholderId returns the EncryptedCardholderId field if non-nil, zero value otherwise.

### GetEncryptedCardholderIdOk

`func (o *Insurance) GetEncryptedCardholderIdOk() (*string, bool)`

GetEncryptedCardholderIdOk returns a tuple with the EncryptedCardholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedCardholderId

`func (o *Insurance) SetEncryptedCardholderId(v string)`

SetEncryptedCardholderId sets EncryptedCardholderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


