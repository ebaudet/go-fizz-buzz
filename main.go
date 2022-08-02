package main

import (
	"log"

	"github.com/ebaudet/go-fizz-buzz/api"
)

const (
	serverAddress = "localhost:8080"
)

func main() {
	server := api.NewServer()

	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
