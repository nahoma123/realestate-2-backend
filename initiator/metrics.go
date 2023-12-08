package initiator

import (
	"visitor_management/platform/logger"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitMetricsRoute(group *gin.Engine, log logger.Logger) {
	group.GET("/metrics", func(context *gin.Context) {
		promhttp.Handler().ServeHTTP(context.Writer, context.Request)
	})
}
