package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type MaintenanceRequest struct {
	Id                   uint   `gorm:"primaryKey" json:"id,omitempty"`
	MaintenanceRequestId string `gorm:"column:maintenance_request_id" json:"maintenance_request_id,omitempty"`
	PropertyId           string `gorm:"column:property_id" json:"property_id,omitempty"`
	RequestedBy          string `gorm:"type:text;column:requested_by" json:"requested_by,omitempty"`
	ProblemDetails       string `gorm:"type:text;column:problem_details" json:"problem_details,omitempty"`
	ProblemDescription   string `gorm:"type:text;column:problem_description" json:"problem_description,omitempty"`
	AlarmInformation     string `gorm:"column:alarm_information" json:"alarm_information,omitempty"`
	PetInformation       string `gorm:"column:pet_information" json:"pet_information,omitempty"`
	FurtherNotes         string `gorm:"column:further_notes" json:"further_notes,omitempty"`
	VulnerableOccupier   bool   `gorm:"column:vulnerable_occupier" json:"vulnerable_occupier,omitempty"`
	AgreeTerms           bool   `gorm:"column:agree_terms" json:"agree_terms,omitempty"`
	AgreePrivacyNotice   bool   `gorm:"column:agree_privacy_notice" json:"agree_privacy_notice,omitempty"`

	ImageUrl string `gorm:"column:image_url" json:"image_url"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}

func (p MaintenanceRequest) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.PropertyId, validation.Required),
	)
}

func (p MaintenanceRequest) ValidateUpdate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.MaintenanceRequestId, validation.Required),
	)
}
