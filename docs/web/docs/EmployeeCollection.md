# EmployeeCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Object**](Object.md) |  | 
**Links** | [**EmployeeCollectionLinks**](EmployeeCollectionLinks.md) |  | 
**Meta** | [**EmployeeCollectionMeta**](EmployeeCollectionMeta.md) |  | 

## Methods

### NewEmployeeCollection

`func NewEmployeeCollection(data []Object, links EmployeeCollectionLinks, meta EmployeeCollectionMeta, ) *EmployeeCollection`

NewEmployeeCollection instantiates a new EmployeeCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeCollectionWithDefaults

`func NewEmployeeCollectionWithDefaults() *EmployeeCollection`

NewEmployeeCollectionWithDefaults instantiates a new EmployeeCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EmployeeCollection) GetData() []Object`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmployeeCollection) GetDataOk() (*[]Object, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmployeeCollection) SetData(v []Object)`

SetData sets Data field to given value.


### GetLinks

`func (o *EmployeeCollection) GetLinks() EmployeeCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EmployeeCollection) GetLinksOk() (*EmployeeCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EmployeeCollection) SetLinks(v EmployeeCollectionLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *EmployeeCollection) GetMeta() EmployeeCollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EmployeeCollection) GetMetaOk() (*EmployeeCollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EmployeeCollection) SetMeta(v EmployeeCollectionMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


