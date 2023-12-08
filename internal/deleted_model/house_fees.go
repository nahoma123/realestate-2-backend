package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HouseFee struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HouseFeeID string             `bson:"house_fee_id,omitempty" json:"house_fee_id,omitempty"`
	HouseID    string             `bson:"house_id,omitempty" json:"house_id,omitempty"`
	Name       string             `bson:"name,omitempty" json:"name"`
	Amount     float64            `bson:"amount,omitempty" json:"amount"`
	Duration   string             `bson:"duration,omitempty" json:"duration,omitempty"`
	Currency   string             `bson:"currency,omitempty" json:"currency,omitempty"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

func (es HouseFee) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.Amount, validation.Required.Error("amount is required")),
		validation.Field(&es.Currency, validation.Required.Error("currency is required")),
		validation.Field(&es.Name, validation.Required.Error("name is required")),
		validation.Field(&es.HouseID, validation.Required.Error("house_id is required")),
	)
}
