package estate

import (
	"visitor_management/internal/storage"
	"visitor_management/internal/storage/persistence"
	"visitor_management/platform/logger"
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
