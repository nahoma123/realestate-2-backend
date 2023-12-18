package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type InspectionResult struct {
	Id                 uint   `gorm:"primaryKey" json:"id,omitempty"`
	InspectionResultId string `gorm:"column:property_id" json:"property_id,omitempty"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}

func (p InspectionResult) Validate() error {
	return validation.ValidateStruct(&p)
}

func (p InspectionResult) ValidateUpdate() error {
	return validation.ValidateStruct(&p)
}
