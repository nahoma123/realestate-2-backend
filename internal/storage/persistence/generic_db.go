package persistence

import (
	"context"
	"fmt"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/storage"
	"visitor_management/platform/logger"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type generic struct {
	db *gorm.DB
}

func InitGenericDB(db *gorm.DB) storage.GenericStorage {
	return &generic{
		db: db,
	}
}

func (g generic) GetAll(ctx context.Context, tableName string, data interface{}, filterPagination *constant.FilterPagination) (interface{}, error) {
	var results []map[string]interface{}

	err := constant.GetResults(ctx, g.db, tableName, filterPagination, &results)
	if err != nil {
		return nil, err
	}

	var models []interface{}
	for _, result := range results {
		models = append(models, result)
	}

	return models, nil
}

// func handleCursor(rows *gorm.Rows) interface{} {
// 	var results []map[string]interface{}

// 	for rows.Next() {
// 		var result map[string]interface{}

// 		err := rows.Scan(&result)
// 		if err != nil {
// 			// handle error
// 		}

// 		results = append(results, result)
// 	}

// 	if err := rows.Error; err != nil {
// 		// handle error
// 	}

//		return results
//	}
func (g generic) UpdateOne(ctx context.Context, tableName string, updateData interface{}, field string, value interface{}) error {
	filter := fmt.Sprintf("%s = ?", field)
	update := g.db.Table(tableName).Where(filter, value).Updates(updateData)
	if update.Error != nil {
		logger.Log().Error(ctx, update.Error.Error())
		return errors.ErrInternalServerError.New("unknown error occurred")
	}

	if update.RowsAffected == 0 {
		return errors.ErrNoRecordFound.New(fmt.Sprintf("%s not found", tableName))
	}

	return nil
}

func (g generic) CreateOne(ctx context.Context, tableName string, data interface{}) error {
	create := g.db.Table(tableName).Create(data)
	if create.Error != nil {
		logger.Log().Error(ctx, create.Error.Error())

		if pqErr, ok := create.Error.(*pq.Error); ok && pqErr.Code == "23505" {
			// Handle duplicate key error
			return errors.ErrDataExists.Wrap(fmt.Errorf("%s", fmt.Sprintf(errors.TemplateAlreadyRegistered, tableName)), fmt.Sprintf(errors.TemplateAlreadyRegistered, tableName))
		}

		return errors.ErrInternalServerError.New("unknown error occurred")
	}

	return nil
}

func (g generic) GetOne(ctx context.Context, tableName string, data interface{}, field string, value interface{}) error {
	filter := fmt.Sprintf("%s = ?", field)
	result := g.db.Table(tableName).Where(filter, value).First(data)
	if result.Error != nil {
		logger.Log().Error(ctx, result.Error.Error())
		if result.Error == gorm.ErrRecordNotFound {
			return errors.ErrNoRecordFound.New(fmt.Sprintf("%s not found", tableName))
		}
		return errors.ErrInternalServerError.New("unknown error occurred")
	}
	return nil
}

func (g generic) DeleteOne(ctx context.Context, tableName string, field string, value interface{}) error {
	filter := fmt.Sprintf("%s = ?", field)
	result := g.db.Table(tableName).Where(filter, value).Delete(nil)
	if result.Error != nil {
		logger.Log().Error(ctx, result.Error.Error())
		if result.RowsAffected == 0 {
			return errors.ErrNoRecordFound.New(fmt.Sprintf("%s not found", tableName))
		}
		return errors.ErrInternalServerError.New("unknown error occurred")
	}
	return nil
}
