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

// checks if the OrganizationDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationDataAttributes{}

// OrganizationDataAttributes struct for OrganizationDataAttributes
type OrganizationDataAttributes struct {
	// Team name
	Name string `json:"name"`
	// Link to photo
	Logo string `json:"logo"`
	// Verified status
	Verified string `json:"verified"`
	// Team description
	Desc string `json:"desc"`
	// Type of Organization
	Sort string `json:"sort"`
	// Country of registration
	Country string `json:"country"`
	// City of HQ
	City *string `json:"city,omitempty"`
	Links *Object `json:"links,omitempty"`
	ComplicatedStatus *Object `json:"complicated_status,omitempty"`
	// Team creation timestamp
	CreatedAt time.Time `json:"created_at"`
}

type _OrganizationDataAttributes OrganizationDataAttributes

// NewOrganizationDataAttributes instantiates a new OrganizationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDataAttributes(name string, logo string, verified string, desc string, sort string, country string, createdAt time.Time) *OrganizationDataAttributes {
	this := OrganizationDataAttributes{}
	this.Name = name
	this.Logo = logo
	this.Verified = verified
	this.Desc = desc
	this.Sort = sort
	this.Country = country
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationDataAttributesWithDefaults instantiates a new OrganizationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDataAttributesWithDefaults() *OrganizationDataAttributes {
	this := OrganizationDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationDataAttributes) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value
func (o *OrganizationDataAttributes) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *OrganizationDataAttributes) SetLogo(v string) {
	o.Logo = v
}

// GetVerified returns the Verified field value
func (o *OrganizationDataAttributes) GetVerified() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetVerifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *OrganizationDataAttributes) SetVerified(v string) {
	o.Verified = v
}

// GetDesc returns the Desc field value
func (o *OrganizationDataAttributes) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *OrganizationDataAttributes) SetDesc(v string) {
	o.Desc = v
}

// GetSort returns the Sort field value
func (o *OrganizationDataAttributes) GetSort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetSortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value
func (o *OrganizationDataAttributes) SetSort(v string) {
	o.Sort = v
}

// GetCountry returns the Country field value
func (o *OrganizationDataAttributes) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *OrganizationDataAttributes) SetCountry(v string) {
	o.Country = v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrganizationDataAttributes) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrganizationDataAttributes) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrganizationDataAttributes) SetCity(v string) {
	o.City = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrganizationDataAttributes) GetLinks() Object {
	if o == nil || IsNil(o.Links) {
		var ret Object
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetLinksOk() (*Object, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrganizationDataAttributes) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Object and assigns it to the Links field.
func (o *OrganizationDataAttributes) SetLinks(v Object) {
	o.Links = &v
}

// GetComplicatedStatus returns the ComplicatedStatus field value if set, zero value otherwise.
func (o *OrganizationDataAttributes) GetComplicatedStatus() Object {
	if o == nil || IsNil(o.ComplicatedStatus) {
		var ret Object
		return ret
	}
	return *o.ComplicatedStatus
}

// GetComplicatedStatusOk returns a tuple with the ComplicatedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetComplicatedStatusOk() (*Object, bool) {
	if o == nil || IsNil(o.ComplicatedStatus) {
		return nil, false
	}
	return o.ComplicatedStatus, true
}

// HasComplicatedStatus returns a boolean if a field has been set.
func (o *OrganizationDataAttributes) HasComplicatedStatus() bool {
	if o != nil && !IsNil(o.ComplicatedStatus) {
		return true
	}

	return false
}

// SetComplicatedStatus gets a reference to the given Object and assigns it to the ComplicatedStatus field.
func (o *OrganizationDataAttributes) SetComplicatedStatus(v Object) {
	o.ComplicatedStatus = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o OrganizationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["logo"] = o.Logo
	toSerialize["verified"] = o.Verified
	toSerialize["desc"] = o.Desc
	toSerialize["sort"] = o.Sort
	toSerialize["country"] = o.Country
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.ComplicatedStatus) {
		toSerialize["complicated_status"] = o.ComplicatedStatus
	}
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *OrganizationDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"logo",
		"verified",
		"desc",
		"sort",
		"country",
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

	varOrganizationDataAttributes := _OrganizationDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationDataAttributes)

	if err != nil {
		return err
	}

	*o = OrganizationDataAttributes(varOrganizationDataAttributes)

	return err
}

type NullableOrganizationDataAttributes struct {
	value *OrganizationDataAttributes
	isSet bool
}

func (v NullableOrganizationDataAttributes) Get() *OrganizationDataAttributes {
	return v.value
}

func (v *NullableOrganizationDataAttributes) Set(val *OrganizationDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDataAttributes(val *OrganizationDataAttributes) *NullableOrganizationDataAttributes {
	return &NullableOrganizationDataAttributes{value: val, isSet: true}
}

func (v NullableOrganizationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


