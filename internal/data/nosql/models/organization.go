package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Logo      string             `bson:"logo" json:"logo"`
	Verified  bool               `bson:"verified" json:"verified"`
	Desc      string             `bson:"desc" json:"desc"`
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

type Links struct {
	Twitter   *string `json:"twitter"`
	Instagram *string `json:"instagram"`
	Facebook  *string `json:"facebook"`
	TikTok    *string `json:"tiktok"`
	Linkedin  *string `json:"linkedin"`
	Telegram  *string `json:"telegram"`
	Discord   *string `json:"discord"`
}
