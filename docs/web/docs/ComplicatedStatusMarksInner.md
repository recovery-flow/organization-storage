# ComplicatedStatusMarksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | Timestamp of the mark | 
**Reason** | **string** | Reason for the compliance mark | 

## Methods

### NewComplicatedStatusMarksInner

`func NewComplicatedStatusMarksInner(timestamp time.Time, reason string, ) *ComplicatedStatusMarksInner`

NewComplicatedStatusMarksInner instantiates a new ComplicatedStatusMarksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplicatedStatusMarksInnerWithDefaults

`func NewComplicatedStatusMarksInnerWithDefaults() *ComplicatedStatusMarksInner`

NewComplicatedStatusMarksInnerWithDefaults instantiates a new ComplicatedStatusMarksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ComplicatedStatusMarksInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ComplicatedStatusMarksInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ComplicatedStatusMarksInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetReason

`func (o *ComplicatedStatusMarksInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ComplicatedStatusMarksInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ComplicatedStatusMarksInner) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


