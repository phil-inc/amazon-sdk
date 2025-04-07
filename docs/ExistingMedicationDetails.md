# ExistingMedicationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingMedications** | Pointer to [**[]ExistingMedication**](ExistingMedication.md) | List of existing medications for the patient. | [optional] 
**ExistingMedicationsDescriptor** | Pointer to [**ExistingMedicationsDescriptor**](ExistingMedicationsDescriptor.md) |  | [optional] 

## Methods

### NewExistingMedicationDetails

`func NewExistingMedicationDetails() *ExistingMedicationDetails`

NewExistingMedicationDetails instantiates a new ExistingMedicationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExistingMedicationDetailsWithDefaults

`func NewExistingMedicationDetailsWithDefaults() *ExistingMedicationDetails`

NewExistingMedicationDetailsWithDefaults instantiates a new ExistingMedicationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingMedications

`func (o *ExistingMedicationDetails) GetExistingMedications() []ExistingMedication`

GetExistingMedications returns the ExistingMedications field if non-nil, zero value otherwise.

### GetExistingMedicationsOk

`func (o *ExistingMedicationDetails) GetExistingMedicationsOk() (*[]ExistingMedication, bool)`

GetExistingMedicationsOk returns a tuple with the ExistingMedications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMedications

`func (o *ExistingMedicationDetails) SetExistingMedications(v []ExistingMedication)`

SetExistingMedications sets ExistingMedications field to given value.

### HasExistingMedications

`func (o *ExistingMedicationDetails) HasExistingMedications() bool`

HasExistingMedications returns a boolean if a field has been set.

### GetExistingMedicationsDescriptor

`func (o *ExistingMedicationDetails) GetExistingMedicationsDescriptor() ExistingMedicationsDescriptor`

GetExistingMedicationsDescriptor returns the ExistingMedicationsDescriptor field if non-nil, zero value otherwise.

### GetExistingMedicationsDescriptorOk

`func (o *ExistingMedicationDetails) GetExistingMedicationsDescriptorOk() (*ExistingMedicationsDescriptor, bool)`

GetExistingMedicationsDescriptorOk returns a tuple with the ExistingMedicationsDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingMedicationsDescriptor

`func (o *ExistingMedicationDetails) SetExistingMedicationsDescriptor(v ExistingMedicationsDescriptor)`

SetExistingMedicationsDescriptor sets ExistingMedicationsDescriptor field to given value.

### HasExistingMedicationsDescriptor

`func (o *ExistingMedicationDetails) HasExistingMedicationsDescriptor() bool`

HasExistingMedicationsDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


