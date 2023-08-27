package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jackpot-mab/arm-selector/controller"
	"jackpot-mab/arm-selector/docs"
	"net/http"
)

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "jackpot-mab:arm-selector")
}

func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/arm")
		{
			eg.POST("/selection", controller.ArmSelectionController)
		}
	}

	router.GET("/", healthCheck)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8090")
}
