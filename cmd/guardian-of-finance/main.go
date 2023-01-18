package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"guardian-of-finance-api/internal/app/service"
	"log"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.Print("Server is starting...")

	router := gin.New()

	router.Use(cors.Default())

	router.GET("/costs", service.CostsHandler)
	router.POST("/costs", service.PostOperation)
	router.DELETE("/costs/:id", service.DeleteOperation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
