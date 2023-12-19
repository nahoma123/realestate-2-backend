package persistence

import (
	"context"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/storage"
	"visitor_management/platform/logger"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type EstateStorage struct {
	db  *gorm.DB
	gnr storage.GenericStorage
}

func InitRlEstateDB(db *gorm.DB, gnr storage.GenericStorage) EstateStorage {
	return EstateStorage{
		db:  db,
		gnr: gnr,
	}
}

func (re EstateStorage) AddEvaluation(ctx context.Context, vl *model.RealEstate) (*model.RealEstate, error) {
	vl.CreatedAt = time.Now()
	vl.UpdatedAt = time.Now()

	id, _ := uuid.NewV4()

	vl.RealEstateId = id.String()

	err := re.gnr.CreateOne(ctx, constant.DbRealEstate, vl)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if gorm.ErrDuplicatedKey == err {
			return nil, errors.ErrDataExists.Wrap(err, errors.EstateIsAlreadyRegistered)
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return vl, err
}

func (re EstateStorage) AddProperty(ctx context.Context, property *model.Property) (*model.Property, error) {
	property.CreatedAt = time.Now()
	property.UpdatedAt = time.Now()

	id, _ := uuid.NewV4()

	property.PropertyId = id.String()

	err := re.gnr.CreateOne(ctx, constant.DbProperties, property)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if gorm.ErrDuplicatedKey == err {
			return nil, errors.ErrDataExists.Wrap(err, errors.PropertyIsAlreadyRegistered)
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return property, err
}

func (re EstateStorage) AddInspectionResult(ctx context.Context, inspRes *model.InspectionResult) (*model.InspectionResult, error) {
	inspRes.CreatedAt = time.Now()
	inspRes.UpdatedAt = time.Now()

	id, _ := uuid.NewV4()

	inspRes.InspectionResultId = id.String()

	err := re.gnr.CreateOne(ctx, constant.DbInspectionResults, inspRes)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if gorm.ErrDuplicatedKey == err {
			return nil, errors.ErrDataExists.Wrap(err, errors.InspectionResultIsAlreadyRegistered)
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return inspRes, err
}
