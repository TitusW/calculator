package routes

import (
	controller "golang-calculator/controllers"

	"github.com/gin-gonic/gin"
)

func CalculatorRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/tambah", controller.Tambah())
	incomingRoutes.POST("/kurang", controller.Kurang())
	incomingRoutes.POST("/kali", controller.Kali())
	incomingRoutes.POST("/bagi", controller.Bagi())
}
