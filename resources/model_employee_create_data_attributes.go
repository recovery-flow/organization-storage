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

// checks if the EmployeeCreateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeCreateDataAttributes{}

// EmployeeCreateDataAttributes struct for EmployeeCreateDataAttributes
type EmployeeCreateDataAttributes struct {
	// Organization ID
	OrgId string `json:"org_id"`
	// User ID
	UserId string `json:"user_id"`
	// first name of employee
	FirstName string `json:"first_name"`
	// second name of employee
	SecondName string `json:"second_name"`
	// third name of employee
	ThirdName string `json:"third_name"`
	// name of employee
	DisplayName string `json:"display_name"`
	// position in the company
	Position string `json:"position"`
	// Description
	Desc string `json:"desc"`
	// User role
	Role string `json:"role"`
}

type _EmployeeCreateDataAttributes EmployeeCreateDataAttributes

// NewEmployeeCreateDataAttributes instantiates a new EmployeeCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeCreateDataAttributes(orgId string, userId string, firstName string, secondName string, thirdName string, displayName string, position string, desc string, role string) *EmployeeCreateDataAttributes {
	this := EmployeeCreateDataAttributes{}
	this.OrgId = orgId
	this.UserId = userId
	this.FirstName = firstName
	this.SecondName = secondName
	this.ThirdName = thirdName
	this.DisplayName = displayName
	this.Position = position
	this.Desc = desc
	this.Role = role
	return &this
}

// NewEmployeeCreateDataAttributesWithDefaults instantiates a new EmployeeCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeCreateDataAttributesWithDefaults() *EmployeeCreateDataAttributes {
	this := EmployeeCreateDataAttributes{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *EmployeeCreateDataAttributes) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *EmployeeCreateDataAttributes) SetOrgId(v string) {
	o.OrgId = v
}

// GetUserId returns the UserId field value
func (o *EmployeeCreateDataAttributes) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *EmployeeCreateDataAttributes) SetUserId(v string) {
	o.UserId = v
}

// GetFirstName returns the FirstName field value
func (o *EmployeeCreateDataAttributes) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *EmployeeCreateDataAttributes) SetFirstName(v string) {
	o.FirstName = v
}

// GetSecondName returns the SecondName field value
func (o *EmployeeCreateDataAttributes) GetSecondName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondName
}

// GetSecondNameOk returns a tuple with the SecondName field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetSecondNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondName, true
}

// SetSecondName sets field value
func (o *EmployeeCreateDataAttributes) SetSecondName(v string) {
	o.SecondName = v
}

// GetThirdName returns the ThirdName field value
func (o *EmployeeCreateDataAttributes) GetThirdName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThirdName
}

// GetThirdNameOk returns a tuple with the ThirdName field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetThirdNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdName, true
}

// SetThirdName sets field value
func (o *EmployeeCreateDataAttributes) SetThirdName(v string) {
	o.ThirdName = v
}

// GetDisplayName returns the DisplayName field value
func (o *EmployeeCreateDataAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EmployeeCreateDataAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPosition returns the Position field value
func (o *EmployeeCreateDataAttributes) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *EmployeeCreateDataAttributes) SetPosition(v string) {
	o.Position = v
}

// GetDesc returns the Desc field value
func (o *EmployeeCreateDataAttributes) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *EmployeeCreateDataAttributes) SetDesc(v string) {
	o.Desc = v
}

// GetRole returns the Role field value
func (o *EmployeeCreateDataAttributes) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *EmployeeCreateDataAttributes) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *EmployeeCreateDataAttributes) SetRole(v string) {
	o.Role = v
}

func (o EmployeeCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeCreateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["user_id"] = o.UserId
	toSerialize["first_name"] = o.FirstName
	toSerialize["second_name"] = o.SecondName
	toSerialize["third_name"] = o.ThirdName
	toSerialize["display_name"] = o.DisplayName
	toSerialize["position"] = o.Position
	toSerialize["desc"] = o.Desc
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *EmployeeCreateDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
		"user_id",
		"first_name",
		"second_name",
		"third_name",
		"display_name",
		"position",
		"desc",
		"role",
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

	varEmployeeCreateDataAttributes := _EmployeeCreateDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployeeCreateDataAttributes)

	if err != nil {
		return err
	}

	*o = EmployeeCreateDataAttributes(varEmployeeCreateDataAttributes)

	return err
}

type NullableEmployeeCreateDataAttributes struct {
	value *EmployeeCreateDataAttributes
	isSet bool
}

func (v NullableEmployeeCreateDataAttributes) Get() *EmployeeCreateDataAttributes {
	return v.value
}

func (v *NullableEmployeeCreateDataAttributes) Set(val *EmployeeCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeCreateDataAttributes(val *EmployeeCreateDataAttributes) *NullableEmployeeCreateDataAttributes {
	return &NullableEmployeeCreateDataAttributes{value: val, isSet: true}
}

func (v NullableEmployeeCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


