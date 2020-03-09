package main

import (
	"blog/middleware"
	"blog/routes"
	"blog/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func setupRouter() *gin.Engine {

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r = routes.Groups(r)

	front := r.Group("/front")
	{
		front.GET("ping", func(context *gin.Context) {
			helper.Success(context, 200, gin.H{"msg": "pong"})
		})
	}

	backend := r.Group("/backend")
	{
		backend.GET("ping", func(context *gin.Context) {
			helper.Success(context, 200, gin.H{"msg": "pong"})
		})
		routes.Backend(backend)
	}

	backend.Use(middleware.AccessTokenMiddleware())

	return r
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println(secretKey)

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
