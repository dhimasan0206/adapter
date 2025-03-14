# UpdateStockLevel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**UpdateStockLevel200ResponseError**](UpdateStockLevel200ResponseError.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]UpdateStockLevel200ResponseResultsInner**](UpdateStockLevel200ResponseResultsInner.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Store** | Pointer to [**UpdateStockLevelRequestStore**](UpdateStockLevelRequestStore.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 

## Methods

### NewUpdateStockLevel200Response

`func NewUpdateStockLevel200Response() *UpdateStockLevel200Response`

NewUpdateStockLevel200Response instantiates a new UpdateStockLevel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStockLevel200ResponseWithDefaults

`func NewUpdateStockLevel200ResponseWithDefaults() *UpdateStockLevel200Response`

NewUpdateStockLevel200ResponseWithDefaults instantiates a new UpdateStockLevel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UpdateStockLevel200Response) GetError() UpdateStockLevel200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateStockLevel200Response) GetErrorOk() (*UpdateStockLevel200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateStockLevel200Response) SetError(v UpdateStockLevel200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *UpdateStockLevel200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *UpdateStockLevel200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStockLevel200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStockLevel200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateStockLevel200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateStockLevel200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateStockLevel200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateStockLevel200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateStockLevel200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *UpdateStockLevel200Response) GetResults() []UpdateStockLevel200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UpdateStockLevel200Response) GetResultsOk() (*[]UpdateStockLevel200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UpdateStockLevel200Response) SetResults(v []UpdateStockLevel200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *UpdateStockLevel200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateStockLevel200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateStockLevel200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateStockLevel200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateStockLevel200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStore

`func (o *UpdateStockLevel200Response) GetStore() UpdateStockLevelRequestStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *UpdateStockLevel200Response) GetStoreOk() (*UpdateStockLevelRequestStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *UpdateStockLevel200Response) SetStore(v UpdateStockLevelRequestStore)`

SetStore sets Store field to given value.

### HasStore

`func (o *UpdateStockLevel200Response) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateStockLevel200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateStockLevel200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateStockLevel200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateStockLevel200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTimestamp

`func (o *UpdateStockLevel200Response) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UpdateStockLevel200Response) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UpdateStockLevel200Response) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UpdateStockLevel200Response) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


