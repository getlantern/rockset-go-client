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

// SourceKafka struct for SourceKafka
type SourceKafka struct {
	// The Kafka topic to be tailed
	KafkaTopicName string `json:"kafka_topic_name"`
	Status *StatusKafka `json:"status,omitempty"`
	UseV3 *bool `json:"use_v3,omitempty"`
	OffsetResetPolicy *string `json:"offset_reset_policy,omitempty"`
}

// NewSourceKafka instantiates a new SourceKafka object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceKafka(kafkaTopicName string) *SourceKafka {
	this := SourceKafka{}
	this.KafkaTopicName = kafkaTopicName
	return &this
}

// NewSourceKafkaWithDefaults instantiates a new SourceKafka object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceKafkaWithDefaults() *SourceKafka {
	this := SourceKafka{}
	return &this
}

// GetKafkaTopicName returns the KafkaTopicName field value
func (o *SourceKafka) GetKafkaTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KafkaTopicName
}

// GetKafkaTopicNameOk returns a tuple with the KafkaTopicName field value
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetKafkaTopicNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KafkaTopicName, true
}

// SetKafkaTopicName sets field value
func (o *SourceKafka) SetKafkaTopicName(v string) {
	o.KafkaTopicName = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SourceKafka) GetStatus() StatusKafka {
	if o == nil || o.Status == nil {
		var ret StatusKafka
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetStatusOk() (*StatusKafka, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SourceKafka) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusKafka and assigns it to the Status field.
func (o *SourceKafka) SetStatus(v StatusKafka) {
	o.Status = &v
}

// GetUseV3 returns the UseV3 field value if set, zero value otherwise.
func (o *SourceKafka) GetUseV3() bool {
	if o == nil || o.UseV3 == nil {
		var ret bool
		return ret
	}
	return *o.UseV3
}

// GetUseV3Ok returns a tuple with the UseV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetUseV3Ok() (*bool, bool) {
	if o == nil || o.UseV3 == nil {
		return nil, false
	}
	return o.UseV3, true
}

// HasUseV3 returns a boolean if a field has been set.
func (o *SourceKafka) HasUseV3() bool {
	if o != nil && o.UseV3 != nil {
		return true
	}

	return false
}

// SetUseV3 gets a reference to the given bool and assigns it to the UseV3 field.
func (o *SourceKafka) SetUseV3(v bool) {
	o.UseV3 = &v
}

// GetOffsetResetPolicy returns the OffsetResetPolicy field value if set, zero value otherwise.
func (o *SourceKafka) GetOffsetResetPolicy() string {
	if o == nil || o.OffsetResetPolicy == nil {
		var ret string
		return ret
	}
	return *o.OffsetResetPolicy
}

// GetOffsetResetPolicyOk returns a tuple with the OffsetResetPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceKafka) GetOffsetResetPolicyOk() (*string, bool) {
	if o == nil || o.OffsetResetPolicy == nil {
		return nil, false
	}
	return o.OffsetResetPolicy, true
}

// HasOffsetResetPolicy returns a boolean if a field has been set.
func (o *SourceKafka) HasOffsetResetPolicy() bool {
	if o != nil && o.OffsetResetPolicy != nil {
		return true
	}

	return false
}

// SetOffsetResetPolicy gets a reference to the given string and assigns it to the OffsetResetPolicy field.
func (o *SourceKafka) SetOffsetResetPolicy(v string) {
	o.OffsetResetPolicy = &v
}

func (o SourceKafka) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kafka_topic_name"] = o.KafkaTopicName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UseV3 != nil {
		toSerialize["use_v3"] = o.UseV3
	}
	if o.OffsetResetPolicy != nil {
		toSerialize["offset_reset_policy"] = o.OffsetResetPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableSourceKafka struct {
	value *SourceKafka
	isSet bool
}

func (v NullableSourceKafka) Get() *SourceKafka {
	return v.value
}

func (v *NullableSourceKafka) Set(val *SourceKafka) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceKafka) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceKafka) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceKafka(val *SourceKafka) *NullableSourceKafka {
	return &NullableSourceKafka{value: val, isSet: true}
}

func (v NullableSourceKafka) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceKafka) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


