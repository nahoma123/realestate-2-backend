package estate

import (
	"net/http"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	estateM "visitor_management/internal/module/estate"
	"visitor_management/platform/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommunicationHandler struct {
	logger              logger.Logger
	CommunicationModule estateM.CommModule
}

func InitComm(logger logger.Logger, commModule estateM.CommModule) CommunicationHandler {
	return CommunicationHandler{
		logger,
		commModule,
	}
}

func (o CommunicationHandler) AddMessage(ctx *gin.Context) {
	message := &model.Message{}
	err := ctx.ShouldBind(&message)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	message.PropertyId = ctx.Param("property_id")

	message, err = o.CommunicationModule.AddMessage(ctx, message)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, message, nil)
}

func (es CommunicationHandler) GetMessages(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)
	// propertyId := ctx.Param("property_id")

	// ftl = constant.AddFilter(*ftl, "property_id", propertyId, "=")

	valuations, err := es.CommunicationModule.GetMessages(ctx, ftl)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	constant.SuccessResponse(ctx, http.StatusOK, valuations, ftl)
}

func (es CommunicationHandler) GetLandLordMessages(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)
	propertyId := ctx.Param("property_id")

	ftl = constant.AddFilter(*ftl, "property_id", propertyId, "=")
	ftl = constant.AddFilter(*ftl, "admin_approved", "true", "=")

	valuations, err := es.CommunicationModule.GetMessages(ctx, ftl)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	constant.SuccessResponse(ctx, http.StatusOK, valuations, ftl)
}
func (es CommunicationHandler) AdminApproveMessage(ctx *gin.Context) {
	message := &model.Message{
		MessageId:     ctx.Param("message_id"),
		AdminApproved: true,
	}

	valuations, err := es.CommunicationModule.UpdateMessage(ctx, message.MessageId, message)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	constant.SuccessResponse(ctx, http.StatusOK, valuations, nil)
}
