package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"monitoring-entry-task/api"
)

func InitRouter()  {
	r := gin.Default()
	LoadDefault(r)

	r.Run(":9094")
}

func LoadDefault(e * gin.Engine)  {
	e.GET("/hello", api.Hello)
	e.GET("/metrics", gin.WrapF(promhttp.Handler().ServeHTTP))
}