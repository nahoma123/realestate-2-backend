package estate

import (
	"context"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/platform/logger"

	"go.uber.org/zap"
)

func (re EstateModule) GetMaintenanceRequests(ctx context.Context, filterPagination *constant.FilterPagination) (interface{}, error) {
	maintenanceRequests, err := re.gnr.GetAll(ctx, string(constant.DbMaintenanceRequests), nil, filterPagination)
	if err != nil {
		return nil, err
	}

	return maintenanceRequests, nil
}

func (es EstateModule) AddMaintenanceRequest(ctx context.Context, inpRst *model.MaintenanceRequest) (*model.MaintenanceRequest, error) {
	if err := inpRst.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	inpRst, err := es.rlEst.AddMaintenaceRequest(ctx, inpRst)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return inpRst, nil
}

func (re EstateModule) UpdateMaintenanceRequest(ctx context.Context, maintenanceRequestId string, maintenance_request *model.MaintenanceRequest) error {
	if err := maintenance_request.ValidateUpdate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		logger.Log().Error(ctx, err.Error())
		return err
	}
	maintenance_request.UpdatedAt = time.Now()

	err := re.gnr.UpdateOne(ctx, string(constant.DbMaintenanceRequests), maintenance_request, "maintenance_request_id", maintenanceRequestId)
	if err != nil {
		return err
	}

	// err = re.gnr.GetOne(ctx, string(constant.DbMaintenanceRequests), maintenance_request, "property_id", maintenanceRequestId)
	// if err != nil {
	// 	return err
	// }

	return nil
}
