# EmployeeCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | Organization ID | 
**UserId** | **string** | User ID | 
**FirstName** | **string** | first name of employee | 
**SecondName** | **string** | second name of employee | 
**ThirdName** | Pointer to **string** | third name of employee | [optional] 
**DisplayName** | **string** | name of employee | 
**Position** | **string** | position in the company | 
**Desc** | **string** | Description | 
**Role** | **string** | User role | 

## Methods

### NewEmployeeCreateDataAttributes

`func NewEmployeeCreateDataAttributes(orgId string, userId string, firstName string, secondName string, displayName string, position string, desc string, role string, ) *EmployeeCreateDataAttributes`

NewEmployeeCreateDataAttributes instantiates a new EmployeeCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeCreateDataAttributesWithDefaults

`func NewEmployeeCreateDataAttributesWithDefaults() *EmployeeCreateDataAttributes`

NewEmployeeCreateDataAttributesWithDefaults instantiates a new EmployeeCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *EmployeeCreateDataAttributes) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EmployeeCreateDataAttributes) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EmployeeCreateDataAttributes) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetUserId

`func (o *EmployeeCreateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EmployeeCreateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EmployeeCreateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFirstName

`func (o *EmployeeCreateDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployeeCreateDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployeeCreateDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetSecondName

`func (o *EmployeeCreateDataAttributes) GetSecondName() string`

GetSecondName returns the SecondName field if non-nil, zero value otherwise.

### GetSecondNameOk

`func (o *EmployeeCreateDataAttributes) GetSecondNameOk() (*string, bool)`

GetSecondNameOk returns a tuple with the SecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondName

`func (o *EmployeeCreateDataAttributes) SetSecondName(v string)`

SetSecondName sets SecondName field to given value.


### GetThirdName

`func (o *EmployeeCreateDataAttributes) GetThirdName() string`

GetThirdName returns the ThirdName field if non-nil, zero value otherwise.

### GetThirdNameOk

`func (o *EmployeeCreateDataAttributes) GetThirdNameOk() (*string, bool)`

GetThirdNameOk returns a tuple with the ThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdName

`func (o *EmployeeCreateDataAttributes) SetThirdName(v string)`

SetThirdName sets ThirdName field to given value.

### HasThirdName

`func (o *EmployeeCreateDataAttributes) HasThirdName() bool`

HasThirdName returns a boolean if a field has been set.

### GetDisplayName

`func (o *EmployeeCreateDataAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EmployeeCreateDataAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EmployeeCreateDataAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPosition

`func (o *EmployeeCreateDataAttributes) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EmployeeCreateDataAttributes) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EmployeeCreateDataAttributes) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetDesc

`func (o *EmployeeCreateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *EmployeeCreateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *EmployeeCreateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetRole

`func (o *EmployeeCreateDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EmployeeCreateDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EmployeeCreateDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


