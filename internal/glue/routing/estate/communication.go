package estate

import (
	"visitor_management/internal/glue/routing"
	"visitor_management/internal/handler/middleware"
	"visitor_management/internal/handler/rest/estate"

	"github.com/gin-gonic/gin"
)

func InitCommunicationRoute(router *gin.RouterGroup, handler estate.CommunicationHandler, authMiddleware middleware.AuthMiddleware) {
	comRoutes := []routing.Router{
		// {
		// 	Method:      "POST",
		// 	Path:        "/estates/add_valuation",
		// 	Handler:     handler.AddValuation,
		// 	Middlewares: []gin.HandlerFunc{},
		// },
	}
	routing.RegisterRoutes(router, comRoutes)
}
