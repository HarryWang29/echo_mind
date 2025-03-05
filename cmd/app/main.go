package main

import (
	"github.com/HarryWang29/echo_mind/internal/infra/di"
	"log"
)

func main() {
	app, err := di.InjectAll()
	if err != nil {
		log.Fatal(err)
	}
	err = app.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
