# AccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountsRequestAccountsInner**](AccountsRequestAccountsInner.md) |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountsRequest

`func NewAccountsRequest() *AccountsRequest`

NewAccountsRequest instantiates a new AccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsRequestWithDefaults

`func NewAccountsRequestWithDefaults() *AccountsRequest`

NewAccountsRequestWithDefaults instantiates a new AccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountsRequest) GetAccounts() []AccountsRequestAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountsRequest) GetAccountsOk() (*[]AccountsRequestAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountsRequest) SetAccounts(v []AccountsRequestAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AccountsRequest) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAction

`func (o *AccountsRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccountsRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccountsRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AccountsRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *AccountsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *AccountsRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AccountsRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AccountsRequest) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AccountsRequest) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


