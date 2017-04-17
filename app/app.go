package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App interface {
	Run()
}

type app struct {
	config AppConfig
	router *httprouter.Router
}

// AppConfig holds various configuration settings for the app to read on startup
type AppConfig struct {
	IsDev   bool
	Port    string
	CertKey string // Supplied if using SSL
	PrivKey string // Supplied if using SSL
	DbURI   string
}

func New(c AppConfig) App {
	r := httprouter.New()
	attachRoutes(r)

	return app{
		config: c,
		router: r,
	}
}

func (a *app) Run() {
	if a.config.IsDev {
		fmt.Printf("Listening on port :%s", a.config.Port)
		err := http.ListenAndServe(":"+a.config.Port, a.router)
		panic(err)
	} else {
		go http.ListenAndServeTLS(":443", a.config.CertKey, a.config.PrivKey, a.router)
	}
}
