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

// checks if the UpdateStockLevelRequestRequestsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStockLevelRequestRequestsInner{}

// UpdateStockLevelRequestRequestsInner struct for UpdateStockLevelRequestRequestsInner
type UpdateStockLevelRequestRequestsInner struct {
	ChannelLocationCode *string `json:"channel_location_code,omitempty"`
	ChannelLocationId *string `json:"channel_location_id,omitempty"`
	ChannelProductId *string `json:"channel_product_id,omitempty"`
	ChannelProductVariantId *string `json:"channel_product_variant_id,omitempty"`
	ChannelProductVariantSku *string `json:"channel_product_variant_sku,omitempty"`
	DeltaStock *float32 `json:"delta_stock,omitempty"`
	ExactStock *float32 `json:"exact_stock,omitempty"`
	LocationCode *string `json:"location_code,omitempty"`
	LocationId *string `json:"location_id,omitempty"`
	ProductId *string `json:"product_id,omitempty"`
	ProductVariantId *string `json:"product_variant_id,omitempty"`
	ProductVariantSku *string `json:"product_variant_sku,omitempty"`
}

// NewUpdateStockLevelRequestRequestsInner instantiates a new UpdateStockLevelRequestRequestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStockLevelRequestRequestsInner() *UpdateStockLevelRequestRequestsInner {
	this := UpdateStockLevelRequestRequestsInner{}
	return &this
}

// NewUpdateStockLevelRequestRequestsInnerWithDefaults instantiates a new UpdateStockLevelRequestRequestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStockLevelRequestRequestsInnerWithDefaults() *UpdateStockLevelRequestRequestsInner {
	this := UpdateStockLevelRequestRequestsInner{}
	return &this
}

// GetChannelLocationCode returns the ChannelLocationCode field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelLocationCode() string {
	if o == nil || IsNil(o.ChannelLocationCode) {
		var ret string
		return ret
	}
	return *o.ChannelLocationCode
}

// GetChannelLocationCodeOk returns a tuple with the ChannelLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelLocationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelLocationCode) {
		return nil, false
	}
	return o.ChannelLocationCode, true
}

// HasChannelLocationCode returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasChannelLocationCode() bool {
	if o != nil && !IsNil(o.ChannelLocationCode) {
		return true
	}

	return false
}

// SetChannelLocationCode gets a reference to the given string and assigns it to the ChannelLocationCode field.
func (o *UpdateStockLevelRequestRequestsInner) SetChannelLocationCode(v string) {
	o.ChannelLocationCode = &v
}

// GetChannelLocationId returns the ChannelLocationId field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelLocationId() string {
	if o == nil || IsNil(o.ChannelLocationId) {
		var ret string
		return ret
	}
	return *o.ChannelLocationId
}

// GetChannelLocationIdOk returns a tuple with the ChannelLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelLocationId) {
		return nil, false
	}
	return o.ChannelLocationId, true
}

// HasChannelLocationId returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasChannelLocationId() bool {
	if o != nil && !IsNil(o.ChannelLocationId) {
		return true
	}

	return false
}

// SetChannelLocationId gets a reference to the given string and assigns it to the ChannelLocationId field.
func (o *UpdateStockLevelRequestRequestsInner) SetChannelLocationId(v string) {
	o.ChannelLocationId = &v
}

// GetChannelProductId returns the ChannelProductId field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelProductId() string {
	if o == nil || IsNil(o.ChannelProductId) {
		var ret string
		return ret
	}
	return *o.ChannelProductId
}

// GetChannelProductIdOk returns a tuple with the ChannelProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelProductId) {
		return nil, false
	}
	return o.ChannelProductId, true
}

// HasChannelProductId returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasChannelProductId() bool {
	if o != nil && !IsNil(o.ChannelProductId) {
		return true
	}

	return false
}

// SetChannelProductId gets a reference to the given string and assigns it to the ChannelProductId field.
func (o *UpdateStockLevelRequestRequestsInner) SetChannelProductId(v string) {
	o.ChannelProductId = &v
}

