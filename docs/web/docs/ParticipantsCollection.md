# ParticipantsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Participant**](Participant.md) |  | 
**Links** | [**ParticipantsCollectionLinks**](ParticipantsCollectionLinks.md) |  | 
**Meta** | [**ParticipantsCollectionMeta**](ParticipantsCollectionMeta.md) |  | 

## Methods

### NewParticipantsCollection

`func NewParticipantsCollection(data []Participant, links ParticipantsCollectionLinks, meta ParticipantsCollectionMeta, ) *ParticipantsCollection`

NewParticipantsCollection instantiates a new ParticipantsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantsCollectionWithDefaults

`func NewParticipantsCollectionWithDefaults() *ParticipantsCollection`

NewParticipantsCollectionWithDefaults instantiates a new ParticipantsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ParticipantsCollection) GetData() []Participant`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ParticipantsCollection) GetDataOk() (*[]Participant, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ParticipantsCollection) SetData(v []Participant)`

SetData sets Data field to given value.


### GetLinks

`func (o *ParticipantsCollection) GetLinks() ParticipantsCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ParticipantsCollection) GetLinksOk() (*ParticipantsCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ParticipantsCollection) SetLinks(v ParticipantsCollectionLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *ParticipantsCollection) GetMeta() ParticipantsCollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ParticipantsCollection) GetMetaOk() (*ParticipantsCollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ParticipantsCollection) SetMeta(v ParticipantsCollectionMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


