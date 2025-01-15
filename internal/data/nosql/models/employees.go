package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/roles"
)

type Employee struct {
	UserID uuid.UUID     `bson:"user_id" json:"user_id"`
	Name   string        `bson:"name" json:"name"`
	Desc   string        `bson:"desc" json:"desc"`
	Role   roles.OrgRole `bson:"role" json:"role"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
