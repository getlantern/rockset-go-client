/*
REST API

Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FieldPartition struct for FieldPartition
type FieldPartition struct {
	// The name of a field, parsed as a SQL qualified name
	FieldName *string `json:"field_name,omitempty"`
	// The type of partitions on a field
	Type *string `json:"type,omitempty"`
	// The values for partitioning of a field
	Keys *[]string `json:"keys,omitempty"`
}

// NewFieldPartition instantiates a new FieldPartition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldPartition() *FieldPartition {
	this := FieldPartition{}
	return &this
}

// NewFieldPartitionWithDefaults instantiates a new FieldPartition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldPartitionWithDefaults() *FieldPartition {
	this := FieldPartition{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *FieldPartition) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldPartition) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *FieldPartition) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *FieldPartition) SetFieldName(v string) {
	o.FieldName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FieldPartition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldPartition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FieldPartition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FieldPartition) SetType(v string) {
	o.Type = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *FieldPartition) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldPartition) GetKeysOk() (*[]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *FieldPartition) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *FieldPartition) SetKeys(v []string) {
	o.Keys = &v
}

func (o FieldPartition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldName != nil {
		toSerialize["field_name"] = o.FieldName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableFieldPartition struct {
	value *FieldPartition
	isSet bool
}

func (v NullableFieldPartition) Get() *FieldPartition {
	return v.value
}

func (v *NullableFieldPartition) Set(val *FieldPartition) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldPartition) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldPartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldPartition(val *FieldPartition) *NullableFieldPartition {
	return &NullableFieldPartition{value: val, isSet: true}
}

func (v NullableFieldPartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldPartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


