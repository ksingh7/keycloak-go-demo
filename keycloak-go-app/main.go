package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

var router *gin.Engine

func goDotEnvVariables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Msgf("%v", "Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()
	initializeRoutes()

	router.Run("localhost:8081")
}
