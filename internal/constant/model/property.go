package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
)

type Property struct {
	Id         uint   `gorm:"primaryKey" json:"id,omitempty"`
	PropertyId string `gorm:"column:property_id" json:"property_id,omitempty"`
	Status     string `gorm:"status,omitempty" json:"status,omitempty"`

	UserID uuid.UUID `gorm:"column:user_id" json:"user_id,omitempty"`
	User   *User     `gorm:"-" json:"user,omitempty"`

	Amount     float64 `gorm:"amount,omitempty" json:"amount,omitempty"`
	Address    string  `gorm:"address,omitempty" json:"address,omitempty"`
	Latitude   float64 `gorm:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude  float64 `gorm:"longitude,omitempty" json:"longitude,omitempty"`
	PostalCode string  `gorm:"postal_code,omitempty" json:"postal_code,omitempty"`

	PropertyType      string      `gorm:"property_type,omitempty" json:"property_type,omitempty"`
	Images            StringArray `gorm:"type:VARCHAR(255)"" json:"images,omitempty"`
	ReceptionNumber   int         `gorm:"reception_number,omitempty" json:"reception_number,omitempty"`
	BedNumber         int         `gorm:"bed_number,omitempty" json:"bed_number,omitempty"`
	BathNumber        int         `gorm:"bath_number,omitempty" json:"bath_number,omitempty"`
	PropertyDetails   string      `gorm:"property_details,omitempty" json:"property_details,omitempty"`
	EPC               string      `gorm:"epc,omitempty" json:"epc,omitempty"`
	IsStudentProperty bool        `gorm:"is_student_property,omitempty" json:"is_student_property,omitempty"`
	Features          StringArray `gorm:"type:VARCHAR(255)"" json:"features,omitempty"`
	Furnished         string      `gorm:"furnished,omitempty" json:"furnished,omitempty"`

	NextInspectionDate time.Time `gorm:"next_inspection_date,omitempty" json:"next_inspection_date,omitempty"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}
type StringArray []string

func (o *StringArray) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*o = strings.Split(string(bytes), ",")
	return nil
}
func (o StringArray) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}
	return strings.Join(o, ","), nil
}

func (p Property) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Amount, validation.Required),
		validation.Field(&p.Address, validation.Required),
		validation.Field(&p.Latitude, validation.Required),
		validation.Field(&p.Longitude, validation.Required),
		validation.Field(&p.PropertyType, validation.Required),
		validation.Field(&p.Images, validation.Required),
		validation.Field(&p.ReceptionNumber, validation.Required),
		validation.Field(&p.BedNumber, validation.Required),
		validation.Field(&p.BathNumber, validation.Required),
		validation.Field(&p.PropertyDetails, validation.Required),
		validation.Field(&p.EPC, validation.Required),
		validation.Field(&p.Features, validation.Required),
		validation.Field(&p.Furnished, validation.Required),
	)
}

func (p Property) ValidateUpdate() error {
	return validation.ValidateStruct(&p)
}
