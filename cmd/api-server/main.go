package main

import (
	"log"

	apiserver "github.com/winterochek/go-app/internal/app/api-server"
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
