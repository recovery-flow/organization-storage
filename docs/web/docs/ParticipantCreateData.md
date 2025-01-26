# ParticipantCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**ParticipantCreateDataAttributes**](ParticipantCreateDataAttributes.md) |  | 

## Methods

### NewParticipantCreateData

`func NewParticipantCreateData(type_ string, attributes ParticipantCreateDataAttributes, ) *ParticipantCreateData`

NewParticipantCreateData instantiates a new ParticipantCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantCreateDataWithDefaults

`func NewParticipantCreateDataWithDefaults() *ParticipantCreateData`

NewParticipantCreateDataWithDefaults instantiates a new ParticipantCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParticipantCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParticipantCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParticipantCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ParticipantCreateData) GetAttributes() ParticipantCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParticipantCreateData) GetAttributesOk() (*ParticipantCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParticipantCreateData) SetAttributes(v ParticipantCreateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


