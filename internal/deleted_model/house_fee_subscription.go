package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HouseFeeSubscription struct {
	Id                     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HouseFeeSubscriptionID string             `bson:"house_fee_subscription_id,omitempty" json:"house_fee_subscription_id,omitempty"`
	ResidentID             string             `bson:"resident_id,omitempty" json:"resident_id,omitempty"`
	HouseFeeID             string             `bson:"house_fee_id,omitempty" json:"house_fee_id,omitempty"`

	StartDate time.Time `bson:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate   time.Time `bson:"end_date,omitempty" json:"end_date,omitempty"`

	NextBillingDate time.Time `bson:"next_billing_date,omitempty" json:"next_billing_date,omitempty"`
	LastBillingDate time.Time `bson:"last_billing_date,omitempty" json:"last_billing_date,omitempty"`

	Frequency   string `bson:"frequency,omitempty" json:"frequency,omitempty"`
	DayOfMonth  int    `bson:"day_of_month,omitempty" json:"day_of_month,omitempty"`
	MonthOfYear int    `bson:"month_of_year,omitempty" json:"month_of_year,omitempty"`

	Status    string    `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (es HouseFeeSubscription) Validate() error {
	if es.Frequency != "yearly" && es.MonthOfYear != 0 {
		return validation.ErrDateOutOfRange.SetMessage("month_of_the_year can only be set if frequency is in yearly")
	} else if es.Frequency == "yearly" && es.MonthOfYear == 0 {
		return validation.ErrDateOutOfRange.SetMessage("month_of_the_year should be set if frequency is in yearly")
	}

	return validation.ValidateStruct(&es,
		validation.Field(&es.Frequency, validation.Required, validation.In("monthly", "yearly").Error("Frequency must be monthly or yearly")),
		validation.Field(&es.ResidentID, validation.Required.Error("resident_id is required")),
		validation.Field(&es.HouseFeeID, validation.Required.Error("house_fee_id is required")),
		validation.Field(&es.StartDate, validation.Required.Error("start_date is required")),
		// validation.Field(&es.EndDate, validation.Required.Error("end_date is required")),
		validation.Field(&es.Frequency, validation.Required.Error("frequency is required")),
	)
}

func (es *HouseFeeSubscription) InitializePaymentDates() {
	if es.Frequency == "monthly" {
		es.NextBillingDate = es.StartDate.AddDate(0, 1, 0) // Add 1 month to the start date
		es.NextBillingDate = time.Date(es.NextBillingDate.Year(), es.NextBillingDate.Month(), es.DayOfMonth, 0, 0, 0, 0, es.NextBillingDate.Location())

		es.LastBillingDate = es.StartDate
		es.LastBillingDate = time.Date(es.LastBillingDate.Year(), es.LastBillingDate.Month(), es.DayOfMonth, 0, 0, 0, 0, es.LastBillingDate.Location())
	} else if es.Frequency == "yearly" {
		es.NextBillingDate = es.StartDate.AddDate(1, 0, 0) // Add 1 year to the start date
		es.NextBillingDate = time.Date(es.NextBillingDate.Year(), time.Month(es.MonthOfYear), es.DayOfMonth, 0, 0, 0, 0, es.NextBillingDate.Location())

		es.LastBillingDate = es.StartDate
		es.LastBillingDate = time.Date(es.LastBillingDate.Year(), time.Month(es.MonthOfYear), es.DayOfMonth, 0, 0, 0, 0, es.LastBillingDate.Location())
	}
}
