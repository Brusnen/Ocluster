package server

import (
	"OCluster/server/runners"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	runners.RunnerRouter(router)
	return router
}

func RunServer() {
	InitRoutes().Run(":8000")
}
