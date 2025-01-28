# OrganizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | organization id | 
**Type** | **string** |  | 
**Attributes** | [**OrganizationDataAttributes**](OrganizationDataAttributes.md) |  | 
**Relationships** | [**OrganizationDataRelationships**](OrganizationDataRelationships.md) |  | 

## Methods

### NewOrganizationData

`func NewOrganizationData(id string, type_ string, attributes OrganizationDataAttributes, relationships OrganizationDataRelationships, ) *OrganizationData`

NewOrganizationData instantiates a new OrganizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataWithDefaults

`func NewOrganizationDataWithDefaults() *OrganizationData`

NewOrganizationDataWithDefaults instantiates a new OrganizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *OrganizationData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrganizationData) GetAttributes() OrganizationDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrganizationData) GetAttributesOk() (*OrganizationDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrganizationData) SetAttributes(v OrganizationDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrganizationData) GetRelationships() OrganizationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrganizationData) GetRelationshipsOk() (*OrganizationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrganizationData) SetRelationships(v OrganizationDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


