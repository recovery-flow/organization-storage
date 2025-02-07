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

// checks if the OrganizationCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCreate{}

// OrganizationCreate struct for OrganizationCreate
type OrganizationCreate struct {
	Data OrganizationCreateData `json:"data"`
}

type _OrganizationCreate OrganizationCreate

// NewOrganizationCreate instantiates a new OrganizationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCreate(data OrganizationCreateData) *OrganizationCreate {
	this := OrganizationCreate{}
	this.Data = data
	return &this
}

// NewOrganizationCreateWithDefaults instantiates a new OrganizationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCreateWithDefaults() *OrganizationCreate {
	this := OrganizationCreate{}
	return &this
}

// GetData returns the Data field value
func (o *OrganizationCreate) GetData() OrganizationCreateData {
	if o == nil {
		var ret OrganizationCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrganizationCreate) GetDataOk() (*OrganizationCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrganizationCreate) SetData(v OrganizationCreateData) {
	o.Data = v
}

func (o OrganizationCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *OrganizationCreate) UnmarshalJSON(data []byte) (err error) {
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

	varOrganizationCreate := _OrganizationCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationCreate)

	if err != nil {
		return err
	}

	*o = OrganizationCreate(varOrganizationCreate)

	return err
}

type NullableOrganizationCreate struct {
	value *OrganizationCreate
	isSet bool
}

func (v NullableOrganizationCreate) Get() *OrganizationCreate {
	return v.value
}

func (v *NullableOrganizationCreate) Set(val *OrganizationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCreate(val *OrganizationCreate) *NullableOrganizationCreate {
	return &NullableOrganizationCreate{value: val, isSet: true}
}

func (v NullableOrganizationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


