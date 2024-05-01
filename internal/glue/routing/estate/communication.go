package estate

import (
	"visitor_management/internal/glue/routing"
	"visitor_management/internal/handler/middleware"
	"visitor_management/internal/handler/rest/estate"

	"github.com/gin-gonic/gin"
)

func InitCommunicationRoute(router *gin.RouterGroup, handler estate.CommunicationHandler, authMiddleware middleware.AuthMiddleware) {
	comRoutes := []routing.Router{
		{
			Method:  "PUT",
			Path:    "/properties/:property_id/compliances",
			Handler: handler.AddCompliance,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(false),
			},
		},
		{
			Method:  "POST",
			Path:    "/compliances",
			Handler: handler.AddCompliance,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(false),
			},
		},
		{
			Method:  "PATCH",
			Path:    "/compliances/:compliance_id",
			Handler: handler.UpdateCompliance,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(false),
			},
		},
		{
			Method:  "GET",
			Path:    "/compliances",
			Handler: handler.GetCompliances,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:  "GET",
			Path:    "properties/:property_id/compliances",
			Handler: handler.GetPropertyCompliance,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},

		//
		{
			Method:  "POST",
			Path:    "/tenant/properties/:property_id/messages",
			Handler: handler.AddMessage,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(false),
			},
		}, {
			Method:  "GET",
			Path:    "/messages",
			Handler: handler.GetMessages,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		//
		{
			Method:  "GET",
			Path:    "/landlord/properties/:property_id/messages",
			Handler: handler.GetLandLordMessages,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:  "POST",
			Path:    "/admin/messages/:message_id/approve",
			Handler: handler.AdminApproveMessage,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
	}
	routing.RegisterRoutes(router, comRoutes)
}
