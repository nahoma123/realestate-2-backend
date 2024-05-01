package estate

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	estateM "visitor_management/internal/module/estate"
	"visitor_management/platform/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EstateHandler struct {
	logger       logger.Logger
	EstateModule estateM.EstateModule
}

func InitEstate(logger logger.Logger, estateModule estateM.EstateModule) EstateHandler {
	return EstateHandler{
		logger,
		estateModule,
	}
}

func (o EstateHandler) AddValuation(ctx *gin.Context) {
	estate := &model.RealEstate{}
	err := ctx.ShouldBind(&estate)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	estate, err = o.EstateModule.AddValuation(ctx, estate)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, estate, nil)
}

func (o EstateHandler) UpdateValuation(ctx *gin.Context) {
	valuation := &model.RealEstate{}
	err := ctx.ShouldBind(&valuation)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	valuationId := ctx.Param("estate_id")
	valuation.RealEstateId = valuationId

	err = o.EstateModule.UpdateValuation(ctx, valuationId, valuation)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, valuation, nil)
}

func (o EstateHandler) AddProperty(ctx *gin.Context) {
	property := &model.Property{}
	err := ctx.ShouldBind(&property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	property, err = o.EstateModule.AddProperty(ctx, property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, property, nil)
}

func (o EstateHandler) UploadImage(ctx *gin.Context) {
	// Get the file from the request
	file, err := ctx.FormFile("image")
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	extension := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	filePath := "uploads/" + fileName

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, fmt.Sprintf("/files/%s", fileName), nil)
}

func (o EstateHandler) GetImage(ctx *gin.Context) {
	// Get the filename from the request params
	fileName := ctx.Param("filename")

	// Construct the file path
	filePath := filepath.Join("uploads", fileName)

	// Serve the file
	ctx.File(filePath)
}

func (o EstateHandler) UpdateProperty(ctx *gin.Context) {
	property := &model.Property{}
	err := ctx.ShouldBind(&property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	propertyId := ctx.Param("property_id")
	property.PropertyId = propertyId

	err = o.EstateModule.UpdateProperty(ctx, propertyId, property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, property, nil)
}

func (o EstateHandler) GetProperties(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)

	properties, err := o.EstateModule.GetProperties(ctx, ftl)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, properties, ftl)

}

func (o EstateHandler) GetValuations(ctx *gin.Context) {
	ftl := constant.ParseFilterPagination(ctx)

	valuations, err := o.EstateModule.GetValuations(ctx, ftl)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, valuations, ftl)

}

func (o EstateHandler) GetRentDetails(ctx *gin.Context) {
	propertyId := ctx.Param("property_id")

	valuations, err := o.EstateModule.GetRentDetails(ctx, propertyId)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, valuations, nil)
}

func (o EstateHandler) ConfirmPaymentRent(ctx *gin.Context) {
	property := &model.Property{}
	err := ctx.ShouldBind(&property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	propertyId := ctx.Param("property_id")
	property.PropertyId = propertyId
	property.RentLastPaid = time.Now()

	err = o.EstateModule.UpdateProperty(ctx, propertyId, property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "rent confirmed", nil)
}

func (o EstateHandler) RentProperty(ctx *gin.Context) {
	property := &model.Property{}
	err := ctx.ShouldBind(&property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	propertyId := ctx.Param("property_id")
	property.PropertyId = propertyId
	property.RentLastPaid = time.Now()

	err = o.EstateModule.RentProperty(ctx, propertyId, property)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "rent confirmed", nil)
}
