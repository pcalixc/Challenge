package main

import (
	"backend/pkg/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loadingh .env file")
	}

	s := server.New()
	s.Run()
}
