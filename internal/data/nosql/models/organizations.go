package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Logo      string             `bson:"logo,omitempty" json:"logo,omitempty"`
	Verified  bool               `bson:"verified" json:"verified"`
	Desc      string             `bson:"desc" json:"desc"`
	Country   string             `bson:"country" json:"country"`
	City      *string            `bson:"city,omitempty" json:"city,omitempty"`
	Sort      SortOfOrg          `bson:"sort" json:"sort"`
	Employees []Employee         `bson:"employees" json:"employees"`
	Status    *Status            `bson:"complianceStatus,omitempty" json:"complianceStatus,omitempty"`
	Links     *Links             `bson:"links,omitempty" json:"links,omitempty"`

	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

type SortOfOrg string

const (
	FoundationSort  SortOfOrg = "foundation"
	CompanySort     SortOfOrg = "company"
	CorporationSort SortOfOrg = "corporation"
)
