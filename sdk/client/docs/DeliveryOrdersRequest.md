# DeliveryOrdersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**DeliveryOrders** | Pointer to [**[]DeliveryOrdersRequestDeliveryOrdersInner**](DeliveryOrdersRequestDeliveryOrdersInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewDeliveryOrdersRequest

`func NewDeliveryOrdersRequest() *DeliveryOrdersRequest`

NewDeliveryOrdersRequest instantiates a new DeliveryOrdersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOrdersRequestWithDefaults

`func NewDeliveryOrdersRequestWithDefaults() *DeliveryOrdersRequest`

NewDeliveryOrdersRequestWithDefaults instantiates a new DeliveryOrdersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DeliveryOrdersRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeliveryOrdersRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeliveryOrdersRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeliveryOrdersRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDeliveryOrders

`func (o *DeliveryOrdersRequest) GetDeliveryOrders() []DeliveryOrdersRequestDeliveryOrdersInner`

GetDeliveryOrders returns the DeliveryOrders field if non-nil, zero value otherwise.

### GetDeliveryOrdersOk

`func (o *DeliveryOrdersRequest) GetDeliveryOrdersOk() (*[]DeliveryOrdersRequestDeliveryOrdersInner, bool)`

GetDeliveryOrdersOk returns a tuple with the DeliveryOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOrders

`func (o *DeliveryOrdersRequest) SetDeliveryOrders(v []DeliveryOrdersRequestDeliveryOrdersInner)`

SetDeliveryOrders sets DeliveryOrders field to given value.

### HasDeliveryOrders

`func (o *DeliveryOrdersRequest) HasDeliveryOrders() bool`

HasDeliveryOrders returns a boolean if a field has been set.

### GetId

`func (o *DeliveryOrdersRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryOrdersRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryOrdersRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryOrdersRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *DeliveryOrdersRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DeliveryOrdersRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DeliveryOrdersRequest) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *DeliveryOrdersRequest) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