// GetChannelProductVariantId returns the ChannelProductVariantId field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelProductVariantId() string {
	if o == nil || IsNil(o.ChannelProductVariantId) {
		var ret string
		return ret
	}
	return *o.ChannelProductVariantId
}

// GetChannelProductVariantIdOk returns a tuple with the ChannelProductVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelProductVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelProductVariantId) {
		return nil, false
	}
	return o.ChannelProductVariantId, true
}

// HasChannelProductVariantId returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasChannelProductVariantId() bool {
	if o != nil && !IsNil(o.ChannelProductVariantId) {
		return true
	}

	return false
}

// SetChannelProductVariantId gets a reference to the given string and assigns it to the ChannelProductVariantId field.
func (o *UpdateStockLevelRequestRequestsInner) SetChannelProductVariantId(v string) {
	o.ChannelProductVariantId = &v
}

// GetChannelProductVariantSku returns the ChannelProductVariantSku field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelProductVariantSku() string {
	if o == nil || IsNil(o.ChannelProductVariantSku) {
		var ret string
		return ret
	}
	return *o.ChannelProductVariantSku
}

// GetChannelProductVariantSkuOk returns a tuple with the ChannelProductVariantSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetChannelProductVariantSkuOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelProductVariantSku) {
		return nil, false
	}
	return o.ChannelProductVariantSku, true
}

// HasChannelProductVariantSku returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasChannelProductVariantSku() bool {
	if o != nil && !IsNil(o.ChannelProductVariantSku) {
		return true
	}

	return false
}

// SetChannelProductVariantSku gets a reference to the given string and assigns it to the ChannelProductVariantSku field.
func (o *UpdateStockLevelRequestRequestsInner) SetChannelProductVariantSku(v string) {
	o.ChannelProductVariantSku = &v
}

// GetDeltaStock returns the DeltaStock field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetDeltaStock() float32 {
	if o == nil || IsNil(o.DeltaStock) {
		var ret float32
		return ret
	}
	return *o.DeltaStock
}

// GetDeltaStockOk returns a tuple with the DeltaStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetDeltaStockOk() (*float32, bool) {
	if o == nil || IsNil(o.DeltaStock) {
		return nil, false
	}
	return o.DeltaStock, true
}

// HasDeltaStock returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasDeltaStock() bool {
	if o != nil && !IsNil(o.DeltaStock) {
		return true
	}

	return false
}

// SetDeltaStock gets a reference to the given float32 and assigns it to the DeltaStock field.
func (o *UpdateStockLevelRequestRequestsInner) SetDeltaStock(v float32) {
	o.DeltaStock = &v
}

// GetExactStock returns the ExactStock field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetExactStock() float32 {
	if o == nil || IsNil(o.ExactStock) {
		var ret float32
		return ret
	}
	return *o.ExactStock
}

// GetExactStockOk returns a tuple with the ExactStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetExactStockOk() (*float32, bool) {
	if o == nil || IsNil(o.ExactStock) {
		return nil, false
	}
	return o.ExactStock, true
}

// HasExactStock returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasExactStock() bool {
	if o != nil && !IsNil(o.ExactStock) {
		return true
	}

	return false
}

// SetExactStock gets a reference to the given float32 and assigns it to the ExactStock field.
func (o *UpdateStockLevelRequestRequestsInner) SetExactStock(v float32) {
	o.ExactStock = &v
}

// GetLocationCode returns the LocationCode field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetLocationCode() string {
	if o == nil || IsNil(o.LocationCode) {
		var ret string
		return ret
	}
	return *o.LocationCode
}

// GetLocationCodeOk returns a tuple with the LocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetLocationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LocationCode) {
		return nil, false
	}
	return o.LocationCode, true
}

// HasLocationCode returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasLocationCode() bool {
	if o != nil && !IsNil(o.LocationCode) {
		return true
	}

	return false
}

