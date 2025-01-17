# LinksData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | organization id | 
**Type** | **string** |  | 
**Attributes** | [**LinksDataAttributes**](LinksDataAttributes.md) |  | 

## Methods

### NewLinksData

`func NewLinksData(id string, type_ string, attributes LinksDataAttributes, ) *LinksData`

NewLinksData instantiates a new LinksData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksDataWithDefaults

`func NewLinksDataWithDefaults() *LinksData`

NewLinksDataWithDefaults instantiates a new LinksData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinksData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinksData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinksData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *LinksData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinksData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinksData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *LinksData) GetAttributes() LinksDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LinksData) GetAttributesOk() (*LinksDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LinksData) SetAttributes(v LinksDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


