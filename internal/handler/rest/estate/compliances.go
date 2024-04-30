package estate

import (
	"net/http"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (o CommunicationHandler) AddCompliance(ctx *gin.Context) {
	compliance := &model.Compliance{}
	err := ctx.ShouldBind(&compliance)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	compliance.PropertyId = ctx.Param("property_id")

	compliance, err = o.CommunicationModule.AddCompliance(ctx, compliance)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, compliance, nil)
}

func (o CommunicationHandler) UpdateCompliance(ctx *gin.Context) {
	compliance := &model.Compliance{}
	err := ctx.ShouldBind(&compliance)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	compliance.ComplianceId = ctx.Param("compliance_id")

	compliance, err = o.CommunicationModule.UpdateCompliance(ctx, compliance.ComplianceId, compliance)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, compliance, nil)
}

func (es CommunicationHandler) GetCompliances(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)
	// propertyId := ctx.Param("property_id")

	// ftl = constant.AddFilter(*ftl, "property_id", propertyId, "=")

	valuations, err := es.CommunicationModule.GetCompliances(ctx, ftl)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	constant.SuccessResponse(ctx, http.StatusOK, valuations, ftl)
}

func (es CommunicationHandler) GetPropertyCompliance(ctx *gin.Context) {
	propertyId := ctx.Param("property_id")

	valuations, err := es.CommunicationModule.GetCompliance(ctx, propertyId)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, valuations, nil)
}
