# PutOrder400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human-readable message describing the invalid input exception. Usually this is due to invalid or out-of-range value was supplied for an input parameter | 
**ErrorCode** | [**InvalidInputErrorCode**](InvalidInputErrorCode.md) |  | 
**ErrorFieldList** | Pointer to [**[]ErrorField**](ErrorField.md) | List of fields responsible for the error, if applicable. | [optional] 

## Methods

### NewPutOrder400Response

`func NewPutOrder400Response(message string, errorCode InvalidInputErrorCode, ) *PutOrder400Response`

NewPutOrder400Response instantiates a new PutOrder400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutOrder400ResponseWithDefaults

`func NewPutOrder400ResponseWithDefaults() *PutOrder400Response`

NewPutOrder400ResponseWithDefaults instantiates a new PutOrder400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PutOrder400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PutOrder400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PutOrder400Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorCode

`func (o *PutOrder400Response) GetErrorCode() InvalidInputErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PutOrder400Response) GetErrorCodeOk() (*InvalidInputErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PutOrder400Response) SetErrorCode(v InvalidInputErrorCode)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorFieldList

`func (o *PutOrder400Response) GetErrorFieldList() []ErrorField`

GetErrorFieldList returns the ErrorFieldList field if non-nil, zero value otherwise.

### GetErrorFieldListOk

`func (o *PutOrder400Response) GetErrorFieldListOk() (*[]ErrorField, bool)`

GetErrorFieldListOk returns a tuple with the ErrorFieldList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFieldList

`func (o *PutOrder400Response) SetErrorFieldList(v []ErrorField)`

SetErrorFieldList sets ErrorFieldList field to given value.

### HasErrorFieldList

`func (o *PutOrder400Response) HasErrorFieldList() bool`

HasErrorFieldList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


