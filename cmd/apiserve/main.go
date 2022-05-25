package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/krizons/rest-api-golang/internal/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("../../config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		s.Close()
		log.Fatal(err)
	}
	/*config := apiserver.NewConfig()
	_, err := toml.DecodeFile("./config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}*/
}
