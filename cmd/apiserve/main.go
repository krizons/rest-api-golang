package main

import (
	"log"

	"github.com/krizons/rest-api-golang/internal/apiserver"

	"github.com/BurntSushi/toml"
)

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("./config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
}
