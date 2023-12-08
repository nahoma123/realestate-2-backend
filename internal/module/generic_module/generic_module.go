package user

import (
	"context"
	"visitor_management/internal/constant"
	"visitor_management/internal/module"
	"visitor_management/internal/storage"
	"visitor_management/platform/logger"
)

type gen struct {
	logger  logger.Logger
	generic storage.GenericStorage
}

func InitGenericModule(logger logger.Logger, genStorage storage.GenericStorage) module.GenericModule {
	return &gen{
		logger,
		genStorage,
	}
}

func (g gen) GetAny(cxt context.Context, colName string, model interface{}, filterPagination *constant.FilterPagination) ([]interface{}, error) {
	// return g.generic.GetAll(cxt, colName, model, filterPagination)
	return nil, nil
}

func (g gen) UpdateOne(ctx constant.Context, id string) (interface{}, error) {
	return nil, nil
}
