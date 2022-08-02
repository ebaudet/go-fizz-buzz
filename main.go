package main

import (
	"log"

	"github.com/ebaudet/go-fizz-buzz/api"
	"github.com/ebaudet/go-fizz-buzz/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server := api.NewServer()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
