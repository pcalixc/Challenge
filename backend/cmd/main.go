package main

import (
	"backend/pkg/server"
	"log"
	"os"
	"strconv"

	//"net/http/pprof"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loadingh .env file")
	}

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error loadingg .env file")
	}
	s := server.New(server.Options{
		Host: host,
		Port: port,
	})
	s.Start()
}
