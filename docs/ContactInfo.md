# ContactInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactSystemType** | [**ContactSystemType**](ContactSystemType.md) |  | 
**CountryCode** | **float32** | Country code of the phone number | 
**PhoneNumber** | **string** | Phone number of the patient. | 

## Methods

### NewContactInfo

`func NewContactInfo(contactSystemType ContactSystemType, countryCode float32, phoneNumber string, ) *ContactInfo`

NewContactInfo instantiates a new ContactInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactInfoWithDefaults

`func NewContactInfoWithDefaults() *ContactInfo`

NewContactInfoWithDefaults instantiates a new ContactInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactSystemType

`func (o *ContactInfo) GetContactSystemType() ContactSystemType`

GetContactSystemType returns the ContactSystemType field if non-nil, zero value otherwise.

### GetContactSystemTypeOk

`func (o *ContactInfo) GetContactSystemTypeOk() (*ContactSystemType, bool)`

GetContactSystemTypeOk returns a tuple with the ContactSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSystemType

`func (o *ContactInfo) SetContactSystemType(v ContactSystemType)`

SetContactSystemType sets ContactSystemType field to given value.


### GetCountryCode

`func (o *ContactInfo) GetCountryCode() float32`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ContactInfo) GetCountryCodeOk() (*float32, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ContactInfo) SetCountryCode(v float32)`

SetCountryCode sets CountryCode field to given value.


### GetPhoneNumber

`func (o *ContactInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


