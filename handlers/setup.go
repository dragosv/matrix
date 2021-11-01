package handlers

import "github.com/gin-gonic/gin"
import "github.com/dragosv/matrix/commands"
import "github.com/dragosv/matrix/processing"

func SetupMatrixHandlers(router *gin.Engine) {
	controller := MatrixController{operations: &commands.Matrix{}, reader: &processing.Matrix{}}

	router.POST("/echo", controller.EchoHandler)
	router.POST("/invert", controller.InvertHandler)
	router.POST("/flatten", controller.FlattenHandler)
	router.POST("/sum", controller.SumHandler)
	router.POST("/multiply", controller.MultiplyHandler)
}
