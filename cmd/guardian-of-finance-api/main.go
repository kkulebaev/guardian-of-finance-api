package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"guardian-of-finance-api/internal/app/handler"
	"log"
)

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

func main() {
	// Создание и подключение к базе данных
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}

	handler.InitRoutes()
}
