package estate

import (
	"context"
	"strings"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/storage"
	"visitor_management/internal/storage/persistence"
	"visitor_management/platform/logger"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type CommModule struct {
	logger logger.Logger
	comStr persistence.CommunicationStorage
	gnr    storage.GenericStorage
}

func InitComm(logger logger.Logger, gnr storage.GenericStorage, comStr persistence.CommunicationStorage) CommModule {
	return CommModule{
		logger,
		comStr,
		gnr,
	}
}

func (comm CommModule) AddCompliance(ctx context.Context, comp *model.Compliance) (*model.Compliance, error) {
	if err := comp.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		comm.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	id, _ := uuid.NewV4()
	comp.ComplianceId = id.String()

	co := &model.Compliance{}
	err := comm.gnr.GetOne(ctx, string(constant.DbCompliances), co, "property_id", comp.PropertyId)
	if err != nil && strings.Contains(err.Error(), "no record found") {
		err = comm.gnr.CreateOne(ctx, string(constant.DbCompliances), comp)
		if err != nil {
			comm.logger.Warn(ctx, err.Error())
			return nil, err
		}
		return comp, nil
	}

	err = comm.gnr.UpdateOne(ctx, string(constant.DbCompliances), comp, "property_id", comp.PropertyId)
	if err != nil {
		comm.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return comp, nil
}

func (comm CommModule) UpdateCompliance(ctx context.Context, complianceId string, compliance *model.Compliance) (*model.Compliance, error) {
	if err := compliance.ValidateUpdate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		comm.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}
	compliance.ComplianceId = complianceId
	err := comm.gnr.UpdateOne(ctx, string(constant.DbCompliances), compliance, "compliance_id", complianceId)
	if err != nil {
		comm.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return compliance, nil
}

func (comm CommModule) GetCompliances(ctx context.Context, filterPagination *constant.FilterPagination) (interface{}, error) {
	compliances, err := comm.gnr.GetAll(ctx, string(constant.DbCompliances), nil, filterPagination)
	if err != nil {
		return nil, err
	}

	return compliances, nil
}

func (comm CommModule) GetCompliance(ctx context.Context, propId string) (*model.Compliance, error) {
	compliance := &model.Compliance{}
	err := comm.gnr.GetOne(ctx, string(constant.DbCompliances), compliance, "property_id", propId)
	if err != nil {
		return nil, err
	}

	return compliance, nil
}
