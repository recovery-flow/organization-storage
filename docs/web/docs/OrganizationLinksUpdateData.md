# OrganizationLinksUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Organization ID | 
**Type** | **string** |  | 
**Attributes** | [**OrganizationLinks**](OrganizationLinks.md) |  | 

## Methods

### NewOrganizationLinksUpdateData

`func NewOrganizationLinksUpdateData(id string, type_ string, attributes OrganizationLinks, ) *OrganizationLinksUpdateData`

NewOrganizationLinksUpdateData instantiates a new OrganizationLinksUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationLinksUpdateDataWithDefaults

`func NewOrganizationLinksUpdateDataWithDefaults() *OrganizationLinksUpdateData`

NewOrganizationLinksUpdateDataWithDefaults instantiates a new OrganizationLinksUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationLinksUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationLinksUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationLinksUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *OrganizationLinksUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationLinksUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationLinksUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrganizationLinksUpdateData) GetAttributes() OrganizationLinks`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrganizationLinksUpdateData) GetAttributesOk() (*OrganizationLinks, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrganizationLinksUpdateData) SetAttributes(v OrganizationLinks)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


