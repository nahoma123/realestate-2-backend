package estate

import (
	"visitor_management/internal/glue/routing"
	"visitor_management/internal/handler/middleware"
	"visitor_management/internal/handler/rest/estate"

	"github.com/gin-gonic/gin"
)

func InitEstateRoute(router *gin.RouterGroup, handler estate.EstateHandler, authMiddleware middleware.AuthMiddleware) {
	estateRoutes := []routing.Router{
		{
			Method:      "POST",
			Path:        "/estates/add_valuation",
			Handler:     handler.AddValuation,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:  "PATCH",
			Path:    "/estates/:estate_id",
			Handler: handler.UpdateValuation,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},

		{
			Method:  "POST",
			Path:    "/properties",
			Handler: handler.AddProperty,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(true),
			},
		},
		{
			Method:      "POST",
			Path:        "/upload_image",
			Handler:     handler.UploadImage,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "GET",
			Path:        "/files/:filename",
			Handler:     handler.GetImage,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:  "PATCH",
			Path:    "/properties/:property_id",
			Handler: handler.UpdateProperty,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(true),
			},
		},
		{
			Method:      "GET",
			Path:        "/properties",
			Handler:     handler.GetProperties,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "GET",
			Path:        "/valuations",
			Handler:     handler.GetValuations,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:  "GET",
			Path:    "/inspection_results",
			Handler: handler.GetInspectionResult,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:  "POST",
			Path:    "/inspection_results",
			Handler: handler.AddInspectionResult,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(true),
			},
		},
		{
			Method:  "PATCH",
			Path:    "/inspection_results/:inspection_result_id",
			Handler: handler.UpdateInspectionResult,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(true),
			},
		},

		{
			Method:  "GET",
			Path:    "/maintenance_requests",
			Handler: handler.GetMaintenanceRequest,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:  "POST",
			Path:    "/maintenance_requests",
			Handler: handler.AddMaintenanceRequest,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(false),
			},
		},
		{
			Method:  "PATCH",
			Path:    "/maintenance_requests/:maintenance_request_id",
			Handler: handler.UpdateMaintenanceRequest,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(false),
			},
		},
	}
	routing.RegisterRoutes(router, estateRoutes)
}
