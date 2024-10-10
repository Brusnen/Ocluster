package runners

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunnerRouter(router *gin.Engine) {
	runnerRouter := router.Group("/api/runners")
	{
		runnerRouter.GET("/register_runner", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "ok"})
		})
	}
}
