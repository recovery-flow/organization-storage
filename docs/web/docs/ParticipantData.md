# ParticipantData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | user id | 
**Type** | **string** |  | 
**Attributes** | [**ParticipantDataAttributes**](ParticipantDataAttributes.md) |  | 
**Relationships** | [**ParticipantDataRelationships**](ParticipantDataRelationships.md) |  | 

## Methods

### NewParticipantData

`func NewParticipantData(id string, type_ string, attributes ParticipantDataAttributes, relationships ParticipantDataRelationships, ) *ParticipantData`

NewParticipantData instantiates a new ParticipantData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantDataWithDefaults

`func NewParticipantDataWithDefaults() *ParticipantData`

NewParticipantDataWithDefaults instantiates a new ParticipantData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParticipantData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParticipantData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParticipantData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ParticipantData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParticipantData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParticipantData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ParticipantData) GetAttributes() ParticipantDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParticipantData) GetAttributesOk() (*ParticipantDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParticipantData) SetAttributes(v ParticipantDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ParticipantData) GetRelationships() ParticipantDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ParticipantData) GetRelationshipsOk() (*ParticipantDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ParticipantData) SetRelationships(v ParticipantDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


