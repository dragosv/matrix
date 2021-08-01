package handlers

import "github.com/gin-gonic/gin"

func SetupMatrixHandlers(router *gin.Engine) {
	router.POST("/echo", EchoHandler)
	router.POST("/invert", InvertHandler)
	router.POST("/flatten", FlattenHandler)
	router.POST("/sum", SumHandler)
	router.POST("/multiply", MultiplyHandler)
}
