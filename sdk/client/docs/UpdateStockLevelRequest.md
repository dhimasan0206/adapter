# UpdateStockLevelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Requests** | Pointer to [**[]UpdateStockLevelRequestRequestsInner**](UpdateStockLevelRequestRequestsInner.md) |  | [optional] 
**Store** | Pointer to [**UpdateStockLevelRequestStore**](UpdateStockLevelRequestStore.md) |  | [optional] 

## Methods

### NewUpdateStockLevelRequest

`func NewUpdateStockLevelRequest() *UpdateStockLevelRequest`

NewUpdateStockLevelRequest instantiates a new UpdateStockLevelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStockLevelRequestWithDefaults

`func NewUpdateStockLevelRequestWithDefaults() *UpdateStockLevelRequest`

NewUpdateStockLevelRequestWithDefaults instantiates a new UpdateStockLevelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateStockLevelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStockLevelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStockLevelRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateStockLevelRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequests

`func (o *UpdateStockLevelRequest) GetRequests() []UpdateStockLevelRequestRequestsInner`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UpdateStockLevelRequest) GetRequestsOk() (*[]UpdateStockLevelRequestRequestsInner, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UpdateStockLevelRequest) SetRequests(v []UpdateStockLevelRequestRequestsInner)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *UpdateStockLevelRequest) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetStore

`func (o *UpdateStockLevelRequest) GetStore() UpdateStockLevelRequestStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *UpdateStockLevelRequest) GetStoreOk() (*UpdateStockLevelRequestStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *UpdateStockLevelRequest) SetStore(v UpdateStockLevelRequestStore)`

SetStore sets Store field to given value.

### HasStore

`func (o *UpdateStockLevelRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


