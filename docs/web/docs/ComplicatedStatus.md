# ComplicatedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | Current compliance state of the organization | 
**Stamp** | **[]string** |  | 
**From** | **time.Time** | Date from which the compliance state is effective | 

## Methods

### NewComplicatedStatus

`func NewComplicatedStatus(state string, stamp []string, from time.Time, ) *ComplicatedStatus`

NewComplicatedStatus instantiates a new ComplicatedStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplicatedStatusWithDefaults

`func NewComplicatedStatusWithDefaults() *ComplicatedStatus`

NewComplicatedStatusWithDefaults instantiates a new ComplicatedStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ComplicatedStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ComplicatedStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ComplicatedStatus) SetState(v string)`

SetState sets State field to given value.


### GetStamp

`func (o *ComplicatedStatus) GetStamp() []string`

GetStamp returns the Stamp field if non-nil, zero value otherwise.

### GetStampOk

`func (o *ComplicatedStatus) GetStampOk() (*[]string, bool)`

GetStampOk returns a tuple with the Stamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamp

`func (o *ComplicatedStatus) SetStamp(v []string)`

SetStamp sets Stamp field to given value.


### GetFrom

`func (o *ComplicatedStatus) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ComplicatedStatus) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ComplicatedStatus) SetFrom(v time.Time)`

SetFrom sets From field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


