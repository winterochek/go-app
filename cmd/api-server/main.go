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

	c := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(c); err != nil {
		log.Fatal(err)
	}
}
