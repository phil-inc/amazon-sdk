# PreconditionFailedExceptionResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human-readable message describing the precondition failure exception describing the precondition failure exception that was caught by the Amazon Pharmacy APIs | 
**ErrorCode** | [**PreconditionErrorCode**](PreconditionErrorCode.md) |  | 
**ErrorFieldList** | Pointer to [**[]ErrorField**](ErrorField.md) | List of fields responsible for the error, if applicable. | [optional] 

## Methods

### NewPreconditionFailedExceptionResponseContent

`func NewPreconditionFailedExceptionResponseContent(message string, errorCode PreconditionErrorCode, ) *PreconditionFailedExceptionResponseContent`

NewPreconditionFailedExceptionResponseContent instantiates a new PreconditionFailedExceptionResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreconditionFailedExceptionResponseContentWithDefaults

`func NewPreconditionFailedExceptionResponseContentWithDefaults() *PreconditionFailedExceptionResponseContent`

NewPreconditionFailedExceptionResponseContentWithDefaults instantiates a new PreconditionFailedExceptionResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PreconditionFailedExceptionResponseContent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PreconditionFailedExceptionResponseContent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PreconditionFailedExceptionResponseContent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorCode

`func (o *PreconditionFailedExceptionResponseContent) GetErrorCode() PreconditionErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PreconditionFailedExceptionResponseContent) GetErrorCodeOk() (*PreconditionErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PreconditionFailedExceptionResponseContent) SetErrorCode(v PreconditionErrorCode)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorFieldList

`func (o *PreconditionFailedExceptionResponseContent) GetErrorFieldList() []ErrorField`

GetErrorFieldList returns the ErrorFieldList field if non-nil, zero value otherwise.

### GetErrorFieldListOk

`func (o *PreconditionFailedExceptionResponseContent) GetErrorFieldListOk() (*[]ErrorField, bool)`

GetErrorFieldListOk returns a tuple with the ErrorFieldList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFieldList

`func (o *PreconditionFailedExceptionResponseContent) SetErrorFieldList(v []ErrorField)`

SetErrorFieldList sets ErrorFieldList field to given value.

### HasErrorFieldList

`func (o *PreconditionFailedExceptionResponseContent) HasErrorFieldList() bool`

HasErrorFieldList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


