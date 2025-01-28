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

// checks if the ParticipantCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParticipantCollection{}

// ParticipantCollection struct for ParticipantCollection
type ParticipantCollection struct {
	Data []ParticipantData `json:"data"`
	Links Links `json:"links"`
}

type _ParticipantCollection ParticipantCollection

// NewParticipantCollection instantiates a new ParticipantCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipantCollection(data []ParticipantData, links Links) *ParticipantCollection {
	this := ParticipantCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewParticipantCollectionWithDefaults instantiates a new ParticipantCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantCollectionWithDefaults() *ParticipantCollection {
	this := ParticipantCollection{}
	return &this
}

// GetData returns the Data field value
func (o *ParticipantCollection) GetData() []ParticipantData {
	if o == nil {
		var ret []ParticipantData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParticipantCollection) GetDataOk() ([]ParticipantData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ParticipantCollection) SetData(v []ParticipantData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ParticipantCollection) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ParticipantCollection) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ParticipantCollection) SetLinks(v Links) {
	o.Links = v
}

func (o ParticipantCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParticipantCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *ParticipantCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varParticipantCollection := _ParticipantCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParticipantCollection)

	if err != nil {
		return err
	}

	*o = ParticipantCollection(varParticipantCollection)

	return err
}

type NullableParticipantCollection struct {
	value *ParticipantCollection
	isSet bool
}

func (v NullableParticipantCollection) Get() *ParticipantCollection {
	return v.value
}

func (v *NullableParticipantCollection) Set(val *ParticipantCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantCollection(val *ParticipantCollection) *NullableParticipantCollection {
	return &NullableParticipantCollection{value: val, isSet: true}
}

func (v NullableParticipantCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


