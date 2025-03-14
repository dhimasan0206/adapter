# ProductsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Products** | Pointer to [**[]ProductsRequestProductsInner**](ProductsRequestProductsInner.md) |  | [optional] 

## Methods

### NewProductsRequest

`func NewProductsRequest() *ProductsRequest`

NewProductsRequest instantiates a new ProductsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsRequestWithDefaults

`func NewProductsRequestWithDefaults() *ProductsRequest`

NewProductsRequestWithDefaults instantiates a new ProductsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProductsRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProductsRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProductsRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProductsRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *ProductsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ProductsRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ProductsRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ProductsRequest) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ProductsRequest) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetProducts

`func (o *ProductsRequest) GetProducts() []ProductsRequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductsRequest) GetProductsOk() (*[]ProductsRequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductsRequest) SetProducts(v []ProductsRequestProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ProductsRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


