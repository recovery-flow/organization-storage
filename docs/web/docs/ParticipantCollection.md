# ParticipantCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ParticipantData**](ParticipantData.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewParticipantCollection

`func NewParticipantCollection(data []ParticipantData, links Links, ) *ParticipantCollection`

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

`func (o *ParticipantCollection) GetData() []ParticipantData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ParticipantCollection) GetDataOk() (*[]ParticipantData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ParticipantCollection) SetData(v []ParticipantData)`

SetData sets Data field to given value.


### GetLinks

`func (o *ParticipantCollection) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ParticipantCollection) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ParticipantCollection) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


