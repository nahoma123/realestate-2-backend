package storage

import (
	"context"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/model"
)

type DatabaseCollection string

// database collection constants
const (
	Users                      DatabaseCollection = "users"
	Estate                     DatabaseCollection = "estates"
	FacultyServiceSubscription DatabaseCollection = "faculty_services_subscriptions"
	Guest                      DatabaseCollection = "guests"
	Invite                     DatabaseCollection = "invites"
	FacultyService             DatabaseCollection = "faculty_services"
	Resident                   DatabaseCollection = "residents"
	HouseFee                   DatabaseCollection = "house_fees"
	HouseFeeSubscription       DatabaseCollection = "house_fee_subscriptions"
	Staff                      DatabaseCollection = "staffs"
	HouseOwner                 DatabaseCollection = "house_owners"
	House                      DatabaseCollection = "houses"
	BookingLog                 DatabaseCollection = "booking_logs"
	EstateConfiguration        DatabaseCollection = "estate_configurations"
	HouseOwnerInvite           DatabaseCollection = "house_owner_invites"
	Default                    DatabaseCollection = "defaults"
)

type UserStorage interface {
	Create(ctx context.Context, estate *model.User) (*model.User, error)
	Update(ctx context.Context, estate *model.User) (*model.User, error)
	Get(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, string string) (*model.User, error)
	GetAll(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.User, error)
}

type GenericStorage interface {
	// UpdateOne(ctx constant.Context, id string) error
	DeleteOne(ctx context.Context, tableName string, field string, value interface{}) error
	GetOne(ctx context.Context, tableName string, data interface{}, field string, value interface{}) error
	CreateOne(ctx context.Context, tableName string, data interface{}) error
	UpdateOne(ctx context.Context, tableName string, updateData interface{}, field string, value interface{}) error
	GetAll(ctx context.Context, tableName string, data interface{}, filterPagination *constant.FilterPagination) (interface{}, error)
	// GetAny(cxt context.Context, colName string, model interface{}, filterPagination *constant.FilterPagination) ([]interface{}, error)
}
