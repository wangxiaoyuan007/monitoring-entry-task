package middleware

import (
	"github.com/gin-gonic/gin"
	"monitoring-entry-task/common"
	"time"
)


func Report() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		now := time.Now()
		path := ctx.Request.URL.Path
		//上报请求次数
		common.HttpRequestsTotal.WithLabelValues("200", path).Inc()

		//执行请求方法
		ctx.Next()

		//上报请求时延
		common.HttpRequestDuration.WithLabelValues("200", path).Observe(time.Since(now).Seconds())
	}
}
