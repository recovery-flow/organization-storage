# ParticipantsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ParticipantData**](ParticipantData.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewParticipantsCollection

`func NewParticipantsCollection(data []ParticipantData, links Links, ) *ParticipantsCollection`

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

`func (o *ParticipantsCollection) GetData() []ParticipantData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ParticipantsCollection) GetDataOk() (*[]ParticipantData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ParticipantsCollection) SetData(v []ParticipantData)`

SetData sets Data field to given value.


### GetLinks

`func (o *ParticipantsCollection) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ParticipantsCollection) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ParticipantsCollection) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


