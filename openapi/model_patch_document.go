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

// PatchDocument struct for PatchDocument
type PatchDocument struct {
	// Unique ID of the document to be patched.
	Id string `json:"_id"`
	// List of patch operations.
	Patch []PatchOperation `json:"patch"`
}

// NewPatchDocument instantiates a new PatchDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchDocument(id string, patch []PatchOperation) *PatchDocument {
	this := PatchDocument{}
	this.Id = id
	this.Patch = patch
	return &this
}

// NewPatchDocumentWithDefaults instantiates a new PatchDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchDocumentWithDefaults() *PatchDocument {
	this := PatchDocument{}
	return &this
}

// GetId returns the Id field value
func (o *PatchDocument) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchDocument) SetId(v string) {
	o.Id = v
}

// GetPatch returns the Patch field value
func (o *PatchDocument) GetPatch() []PatchOperation {
	if o == nil {
		var ret []PatchOperation
		return ret
	}

	return o.Patch
}

// GetPatchOk returns a tuple with the Patch field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetPatchOk() (*[]PatchOperation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Patch, true
}

// SetPatch sets field value
func (o *PatchDocument) SetPatch(v []PatchOperation) {
	o.Patch = v
}

func (o PatchDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_id"] = o.Id
	}
	if true {
		toSerialize["patch"] = o.Patch
	}
	return json.Marshal(toSerialize)
}

type NullablePatchDocument struct {
	value *PatchDocument
	isSet bool
}

func (v NullablePatchDocument) Get() *PatchDocument {
	return v.value
}

func (v *NullablePatchDocument) Set(val *PatchDocument) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchDocument) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchDocument(val *PatchDocument) *NullablePatchDocument {
	return &NullablePatchDocument{value: val, isSet: true}
}

func (v NullablePatchDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


