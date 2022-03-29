package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()
	initializeRoutes()

	router.Run("localhost:8081")
}
