package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// ArmSelectionController godoc
// @Summary select arm
// @Schemes
// @Description select arm
// @Tags arm-selector
// @Accept json
// @Produce json
// @Success 200 {string} Arm
// @Router /select [get]
func ArmSelectionController(g *gin.Context) {
	g.JSON(http.StatusOK, "test")
}
