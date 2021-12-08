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

// InvertedIndexGroupEncodingOptions struct for InvertedIndexGroupEncodingOptions
type InvertedIndexGroupEncodingOptions struct {
	FormatVersion *int64 `json:"format_version,omitempty"`
	GroupSize *int64 `json:"group_size,omitempty"`
	RestartLength *int64 `json:"restart_length,omitempty"`
	EventTimeCodec *string `json:"event_time_codec,omitempty"`
	DocIdCodec *string `json:"doc_id_codec,omitempty"`
}

// NewInvertedIndexGroupEncodingOptions instantiates a new InvertedIndexGroupEncodingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvertedIndexGroupEncodingOptions() *InvertedIndexGroupEncodingOptions {
	this := InvertedIndexGroupEncodingOptions{}
	return &this
}

// NewInvertedIndexGroupEncodingOptionsWithDefaults instantiates a new InvertedIndexGroupEncodingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvertedIndexGroupEncodingOptionsWithDefaults() *InvertedIndexGroupEncodingOptions {
	this := InvertedIndexGroupEncodingOptions{}
	return &this
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *InvertedIndexGroupEncodingOptions) GetFormatVersion() int64 {
	if o == nil || o.FormatVersion == nil {
		var ret int64
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvertedIndexGroupEncodingOptions) GetFormatVersionOk() (*int64, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *InvertedIndexGroupEncodingOptions) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given int64 and assigns it to the FormatVersion field.
func (o *InvertedIndexGroupEncodingOptions) SetFormatVersion(v int64) {
	o.FormatVersion = &v
}

// GetGroupSize returns the GroupSize field value if set, zero value otherwise.
func (o *InvertedIndexGroupEncodingOptions) GetGroupSize() int64 {
	if o == nil || o.GroupSize == nil {
		var ret int64
		return ret
	}
	return *o.GroupSize
}

// GetGroupSizeOk returns a tuple with the GroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvertedIndexGroupEncodingOptions) GetGroupSizeOk() (*int64, bool) {
	if o == nil || o.GroupSize == nil {
		return nil, false
	}
	return o.GroupSize, true
}

// HasGroupSize returns a boolean if a field has been set.
func (o *InvertedIndexGroupEncodingOptions) HasGroupSize() bool {
	if o != nil && o.GroupSize != nil {
		return true
	}

	return false
}

// SetGroupSize gets a reference to the given int64 and assigns it to the GroupSize field.
func (o *InvertedIndexGroupEncodingOptions) SetGroupSize(v int64) {
	o.GroupSize = &v
}

// GetRestartLength returns the RestartLength field value if set, zero value otherwise.
func (o *InvertedIndexGroupEncodingOptions) GetRestartLength() int64 {
	if o == nil || o.RestartLength == nil {
		var ret int64
		return ret
	}
	return *o.RestartLength
}

// GetRestartLengthOk returns a tuple with the RestartLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvertedIndexGroupEncodingOptions) GetRestartLengthOk() (*int64, bool) {
	if o == nil || o.RestartLength == nil {
		return nil, false
	}
	return o.RestartLength, true
}

// HasRestartLength returns a boolean if a field has been set.
func (o *InvertedIndexGroupEncodingOptions) HasRestartLength() bool {
	if o != nil && o.RestartLength != nil {
		return true
	}

	return false
}

// SetRestartLength gets a reference to the given int64 and assigns it to the RestartLength field.
func (o *InvertedIndexGroupEncodingOptions) SetRestartLength(v int64) {
	o.RestartLength = &v
}

// GetEventTimeCodec returns the EventTimeCodec field value if set, zero value otherwise.
func (o *InvertedIndexGroupEncodingOptions) GetEventTimeCodec() string {
	if o == nil || o.EventTimeCodec == nil {
		var ret string
		return ret
	}
	return *o.EventTimeCodec
}

// GetEventTimeCodecOk returns a tuple with the EventTimeCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvertedIndexGroupEncodingOptions) GetEventTimeCodecOk() (*string, bool) {
	if o == nil || o.EventTimeCodec == nil {
		return nil, false
	}
	return o.EventTimeCodec, true
}

// HasEventTimeCodec returns a boolean if a field has been set.
func (o *InvertedIndexGroupEncodingOptions) HasEventTimeCodec() bool {
	if o != nil && o.EventTimeCodec != nil {
		return true
	}

	return false
}

// SetEventTimeCodec gets a reference to the given string and assigns it to the EventTimeCodec field.
func (o *InvertedIndexGroupEncodingOptions) SetEventTimeCodec(v string) {
	o.EventTimeCodec = &v
}

// GetDocIdCodec returns the DocIdCodec field value if set, zero value otherwise.
func (o *InvertedIndexGroupEncodingOptions) GetDocIdCodec() string {
	if o == nil || o.DocIdCodec == nil {
		var ret string
		return ret
	}
	return *o.DocIdCodec
}

// GetDocIdCodecOk returns a tuple with the DocIdCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvertedIndexGroupEncodingOptions) GetDocIdCodecOk() (*string, bool) {
	if o == nil || o.DocIdCodec == nil {
		return nil, false
	}
	return o.DocIdCodec, true
}

// HasDocIdCodec returns a boolean if a field has been set.
func (o *InvertedIndexGroupEncodingOptions) HasDocIdCodec() bool {
	if o != nil && o.DocIdCodec != nil {
		return true
	}

	return false
}

// SetDocIdCodec gets a reference to the given string and assigns it to the DocIdCodec field.
func (o *InvertedIndexGroupEncodingOptions) SetDocIdCodec(v string) {
	o.DocIdCodec = &v
}

func (o InvertedIndexGroupEncodingOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FormatVersion != nil {
		toSerialize["format_version"] = o.FormatVersion
	}
	if o.GroupSize != nil {
		toSerialize["group_size"] = o.GroupSize
	}
	if o.RestartLength != nil {
		toSerialize["restart_length"] = o.RestartLength
	}
	if o.EventTimeCodec != nil {
		toSerialize["event_time_codec"] = o.EventTimeCodec
	}
	if o.DocIdCodec != nil {
		toSerialize["doc_id_codec"] = o.DocIdCodec
	}
	return json.Marshal(toSerialize)
}

type NullableInvertedIndexGroupEncodingOptions struct {
	value *InvertedIndexGroupEncodingOptions
	isSet bool
}

func (v NullableInvertedIndexGroupEncodingOptions) Get() *InvertedIndexGroupEncodingOptions {
	return v.value
}

func (v *NullableInvertedIndexGroupEncodingOptions) Set(val *InvertedIndexGroupEncodingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInvertedIndexGroupEncodingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInvertedIndexGroupEncodingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvertedIndexGroupEncodingOptions(val *InvertedIndexGroupEncodingOptions) *NullableInvertedIndexGroupEncodingOptions {
	return &NullableInvertedIndexGroupEncodingOptions{value: val, isSet: true}
}

func (v NullableInvertedIndexGroupEncodingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvertedIndexGroupEncodingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


