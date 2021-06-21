package main

import (
	"monitoring-entry-task/common"
	"monitoring-entry-task/router"
)

func main()  {
	common.Registry()
	router.InitRouter()
}