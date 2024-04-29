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
		// 	Path:        "/properties/:property_id/rent_details",
		// 	Handler:     handler.,
		// 	Middlewares: []gin.HandlerFunc{},
		// },
	}
	routing.RegisterRoutes(router, comRoutes)
}
