# Prescriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Npi** | **string** | A National Provider System (NPS) assigned identifier for the prescriber. | 
**PrescriberName** | [**Name**](Name.md) |  | 
**PrescriberAddress** | [**Address**](Address.md) |  | 
**PrimaryTelephone** | **string** | The primary telephone number of the prescriber. | 
**DeaNumber** | Pointer to **string** | The Drug Enforcement Administration (DEA) assigned number to the prescriber. | [optional] 
**StateLicenseNumber** | Pointer to **string** | The prescriber&#39;s state license number. | [optional] 
**BusinessName** | Pointer to **string** | The business name associated with the prescriber. | [optional] 
**Fax** | Pointer to **string** | The fax number of the prescriber. | [optional] 

## Methods

### NewPrescriber

`func NewPrescriber(npi string, prescriberName Name, prescriberAddress Address, primaryTelephone string, ) *Prescriber`

NewPrescriber instantiates a new Prescriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrescriberWithDefaults

`func NewPrescriberWithDefaults() *Prescriber`

NewPrescriberWithDefaults instantiates a new Prescriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNpi

`func (o *Prescriber) GetNpi() string`

GetNpi returns the Npi field if non-nil, zero value otherwise.

### GetNpiOk

`func (o *Prescriber) GetNpiOk() (*string, bool)`

GetNpiOk returns a tuple with the Npi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpi

`func (o *Prescriber) SetNpi(v string)`

SetNpi sets Npi field to given value.


### GetPrescriberName

`func (o *Prescriber) GetPrescriberName() Name`

GetPrescriberName returns the PrescriberName field if non-nil, zero value otherwise.

### GetPrescriberNameOk

`func (o *Prescriber) GetPrescriberNameOk() (*Name, bool)`

GetPrescriberNameOk returns a tuple with the PrescriberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriberName

`func (o *Prescriber) SetPrescriberName(v Name)`

SetPrescriberName sets PrescriberName field to given value.


### GetPrescriberAddress

`func (o *Prescriber) GetPrescriberAddress() Address`

GetPrescriberAddress returns the PrescriberAddress field if non-nil, zero value otherwise.

### GetPrescriberAddressOk

`func (o *Prescriber) GetPrescriberAddressOk() (*Address, bool)`

GetPrescriberAddressOk returns a tuple with the PrescriberAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescriberAddress

`func (o *Prescriber) SetPrescriberAddress(v Address)`

SetPrescriberAddress sets PrescriberAddress field to given value.


### GetPrimaryTelephone

`func (o *Prescriber) GetPrimaryTelephone() string`

GetPrimaryTelephone returns the PrimaryTelephone field if non-nil, zero value otherwise.

### GetPrimaryTelephoneOk

`func (o *Prescriber) GetPrimaryTelephoneOk() (*string, bool)`

GetPrimaryTelephoneOk returns a tuple with the PrimaryTelephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTelephone

`func (o *Prescriber) SetPrimaryTelephone(v string)`

SetPrimaryTelephone sets PrimaryTelephone field to given value.


### GetDeaNumber

`func (o *Prescriber) GetDeaNumber() string`

GetDeaNumber returns the DeaNumber field if non-nil, zero value otherwise.

### GetDeaNumberOk

`func (o *Prescriber) GetDeaNumberOk() (*string, bool)`

GetDeaNumberOk returns a tuple with the DeaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaNumber

`func (o *Prescriber) SetDeaNumber(v string)`

SetDeaNumber sets DeaNumber field to given value.

### HasDeaNumber

`func (o *Prescriber) HasDeaNumber() bool`

HasDeaNumber returns a boolean if a field has been set.

### GetStateLicenseNumber

`func (o *Prescriber) GetStateLicenseNumber() string`

GetStateLicenseNumber returns the StateLicenseNumber field if non-nil, zero value otherwise.

### GetStateLicenseNumberOk

`func (o *Prescriber) GetStateLicenseNumberOk() (*string, bool)`

GetStateLicenseNumberOk returns a tuple with the StateLicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateLicenseNumber

`func (o *Prescriber) SetStateLicenseNumber(v string)`

SetStateLicenseNumber sets StateLicenseNumber field to given value.

### HasStateLicenseNumber

`func (o *Prescriber) HasStateLicenseNumber() bool`

HasStateLicenseNumber returns a boolean if a field has been set.

### GetBusinessName

`func (o *Prescriber) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *Prescriber) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *Prescriber) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *Prescriber) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetFax

`func (o *Prescriber) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Prescriber) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Prescriber) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Prescriber) HasFax() bool`

HasFax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


