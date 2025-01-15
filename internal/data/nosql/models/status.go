package models

import "time"

type Status struct {
	State OrgState  `bson:"state" json:"state"`
	Stamp []Stamp   `bson:"marks" json:"marks"`
	From  time.Time `bson:"from" json:"from"`
}

type OrgState string

const (
	OpenCaseState                    OrgState = "open_case"
	UnderUserInvestigationState      OrgState = "under_user_investigation"
	UnderInternalInvestigationState  OrgState = "under_internal_investigation"
	UnderMultipleInvestigationsState OrgState = "under_multiple_investigations"
	PermanentState                   OrgState = "permanent_ban"
)

type Stamp string

const (
	CorruptionStamp         Stamp = "corruption"
	DocumentProblemsStamp   Stamp = "document_problems"
	QuestionableOriginStamp Stamp = "questionable_origin"
	MoneyLaunderingStamp    Stamp = "money_laundering"
)
