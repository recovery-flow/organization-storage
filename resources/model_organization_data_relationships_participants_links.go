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

// checks if the OrganizationDataRelationshipsParticipantsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationDataRelationshipsParticipantsLinks{}

// OrganizationDataRelationshipsParticipantsLinks struct for OrganizationDataRelationshipsParticipantsLinks
type OrganizationDataRelationshipsParticipantsLinks struct {
	// Link to participants
	Self string `json:"self"`
	// Link to participants
	Related string `json:"related"`
}

type _OrganizationDataRelationshipsParticipantsLinks OrganizationDataRelationshipsParticipantsLinks

// NewOrganizationDataRelationshipsParticipantsLinks instantiates a new OrganizationDataRelationshipsParticipantsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDataRelationshipsParticipantsLinks(self string, related string) *OrganizationDataRelationshipsParticipantsLinks {
	this := OrganizationDataRelationshipsParticipantsLinks{}
	this.Self = self
	this.Related = related
	return &this
}

// NewOrganizationDataRelationshipsParticipantsLinksWithDefaults instantiates a new OrganizationDataRelationshipsParticipantsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDataRelationshipsParticipantsLinksWithDefaults() *OrganizationDataRelationshipsParticipantsLinks {
	this := OrganizationDataRelationshipsParticipantsLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *OrganizationDataRelationshipsParticipantsLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataRelationshipsParticipantsLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *OrganizationDataRelationshipsParticipantsLinks) SetSelf(v string) {
	o.Self = v
}

// GetRelated returns the Related field value
func (o *OrganizationDataRelationshipsParticipantsLinks) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataRelationshipsParticipantsLinks) GetRelatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *OrganizationDataRelationshipsParticipantsLinks) SetRelated(v string) {
	o.Related = v
}

func (o OrganizationDataRelationshipsParticipantsLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationDataRelationshipsParticipantsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["related"] = o.Related
	return toSerialize, nil
}

func (o *OrganizationDataRelationshipsParticipantsLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
		"related",
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

	varOrganizationDataRelationshipsParticipantsLinks := _OrganizationDataRelationshipsParticipantsLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationDataRelationshipsParticipantsLinks)

	if err != nil {
		return err
	}

	*o = OrganizationDataRelationshipsParticipantsLinks(varOrganizationDataRelationshipsParticipantsLinks)

	return err
}

type NullableOrganizationDataRelationshipsParticipantsLinks struct {
	value *OrganizationDataRelationshipsParticipantsLinks
	isSet bool
}

func (v NullableOrganizationDataRelationshipsParticipantsLinks) Get() *OrganizationDataRelationshipsParticipantsLinks {
	return v.value
}

func (v *NullableOrganizationDataRelationshipsParticipantsLinks) Set(val *OrganizationDataRelationshipsParticipantsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDataRelationshipsParticipantsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDataRelationshipsParticipantsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDataRelationshipsParticipantsLinks(val *OrganizationDataRelationshipsParticipantsLinks) *NullableOrganizationDataRelationshipsParticipantsLinks {
	return &NullableOrganizationDataRelationshipsParticipantsLinks{value: val, isSet: true}
}

func (v NullableOrganizationDataRelationshipsParticipantsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDataRelationshipsParticipantsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


