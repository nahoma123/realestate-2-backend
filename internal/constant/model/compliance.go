package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Compliance struct {
	Id           uint   `gorm:"primaryKey" json:"id,omitempty"`
	ComplianceId string `json:"compliance_id"`
	PropertyId   string `json:"property_id"`

	GasSafetyDueDate   time.Time `json:"gas_safety_due_date"`
	ElectricityDueDate time.Time `json:"electricity_due_date"`

	GasImages         StringArray `gorm:"type:VARCHAR(500)" json:"gas_images"`
	ElectricityImages StringArray `gorm:"type:VARCHAR(500)" json:"electricity_images"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}

func (p Compliance) Validate() error {

	return validation.ValidateStruct(&p)
}

func (p Compliance) ValidateUpdate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ComplianceId, validation.Required))
}
