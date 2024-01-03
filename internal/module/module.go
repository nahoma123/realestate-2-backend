package module

import (
	"context"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/model"

	"github.com/golang-jwt/jwt/v4"
)

type AuthDetail struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

type UserModule interface {
	VerifyUserStatus(ctx context.Context, phone string) error
	VerifyToken(signingMethod jwt.SigningMethod, tokenString string) (bool, *jwt.MapClaims, error)
	GetUserStatus(ctx context.Context, Id string) (string, error)
	GetUser(ctx context.Context, Id string) (*model.User, error)
	Login(ctx context.Context, lType, email, password string) (*AuthDetail, error)
	RegisterUser(ctx context.Context, profile *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, profile *model.User) (*model.User, error)
	GetAll(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.User, error)

	CreatePasswordResetRequest(ctx context.Context, userId string) error
	VerifyResetCode(ctx context.Context, userCode int, userId, newPassword string) error

	ForgotPasswordResetRequest(ctx context.Context, email string) error
}

type EstateModule interface {
	CreateEstate(ctx context.Context, valuation *model.RealEstate) (*model.RealEstate, error)
}

type GenericModule interface {
	UpdateOne(ctx constant.Context, id string) (interface{}, error)
	GetAny(cxt context.Context, colName string, model interface{}, filterPagination *constant.FilterPagination) ([]interface{}, error)
}
