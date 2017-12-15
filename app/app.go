package app

import (
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/go-starter/config"
)

type App interface {
	Run()
}

type app struct {
	router *httprouter.Router
}

func New() (App, error) {
	// Create router and attach routes
	r := httprouter.New()
	attachRoutes(r)

	return &app{
		router: r,
	}, nil
}

func (a *app) Run() {
	if config.Get(config.IsDev) != "" {
		fmt.Printf("Listening on port :%s\n", config.Get(config.Port))
		err := http.ListenAndServe(":"+config.Get(config.Port), a.router)
		panic(err)
	} else {
		err := http.ListenAndServeTLS(":443", config.Get(config.CertKey), config.Get(config.PrivKey), a.router)
		panic(err)
		//TODO(suyash): Add redirect to https
	}
}
