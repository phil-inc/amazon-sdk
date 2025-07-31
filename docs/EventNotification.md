# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to [**EventType**](EventType.md) |  | [optional] 
**EventDatetime** | Pointer to **string** | The date and time of the event | [optional] 
**EntityType** | Pointer to [**EntityType**](EntityType.md) |  | [optional] 
**EntityId** | Pointer to **string** | The entity Id of the entity. This will map to the IDs during the Put operation of the entity. For example    this will map to the partnerPrescriptionId in the Prescription /PUT operation | [optional] 

## Methods

### NewEventNotification

`func NewEventNotification() *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *EventNotification) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventNotification) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventNotification) SetEventType(v EventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventNotification) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventDatetime

`func (o *EventNotification) GetEventDatetime() string`

GetEventDatetime returns the EventDatetime field if non-nil, zero value otherwise.

### GetEventDatetimeOk

`func (o *EventNotification) GetEventDatetimeOk() (*string, bool)`

GetEventDatetimeOk returns a tuple with the EventDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDatetime

`func (o *EventNotification) SetEventDatetime(v string)`

SetEventDatetime sets EventDatetime field to given value.

### HasEventDatetime

`func (o *EventNotification) HasEventDatetime() bool`

HasEventDatetime returns a boolean if a field has been set.

### GetEntityType

`func (o *EventNotification) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EventNotification) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EventNotification) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *EventNotification) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityId

`func (o *EventNotification) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EventNotification) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EventNotification) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EventNotification) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


