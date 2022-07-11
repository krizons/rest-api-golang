package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/krizons/rest-api-golang/internal/apiserver"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("../../config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(ctx); err != nil && err != http.ErrServerClosed {

		log.Fatal(err)
	}

	/*config := apiserver.NewConfig()
	_, err := toml.DecodeFile("./config/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}*/
}
