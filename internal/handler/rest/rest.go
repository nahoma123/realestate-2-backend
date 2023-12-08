package rest

import (
	"github.com/gin-gonic/gin"
)

type User interface {
	Register(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	Login(ctx *gin.Context)

	CreatePasswordResetRequest(ctx *gin.Context)
	VerifyResetCode(ctx *gin.Context)

	VerifyForgotPassword(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
}

type Estate interface {
	AddEstate(ctx *gin.Context)
	UpdateEstate(ctx *gin.Context)
	AddEstateService(ctx *gin.Context)
	AddEstateStaff(ctx *gin.Context)
	AddHouseOwner(ctx *gin.Context)
	AddHouse(ctx *gin.Context)
	GetEstateHouses(ctx *gin.Context)
	GetOwnerHouses(ctx *gin.Context)
	GetFacultyServices(ctx *gin.Context)
	GetEstateStaffs(ctx *gin.Context)
	AssignResident(ctx *gin.Context)
	GetResident(ctx *gin.Context)
	GetResidents(ctx *gin.Context)
	GetHouseFees(ctx *gin.Context)
	CreateHouseFee(ctx *gin.Context)

	GetHouseFeeSubscription(ctx *gin.Context)
	AddHouseFeeSubscription(ctx *gin.Context)

	GetFacultyServiceSubscriptions(ctx *gin.Context)
	AddFacultyServiceSubscription(ctx *gin.Context)

	AddGuest(ctx *gin.Context)
	GetGuests(ctx *gin.Context)

	AddInvites(ctx *gin.Context)
	GetInvites(ctx *gin.Context)

	VerifyCode(ctx *gin.Context)
	UpdateInvite(ctx *gin.Context)
	GetMonthlyReport(ctx *gin.Context)
	FacultySubscriptionFeeReport(ctx *gin.Context)
	InvitesReport(ctx *gin.Context)
	LogBooking(ctx *gin.Context)

	UpdateEstateConfiguration(ctx *gin.Context)
	AddEstateConfiguration(ctx *gin.Context)
	GetEstateConfiguration(ctx *gin.Context)
}
