package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FacultyService struct {
	Id               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EstateID         string             `bson:"estate_id,omitempty" json:"estate_id"`
	FacultyServiceID string             `bson:"faculty_service_id,omitempty" json:"faculty_service_id"`
	Name             string             `bson:"name,omitempty" json:"name"`
	Amount           float64            `bson:"amount,omitempty" json:"amount"`
	Duration         string             `bson:"duration,omitempty" json:"duration,omitempty"`
	Currency         string             `bson:"currency,omitempty" json:"currency,omitempty"`
	Status           string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

func (es FacultyService) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.Amount, validation.Required.Error("amount is required")),
		validation.Field(&es.Currency, validation.Required.Error("currency is required")),
		validation.Field(&es.Name, validation.Required.Error("name is required")),
	)
}
