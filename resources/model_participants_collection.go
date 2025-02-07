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

// checks if the ParticipantsCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParticipantsCollection{}

// ParticipantsCollection struct for ParticipantsCollection
type ParticipantsCollection struct {
	Data []ParticipantData `json:"data"`
	Links Links `json:"links"`
}

type _ParticipantsCollection ParticipantsCollection

// NewParticipantsCollection instantiates a new ParticipantsCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipantsCollection(data []ParticipantData, links Links) *ParticipantsCollection {
	this := ParticipantsCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewParticipantsCollectionWithDefaults instantiates a new ParticipantsCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantsCollectionWithDefaults() *ParticipantsCollection {
	this := ParticipantsCollection{}
	return &this
}

// GetData returns the Data field value
func (o *ParticipantsCollection) GetData() []ParticipantData {
	if o == nil {
		var ret []ParticipantData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParticipantsCollection) GetDataOk() ([]ParticipantData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ParticipantsCollection) SetData(v []ParticipantData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ParticipantsCollection) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ParticipantsCollection) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ParticipantsCollection) SetLinks(v Links) {
	o.Links = v
}

func (o ParticipantsCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParticipantsCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *ParticipantsCollection) UnmarshalJSON(data []byte) (err error) {
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

	varParticipantsCollection := _ParticipantsCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParticipantsCollection)

	if err != nil {
		return err
	}

	*o = ParticipantsCollection(varParticipantsCollection)

	return err
}

type NullableParticipantsCollection struct {
	value *ParticipantsCollection
	isSet bool
}

func (v NullableParticipantsCollection) Get() *ParticipantsCollection {
	return v.value
}

func (v *NullableParticipantsCollection) Set(val *ParticipantsCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantsCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantsCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantsCollection(val *ParticipantsCollection) *NullableParticipantsCollection {
	return &NullableParticipantsCollection{value: val, isSet: true}
}

func (v NullableParticipantsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantsCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


