package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HouseOwner struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HouseOwnerID       string             `bson:"house_owner_id,omitempty" json:"house_owner_id,omitempty"`
	EstateID           string             `bson:"estate_id,omitempty" json:"estate_id,omitempty"`
	HouseOwnerCategory string             `bson:"house_owner_category,omitempty" json:"house_owner_category,omitempty"`
	LivingInTheEstate  bool               `bson:"living_in_the_estate,omitempty" json:"living_in_the_estate,omitempty"`
	UserID             string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Status             string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

func (es HouseOwner) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.EstateID, validation.Required.Error("estate_id is required")),
		validation.Field(&es.UserID, validation.Required.Error("user_id is required")),
		validation.Field(&es.HouseOwnerCategory, validation.Required.Error("house_owner_category is required")),
	)
}

type House struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HouseID        string             `bson:"house_id,omitempty" json:"house_id,omitempty"`
	EstateID       string             `bson:"estate_id,omitempty" json:"estate_id,omitempty"`
	HouseOwnerID   string             `bson:"house_owner_id,omitempty" json:"house_owner_id,omitempty"`
	ApartmentID    string             `bson:"apartment_number,omitempty" json:"apartment_number,omitempty"`
	FloorNumber    string             `bson:"floor_number,omitempty" json:"floor_number,omitempty"`
	BuildingNumber string             `bson:"building_number,omitempty" json:"building_number,omitempty"`
	Status         string             `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

func (es House) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.EstateID, validation.Required.Error("estate_id is required")),
		validation.Field(&es.HouseOwnerID, validation.Required.Error("house_owner_id is required")),
		validation.Field(&es.ApartmentID, validation.Required.Error("apartment_number is required")),
		validation.Field(&es.FloorNumber, validation.Required.Error("floor_number is required")),
		validation.Field(&es.BuildingNumber, validation.Required.Error("building_number is required")),
	)
}
