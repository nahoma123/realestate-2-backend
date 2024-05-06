package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Message struct {
	Id         uint   `gorm:"primaryKey" json:"id,omitempty"`
	MessageId  string `json:"message_id"`
	TenantId   string `json:"tenant_id"`
	LandlordID string `gorm:"type:text;column:landlord_id" json:"landlord_id,omitempty"`
	PropertyId string `json:"property_id"`

	LandLord *User `gorm:"foreignKey:LandlordID;references:UserID" json:"landlord,omitempty"`
	Tenant   *User `gorm:"foreignKey:TenantId;references:UserID" json:"tenant,omitempty"`

	AdminApproved bool `json:"admin_approved"`

	IsReadAdmin    bool   `json:"is_read_admin"`
	IsReadLandLord bool   `json:"is_read_landlord"`
	Text           string `json:"text,omitempty"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}

func (p Message) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Text,
			validation.Required),
		validation.Field(&p.PropertyId,
			validation.Required),
	)
}

func (p Message) ValidateUpdate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.MessageId, validation.Required),
	)
}
