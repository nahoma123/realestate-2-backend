package estate

import (
	"net/http"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (es EstateHandler) GetInspectionResult(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)

	valuations, err := es.EstateModule.GetInspectionResults(ctx, ftl)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	constant.SuccessResponse(ctx, http.StatusOK, valuations, ftl)
}

func (es EstateHandler) AddInspectionResult(ctx *gin.Context) {
	ispRst := &model.InspectionResult{}
	err := ctx.ShouldBind(&ispRst)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	ispRst, err = es.EstateModule.AddInspectionResult(ctx, ispRst)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, ispRst, nil)
}

func (o EstateHandler) UpdateInspectionResult(ctx *gin.Context) {
	inspectionResult := &model.InspectionResult{}
	err := ctx.ShouldBind(&inspectionResult)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	inspectionResultId := ctx.Param("inspection_result_id")
	inspectionResult.InspectionResultId = inspectionResultId

	err = o.EstateModule.UpdateInspectionResult(ctx, inspectionResultId, inspectionResult)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, inspectionResult, nil)
}
