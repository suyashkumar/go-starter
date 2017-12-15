package main

import (
	"github.com/suyashkumar/go-starter/app"
	"github.com/suyashkumar/go-starter/log"
)

func main() {
	log.Configure()
	app, err := app.New()
	if err != nil {
		panic(err)
	}
	app.Run()
}
