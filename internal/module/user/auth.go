package user

import (
	"context"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/module"
	"visitor_management/internal/storage"
	"visitor_management/internal/storage/persistence"
	"visitor_management/platform/logger"
)

type UserModuleWrapper struct {
	*user
}

type user struct {
	logger      logger.Logger
	userStorage storage.UserStorage
	generic     storage.GenericStorage
}

func InitOAuth(logger logger.Logger, generic storage.GenericStorage, userStorage storage.UserStorage) UserModuleWrapper {
	return UserModuleWrapper{&user{
		logger,
		userStorage,
		generic,
	}}
}

func (o *user) VerifyUserStatus(ctx context.Context, phone string) error {

	// logic from other microservice
	return nil
}

func (o *user) GetUserStatus(ctx context.Context, Id string) (string, error) {
	//
	return "", nil
}

func (o *user) Login(ctx context.Context, email, password string) (*module.AuthDetail, error) {
	user, err := o.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if persistence.CheckPasswordHash(password, user.Password) {
		token, err := o.GenerateJWT(user)
		if err != nil {
			logger.Log().Error(ctx, err.Error())
			return nil, errors.ErrInternalServerError.New("unable to generate token")
		}
		return &module.AuthDetail{
			User:  user,
			Token: token,
		}, nil
	}
	return nil, errors.ErrAuthError.New("email and password do not match")
}

func (o *user) GetAll(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.User, error) {
	fsc, err := o.userStorage.GetAll(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}
