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

func (re EstateModule) GetInspectionResults(ctx context.Context, filterPagination *constant.FilterPagination) (interface{}, error) {
	inspectionResults, err := re.gnr.GetAll(ctx, string(constant.DbInspectionResults), nil, filterPagination)
	if err != nil {
		return nil, err
	}

	return inspectionResults, nil
}

func (es EstateModule) AddInspectionResult(ctx context.Context, inpRst *model.InspectionResult) (*model.InspectionResult, error) {
	if err := inpRst.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	inpRst, err := es.rlEst.AddInspectionResult(ctx, inpRst)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return inpRst, nil
}

func (re EstateModule) UpdateInspectionResult(ctx context.Context, inspectionResultId string, inspection_result *model.InspectionResult) error {
	if err := inspection_result.ValidateUpdate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		logger.Log().Error(ctx, err.Error())
		return err
	}
	inspection_result.UpdatedAt = time.Now()

	err := re.gnr.UpdateOne(ctx, string(constant.DbInspectionResults), inspection_result, "inspection_result_id", inspectionResultId)
	if err != nil {
		return err
	}

	// err = re.gnr.GetOne(ctx, string(constant.DbInspectionResults), inspection_result, "property_id", inspectionResultId)
	// if err != nil {
	// 	return err
	// }

	return nil
}
