# TransferToPharmacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NcpdpId** | **string** | An NCPDP-assigned national provider identification number that assists pharmacies in their interactions with federal agencies and third party providers. | 
**Npi** | **string** | A National Provider System (NPS) assigned identifier for health care providers. | 
**PharmacyName** | **string** | The pharmacy this prescription needs transfer to. | 
**DeaNumber** | **string** | The Drug Enforcement Administration (DEA) assigned number to all businesses that manufacture or distribute controlled pharmaceuticals,         all health professionals who dispense, administer, or prescribe controlled pharmaceuticals and all pharmacies that dispense prescriptions. | 
**PharmacistName** | [**Name**](Name.md) |  | 
**PharmacyAddress** | [**Address**](Address.md) |  | 
**PrimaryTelephone** | **string** | Pharmacy telephone number. | 
**Fax** | **string** | Pharmacy fax number. | 

## Methods

### NewTransferToPharmacy

`func NewTransferToPharmacy(ncpdpId string, npi string, pharmacyName string, deaNumber string, pharmacistName Name, pharmacyAddress Address, primaryTelephone string, fax string, ) *TransferToPharmacy`

NewTransferToPharmacy instantiates a new TransferToPharmacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToPharmacyWithDefaults

`func NewTransferToPharmacyWithDefaults() *TransferToPharmacy`

NewTransferToPharmacyWithDefaults instantiates a new TransferToPharmacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNcpdpId

`func (o *TransferToPharmacy) GetNcpdpId() string`

GetNcpdpId returns the NcpdpId field if non-nil, zero value otherwise.

### GetNcpdpIdOk

`func (o *TransferToPharmacy) GetNcpdpIdOk() (*string, bool)`

GetNcpdpIdOk returns a tuple with the NcpdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcpdpId

`func (o *TransferToPharmacy) SetNcpdpId(v string)`

SetNcpdpId sets NcpdpId field to given value.


### GetNpi

`func (o *TransferToPharmacy) GetNpi() string`

GetNpi returns the Npi field if non-nil, zero value otherwise.

### GetNpiOk

`func (o *TransferToPharmacy) GetNpiOk() (*string, bool)`

GetNpiOk returns a tuple with the Npi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpi

`func (o *TransferToPharmacy) SetNpi(v string)`

SetNpi sets Npi field to given value.


### GetPharmacyName

`func (o *TransferToPharmacy) GetPharmacyName() string`

GetPharmacyName returns the PharmacyName field if non-nil, zero value otherwise.

### GetPharmacyNameOk

`func (o *TransferToPharmacy) GetPharmacyNameOk() (*string, bool)`

GetPharmacyNameOk returns a tuple with the PharmacyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmacyName

`func (o *TransferToPharmacy) SetPharmacyName(v string)`

SetPharmacyName sets PharmacyName field to given value.


### GetDeaNumber

`func (o *TransferToPharmacy) GetDeaNumber() string`

GetDeaNumber returns the DeaNumber field if non-nil, zero value otherwise.

### GetDeaNumberOk

`func (o *TransferToPharmacy) GetDeaNumberOk() (*string, bool)`

GetDeaNumberOk returns a tuple with the DeaNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaNumber

`func (o *TransferToPharmacy) SetDeaNumber(v string)`

SetDeaNumber sets DeaNumber field to given value.


### GetPharmacistName

`func (o *TransferToPharmacy) GetPharmacistName() Name`

GetPharmacistName returns the PharmacistName field if non-nil, zero value otherwise.

### GetPharmacistNameOk

`func (o *TransferToPharmacy) GetPharmacistNameOk() (*Name, bool)`

GetPharmacistNameOk returns a tuple with the PharmacistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmacistName

`func (o *TransferToPharmacy) SetPharmacistName(v Name)`

SetPharmacistName sets PharmacistName field to given value.


### GetPharmacyAddress

`func (o *TransferToPharmacy) GetPharmacyAddress() Address`

GetPharmacyAddress returns the PharmacyAddress field if non-nil, zero value otherwise.

### GetPharmacyAddressOk

`func (o *TransferToPharmacy) GetPharmacyAddressOk() (*Address, bool)`

GetPharmacyAddressOk returns a tuple with the PharmacyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmacyAddress

`func (o *TransferToPharmacy) SetPharmacyAddress(v Address)`

SetPharmacyAddress sets PharmacyAddress field to given value.


### GetPrimaryTelephone

`func (o *TransferToPharmacy) GetPrimaryTelephone() string`

GetPrimaryTelephone returns the PrimaryTelephone field if non-nil, zero value otherwise.

### GetPrimaryTelephoneOk

`func (o *TransferToPharmacy) GetPrimaryTelephoneOk() (*string, bool)`

GetPrimaryTelephoneOk returns a tuple with the PrimaryTelephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTelephone

`func (o *TransferToPharmacy) SetPrimaryTelephone(v string)`

SetPrimaryTelephone sets PrimaryTelephone field to given value.


### GetFax

`func (o *TransferToPharmacy) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *TransferToPharmacy) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *TransferToPharmacy) SetFax(v string)`

SetFax sets Fax field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


