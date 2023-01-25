package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"guardian-of-finance-api/internal/app/service"
	"log"
	"os"
)

func InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	log.Print("Server is starting...")

	router := gin.New()

	router.Use(cors.Default())

	router.GET("/costs", service.GetListOperation)
	router.POST("/costs", service.CreateOperation)
	router.DELETE("/costs/:id", service.DeleteOperation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

	return router
}
