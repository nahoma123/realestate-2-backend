package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PropertyRentDetails struct {
	Id         uint   `gorm:"primaryKey" json:"id"`
	PropertyId string `gorm:"column:property_id" json:"property_id"`

	LandlordID string `gorm:"type:text;column:landlord_id" json:"landlord_id"`
	TenantID   string `gorm:"type:text;column:tenant_id" json:"tenant_id"`

	Landlord *User `gorm:"foreignKey:LandlordID;references:UserID" json:"landlord"`
	Tenant   *User `gorm:"foreignKey:TenantID;references:UserID" json:"tenant"`

	CurrentRentAmount     float64   `gorm:"rent_amount" json:"rent_amount"`
	RentLastPaid          time.Time `gorm:"rent_last_paid" json:"rent_last_paid"`
	CurrentRentLeasedDate time.Time `gorm:"current_rent_leased_date,omitempty" json:"current_rent_leased_date"`

	RentStatus string `gorm:"-" json:"rent_status"`
}

type Property struct {
	Id         uint   `gorm:"primaryKey" json:"id,omitempty"`
	PropertyId string `gorm:"column:property_id" json:"property_id,omitempty"`
	Status     string `gorm:"status,omitempty" json:"status,omitempty"`

	LandlordID string `gorm:"type:text;column:landlord_id" json:"landlord_id,omitempty"`
	TenantID   string `gorm:"type:text;column:tenant_id" json:"tenant_id,omitempty"`

	Landlord *User `gorm:"foreignKey:LandlordID;references:UserID" json:"landlord,omitempty"`
	Tenant   *User `gorm:"foreignKey:TenantID;references:UserID" json:"tenant,omitempty"`

	Amount     float64 `gorm:"amount,omitempty" json:"amount,omitempty"`
	Address    string  `gorm:"address,omitempty" json:"address,omitempty"`
	Latitude   float64 `gorm:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude  float64 `gorm:"longitude,omitempty" json:"longitude,omitempty"`
	PostalCode string  `gorm:"postal_code,omitempty" json:"postal_code,omitempty"`

	PropertyType      string      `gorm:"property_type,omitempty" json:"property_type,omitempty"`
	Images            StringArray `gorm:"type:VARCHAR(500)" json:"images,omitempty"`
	ReceptionNumber   int         `gorm:"reception_number,omitempty" json:"reception_number,omitempty"`
	BedNumber         int         `gorm:"bed_number,omitempty" json:"bed_number,omitempty"`
	BathNumber        int         `gorm:"bath_number,omitempty" json:"bath_number,omitempty"`
	PropertyDetails   string      `gorm:"property_details,omitempty" json:"property_details,omitempty"`
	EPC               string      `gorm:"epc,omitempty" json:"epc,omitempty"`
	IsStudentProperty bool        `gorm:"is_student_property,omitempty" json:"is_student_property,omitempty"`
	Features          StringArray `gorm:"type:VARCHAR(500)" json:"features,omitempty"`
	Furnished         string      `gorm:"furnished,omitempty" json:"furnished,omitempty"`

	Inspected          bool      `gorm:"inspected,omitempty" json:"inspected,omitempty"`
	NextInspectionDate time.Time `gorm:"next_inspection_date,omitempty" json:"next_inspection_date,omitempty"`

	CurrentRentAmount     float64   `gorm:"current_rent_amount,omitempty" json:"current_rent_amount,omitempty"`
	RentLastPaid          time.Time `gorm:"rent_last_paid,omitempty" json:"rent_last_paid"`
	CurrentRentLeasedDate time.Time `gorm:"current_rent_leased_date,omitempty" json:"current_rent_leased_date"`

	CreatedAt time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}

func (p *PropertyRentDetails) CalculateRentStatus() string {
	// Get the current time
	now := time.Now()

	// Calculate the number of months since the lease date
	monthsSinceLease := now.Year()*12 + int(now.Month()) - (p.CurrentRentLeasedDate.Year()*12 + int(p.CurrentRentLeasedDate.Month()))

	// Calculate the number of months since the last rent payment
	monthsSinceLastPayment := now.Year()*12 + int(now.Month()) - (p.RentLastPaid.Year()*12 + int(p.RentLastPaid.Month()))

	// If the number of months since the last payment is less than the number of months since the lease date, rent is paid
	if monthsSinceLastPayment <= monthsSinceLease {
		return "Paid"
	}

	// Otherwise, rent is unpaid
	return "Unpaid"
}

type StringArray []string

func (o *StringArray) Scan(src interface{}) error {
	fmt.Println("Type of variable1:", reflect.TypeOf(src))

	switch src := src.(type) {
	case []byte:
		*o = strings.Split(string(src), ",")
	case string:
		*o = strings.Split(src, ",")
	default:
		return errors.New("src value is not a string or []byte")
	}

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

func (p Property) ValidatePropertyRent() error {
	return validation.ValidateStruct(&p,
		// validation.Field(&p.CurrentRentAmount, validation.Required),
		validation.Field(&p.LandlordID, validation.Required),
		validation.Field(&p.TenantID, validation.Required),
		validation.Field(&p.RentLastPaid, validation.Required),
	)
}
