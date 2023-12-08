package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RealEstate struct {
	Id           uint   `gorm:"primaryKey" json:"id,omitempty"`
	RealEstateId string `gorm:"column:real_estate_id" json:"real_estate_id,omitempty"`
	Email        string `gorm:"column:email" json:"email,omitempty"`
	Address      string `gorm:"column:address" json:"address,omitempty"`
	PhoneNumber  string `gorm:"column:phone_number" json:"phone_number,omitempty"`
	WhyJoined    int    `gorm:"column:why_joined" json:"why_joined,omitempty"`

	PreferredTime time.Time `gorm:"preferred_time,omitempty" json:"preferred_time,omitempty"`
	Status        string    `gorm:"status,omitempty" json:"status,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (es RealEstate) Validate() error {
	now := time.Now()

	return validation.ValidateStruct(&es,
		validation.Field(&es.PreferredTime, validation.Required.Error("preferred_time is required"), validation.Min(now).Error("preferred_time must be greater than or equal to now")),
	)
}

func (es RealEstate) ValidateUpdate() error {
	return validation.ValidateStruct(&es)
}

// func (es *Invite) GenerateInviteCode() string {
// 	b := make([]byte, 16)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		fmt.Println("Error generating invite code:", err)
// 		return ""
// 	}
// 	return "inv_code_" + fmt.Sprintf("%x", b)
// }
