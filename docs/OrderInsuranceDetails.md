# OrderInsuranceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | [**OrderInsurance**](OrderInsurance.md) |  | 
**Secondary** | Pointer to [**OrderInsurance**](OrderInsurance.md) |  | [optional] 

## Methods

### NewOrderInsuranceDetails

`func NewOrderInsuranceDetails(primary OrderInsurance, ) *OrderInsuranceDetails`

NewOrderInsuranceDetails instantiates a new OrderInsuranceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInsuranceDetailsWithDefaults

`func NewOrderInsuranceDetailsWithDefaults() *OrderInsuranceDetails`

NewOrderInsuranceDetailsWithDefaults instantiates a new OrderInsuranceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *OrderInsuranceDetails) GetPrimary() OrderInsurance`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *OrderInsuranceDetails) GetPrimaryOk() (*OrderInsurance, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *OrderInsuranceDetails) SetPrimary(v OrderInsurance)`

SetPrimary sets Primary field to given value.


### GetSecondary

`func (o *OrderInsuranceDetails) GetSecondary() OrderInsurance`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *OrderInsuranceDetails) GetSecondaryOk() (*OrderInsurance, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *OrderInsuranceDetails) SetSecondary(v OrderInsurance)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *OrderInsuranceDetails) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


