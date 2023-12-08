package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Estate struct {
	Id                     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EstateID               string             `bson:"estate_id,omitempty" json:"estate_id"`
	Name                   string             `bson:"name,omitempty" json:"name"`
	Location               string             `bson:"location,omitempty" json:"location"`
	ContactName            string             `bson:"contact_name,omitempty" json:"contact_name"`
	PhoneNumber            string             `bson:"phone_number,omitempty" json:"phone_number"`
	AlternativePhoneNumber string             `bson:"alternative_phone_number,omitempty" json:"alternative_phone_number"`
	Status                 string             `bson:"status,omitempty" json:"status,omitempty"`
	UserID                 string             `bson:"user_id,omitempty" json:"user_id"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`
}

func (es Estate) Validate() error {
	// return validation.ValidateStruct(&u,
	// 	validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
	// 	validation.Field(&u.MiddleName, validation.Required.Error("middle name is required")),
	// 	validation.Field(&u.LastName, validation.Required.Error("last name is required")),
	// 	validation.Field(&u.Phone, validation.Required.Error("phone is required")),
	// 	validation.Field(&u.Email, is.EmailFormat.Error("email is not valid")),
	// 	// validation.Field(&u.Phone, validation.Required.Error("phone is required"), validation.By(validatePhone)),
	// 	validation.Field(&u.Password, validation.When(u.Email != "", validation.Required.Error("password is required"), validation.Length(6, 32).Error("password must be between 6 and 32 characters"))),
	// )
	return nil
}
