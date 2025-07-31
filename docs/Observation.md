# Observation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measurement** | Pointer to [**[]Measurement**](Measurement.md) | Used to record patient&#39;s vital signs such as weight and height | [optional] 
**ObservationNotes** | Pointer to **string** | Additional notes or comments related to the vital sign observations | [optional] 

## Methods

### NewObservation

`func NewObservation() *Observation`

NewObservation instantiates a new Observation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationWithDefaults

`func NewObservationWithDefaults() *Observation`

NewObservationWithDefaults instantiates a new Observation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasurement

`func (o *Observation) GetMeasurement() []Measurement`

GetMeasurement returns the Measurement field if non-nil, zero value otherwise.

### GetMeasurementOk

`func (o *Observation) GetMeasurementOk() (*[]Measurement, bool)`

GetMeasurementOk returns a tuple with the Measurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurement

`func (o *Observation) SetMeasurement(v []Measurement)`

SetMeasurement sets Measurement field to given value.

### HasMeasurement

`func (o *Observation) HasMeasurement() bool`

HasMeasurement returns a boolean if a field has been set.

### GetObservationNotes

`func (o *Observation) GetObservationNotes() string`

GetObservationNotes returns the ObservationNotes field if non-nil, zero value otherwise.

### GetObservationNotesOk

`func (o *Observation) GetObservationNotesOk() (*string, bool)`

GetObservationNotesOk returns a tuple with the ObservationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationNotes

`func (o *Observation) SetObservationNotes(v string)`

SetObservationNotes sets ObservationNotes field to given value.

### HasObservationNotes

`func (o *Observation) HasObservationNotes() bool`

HasObservationNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


