# Pharmacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NcpdpId** | **string** | An NCPDP-assigned national provider identification number that assists pharmacies in their interactions with federal agencies and third party providers. | 
**Npi** | **string** | A National Provider System (NPS) assigned identifier for health care providers. | 
**PharmacyName** | **string** | A string that must contain at least one non-whitespace character, potentially preceded by whitespace.  Here&#39;s how it&#39;s checked: - &#39;^&#39; asserts the beginning of the string. - &#39;\\s*&#39; allows any number of whitespace characters at the start of the string, including none. - &#39;\\S&#39; ensures there is at least one non-whitespace character in the string. - &#39;.*$&#39; allows any characters to follow the non-whitespace character, extending to the end of the string.  This ensures that the string is not entirely whitespace, although it can start with whitespace and can contain any characters after the first non-whitespace character.  Note: This naturally enforces a minimum length of 1 due to 1 non-whitespace character requirement. | 
**DeaNumber** | Pointer to **string** | The Drug Enforcement Administration (DEA) assigned number to all businesses that manufacture or distribute controlled pharmaceuticals,         all health professionals who dispense, administer, or prescribe controlled pharmaceuticals and all pharmacies that dispense prescriptions. | [optional] 
**PharmacistName** | [**Name**](Name.md) |  | 
**PharmacyAddress** | [**Address**](Address.md) |  | 
**PrimaryTelephone** | **string** | Pharmacy telephone number. | 
**Fax** | Pointer to **string** | Pharmacy fax number. | [optional] 

## Methods

### NewPharmacy

`func NewPharmacy(ncpdpId string, npi string, pharmacyName string, pharmacistName Name, pharmacyAddress Address, primaryTelephone string, ) *Pharmacy`

NewPharmacy instantiates a new Pharmacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPharmacyWithDefaults

`func NewPharmacyWithDefaults() *Pharmacy`

NewPharmacyWithDefaults instantiates a new Pharmacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNcpdpId

`func (o *Pharmacy) GetNcpdpId() string`

GetNcpdpId returns the NcpdpId field if non-nil, zero value otherwise.

### GetNcpdpIdOk

`func (o *Pharmacy) GetNcpdpIdOk() (*string, bool)`

GetNcpdpIdOk returns a tuple with the NcpdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcpdpId

`func (o *Pharmacy) SetNcpdpId(v string)`

SetNcpdpId sets NcpdpId field to given value.


### GetNpi

`func (o *Pharmacy) GetNpi() string`

GetNpi returns the Npi field if non-nil, zero value otherwise.

### GetNpiOk

`func (o *Pharmacy) GetNpiOk() (*string, bool)`

GetNpiOk returns a tuple with the Npi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpi

`func (o *Pharmacy) SetNpi(v string)`

SetNpi sets Npi field to given value.


### GetPharmacyName

`func (o *Pharmacy) GetPharmacyName() string`

GetPharmacyName returns the PharmacyName field if non-nil, zero value otherwise.

### GetPharmacyNameOk

`func (o *Pharmacy) GetPharmacyNameOk() (*string, bool)`

GetPharmacyNameOk returns a tuple with the PharmacyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmacyName

`func (o *Pharmacy) SetPharmacyName(v string)`

SetPharmacyName sets PharmacyName field to given value.


### GetDeaNumber

`func (o *Pharmacy) GetDeaNumber() string`

GetDeaNumber returns the DeaNumber field if non-nil, zero value otherwise.

### GetDeaNumberOk

`func (o *Pharmacy) GetDeaNumberOk() (*string, bool)`

GetDeaNumberOk returns a tuple with the DeaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaNumber

`func (o *Pharmacy) SetDeaNumber(v string)`

SetDeaNumber sets DeaNumber field to given value.

### HasDeaNumber

`func (o *Pharmacy) HasDeaNumber() bool`

HasDeaNumber returns a boolean if a field has been set.

### GetPharmacistName

`func (o *Pharmacy) GetPharmacistName() Name`

GetPharmacistName returns the PharmacistName field if non-nil, zero value otherwise.

### GetPharmacistNameOk

`func (o *Pharmacy) GetPharmacistNameOk() (*Name, bool)`

GetPharmacistNameOk returns a tuple with the PharmacistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmacistName

`func (o *Pharmacy) SetPharmacistName(v Name)`

SetPharmacistName sets PharmacistName field to given value.


### GetPharmacyAddress

`func (o *Pharmacy) GetPharmacyAddress() Address`

GetPharmacyAddress returns the PharmacyAddress field if non-nil, zero value otherwise.

### GetPharmacyAddressOk

`func (o *Pharmacy) GetPharmacyAddressOk() (*Address, bool)`

GetPharmacyAddressOk returns a tuple with the PharmacyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmacyAddress

`func (o *Pharmacy) SetPharmacyAddress(v Address)`

SetPharmacyAddress sets PharmacyAddress field to given value.


### GetPrimaryTelephone

`func (o *Pharmacy) GetPrimaryTelephone() string`

GetPrimaryTelephone returns the PrimaryTelephone field if non-nil, zero value otherwise.

### GetPrimaryTelephoneOk

`func (o *Pharmacy) GetPrimaryTelephoneOk() (*string, bool)`

GetPrimaryTelephoneOk returns a tuple with the PrimaryTelephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTelephone

`func (o *Pharmacy) SetPrimaryTelephone(v string)`

SetPrimaryTelephone sets PrimaryTelephone field to given value.


### GetFax

`func (o *Pharmacy) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Pharmacy) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Pharmacy) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Pharmacy) HasFax() bool`

HasFax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


