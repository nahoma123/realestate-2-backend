package initiator

import (
	"context"
	"fmt"
	"visitor_management/internal/storage"
	"visitor_management/internal/storage/persistence"
	"visitor_management/platform/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Persistence struct {
	// TODO implement
	User          storage.UserStorage
	Generic       storage.GenericStorage
	EstateStorage persistence.EstateStorage
}

func createDB(host, user, password, dbname, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateIndexes(log logger.Logger, db *gorm.DB) {
	log.Info(context.Background(), "create indexes")

	err := db.Exec(`
		CREATE UNIQUE INDEX unique_email
		ON users ( email);
	`).Error
	if err != nil {
		log.Debug(context.Background(), fmt.Sprintf("create indexes error: %s", err.Error()))
	}
}

func InitPersistence(db *gorm.DB, log logger.Logger) Persistence {

	userStorage := persistence.InitUserDB(db)
	genericStorage := persistence.InitGenericDB(db)
	estateStorage := persistence.InitRlEstateDB(db, genericStorage)

	return Persistence{
		User:          userStorage,
		Generic:       genericStorage,
		EstateStorage: estateStorage,
	}
}
