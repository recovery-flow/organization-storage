# OrganizationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Team name | 
**Logo** | **string** | link to photo | 
**Desc** | **string** | Team description | 
**Sort** | **string** | Type of Organization | 
**Links** | [**Object**](Object.md) |  | 
**ComplicatedStatus** | [**Object**](Object.md) |  | 
**CreatedAt** | **time.Time** | Team creation timestamp | 

## Methods

### NewOrganizationDataAttributes

`func NewOrganizationDataAttributes(name string, logo string, desc string, sort string, links Object, complicatedStatus Object, createdAt time.Time, ) *OrganizationDataAttributes`

NewOrganizationDataAttributes instantiates a new OrganizationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDataAttributesWithDefaults

`func NewOrganizationDataAttributesWithDefaults() *OrganizationDataAttributes`

NewOrganizationDataAttributesWithDefaults instantiates a new OrganizationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *OrganizationDataAttributes) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *OrganizationDataAttributes) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *OrganizationDataAttributes) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetDesc

`func (o *OrganizationDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *OrganizationDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *OrganizationDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetSort

`func (o *OrganizationDataAttributes) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *OrganizationDataAttributes) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *OrganizationDataAttributes) SetSort(v string)`

SetSort sets Sort field to given value.


### GetLinks

`func (o *OrganizationDataAttributes) GetLinks() Object`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganizationDataAttributes) GetLinksOk() (*Object, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganizationDataAttributes) SetLinks(v Object)`

SetLinks sets Links field to given value.


### GetComplicatedStatus

`func (o *OrganizationDataAttributes) GetComplicatedStatus() Object`

GetComplicatedStatus returns the ComplicatedStatus field if non-nil, zero value otherwise.

### GetComplicatedStatusOk

`func (o *OrganizationDataAttributes) GetComplicatedStatusOk() (*Object, bool)`

GetComplicatedStatusOk returns a tuple with the ComplicatedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplicatedStatus

`func (o *OrganizationDataAttributes) SetComplicatedStatus(v Object)`

SetComplicatedStatus sets ComplicatedStatus field to given value.


### GetCreatedAt

`func (o *OrganizationDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


