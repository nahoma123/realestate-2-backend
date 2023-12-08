package initiator

import (
	"visitor_management/internal/module"
	"visitor_management/internal/module/estate"
	gMod "visitor_management/internal/module/generic_module"
	"visitor_management/internal/module/user"
	"visitor_management/platform/logger"
)

type Module struct {
	// TODO implement
	UserModule    user.UserModuleWrapper
	GenericModule module.GenericModule
	EstateModule  estate.EstateModule
}

func InitModule(persistence Persistence, privateKeyPath string, platformLayer PlatformLayer, log logger.Logger) Module {

	gmod := gMod.InitGenericModule(log, persistence.Generic)
	estateModule := estate.InitEstate(log, persistence.Generic, persistence.EstateStorage)
	return Module{
		UserModule:    user.InitOAuth(log, persistence.Generic, persistence.User),
		GenericModule: gmod,
		EstateModule:  estateModule,
	}
}
