# SalesOrdersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**SalesOrders** | Pointer to [**[]SalesOrdersRequestSalesOrdersInner**](SalesOrdersRequestSalesOrdersInner.md) |  | [optional] 

## Methods

### NewSalesOrdersRequest

`func NewSalesOrdersRequest() *SalesOrdersRequest`

NewSalesOrdersRequest instantiates a new SalesOrdersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesOrdersRequestWithDefaults

`func NewSalesOrdersRequestWithDefaults() *SalesOrdersRequest`

NewSalesOrdersRequestWithDefaults instantiates a new SalesOrdersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SalesOrdersRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SalesOrdersRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SalesOrdersRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SalesOrdersRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *SalesOrdersRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesOrdersRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesOrdersRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SalesOrdersRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *SalesOrdersRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SalesOrdersRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SalesOrdersRequest) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SalesOrdersRequest) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSalesOrders

`func (o *SalesOrdersRequest) GetSalesOrders() []SalesOrdersRequestSalesOrdersInner`

GetSalesOrders returns the SalesOrders field if non-nil, zero value otherwise.

### GetSalesOrdersOk

`func (o *SalesOrdersRequest) GetSalesOrdersOk() (*[]SalesOrdersRequestSalesOrdersInner, bool)`

GetSalesOrdersOk returns a tuple with the SalesOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrders

`func (o *SalesOrdersRequest) SetSalesOrders(v []SalesOrdersRequestSalesOrdersInner)`

SetSalesOrders sets SalesOrders field to given value.

### HasSalesOrders

`func (o *SalesOrdersRequest) HasSalesOrders() bool`

HasSalesOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


