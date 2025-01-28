# ParticipantCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | organization ID | 
**UserId** | **string** | user ID | 
**FirstName** | **string** | first name of participant | 
**SecondName** | **string** | second name of participant | 
**ThirdName** | Pointer to **string** | third name of participant | [optional] 
**DisplayName** | **string** | name of participant | 
**Position** | **string** | position in the company | 
**Desc** | **string** | description of participant | 
**Role** | **string** | participant role | 

## Methods

### NewParticipantCreateDataAttributes

`func NewParticipantCreateDataAttributes(orgId string, userId string, firstName string, secondName string, displayName string, position string, desc string, role string, ) *ParticipantCreateDataAttributes`

NewParticipantCreateDataAttributes instantiates a new ParticipantCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantCreateDataAttributesWithDefaults

`func NewParticipantCreateDataAttributesWithDefaults() *ParticipantCreateDataAttributes`

NewParticipantCreateDataAttributesWithDefaults instantiates a new ParticipantCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ParticipantCreateDataAttributes) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ParticipantCreateDataAttributes) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ParticipantCreateDataAttributes) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetUserId

`func (o *ParticipantCreateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ParticipantCreateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ParticipantCreateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFirstName

`func (o *ParticipantCreateDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ParticipantCreateDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ParticipantCreateDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetSecondName

`func (o *ParticipantCreateDataAttributes) GetSecondName() string`

GetSecondName returns the SecondName field if non-nil, zero value otherwise.

### GetSecondNameOk

`func (o *ParticipantCreateDataAttributes) GetSecondNameOk() (*string, bool)`

GetSecondNameOk returns a tuple with the SecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondName

`func (o *ParticipantCreateDataAttributes) SetSecondName(v string)`

SetSecondName sets SecondName field to given value.


### GetThirdName

`func (o *ParticipantCreateDataAttributes) GetThirdName() string`

GetThirdName returns the ThirdName field if non-nil, zero value otherwise.

### GetThirdNameOk

`func (o *ParticipantCreateDataAttributes) GetThirdNameOk() (*string, bool)`

GetThirdNameOk returns a tuple with the ThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdName

`func (o *ParticipantCreateDataAttributes) SetThirdName(v string)`

SetThirdName sets ThirdName field to given value.

### HasThirdName

`func (o *ParticipantCreateDataAttributes) HasThirdName() bool`

HasThirdName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ParticipantCreateDataAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ParticipantCreateDataAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ParticipantCreateDataAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPosition

`func (o *ParticipantCreateDataAttributes) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ParticipantCreateDataAttributes) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ParticipantCreateDataAttributes) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetDesc

`func (o *ParticipantCreateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ParticipantCreateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ParticipantCreateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetRole

`func (o *ParticipantCreateDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ParticipantCreateDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ParticipantCreateDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


