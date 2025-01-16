# OrganizationUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | organization id | 
**Type** | **string** |  | 
**Attributes** | [**OrganizationCreateDataAttributes**](OrganizationCreateDataAttributes.md) |  | 

## Methods

### NewOrganizationUpdateData

`func NewOrganizationUpdateData(id string, type_ string, attributes OrganizationCreateDataAttributes, ) *OrganizationUpdateData`

NewOrganizationUpdateData instantiates a new OrganizationUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUpdateDataWithDefaults

`func NewOrganizationUpdateDataWithDefaults() *OrganizationUpdateData`

NewOrganizationUpdateDataWithDefaults instantiates a new OrganizationUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *OrganizationUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrganizationUpdateData) GetAttributes() OrganizationCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrganizationUpdateData) GetAttributesOk() (*OrganizationCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrganizationUpdateData) SetAttributes(v OrganizationCreateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


