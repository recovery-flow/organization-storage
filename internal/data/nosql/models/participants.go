package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/roles"
)

type Participant struct {
	UserID      uuid.UUID     `bson:"user_id" json:"user_id"`
	FirstName   string        `bson:"first_name" json:"first_name"`
	SecondName  string        `bson:"second_name" json:"second_name"`
	ThirdName   *string       `bson:"third_name,omitempty" json:"third_name,omitempty"`
	DisplayName string        `bson:"display_name" json:"display_name"`
	Position    string        `bson:"position" json:"position"`
	Desc        string        `bson:"desc" json:"desc"`
	Verified    bool          `bson:"verified" json:"verified"`
	Role        roles.OrgRole `bson:"role" json:"role"`

	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
