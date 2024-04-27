package initiator

import (
	"visitor_management/internal/handler/rest"
	"visitor_management/internal/handler/rest/estate"
	"visitor_management/internal/handler/rest/user"
	"visitor_management/platform/logger"
)

type Handler struct {
	// TODO implement
	user        rest.User
	estate      estate.EstateHandler
	commHandler estate.CommunicationHandler
}

func InitHandler(module Module, log logger.Logger) Handler {
	return Handler{
		// TODO implement
		user:        user.InitUser(log, module.UserModule),
		estate:      estate.InitEstate(log, module.EstateModule),
		commHandler: estate.InitComm(log, module.CommModule),
	}
}
