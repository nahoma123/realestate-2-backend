package user

import (
	"net/http"
	"strings"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	userM "visitor_management/internal/module/user"
	"visitor_management/platform/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandlerWrapper struct {
	*user
}

type user struct {
	logger     logger.Logger
	UserModule userM.UserModuleWrapper
}

func InitUser(logger logger.Logger, userModule userM.UserModuleWrapper) *UserHandlerWrapper {
	return &UserHandlerWrapper{
		&user{
			logger,
			userModule,
		},
	}
}

func (o *UserHandlerWrapper) Register(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	user, err = o.UserModule.RegisterUser(ctx, user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, user, nil)
}

func (o *UserHandlerWrapper) UpdateUser(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.ShouldBind(&user)
	id := ctx.GetString("x-user-id")
	if err != nil || id == "" {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	user.UserID = id
	user, err = o.UserModule.UpdateUser(ctx, user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, user, nil)
}

func (o *UserHandlerWrapper) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := o.UserModule.GetUser(ctx, id)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, user, nil)
}

func (o *UserHandlerWrapper) Login(ctx *gin.Context) {
	url := ctx.Request.URL.Path
	lType := "Admin"
	if strings.Contains(url, "landlord") {
		lType = "Landlord"
	} else if strings.Contains(url, "tenant") {
		lType = "Tenant"
	}

	user := &model.User{}
	err := ctx.Bind(user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	auth, err := o.UserModule.Login(ctx, lType, user.Email, user.Password)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, auth, nil)
}

func (o *UserHandlerWrapper) GetUsers(ctx *gin.Context) {
	ftr := constant.ParseFilterPagination(ctx)

	states, err := o.UserModule.GetAll(ctx, ftr)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, states, ftr)
}

func (o *UserHandlerWrapper) CreatePasswordResetRequest(ctx *gin.Context) {
	userID := ctx.GetString("x-user-id")
	err := o.UserModule.CreatePasswordResetRequest(ctx, userID)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "successfully requested password reset", nil)
}

func (o *UserHandlerWrapper) VerifyResetCode(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.Bind(user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	userID := ctx.GetString("x-user-id")
	user.UserID = userID
	err = o.UserModule.VerifyResetCode(ctx, user.ResetCode, userID, user.Password)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "password changed successful", nil)
}

func (o *UserHandlerWrapper) ForgotPassword(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.Bind(user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	err = o.UserModule.ForgotPasswordResetRequest(ctx, user.Email)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "successfully requested password reset", nil)
}

func (o *UserHandlerWrapper) VerifyForgotPassword(ctx *gin.Context) {
	user := &model.User{}
	err := ctx.Bind(user)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	err = o.UserModule.VerifyForgetPasswordCode(ctx, user.ResetCode, user.Email, user.Password)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "password reset successful", nil)
}
