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

// DeleteCollectionResponse struct for DeleteCollectionResponse
type DeleteCollectionResponse struct {
	Data *Collection `json:"data,omitempty"`
}

// NewDeleteCollectionResponse instantiates a new DeleteCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCollectionResponse() *DeleteCollectionResponse {
	this := DeleteCollectionResponse{}
	return &this
}

// NewDeleteCollectionResponseWithDefaults instantiates a new DeleteCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCollectionResponseWithDefaults() *DeleteCollectionResponse {
	this := DeleteCollectionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteCollectionResponse) GetData() Collection {
	if o == nil || o.Data == nil {
		var ret Collection
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCollectionResponse) GetDataOk() (*Collection, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteCollectionResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Collection and assigns it to the Data field.
func (o *DeleteCollectionResponse) SetData(v Collection) {
	o.Data = &v
}

func (o DeleteCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteCollectionResponse struct {
	value *DeleteCollectionResponse
	isSet bool
}

func (v NullableDeleteCollectionResponse) Get() *DeleteCollectionResponse {
	return v.value
}

func (v *NullableDeleteCollectionResponse) Set(val *DeleteCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCollectionResponse(val *DeleteCollectionResponse) *NullableDeleteCollectionResponse {
	return &NullableDeleteCollectionResponse{value: val, isSet: true}
}

func (v NullableDeleteCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


