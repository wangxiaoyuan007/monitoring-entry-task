package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"monitoring-entry-task/common"
	"net/http"
)
func Hello(ctx *gin.Context) {
	name := ctx.Query("name")
	path := ctx.Request.URL.Path
	if name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNoContent,
			"message": "fail",
		})
		common.HttpRequestsErrorTotal.WithLabelValues(string(http.StatusNoContent), path).Inc()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Hello! " + name,
	})
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
