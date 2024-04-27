package persistence

import (
	"visitor_management/internal/storage"

	"gorm.io/gorm"
)

type CommunicationStorage struct {
	db  *gorm.DB
	gnr storage.GenericStorage
}

func InitCommunicationDB(db *gorm.DB, gnr storage.GenericStorage) CommunicationStorage {
	return CommunicationStorage{
		db:  db,
		gnr: gnr,
	}
}
