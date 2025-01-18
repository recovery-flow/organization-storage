/*
test

example

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the EmployeeDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeDataAttributes{}

// EmployeeDataAttributes struct for EmployeeDataAttributes
type EmployeeDataAttributes struct {
	// first name of employee
	FirstName string `json:"first_name"`
	// second name of employee
	SecondName string `json:"second_name"`
	// third name of employee
	ThirdName *string `json:"third_name,omitempty"`
	// name of employee
	DisplayName string `json:"display_name"`
	// position in the company
	Position string `json:"position"`
	// verified status
	Verified string `json:"verified"`
	// Description
	Desc string `json:"desc"`
	// User role
	Role string `json:"role"`
	// User updated at
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// User created at
	CreatedAt time.Time `json:"created_at"`
}

type _EmployeeDataAttributes EmployeeDataAttributes

// NewEmployeeDataAttributes instantiates a new EmployeeDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeDataAttributes(firstName string, secondName string, displayName string, position string, verified string, desc string, role string, createdAt time.Time) *EmployeeDataAttributes {
	this := EmployeeDataAttributes{}
	this.FirstName = firstName
	this.SecondName = secondName
	this.DisplayName = displayName
	this.Position = position
	this.Verified = verified
	this.Desc = desc
	this.Role = role
	this.CreatedAt = createdAt
	return &this
}

// NewEmployeeDataAttributesWithDefaults instantiates a new EmployeeDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeDataAttributesWithDefaults() *EmployeeDataAttributes {
	this := EmployeeDataAttributes{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *EmployeeDataAttributes) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *EmployeeDataAttributes) SetFirstName(v string) {
	o.FirstName = v
}

// GetSecondName returns the SecondName field value
func (o *EmployeeDataAttributes) GetSecondName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondName
}

// GetSecondNameOk returns a tuple with the SecondName field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetSecondNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondName, true
}

// SetSecondName sets field value
func (o *EmployeeDataAttributes) SetSecondName(v string) {
	o.SecondName = v
}

// GetThirdName returns the ThirdName field value if set, zero value otherwise.
func (o *EmployeeDataAttributes) GetThirdName() string {
	if o == nil || IsNil(o.ThirdName) {
		var ret string
		return ret
	}
	return *o.ThirdName
}

// GetThirdNameOk returns a tuple with the ThirdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetThirdNameOk() (*string, bool) {
	if o == nil || IsNil(o.ThirdName) {
		return nil, false
	}
	return o.ThirdName, true
}

// HasThirdName returns a boolean if a field has been set.
func (o *EmployeeDataAttributes) HasThirdName() bool {
	if o != nil && !IsNil(o.ThirdName) {
		return true
	}

	return false
}

// SetThirdName gets a reference to the given string and assigns it to the ThirdName field.
func (o *EmployeeDataAttributes) SetThirdName(v string) {
	o.ThirdName = &v
}

// GetDisplayName returns the DisplayName field value
func (o *EmployeeDataAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EmployeeDataAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPosition returns the Position field value
func (o *EmployeeDataAttributes) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *EmployeeDataAttributes) SetPosition(v string) {
	o.Position = v
}

// GetVerified returns the Verified field value
func (o *EmployeeDataAttributes) GetVerified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetVerifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *EmployeeDataAttributes) SetVerified(v string) {
	o.Verified = v
}

// GetDesc returns the Desc field value
func (o *EmployeeDataAttributes) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *EmployeeDataAttributes) SetDesc(v string) {
	o.Desc = v
}

// GetRole returns the Role field value
func (o *EmployeeDataAttributes) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *EmployeeDataAttributes) SetRole(v string) {
	o.Role = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EmployeeDataAttributes) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EmployeeDataAttributes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *EmployeeDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EmployeeDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EmployeeDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EmployeeDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o EmployeeDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["first_name"] = o.FirstName
	toSerialize["second_name"] = o.SecondName
	if !IsNil(o.ThirdName) {
		toSerialize["third_name"] = o.ThirdName
	}
	toSerialize["display_name"] = o.DisplayName
	toSerialize["position"] = o.Position
	toSerialize["verified"] = o.Verified
	toSerialize["desc"] = o.Desc
	toSerialize["role"] = o.Role
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *EmployeeDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first_name",
		"second_name",
		"display_name",
		"position",
		"verified",
		"desc",
		"role",
		"created_at",
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

	varEmployeeDataAttributes := _EmployeeDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployeeDataAttributes)

	if err != nil {
		return err
	}

	*o = EmployeeDataAttributes(varEmployeeDataAttributes)

	return err
}

type NullableEmployeeDataAttributes struct {
	value *EmployeeDataAttributes
	isSet bool
}

func (v NullableEmployeeDataAttributes) Get() *EmployeeDataAttributes {
	return v.value
}

func (v *NullableEmployeeDataAttributes) Set(val *EmployeeDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeDataAttributes(val *EmployeeDataAttributes) *NullableEmployeeDataAttributes {
	return &NullableEmployeeDataAttributes{value: val, isSet: true}
}

func (v NullableEmployeeDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


