package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"guardian-of-finance-api/internal/app/handler"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed loading environment variables")
	}
	handler.InitRoutes()
}
