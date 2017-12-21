package app

import (
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/suyashkumar/go-starter/config"
)

type App interface {
	Run()
}

type app struct {
	router *httprouter.Router
}

// New creates and initializes a new web service App
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
		p := config.Get(config.Port)
		logrus.WithField("port", p).Info("Server listening")
		err := http.ListenAndServe(fmt.Sprintf(":%s", p), a.router)
		panic(err)
	} else {
		logrus.Info("Listening on :443")
		err := http.ListenAndServeTLS(":443", config.Get(config.CertKey), config.Get(config.PrivKey), a.router)
		panic(err)
		//TODO(suyash): Add redirect to https
	}
}
