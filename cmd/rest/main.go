package rest

import (
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	config := rest.NewConfig()
	_, err := toml.DecodeFile("./config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
}
