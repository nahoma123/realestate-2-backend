package initiator

import (
	"context"
	"visitor_management/internal/constant/model"
	"visitor_management/platform/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitiateMigration(db *gorm.DB, log logger.Logger) {
	// Auto-migrate the database models

	err := db.AutoMigrate(
		&model.User{},
		&model.RealEstate{},
		&model.Property{},
		&model.InspectionResult{},
		&model.MaintenanceRequest{},
		&model.Contract{})
	if err != nil {
		log.Error(context.Background(), "Failed to perform database migration", zap.Error(err))
		return
	}

	log.Info(context.Background(), "Database migration completed successfully")
}
