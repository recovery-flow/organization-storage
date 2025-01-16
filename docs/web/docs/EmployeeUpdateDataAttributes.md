# EmployeeUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | Organization id | 
**UserId** | **string** | User id | 
**Name** | Pointer to **string** | name of employee | [optional] 
**Desc** | Pointer to **string** | Description | [optional] 
**Role** | Pointer to **string** | User role | [optional] 

## Methods

### NewEmployeeUpdateDataAttributes

`func NewEmployeeUpdateDataAttributes(orgId string, userId string, ) *EmployeeUpdateDataAttributes`

NewEmployeeUpdateDataAttributes instantiates a new EmployeeUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeUpdateDataAttributesWithDefaults

`func NewEmployeeUpdateDataAttributesWithDefaults() *EmployeeUpdateDataAttributes`

NewEmployeeUpdateDataAttributesWithDefaults instantiates a new EmployeeUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *EmployeeUpdateDataAttributes) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EmployeeUpdateDataAttributes) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EmployeeUpdateDataAttributes) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetUserId

`func (o *EmployeeUpdateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmployeeUpdateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmployeeUpdateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *EmployeeUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmployeeUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmployeeUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmployeeUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesc

`func (o *EmployeeUpdateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *EmployeeUpdateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *EmployeeUpdateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *EmployeeUpdateDataAttributes) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetRole

`func (o *EmployeeUpdateDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EmployeeUpdateDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EmployeeUpdateDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EmployeeUpdateDataAttributes) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


