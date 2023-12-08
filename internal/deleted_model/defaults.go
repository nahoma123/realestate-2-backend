package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Default struct {
	Id                       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	DefaultID                string             `bson:"default_id,omitempty" json:"default_id,omitempty"`
	IsBanActive              bool               `bson:"is_ban_active,omitempty" json:"is_ban_active,omitempty"`
	IsPaymentOptionAvailable bool               `bson:"is_payment_option_available,omitempty" json:"is_payment_option_available,omitempty"`
	PaymentAmount            float64            `bson:"payment_amount,omitempty" json:"payment_amount,omitempty"`
	BanStartTime             time.Time          `bson:"ban_start_time,omitempty" json:"ban_start_time,omitempty"`
	BanEndTime               time.Time          `bson:"ban_end_time,omitempty" json:"ban_end_time,omitempty"`
	HouseOwnerDefault        *HouseOwnerDefault `bson:"house_owner_default,omitempty" json:"house_owner_default,omitempty"`
	ResidentDefault          *ResidentDefault   `bson:"resident_default,omitempty" json:"resident_default,omitempty"`
	CreatedAt                time.Time          `json:"created_at"`
	UpdatedAt                time.Time          `json:"updated_at"`
}

type HouseOwnerDefault struct {
	HouseOwnerID string `bson:"house_owner_id,omitempty" json:"house_owner_id,omitempty"`
	Reason       string `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ResidentDefault struct {
	ResidentID string `bson:"resident_id,omitempty" json:"resident_id,omitempty"`
	Reason     string `bson:"reason,omitempty" json:"reason,omitempty"`
}

func (es Default) Validate() error {
	return validation.ValidateStruct(&es) // validation.Field(&es.MaxOwnerDept, validation.Required.Error("max owner dept is required")),
	// validation.Field(&es.InviteDefaultPenaltyAmount, validation.When(es.IsInviteBanEnabled, validation.Required.Error("please provide penalty amount"))),
	// validation.Field(&es.BanDays, validation.When(es.IsInviteBanEnabled, validation.Required.Error("please provide ban days and months amount"))),
	// validation.Field(&es.BanMonths, validation.When(es.IsInviteBanEnabled, validation.Required.Error("please provide ban days and months amount"))),

}

func (es Default) ValidateUpdate() error {
	return validation.ValidateStruct(&es)
}
