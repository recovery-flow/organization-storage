# ParticipantCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Participant**](Participant.md) |  | 
**Links** | [**ParticipantsCollectionLinks**](ParticipantsCollectionLinks.md) |  | 
**Meta** | [**ParticipantsCollectionMeta**](ParticipantsCollectionMeta.md) |  | 

## Methods

### NewParticipantCollection

`func NewParticipantCollection(data []Participant, links ParticipantsCollectionLinks, meta ParticipantsCollectionMeta, ) *ParticipantCollection`

NewParticipantCollection instantiates a new ParticipantCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantCollectionWithDefaults

`func NewParticipantCollectionWithDefaults() *ParticipantCollection`

NewParticipantCollectionWithDefaults instantiates a new ParticipantCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ParticipantCollection) GetData() []Participant`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ParticipantCollection) GetDataOk() (*[]Participant, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ParticipantCollection) SetData(v []Participant)`

SetData sets Data field to given value.


### GetLinks

`func (o *ParticipantCollection) GetLinks() ParticipantsCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ParticipantCollection) GetLinksOk() (*ParticipantsCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ParticipantCollection) SetLinks(v ParticipantsCollectionLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *ParticipantCollection) GetMeta() ParticipantsCollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ParticipantCollection) GetMetaOk() (*ParticipantsCollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ParticipantCollection) SetMeta(v ParticipantsCollectionMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


