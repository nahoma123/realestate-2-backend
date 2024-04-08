package estate

import (
	"net/http"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (es EstateHandler) GetMaintenanceRequest(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)

	valuations, err := es.EstateModule.GetMaintenanceRequests(ctx, ftl)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	constant.SuccessResponse(ctx, http.StatusOK, valuations, ftl)
}

func (es EstateHandler) AddMaintenanceRequest(ctx *gin.Context) {
	ispRst := &model.MaintenanceRequest{}
	err := ctx.ShouldBind(&ispRst)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	ispRst, err = es.EstateModule.AddMaintenanceRequest(ctx, ispRst)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, ispRst, nil)
}

func (o EstateHandler) UpdateMaintenanceRequest(ctx *gin.Context) {
	maintenanceRequest := &model.MaintenanceRequest{}
	err := ctx.ShouldBind(&maintenanceRequest)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	maintenanceRequestId := ctx.Param("maintenance_request_id")
	maintenanceRequest.MaintenanceRequestId = maintenanceRequestId

	err = o.EstateModule.UpdateMaintenanceRequest(ctx, maintenanceRequestId, maintenanceRequest)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, maintenanceRequest, nil)
}
