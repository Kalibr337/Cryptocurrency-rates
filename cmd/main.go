package main

import (
	"log"

	todo "github.com/Kalibr337/Cryptocurrency-rates.git"
	"github.com/Kalibr337/Cryptocurrency-rates.git/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoures()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
