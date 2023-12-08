package model

// for monthly tenant stat
type MonthlyReport struct {
	PaidResidents   int     `json:"paid_residents"`
	UnpaidResidents int     `json:"unpaid_residents"`
	TotalResidents  int     `json:"total_residents"`
	UnpaidAmount    float64 `json:"unpaid_amount"`
}

type TenantReport struct {
	Amount float64 `bson:"amount"`
}

type FacultyServiceMonthlyReport struct {
	FacultyService *FacultyService `bson:"faculty_service" json:"faculty_service"`
	PaidCount      int             `bson:"paid_count" json:"paid_count"`
	UnpaidCount    int             `bson:"unpaid_count" json:"unpaid_count"`
}

type InviteReport struct {
	MonthNumber int `bson:"_id" json:"month_number"`
	Invites     int `bson:"count" json:"invites"`
}
