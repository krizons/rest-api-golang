package main

import (
	"log"

	"rest/internal/apiserver"

	"github.com/BurntSushi/toml"
)

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("./config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
}
