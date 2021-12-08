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

// QueryLambda struct for QueryLambda
type QueryLambda struct {
	// workspace of this Query Lambda
	Workspace *string `json:"workspace,omitempty"`
	// user that created this Query Lambda
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// ISO-8601 date of when Query Lambda was last updated
	LastUpdated *string `json:"last_updated,omitempty"`
	// Query Lambda name
	Name *string `json:"name,omitempty"`
	// number of Query Lambda versions
	VersionCount *int32 `json:"version_count,omitempty"`
	// collections/aliases queried by underlying SQL query
	Collections *[]string `json:"collections,omitempty"`
	LatestVersion *QueryLambdaVersion `json:"latest_version,omitempty"`
}

// NewQueryLambda instantiates a new QueryLambda object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLambda() *QueryLambda {
	this := QueryLambda{}
	return &this
}

// NewQueryLambdaWithDefaults instantiates a new QueryLambda object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLambdaWithDefaults() *QueryLambda {
	this := QueryLambda{}
	return &this
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *QueryLambda) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambda) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *QueryLambda) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *QueryLambda) SetWorkspace(v string) {
	o.Workspace = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *QueryLambda) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambda) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *QueryLambda) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *QueryLambda) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *QueryLambda) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambda) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *QueryLambda) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *QueryLambda) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryLambda) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambda) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueryLambda) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryLambda) SetName(v string) {
	o.Name = &v
}

// GetVersionCount returns the VersionCount field value if set, zero value otherwise.
func (o *QueryLambda) GetVersionCount() int32 {
	if o == nil || o.VersionCount == nil {
		var ret int32
		return ret
	}
	return *o.VersionCount
}

// GetVersionCountOk returns a tuple with the VersionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambda) GetVersionCountOk() (*int32, bool) {
	if o == nil || o.VersionCount == nil {
		return nil, false
	}
	return o.VersionCount, true
}

// HasVersionCount returns a boolean if a field has been set.
func (o *QueryLambda) HasVersionCount() bool {
	if o != nil && o.VersionCount != nil {
		return true
	}

	return false
}

// SetVersionCount gets a reference to the given int32 and assigns it to the VersionCount field.
func (o *QueryLambda) SetVersionCount(v int32) {
	o.VersionCount = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *QueryLambda) GetCollections() []string {
	if o == nil || o.Collections == nil {
		var ret []string
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambda) GetCollectionsOk() (*[]string, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *QueryLambda) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []string and assigns it to the Collections field.
func (o *QueryLambda) SetCollections(v []string) {
	o.Collections = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *QueryLambda) GetLatestVersion() QueryLambdaVersion {
	if o == nil || o.LatestVersion == nil {
		var ret QueryLambdaVersion
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambda) GetLatestVersionOk() (*QueryLambdaVersion, bool) {
	if o == nil || o.LatestVersion == nil {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *QueryLambda) HasLatestVersion() bool {
	if o != nil && o.LatestVersion != nil {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given QueryLambdaVersion and assigns it to the LatestVersion field.
func (o *QueryLambda) SetLatestVersion(v QueryLambdaVersion) {
	o.LatestVersion = &v
}

func (o QueryLambda) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.VersionCount != nil {
		toSerialize["version_count"] = o.VersionCount
	}
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
	}
	if o.LatestVersion != nil {
		toSerialize["latest_version"] = o.LatestVersion
	}
	return json.Marshal(toSerialize)
}

type NullableQueryLambda struct {
	value *QueryLambda
	isSet bool
}

func (v NullableQueryLambda) Get() *QueryLambda {
	return v.value
}

func (v *NullableQueryLambda) Set(val *QueryLambda) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLambda) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLambda) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLambda(val *QueryLambda) *NullableQueryLambda {
	return &NullableQueryLambda{value: val, isSet: true}
}

func (v NullableQueryLambda) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLambda) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


