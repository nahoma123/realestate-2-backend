package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* mongodb collection*/
type Guest struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GuestID     string             `bson:"guest_id,omitempty" json:"guest_id,omitempty"`
	HouseID     string             `bson:"house_id,omitempty" json:"house_id,omitempty"`
	FirstName   string             `bson:"first_name,omitempty" json:"first_name,omitempty"`
	MiddleName  string             `bson:"middle_name,omitempty" json:"middle_name,omitempty"`
	LastName    string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Type        string             `bson:"type,omitempty" json:"type,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Address     string             `bson:"address,omitempty" json:"address,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func (es Guest) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.FirstName, validation.Required.Error("first_name is required")),
		validation.Field(&es.MiddleName, validation.Required.Error("middle_name is required")),
		validation.Field(&es.HouseID, validation.Required.Error("house_id is required")),
		validation.Field(&es.LastName, validation.Required.Error("last_name is required")),
		validation.Field(&es.Type, validation.Required.Error("type is required")),
		validation.Field(&es.PhoneNumber, validation.Required.Error("phone_number is required")),
		validation.Field(&es.Email, validation.Required.Error("email is required")),
	)
}
