package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Staff struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StaffID   string             `bson:"staff_id,omitempty" json:"staff_id,omitempty"`
	EstateID  string             `bson:"estate_id,omitempty" json:"estate_id,omitempty"`
	UserID    string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func (es Staff) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.EstateID, validation.Required.Error("estate_id is required")),
		validation.Field(&es.UserID, validation.Required.Error("user_id is required")),
	)
}
