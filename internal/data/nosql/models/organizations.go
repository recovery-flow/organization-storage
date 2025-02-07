package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Logo         string             `bson:"logo,omitempty" json:"logo,omitempty"`
	Verified     bool               `bson:"verified" json:"verified"`
	Desc         string             `bson:"desc" json:"desc"`
	Country      string             `bson:"country" json:"country"`
	City         *string            `bson:"city,omitempty" json:"city,omitempty"`
	Sort         SortOfOrg          `bson:"sort" json:"sort"`
	Participants []Participant      `bson:"participants" json:"participants"`
	Status       *Status            `bson:"compliance_status,omitempty" json:"compliance_status,omitempty"`
	Links        *Links             `bson:"links,omitempty" json:"links,omitempty"`

	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

type SortOfOrg string

const (
	FoundationSort  SortOfOrg = "foundation"
	CompanySort     SortOfOrg = "company"
	CorporationSort SortOfOrg = "corporation"
)

var ErrInvalidSort = fmt.Errorf("invalid sort of organization")

func StringToSortOfOrg(s string) (SortOfOrg, error) {
	switch s {
	case "foundation":
		return FoundationSort, nil
	case "company":
		return CompanySort, nil
	case "corporation":
		return CorporationSort, nil
	default:
		return FoundationSort, ErrInvalidSort
	}
}
