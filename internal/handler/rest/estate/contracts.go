package estate

import (
	"net/http"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (o EstateHandler) AddContract(ctx *gin.Context) {
	contract := &model.Contract{}
	err := ctx.ShouldBind(&contract)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	contract.PropertyId = ctx.Param("property_id")

	contract, err = o.EstateModule.AddContract(ctx, contract)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, contract, nil)
}

func (o EstateHandler) UpdateContract(ctx *gin.Context) {
	contract := &model.Contract{}
	err := ctx.ShouldBind(&contract)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	contract.ContractId = ctx.Param("contract_id")

	contract, err = o.EstateModule.UpdateContract(ctx, contract.ContractId, contract)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, contract, nil)
}

func (es EstateHandler) GetContracts(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)
	propertyId := ctx.Param("property_id")

	ftl = constant.AddFilter(*ftl, "property_id", propertyId, "=")

	valuations, err := es.EstateModule.GetContracts(ctx, ftl)
	if err != nil {
		es.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}
	constant.SuccessResponse(ctx, http.StatusOK, valuations, ftl)
}
