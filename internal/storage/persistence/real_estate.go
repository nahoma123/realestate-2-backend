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

func (re EstateStorage) AddMaintenaceRequest(ctx context.Context, mainReq *model.MaintenanceRequest) (*model.MaintenanceRequest, error) {
	mainReq.CreatedAt = time.Now()
	mainReq.UpdatedAt = time.Now()

	id, _ := uuid.NewV4()

	mainReq.MaintenanceRequestId = id.String()

	err := re.gnr.CreateOne(ctx, constant.DbMaintenanceRequests, mainReq)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if gorm.ErrDuplicatedKey == err {
			return nil, errors.ErrDataExists.Wrap(err, errors.MaintenanceIsAlreadyRegistered)
		}
		return nil, errors.ErrInternalServerError.New("unknown error occurred")
	}

	return mainReq, err
}

func (re EstateStorage) GetRentDetails(ctx context.Context, propertyId string) (*model.PropertyRentDetails, error) {
	rent := &model.PropertyRentDetails{}
	err := re.db.Table(string(constant.DbProperties)).Preload("Tenant").Preload("Landlord").Where("property_id", propertyId).First(rent).Error
	if err != nil {
		return nil, err
	}

	return rent, nil
}

func (re EstateStorage) ConfirmPayRent(ctx context.Context, propertyId string) (*model.Property, error) {
	rent := &model.Property{}
	rent.RentLastPaid = time.Now()

	err := re.db.Table(string(constant.DbProperties)).Where("property_id=?", propertyId).Save(rent).Error
	if err != nil {
		return nil, err
	}

	return rent, nil
}
