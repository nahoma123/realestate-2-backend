package model

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Contract struct {
	Id         uint   `gorm:"primaryKey" json:"id,omitempty"`
	ContractId string `json:"contract_id"`
	PropertyId string `json:"property_id"`

	DocumentType string    `json:"document_type"`
	DocumentName string    `json:"document_name"`
	UploadedDate time.Time `json:"uploaded_date"`

	ExpiryDate  time.Time `json:"expiry_date"`
	DocumentUrl string    `json:"document_url"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}

func (p Contract) Validate() error {
	validDocumentTypes := []string{"agreement", "lease", "nda", "service agreement", "purchase agreement"} // Replace with your array of valid options

	return validation.ValidateStruct(&p,
		validation.Field(&p.DocumentUrl, validation.Required),
		validation.Field(&p.ExpiryDate, validation.Required),
		validation.Field(&p.DocumentName, validation.Required),
		validation.Field(&p.PropertyId, validation.Required),
		validation.Field(&p.DocumentType,
			validation.Required,
			validation.In(validDocumentTypes[0], validDocumentTypes[1], validDocumentTypes[2],
				validDocumentTypes[3]).Error(fmt.Sprintf("document_type must be either %s, %s, %s, or %s",
				validDocumentTypes[0], validDocumentTypes[1], validDocumentTypes[2], validDocumentTypes[3]))))
}

func (p Contract) ValidateUpdate() error {
	validDocumentTypes := []string{"agreement", "lease", "nda", "service agreement", "purchase agreement"} // Replace with your array of valid options

	return validation.ValidateStruct(&p,
		validation.Field(&p.ContractId, validation.Required),
		validation.Field(&p.DocumentType,
			validation.In(validDocumentTypes[0], validDocumentTypes[1], validDocumentTypes[2],
				validDocumentTypes[3]).Error(fmt.Sprintf("document_type must be either %s, %s, %s, or %s",
				validDocumentTypes[0], validDocumentTypes[1], validDocumentTypes[2], validDocumentTypes[3]))))
}
