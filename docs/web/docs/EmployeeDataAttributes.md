# EmployeeDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of employee | 
**Desc** | **string** | Description | 
**Role** | **string** | User role | 
**CreatedAt** | **time.Time** | User created at | 

## Methods

### NewEmployeeDataAttributes

`func NewEmployeeDataAttributes(name string, desc string, role string, createdAt time.Time, ) *EmployeeDataAttributes`

NewEmployeeDataAttributes instantiates a new EmployeeDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeDataAttributesWithDefaults

`func NewEmployeeDataAttributesWithDefaults() *EmployeeDataAttributes`

NewEmployeeDataAttributesWithDefaults instantiates a new EmployeeDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmployeeDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployeeDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployeeDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *EmployeeDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *EmployeeDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *EmployeeDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetRole

`func (o *EmployeeDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EmployeeDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EmployeeDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *EmployeeDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmployeeDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmployeeDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


