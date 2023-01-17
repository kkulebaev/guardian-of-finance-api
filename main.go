package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	Sum      float64   `json:"sum"`
}

type IOperationDB struct {
	ID       string    `json:"id"`
	User     IUser     `json:"user"`
	Month    string    `json:"month"`
	Category ICategory `json:"category"`
	Sum      float64   `json:"sum"`
}

var (
	operations = []IOperationDB{
		{
			ID: uuid.New().String(),
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

	router.Use(cors.Default())

	router.GET("/costs", costsHandler)
	router.POST("/costs", postOperation)
	router.DELETE("/costs/:id", deleteOperation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func costsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, operations)
}

func postOperation(c *gin.Context) {
	var newOperation IOperation

	err := c.BindJSON(&newOperation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	newOperationDB := IOperationDB{
		ID:       uuid.New().String(),
		User:     newOperation.User,
		Month:    newOperation.Month,
		Category: newOperation.Category,
		Sum:      newOperation.Sum,
	}

	operations = append(operations, newOperationDB)
	c.IndentedJSON(http.StatusCreated, newOperation)
}

func deleteOperation(c *gin.Context) {
	id := c.Param("id")

	for i, operation := range operations {
		if operation.ID == id {
			operations = append(operations[:i], operations[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, operation)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "operation not found"})
}
