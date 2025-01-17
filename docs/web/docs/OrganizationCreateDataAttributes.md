# OrganizationCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Team name | 
**Desc** | **string** | Team description | 
**Sort** | **string** | Type of Organization | 
**Country** | **string** | Country of registration | 
**Owner** | [**OrganizationCreateDataAttributesOwner**](OrganizationCreateDataAttributesOwner.md) |  | 
**City** | Pointer to **string** | City of HQ | [optional] 

## Methods

### NewOrganizationCreateDataAttributes

`func NewOrganizationCreateDataAttributes(name string, desc string, sort string, country string, owner OrganizationCreateDataAttributesOwner, ) *OrganizationCreateDataAttributes`

NewOrganizationCreateDataAttributes instantiates a new OrganizationCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCreateDataAttributesWithDefaults

`func NewOrganizationCreateDataAttributesWithDefaults() *OrganizationCreateDataAttributes`

NewOrganizationCreateDataAttributesWithDefaults instantiates a new OrganizationCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *OrganizationCreateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *OrganizationCreateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *OrganizationCreateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetSort

`func (o *OrganizationCreateDataAttributes) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *OrganizationCreateDataAttributes) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *OrganizationCreateDataAttributes) SetSort(v string)`

SetSort sets Sort field to given value.


### GetCountry

`func (o *OrganizationCreateDataAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationCreateDataAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationCreateDataAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetOwner

`func (o *OrganizationCreateDataAttributes) GetOwner() OrganizationCreateDataAttributesOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OrganizationCreateDataAttributes) GetOwnerOk() (*OrganizationCreateDataAttributesOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OrganizationCreateDataAttributes) SetOwner(v OrganizationCreateDataAttributesOwner)`

SetOwner sets Owner field to given value.


### GetCity

`func (o *OrganizationCreateDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationCreateDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationCreateDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationCreateDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


