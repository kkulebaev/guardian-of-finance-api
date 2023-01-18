package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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

var operations []IOperationDB

func CostsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, operations)
}

func PostOperation(c *gin.Context) {
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
	c.IndentedJSON(http.StatusCreated, newOperationDB)
}

func DeleteOperation(c *gin.Context) {
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
