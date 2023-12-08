package model

import (
	"crypto/rand"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	RESIDENT    = "RESIDENT"
	HOUSE_OWNER = "HOUSE_OWNER"
)

type Invite struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	InviteID     string             `bson:"invite_id,omitempty" json:"invite_id,omitempty"`
	InviteCode   string             `bson:"invite_code,omitempty" json:"invite_code,omitempty"`
	GuestID      string             `bson:"guest_id,omitempty" json:"guest_id,omitempty"`
	ResidentID   string             `bson:"resident_id,omitempty" json:"resident_id,omitempty"`
	HouseOwnerID string             `bson:"house_owner_id,omitempty" json:"house_owner_id,omitempty"`
	InviteBy     string             `bson:"invite_by,omitempty" json:"invite_by,omitempty"`
	HouseID      string             `bson:"house_id,omitempty" json:"house_id,omitempty"`

	StartDate time.Time `bson:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate   time.Time `bson:"end_date,omitempty" json:"end_date,omitempty"`

	Guest      *Guest      `bson:"-" json:"guest,omitempty"`
	Resident   *Resident   `bson:"-" json:"resident,omitempty"`
	HouseOwner *HouseOwner `bson:"-" json:"house_owner,omitempty"`
	House      *House      `bson:"-" json:"house,omitempty"`

	Status    string    `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (es Invite) Validate() error {
	now := time.Now()

	return validation.ValidateStruct(&es,
		validation.Field(&es.ResidentID, validation.When(es.InviteBy == RESIDENT, validation.Required.Error("resident_id is required"))),
		validation.Field(&es.HouseOwnerID, validation.When(es.InviteBy == HOUSE_OWNER, validation.Required.Error("house_owner_id is required"))),
		validation.Field(&es.InviteBy, validation.In(RESIDENT, HOUSE_OWNER).Error(fmt.Sprintf("invite_by must be either %s, %s", RESIDENT, HOUSE_OWNER))),
		validation.Field(&es.HouseID, validation.Required.Error("house_id is required")),
		validation.Field(&es.InviteBy, validation.Required.Error("invite_by is required")),
		validation.Field(&es.StartDate, validation.Required.Error("start_date is required"), validation.Min(now).Error("start_date must be greater than or equal to now")),
		validation.Field(&es.EndDate, validation.Required.Error("end_date is required"), validation.Min(es.StartDate).Error("end_date must be greater than start_date")),
	)
}

func (es *Invite) ValidateUpdate() error {
	now := time.Now()
	return validation.ValidateStruct(es,
		validation.Field(&es.ResidentID, validation.When(es.ResidentID != "", validation.Required)),
		validation.Field(&es.HouseID, validation.When(es.HouseID != "", validation.Required)),
		validation.Field(&es.StartDate, validation.When(!es.StartDate.IsZero(), validation.Min(now).Error("start_date must be greater than or equal to now"))),
		validation.Field(&es.EndDate, validation.When(!es.EndDate.IsZero(), validation.Min(es.StartDate).Error("end_date must be greater than start_date"))),
	)
}

func (es *Invite) GenerateInviteCode() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error generating invite code:", err)
		return ""
	}
	return "inv_code_" + fmt.Sprintf("%x", b)
}
