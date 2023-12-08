package estate

import (
	"context"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/storage"
	"visitor_management/internal/storage/persistence"
	"visitor_management/platform/logger"

	"go.uber.org/zap"
)

type EstateModule struct {
	logger logger.Logger
	rlEst  persistence.EstateStorage
	gnr    storage.GenericStorage
}

func InitEstate(logger logger.Logger, gnr storage.GenericStorage, rlEst persistence.EstateStorage) EstateModule {
	return EstateModule{
		logger,
		rlEst,
		gnr,
	}
}

func (es EstateModule) AddValuation(ctx context.Context, valuation *model.RealEstate) (*model.RealEstate, error) {
	if err := valuation.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	valuation.Status = constant.ActiveRealEstateStatus

	valuation, err := es.rlEst.AddEvaluation(ctx, valuation)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return valuation, nil
}

func (re EstateModule) GetValuations(ctx context.Context, filterPagination *constant.FilterPagination) (interface{}, error) {
	valuations, err := re.gnr.GetAll(ctx, string(constant.DbRealEstate), nil, filterPagination)
	if err != nil {
		return nil, err
	}

	return valuations, nil
}

func (re EstateModule) UpdateValuation(ctx context.Context, valuationId string, valuation *model.RealEstate) error {
	if err := valuation.ValidateUpdate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		logger.Log().Error(ctx, err.Error())
		return err
	}

	valuation.UpdatedAt = time.Now()

	err := re.gnr.UpdateOne(ctx, string(constant.DbRealEstate), valuation, "real_estate_id", valuationId)
	if err != nil {
		return err
	}

	// start prepare to send valuation post through rabbitmq
	err = re.gnr.GetOne(ctx, string(constant.DbRealEstate), valuation, "real_estate_id", valuationId)
	if err != nil {
		return err
	}

	return nil
}

func (es EstateModule) AddProperty(ctx context.Context, property *model.Property) (*model.Property, error) {
	if err := property.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	property.Status = constant.ActivePropertyStatus

	property, err := es.rlEst.AddProperty(ctx, property)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return property, nil
}

func (re EstateModule) UpdateProperty(ctx context.Context, propertyId string, property *model.Property) error {
	if err := property.ValidateUpdate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		logger.Log().Error(ctx, err.Error())
		return err
	}
	property.UpdatedAt = time.Now()

	err := re.gnr.UpdateOne(ctx, string(constant.DbProperties), property, "property_id", propertyId)
	if err != nil {
		return err
	}

	// err = re.gnr.GetOne(ctx, string(constant.DbProperties), property, "property_id", propertyId)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (re EstateModule) GetProperties(ctx context.Context, filterPagination *constant.FilterPagination) (interface{}, error) {
	jbSeeker, err := re.gnr.GetAll(ctx, string(constant.DbProperties), nil, filterPagination)
	if err != nil {
		return nil, err
	}

	return jbSeeker, nil
}
