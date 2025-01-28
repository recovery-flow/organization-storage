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

// checks if the ParticipantData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParticipantData{}

// ParticipantData struct for ParticipantData
type ParticipantData struct {
	// user id
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes ParticipantDataAttributes `json:"attributes"`
	Relationships ParticipantDataRelationships `json:"relationships"`
}

type _ParticipantData ParticipantData

// NewParticipantData instantiates a new ParticipantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipantData(id string, type_ string, attributes ParticipantDataAttributes, relationships ParticipantDataRelationships) *ParticipantData {
	this := ParticipantData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewParticipantDataWithDefaults instantiates a new ParticipantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantDataWithDefaults() *ParticipantData {
	this := ParticipantData{}
	return &this
}

// GetId returns the Id field value
func (o *ParticipantData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParticipantData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParticipantData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ParticipantData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParticipantData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParticipantData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ParticipantData) GetAttributes() ParticipantDataAttributes {
	if o == nil {
		var ret ParticipantDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParticipantData) GetAttributesOk() (*ParticipantDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParticipantData) SetAttributes(v ParticipantDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *ParticipantData) GetRelationships() ParticipantDataRelationships {
	if o == nil {
		var ret ParticipantDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *ParticipantData) GetRelationshipsOk() (*ParticipantDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *ParticipantData) SetRelationships(v ParticipantDataRelationships) {
	o.Relationships = v
}

func (o ParticipantData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParticipantData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *ParticipantData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
		"relationships",
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

	varParticipantData := _ParticipantData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParticipantData)

	if err != nil {
		return err
	}

	*o = ParticipantData(varParticipantData)

	return err
}

type NullableParticipantData struct {
	value *ParticipantData
	isSet bool
}

func (v NullableParticipantData) Get() *ParticipantData {
	return v.value
}

func (v *NullableParticipantData) Set(val *ParticipantData) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantData) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantData(val *ParticipantData) *NullableParticipantData {
	return &NullableParticipantData{value: val, isSet: true}
}

func (v NullableParticipantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


