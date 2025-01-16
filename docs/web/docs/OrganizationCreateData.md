# OrganizationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**OrganizationCreateDataAttributes**](OrganizationCreateDataAttributes.md) |  | 

## Methods

### NewOrganizationCreateData

`func NewOrganizationCreateData(type_ string, attributes OrganizationCreateDataAttributes, ) *OrganizationCreateData`

NewOrganizationCreateData instantiates a new OrganizationCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCreateDataWithDefaults

`func NewOrganizationCreateDataWithDefaults() *OrganizationCreateData`

NewOrganizationCreateDataWithDefaults instantiates a new OrganizationCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrganizationCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrganizationCreateData) GetAttributes() OrganizationCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrganizationCreateData) GetAttributesOk() (*OrganizationCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrganizationCreateData) SetAttributes(v OrganizationCreateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


