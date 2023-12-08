package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstateConfiguration struct {
	Id                    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EstateConfigurationID string             `bson:"estate_configuration_id,omitempty" json:"estate_configuration_id,omitempty"`
	EstateID              string             `bson:"estate_id,omitempty" json:"estate_id,omitempty"`
	MaxOwnerDept          float64            `bson:"max_owner_dept,omitempty" json:"max_owner_dept,omitempty"`

	MaxDefaultedInvites        int       `bson:"max_defaulted_invites,omitempty" json:"max_defaulted_invites,omitempty"`
	IsInviteBanEnabled         bool      `bson:"is_invite_ban_enabled,omitempty" json:"is_invite_ban_enabled,omitempty"`
	InviteDefaultPenaltyAmount float64   `bson:"invite_default_penalty_amount,omitempty" json:"invite_default_penalty_amount,omitempty"`
	BanMonths                  uint32    `bson:"ban_months,omitempty" json:"ban_months,omitempty"`
	BanDays                    uint32    `bson:"ban_days,omitempty" json:"ban_days,omitempty"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

func (es EstateConfiguration) Validate() error {
	return validation.ValidateStruct(&es,
		validation.Field(&es.MaxOwnerDept, validation.Required.Error("max owner dept is required")),
		validation.Field(&es.InviteDefaultPenaltyAmount, validation.When(es.IsInviteBanEnabled, validation.Required.Error("please provide penalty amount"))),
		validation.Field(&es.BanDays, validation.When(es.IsInviteBanEnabled, validation.Required.Error("please provide ban days and months amount"))),
		validation.Field(&es.BanMonths, validation.When(es.IsInviteBanEnabled, validation.Required.Error("please provide ban days and months amount"))),
	)
}

func (es EstateConfiguration) ValidateUpdate() error {
	return validation.ValidateStruct(&es)
}
