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

// checks if the EmployeeCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeCreate{}

// EmployeeCreate struct for EmployeeCreate
type EmployeeCreate struct {
	Data EmployeeCreateData `json:"data"`
}

type _EmployeeCreate EmployeeCreate

// NewEmployeeCreate instantiates a new EmployeeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeCreate(data EmployeeCreateData) *EmployeeCreate {
	this := EmployeeCreate{}
	this.Data = data
	return &this
}

// NewEmployeeCreateWithDefaults instantiates a new EmployeeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeCreateWithDefaults() *EmployeeCreate {
	this := EmployeeCreate{}
	return &this
}

// GetData returns the Data field value
func (o *EmployeeCreate) GetData() EmployeeCreateData {
	if o == nil {
		var ret EmployeeCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreate) GetDataOk() (*EmployeeCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EmployeeCreate) SetData(v EmployeeCreateData) {
	o.Data = v
}

func (o EmployeeCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EmployeeCreate) UnmarshalJSON(data []byte) (err error) {
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

	varEmployeeCreate := _EmployeeCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployeeCreate)

	if err != nil {
		return err
	}

	*o = EmployeeCreate(varEmployeeCreate)

	return err
}

type NullableEmployeeCreate struct {
	value *EmployeeCreate
	isSet bool
}

func (v NullableEmployeeCreate) Get() *EmployeeCreate {
	return v.value
}

func (v *NullableEmployeeCreate) Set(val *EmployeeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeCreate(val *EmployeeCreate) *NullableEmployeeCreate {
	return &NullableEmployeeCreate{value: val, isSet: true}
}

func (v NullableEmployeeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


