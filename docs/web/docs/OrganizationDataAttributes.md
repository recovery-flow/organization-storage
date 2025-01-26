# OrganizationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Team name | 
**Logo** | **string** | Link to photo | 
**Verified** | **string** | Verified status | 
**Desc** | **string** | Team description | 
**Sort** | **string** | Type of Organization | 
**Country** | **string** | Country of registration | 
**City** | Pointer to **string** | City of HQ | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**ComplicatedStatus** | Pointer to [**ComplicatedStatus**](ComplicatedStatus.md) |  | [optional] 
**CreatedAt** | **time.Time** | Team creation timestamp | 

## Methods

### NewOrganizationDataAttributes

`func NewOrganizationDataAttributes(name string, logo string, verified string, desc string, sort string, country string, createdAt time.Time, ) *OrganizationDataAttributes`

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


### GetVerified

`func (o *OrganizationDataAttributes) GetVerified() string`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *OrganizationDataAttributes) GetVerifiedOk() (*string, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *OrganizationDataAttributes) SetVerified(v string)`

SetVerified sets Verified field to given value.


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


### GetCountry

`func (o *OrganizationDataAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationDataAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationDataAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *OrganizationDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetLinks

`func (o *OrganizationDataAttributes) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganizationDataAttributes) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganizationDataAttributes) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganizationDataAttributes) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetComplicatedStatus

`func (o *OrganizationDataAttributes) GetComplicatedStatus() ComplicatedStatus`

GetComplicatedStatus returns the ComplicatedStatus field if non-nil, zero value otherwise.

### GetComplicatedStatusOk

`func (o *OrganizationDataAttributes) GetComplicatedStatusOk() (*ComplicatedStatus, bool)`

GetComplicatedStatusOk returns a tuple with the ComplicatedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplicatedStatus

`func (o *OrganizationDataAttributes) SetComplicatedStatus(v ComplicatedStatus)`

SetComplicatedStatus sets ComplicatedStatus field to given value.

### HasComplicatedStatus

`func (o *OrganizationDataAttributes) HasComplicatedStatus() bool`

HasComplicatedStatus returns a boolean if a field has been set.

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


