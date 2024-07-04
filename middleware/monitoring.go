package middleware

import (
	"net/http"

	"github.com/Lucifer07/e-wallet/server/monitoring"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func MiddlewareMonitoring(c *gin.Context) {
	timer:=prometheus.NewTimer(monitoring.HttpRequestDuration.WithLabelValues(c.Request.URL.Path,c.Request.Method,http.StatusText(c.Writer.Status())))
	c.Next()
	timer.ObserveDuration()
	monitoring.HttpRequestTotal.WithLabelValues(c.Request.URL.Path,c.Request.Method,http.StatusText(c.Writer.Status())).Inc()
}
