package main

import (
	"os"

	"github.com/suyashkumar/go-starter/app"
)

func main() {
	c := app.AppConfig{
		IsDev: true,
		Port:  os.Getenv("PORT"),
	}

	app := app.New(c)
	app.Run()
}
