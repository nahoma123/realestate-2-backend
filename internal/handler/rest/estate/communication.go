package estate

import (
	estateM "visitor_management/internal/module/estate"
	"visitor_management/platform/logger"
)

type CommunicationHandler struct {
	logger              logger.Logger
	CommunicationModule estateM.CommModule
}

func InitComm(logger logger.Logger, commModule estateM.CommModule) CommunicationHandler {
	return CommunicationHandler{
		logger,
		commModule,
	}
}
