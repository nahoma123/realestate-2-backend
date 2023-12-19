package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type InspectionResult struct {
	Id                 uint   `gorm:"primaryKey" json:"id,omitempty"`
	InspectionResultId string `gorm:"column:inspection_result_id" json:"inspection_result_id,omitempty"`
	PropertyId         string `gorm:"column:property_id" json:"property_id,omitempty"`

	InspectionDate time.Time `gorm:"column:inspection_date" json:"inspection_date,omitempty"`
	Satisfactory   bool      `gorm:"column:satisfactory" json:"satisfactory,omitempty"`
	Comment        string    `gorm:"column:comment" json:"comment,omitempty"`
	LandlordView   bool      `gorm:"column:landlord_view" json:"landlord_view,omitempty"`
	TenantView     bool      `gorm:"column:tenant_view" json:"tenant_view,omitempty"`

	Image1        string `gorm:"column:image_1" json:"image_1"`
	Image1Comment string `gorm:"column:image_1_comment" json:"image_1_comment,omitempty"`

	Image2        string `gorm:"column:image_2" json:"image_2"`
	Image2Comment string `gorm:"column:image_2_comment" json:"image_2_comment,omitempty"`

	Image3        string `gorm:"column:image_3" json:"image_3"`
	Image3Comment string `gorm:"column:image_3_comment" json:"image_3_comment,omitempty"`

	Image4        string `gorm:"column:image_4" json:"image_4"`
	Image4Comment string `gorm:"column:image_4_comment" json:"image_4_comment,omitempty"`

	Image5        string `gorm:"column:image_5" json:"image_5"`
	Image5Comment string `gorm:"column:image_5_comment" json:"image_5_comment,omitempty"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}

func (p InspectionResult) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.PropertyId, validation.Required),
	)
}

func (p InspectionResult) ValidateUpdate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.InspectionResultId, validation.Required),
	)
}
