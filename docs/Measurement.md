# Measurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VitalSign** | [**VitalSign**](VitalSign.md) |  | 
**Value** | **string** | The empirical value of the vital sign observed | 
**MeasurementUnitOfMeasure** | **string** | The unit of measurement for the vital sign (e.g., kg, cm) | 
**ObservationDate** | **string** | Date on which observation was recorded | 

## Methods

### NewMeasurement

`func NewMeasurement(vitalSign VitalSign, value string, measurementUnitOfMeasure string, observationDate string, ) *Measurement`

NewMeasurement instantiates a new Measurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementWithDefaults

`func NewMeasurementWithDefaults() *Measurement`

NewMeasurementWithDefaults instantiates a new Measurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVitalSign

`func (o *Measurement) GetVitalSign() VitalSign`

GetVitalSign returns the VitalSign field if non-nil, zero value otherwise.

### GetVitalSignOk

`func (o *Measurement) GetVitalSignOk() (*VitalSign, bool)`

GetVitalSignOk returns a tuple with the VitalSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVitalSign

`func (o *Measurement) SetVitalSign(v VitalSign)`

SetVitalSign sets VitalSign field to given value.


### GetValue

`func (o *Measurement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Measurement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Measurement) SetValue(v string)`

SetValue sets Value field to given value.


### GetMeasurementUnitOfMeasure

`func (o *Measurement) GetMeasurementUnitOfMeasure() string`

GetMeasurementUnitOfMeasure returns the MeasurementUnitOfMeasure field if non-nil, zero value otherwise.

### GetMeasurementUnitOfMeasureOk

`func (o *Measurement) GetMeasurementUnitOfMeasureOk() (*string, bool)`

GetMeasurementUnitOfMeasureOk returns a tuple with the MeasurementUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementUnitOfMeasure

`func (o *Measurement) SetMeasurementUnitOfMeasure(v string)`

SetMeasurementUnitOfMeasure sets MeasurementUnitOfMeasure field to given value.


### GetObservationDate

`func (o *Measurement) GetObservationDate() string`

GetObservationDate returns the ObservationDate field if non-nil, zero value otherwise.

### GetObservationDateOk

`func (o *Measurement) GetObservationDateOk() (*string, bool)`

GetObservationDateOk returns a tuple with the ObservationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDate

`func (o *Measurement) SetObservationDate(v string)`

SetObservationDate sets ObservationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


