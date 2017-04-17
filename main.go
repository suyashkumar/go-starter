package main

import (
	"os"

	"github.com/suyashkumar/go-starter/app"
)

func main() {
	c := app.AppConfig{
		IsDev: true,
		Port:  os.Getenv("PORT"),
		DbURI: os.Getenv("DbURI"),
	}

	app, err := app.New(c)
	if err != nil {
		panic(err)
	}
	app.Run()
}
