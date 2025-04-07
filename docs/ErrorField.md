# ErrorField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Indicates the field name responsible for the error | 
**Message** | Pointer to **string** | Message describing why the field was invalid | [optional] 

## Methods

### NewErrorField

`func NewErrorField(field string, ) *ErrorField`

NewErrorField instantiates a new ErrorField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorFieldWithDefaults

`func NewErrorFieldWithDefaults() *ErrorField`

NewErrorFieldWithDefaults instantiates a new ErrorField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ErrorField) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ErrorField) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ErrorField) SetField(v string)`

SetField sets Field field to given value.


### GetMessage

`func (o *ErrorField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorField) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorField) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


