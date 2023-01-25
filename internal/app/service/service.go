package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
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

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

var cfg = Config{
	Host:     "localhost",
	Port:     "5436",
	Username: "postgres",
	DBName:   "postgres",
	SSLMode:  "disable",
	Password: "qwerty123",
}

const (
	operationTable = "operations"
)

var operations []IOperationDB

func GetListOperation(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, operations)
}

func CreateOperation(c *gin.Context) {
	log.Print("test")
	var newOperation IOperation

	err := c.BindJSON(&newOperation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}
	log.Print("SUCCESS")
	query := fmt.Sprintf("INSERT INTO %s (user_id, month_date, category_id, total) values ($1, $2, $3, $4) RETURNING id;", operationTable)
	log.Print("query: ", query)
	row := db.QueryRow(query, newOperation.User.ID, newOperation.Month, newOperation.Category.ID, newOperation.Sum)
	var id int
	if err := row.Scan(&id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request2"})
	}
	log.Print("id: ", id)

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
