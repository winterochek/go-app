package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	apiserver "github.com/winterochek/go-app/internal/app/api-server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api-server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
