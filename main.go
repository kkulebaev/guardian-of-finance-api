package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type ICategory struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type IUser struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
}

type IOperation struct {
	User     IUser     `json:"user"`
	Month    string    `json:"month"`
	Category ICategory `json:"category"`
	Sum      float32   `json:"sum"`
}

var (
	operations = []IOperation{
		{
			User: IUser{
				ID:   1,
				Name: "Konstantin",
			},
			Month: "2023-01-17T14:23:58.911Z",
			Category: ICategory{
				ID:    "transport",
				Label: "Транспорт",
			},
			Sum: 20000,
		},
	}
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.Print("Server is starting...")

	router := gin.New()

	router.GET("/costs", costsHandler)
	router.POST("/costs", postOperation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func costsHandler(c *gin.Context) {
	var costs = getCosts()

	c.IndentedJSON(http.StatusOK, costs)
}

func getCosts() []IOperation {
	return operations
}

func postOperation(c *gin.Context) {
	var newOperation IOperation

	err := c.BindJSON(&newOperation)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	operations = append(operations, newOperation)
	c.IndentedJSON(http.StatusCreated, newOperation)
}