// SetLocationCode gets a reference to the given string and assigns it to the LocationCode field.
func (o *UpdateStockLevelRequestRequestsInner) SetLocationCode(v string) {
	o.LocationCode = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetLocationId() string {
	if o == nil || IsNil(o.LocationId) {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocationId) {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasLocationId() bool {
	if o != nil && !IsNil(o.LocationId) {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *UpdateStockLevelRequestRequestsInner) SetLocationId(v string) {
	o.LocationId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *UpdateStockLevelRequestRequestsInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetProductVariantId returns the ProductVariantId field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetProductVariantId() string {
	if o == nil || IsNil(o.ProductVariantId) {
		var ret string
		return ret
	}
	return *o.ProductVariantId
}

// GetProductVariantIdOk returns a tuple with the ProductVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetProductVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVariantId) {
		return nil, false
	}
	return o.ProductVariantId, true
}

// HasProductVariantId returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasProductVariantId() bool {
	if o != nil && !IsNil(o.ProductVariantId) {
		return true
	}

	return false
}

// SetProductVariantId gets a reference to the given string and assigns it to the ProductVariantId field.
func (o *UpdateStockLevelRequestRequestsInner) SetProductVariantId(v string) {
	o.ProductVariantId = &v
}

// GetProductVariantSku returns the ProductVariantSku field value if set, zero value otherwise.
func (o *UpdateStockLevelRequestRequestsInner) GetProductVariantSku() string {
	if o == nil || IsNil(o.ProductVariantSku) {
		var ret string
		return ret
	}
	return *o.ProductVariantSku
}

// GetProductVariantSkuOk returns a tuple with the ProductVariantSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockLevelRequestRequestsInner) GetProductVariantSkuOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVariantSku) {
		return nil, false
	}
	return o.ProductVariantSku, true
}

// HasProductVariantSku returns a boolean if a field has been set.
func (o *UpdateStockLevelRequestRequestsInner) HasProductVariantSku() bool {
	if o != nil && !IsNil(o.ProductVariantSku) {
		return true
	}

	return false
}

// SetProductVariantSku gets a reference to the given string and assigns it to the ProductVariantSku field.
func (o *UpdateStockLevelRequestRequestsInner) SetProductVariantSku(v string) {
	o.ProductVariantSku = &v
}

func (o UpdateStockLevelRequestRequestsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStockLevelRequestRequestsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelLocationCode) {
		toSerialize["channel_location_code"] = o.ChannelLocationCode
	}
	if !IsNil(o.ChannelLocationId) {
		toSerialize["channel_location_id"] = o.ChannelLocationId
	}
	if !IsNil(o.ChannelProductId) {
		toSerialize["channel_product_id"] = o.ChannelProductId
	}
	if !IsNil(o.ChannelProductVariantId) {
		toSerialize["channel_product_variant_id"] = o.ChannelProductVariantId
	}
	if !IsNil(o.ChannelProductVariantSku) {
		toSerialize["channel_product_variant_sku"] = o.ChannelProductVariantSku
	}
	if !IsNil(o.DeltaStock) {
		toSerialize["delta_stock"] = o.DeltaStock
	}
	if !IsNil(o.ExactStock) {
		toSerialize["exact_stock"] = o.ExactStock
	}
	if !IsNil(o.LocationCode) {
		toSerialize["location_code"] = o.LocationCode
	}
	if !IsNil(o.LocationId) {
		toSerialize["location_id"] = o.LocationId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ProductVariantId) {
		toSerialize["product_variant_id"] = o.ProductVariantId
	}
	if !IsNil(o.ProductVariantSku) {
		toSerialize["product_variant_sku"] = o.ProductVariantSku
	}
	return toSerialize, nil
}

type NullableUpdateStockLevelRequestRequestsInner struct {
	value *UpdateStockLevelRequestRequestsInner
	isSet bool
}

func (v NullableUpdateStockLevelRequestRequestsInner) Get() *UpdateStockLevelRequestRequestsInner {
	return v.value
}

func (v *NullableUpdateStockLevelRequestRequestsInner) Set(val *UpdateStockLevelRequestRequestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStockLevelRequestRequestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStockLevelRequestRequestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStockLevelRequestRequestsInner(val *UpdateStockLevelRequestRequestsInner) *NullableUpdateStockLevelRequestRequestsInner {
	return &NullableUpdateStockLevelRequestRequestsInner{value: val, isSet: true}
}

func (v NullableUpdateStockLevelRequestRequestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStockLevelRequestRequestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


