# ComplicatedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | Current compliance state of the organization | 
**Marks** | [**[]ComplicatedStatusMarksInner**](ComplicatedStatusMarksInner.md) |  | 
**From** | **time.Time** | Date from which the compliance state is effective | 

## Methods

### NewComplicatedStatus

`func NewComplicatedStatus(state string, marks []ComplicatedStatusMarksInner, from time.Time, ) *ComplicatedStatus`

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


### GetMarks

`func (o *ComplicatedStatus) GetMarks() []ComplicatedStatusMarksInner`

GetMarks returns the Marks field if non-nil, zero value otherwise.

### GetMarksOk

`func (o *ComplicatedStatus) GetMarksOk() (*[]ComplicatedStatusMarksInner, bool)`

GetMarksOk returns a tuple with the Marks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarks

`func (o *ComplicatedStatus) SetMarks(v []ComplicatedStatusMarksInner)`

SetMarks sets Marks field to given value.


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


