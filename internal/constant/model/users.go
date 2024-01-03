package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id uint `gorm:"primaryKey" json:"id,omitempty"`

	Properties []Property `gorm:"foreignKey:LandlordID;references:UserID" json:"landlords,omitempty"`
	Rental     []Property `gorm:"foreignKey:TenantID;references:UserID" json:"rentals,omitempty"`

	UserID     string    `gorm:"unique;column:user_id" json:"user_id,omitempty"`
	FirstName  string    `gorm:"column:first_name" json:"first_name,omitempty"`
	MiddleName string    `gorm:"column:middle_name" json:"middle_name,omitempty"`
	LastName   string    `gorm:"column:last_name" json:"last_name,omitempty"`
	Email      string    `gorm:"column:email" json:"email,omitempty"`
	ResetCode  int       `gorm:"column:reset_code" json:"reset_code,omitempty"`
	Phone      string    `gorm:"column:phone" json:"phone,omitempty"`
	Password   string    `gorm:"column:password" json:"password,omitempty"`
	UserName   string    `gorm:"column:user_name" json:"user_name,omitempty"`
	Gender     string    `gorm:"column:gender" json:"gender,omitempty"`
	Status     string    `gorm:"column:status" json:"status,omitempty"`
	Role       string    `gorm:"column:role" json:"role,omitempty"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.MiddleName, validation.Required.Error("middle name is required")),
		validation.Field(&u.LastName, validation.Required.Error("last name is required")),
		validation.Field(&u.Phone, validation.Required.Error("phone is required")),
		validation.Field(&u.Email, is.EmailFormat.Error("email is not valid")),
		// validation.Field(&u.Phone, validation.Required.Error("phone is required"), validation.By(validatePhone)),
		validation.Field(&u.Password, validation.When(u.Email != "", validation.Required.Error("password is required"), validation.Length(6, 32).Error("password must be between 6 and 32 characters"))),
	)
}

// func validatePhone(phone interface{}) error {
// 	str := phonenumber.Parse(fmt.Sprintf("%v", phone), "ET")
// 	if str == "" {
// 		return fmt.Errorf("invalid phone number")
// 	}
// 	return nil
// }
