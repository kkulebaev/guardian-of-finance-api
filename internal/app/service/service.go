package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"guardian-of-finance-api/internal/app/database"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ICategory struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

type IUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type IOperation struct {
	User     IUser     `json:"user"`
	Month    string    `json:"month"`
	Category ICategory `json:"category"`
	Sum      float64   `json:"sum"`
}

type IOperationDB struct {
	ID       int       `json:"id"`
	User     IUser     `json:"user"`
	Month    string    `json:"month"`
	Category ICategory `json:"category"`
	Sum      float64   `json:"sum"`
}

const (
	operationTable = "operations"
)

var operations []IOperationDB

func GetListOperation(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, operations)
}

func CreateOperation(c *gin.Context) {
	var newOperation IOperation
	var id int

	var cfg = database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  "require",
	}

	err := c.BindJSON(&newOperation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "incorrect post body"})
		return
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, month_date, category_id, total) values ($1, $2, $3, $4) RETURNING id;", operationTable)
	row := db.QueryRow(query, newOperation.User.ID, newOperation.Month, newOperation.Category.ID, newOperation.Sum)
	if err := row.Scan(&id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "incorrect db query"})
	}

	log.Print("created operation id: ", id)
	c.IndentedJSON(http.StatusCreated, row)
}

func DeleteOperation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("failed to parsing id when removing operation: %s", err)
	}

	for i, operation := range operations {
		if operation.ID == id {
			operations = append(operations[:i], operations[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, operation)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "operation not found"})
}
