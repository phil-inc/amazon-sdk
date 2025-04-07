# AllergyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoKnownAllergies** | Pointer to **bool** | Currently supported but on path to deprecation as we will plan to switch to allergiesDescriptor This shape is deprecated since 2024-05-15: Use allergiesDescriptor instead | [optional] 
**Allergies** | Pointer to [**[]Allergy**](Allergy.md) | If noKnownAllergies is false, allergies is required | [optional] 
**AllergiesDescriptor** | Pointer to [**AllergiesDescriptor**](AllergiesDescriptor.md) |  | [optional] 

## Methods

### NewAllergyDetails

`func NewAllergyDetails() *AllergyDetails`

NewAllergyDetails instantiates a new AllergyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllergyDetailsWithDefaults

`func NewAllergyDetailsWithDefaults() *AllergyDetails`

NewAllergyDetailsWithDefaults instantiates a new AllergyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoKnownAllergies

`func (o *AllergyDetails) GetNoKnownAllergies() bool`

GetNoKnownAllergies returns the NoKnownAllergies field if non-nil, zero value otherwise.

### GetNoKnownAllergiesOk

`func (o *AllergyDetails) GetNoKnownAllergiesOk() (*bool, bool)`

GetNoKnownAllergiesOk returns a tuple with the NoKnownAllergies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoKnownAllergies

`func (o *AllergyDetails) SetNoKnownAllergies(v bool)`

SetNoKnownAllergies sets NoKnownAllergies field to given value.

### HasNoKnownAllergies

`func (o *AllergyDetails) HasNoKnownAllergies() bool`

HasNoKnownAllergies returns a boolean if a field has been set.

### GetAllergies

`func (o *AllergyDetails) GetAllergies() []Allergy`

GetAllergies returns the Allergies field if non-nil, zero value otherwise.

### GetAllergiesOk

`func (o *AllergyDetails) GetAllergiesOk() (*[]Allergy, bool)`

GetAllergiesOk returns a tuple with the Allergies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllergies

`func (o *AllergyDetails) SetAllergies(v []Allergy)`

SetAllergies sets Allergies field to given value.

### HasAllergies

`func (o *AllergyDetails) HasAllergies() bool`

HasAllergies returns a boolean if a field has been set.

### GetAllergiesDescriptor

`func (o *AllergyDetails) GetAllergiesDescriptor() AllergiesDescriptor`

GetAllergiesDescriptor returns the AllergiesDescriptor field if non-nil, zero value otherwise.

### GetAllergiesDescriptorOk

`func (o *AllergyDetails) GetAllergiesDescriptorOk() (*AllergiesDescriptor, bool)`

GetAllergiesDescriptorOk returns a tuple with the AllergiesDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllergiesDescriptor

`func (o *AllergyDetails) SetAllergiesDescriptor(v AllergiesDescriptor)`

SetAllergiesDescriptor sets AllergiesDescriptor field to given value.

### HasAllergiesDescriptor

`func (o *AllergyDetails) HasAllergiesDescriptor() bool`

HasAllergiesDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


