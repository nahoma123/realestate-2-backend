package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resident struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ResidentID string             `bson:"resident_id,omitempty" json:"resident_id,omitempty"`
	HouseID    string             `bson:"house_id,omitempty" json:"house_id,omitempty"`
	UserID     string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Estate     *Estate            `bson:"estate,omitempty" json:"estate,omitempty"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

func (es Resident) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.HouseID, validation.Required.Error("estate_id is required")),
		validation.Field(&es.UserID, validation.Required.Error("user_id is required")),
	)
}
