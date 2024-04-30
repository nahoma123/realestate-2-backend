package initiator

import (
	// swager docs import
	// "visitor_management/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	"visitor_management/internal/glue/routing/estate"
	"visitor_management/internal/glue/routing/user"
	"visitor_management/internal/handler/middleware"
	"visitor_management/platform/logger"
)

func InitRouter(router *gin.Engine, group *gin.RouterGroup, handler Handler, module Module, log logger.Logger, publicKeyPath string) {

	router.Use(gin.Logger())

	authMiddleware := middleware.InitAuthMiddleware(module.UserModule, nil)

	docs.SwaggerInfo.BasePath = "/v1"

	// swager docs import
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// auth.InitRoute(group, handler.oauth, authMiddleware)

	user.InitRoute(group, handler.user, authMiddleware)
	estate.InitEstateRoute(group, handler.estate, authMiddleware)
	estate.InitCommunicationRoute(group, handler.commHandler, authMiddleware)
}
