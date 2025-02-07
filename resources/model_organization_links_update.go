/*
test

example

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrganizationLinksUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationLinksUpdate{}

// OrganizationLinksUpdate struct for OrganizationLinksUpdate
type OrganizationLinksUpdate struct {
	Data OrganizationLinksUpdateData `json:"data"`
}

type _OrganizationLinksUpdate OrganizationLinksUpdate

// NewOrganizationLinksUpdate instantiates a new OrganizationLinksUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationLinksUpdate(data OrganizationLinksUpdateData) *OrganizationLinksUpdate {
	this := OrganizationLinksUpdate{}
	this.Data = data
	return &this
}

// NewOrganizationLinksUpdateWithDefaults instantiates a new OrganizationLinksUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationLinksUpdateWithDefaults() *OrganizationLinksUpdate {
	this := OrganizationLinksUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *OrganizationLinksUpdate) GetData() OrganizationLinksUpdateData {
	if o == nil {
		var ret OrganizationLinksUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrganizationLinksUpdate) GetDataOk() (*OrganizationLinksUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrganizationLinksUpdate) SetData(v OrganizationLinksUpdateData) {
	o.Data = v
}

func (o OrganizationLinksUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationLinksUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *OrganizationLinksUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrganizationLinksUpdate := _OrganizationLinksUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationLinksUpdate)

	if err != nil {
		return err
	}

	*o = OrganizationLinksUpdate(varOrganizationLinksUpdate)

	return err
}

type NullableOrganizationLinksUpdate struct {
	value *OrganizationLinksUpdate
	isSet bool
}

func (v NullableOrganizationLinksUpdate) Get() *OrganizationLinksUpdate {
	return v.value
}

func (v *NullableOrganizationLinksUpdate) Set(val *OrganizationLinksUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationLinksUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationLinksUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationLinksUpdate(val *OrganizationLinksUpdate) *NullableOrganizationLinksUpdate {
	return &NullableOrganizationLinksUpdate{value: val, isSet: true}
}

func (v NullableOrganizationLinksUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationLinksUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


