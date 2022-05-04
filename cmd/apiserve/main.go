package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/krizons/rest-api-golang/internal/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("./config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	/*config := apiserver.NewConfig()
	_, err := toml.DecodeFile("./config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}*/
}
