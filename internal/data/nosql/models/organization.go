package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Desc      string             `bson:"desc" json:"desc"`
	Type      OrgType            `bson:"type" json:"type"`
	Employees []Employee         `bson:"employees" json:"employees"`
	Status    *Status            `bson:"complianceStatus,omitempty" json:"complianceStatus,omitempty"`
	Links     *Links             `bson:"links,omitempty" json:"links,omitempty"`

	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

type OrgType string

const (
	FoundationType OrgType = "foundation"
	CompanyType    OrgType = "company"
	Corporation    OrgType = "corporation"
)

type Links struct {
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	TikTok    string `json:"tiktok"`
	Linkedin  string `json:"linkedin"`
	Telegram  string `json:"telegram"`
	Discord   string `json:"discord"`
}
