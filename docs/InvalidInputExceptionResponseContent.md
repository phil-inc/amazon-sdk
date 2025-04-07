# InvalidInputExceptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human-readable message describing the invalid input exception. Usually this is due to invalid or out-of-range value was supplied for an input parameter | 
**ErrorCode** | Pointer to [**InvalidInputErrorCode**](InvalidInputErrorCode.md) |  | [optional] 
**ErrorFieldList** | Pointer to [**[]ErrorField**](ErrorField.md) | List of fields responsible for the error, if applicable. | [optional] 

## Methods

### NewInvalidInputExceptionResponseContent

`func NewInvalidInputExceptionResponseContent(message string, ) *InvalidInputExceptionResponseContent`

NewInvalidInputExceptionResponseContent instantiates a new InvalidInputExceptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidInputExceptionResponseContentWithDefaults

`func NewInvalidInputExceptionResponseContentWithDefaults() *InvalidInputExceptionResponseContent`

NewInvalidInputExceptionResponseContentWithDefaults instantiates a new InvalidInputExceptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InvalidInputExceptionResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidInputExceptionResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidInputExceptionResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorCode

`func (o *InvalidInputExceptionResponseContent) GetErrorCode() InvalidInputErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InvalidInputExceptionResponseContent) GetErrorCodeOk() (*InvalidInputErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InvalidInputExceptionResponseContent) SetErrorCode(v InvalidInputErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *InvalidInputExceptionResponseContent) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorFieldList

`func (o *InvalidInputExceptionResponseContent) GetErrorFieldList() []ErrorField`

GetErrorFieldList returns the ErrorFieldList field if non-nil, zero value otherwise.

### GetErrorFieldListOk

`func (o *InvalidInputExceptionResponseContent) GetErrorFieldListOk() (*[]ErrorField, bool)`

GetErrorFieldListOk returns a tuple with the ErrorFieldList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFieldList

`func (o *InvalidInputExceptionResponseContent) SetErrorFieldList(v []ErrorField)`

SetErrorFieldList sets ErrorFieldList field to given value.

### HasErrorFieldList

`func (o *InvalidInputExceptionResponseContent) HasErrorFieldList() bool`

HasErrorFieldList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


