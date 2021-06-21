package router

import (
	"github.com/gin-gonic/gin"
	"monitoring-entry-task/api"
	"monitoring-entry-task/common"
	"monitoring-entry-task/middleware"
)

func InitRouter()  {
	common.Registry()
	r := gin.Default()
	r.Use(middleware.Report())
	LoadDefault(r)
	r.Run(":9094")
}

func LoadDefault(e * gin.Engine)  {
	e.GET("/hello", api.Hello)
	e.GET("/metrics", api.PrometheusHandler())
}