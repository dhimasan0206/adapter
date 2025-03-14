/*
Integration Adapter

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateStockLevelRequestStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStockLevelRequestStore{}

// UpdateStockLevelRequestStore struct for UpdateStockLevelRequestStore
type UpdateStockLevelRequestStore struct {
	Code *string `json:"code,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewUpdateStockLevelRequestStore instantiates a new UpdateStockLevelRequestStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStockLevelRequestStore() *UpdateStockLevelRequestStore {
	this := UpdateStockLevelRequestStore{}
	return &this
}

// NewUpdateStockLevelRequestStoreWithDefaults instantiates a new UpdateStockLevelRequestStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStockLevelRequestStoreWithDefaults() *UpdateStockLevelRequestStore {
	this := UpdateStockLevelRequestStore{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestStore) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestStore) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestStore) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateStockLevelRequestStore) SetCode(v string) {
	o.Code = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestStore) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestStore) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestStore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateStockLevelRequestStore) SetId(v string) {
	o.Id = &v
}

func (o UpdateStockLevelRequestStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStockLevelRequestStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUpdateStockLevelRequestStore struct {
	value *UpdateStockLevelRequestStore
	isSet bool
}

func (v NullableUpdateStockLevelRequestStore) Get() *UpdateStockLevelRequestStore {
	return v.value
}

func (v *NullableUpdateStockLevelRequestStore) Set(val *UpdateStockLevelRequestStore) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStockLevelRequestStore) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStockLevelRequestStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStockLevelRequestStore(val *UpdateStockLevelRequestStore) *NullableUpdateStockLevelRequestStore {
	return &NullableUpdateStockLevelRequestStore{value: val, isSet: true}
}

func (v NullableUpdateStockLevelRequestStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStockLevelRequestStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


