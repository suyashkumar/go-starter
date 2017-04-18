package app

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/go-starter/handlers"
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

func New(c AppConfig) (App, error) {
	// Create and populate context with DB session:
	db, err := gorm.Open("mysql", c.DbURI)
	if err != nil {
		return nil, err
	}
	ctx := handlers.Context{DB: db}

	// Create router and attach routes
	r := httprouter.New()
	attachRoutes(r, &ctx)

	return &app{
		config: c,
		router: r,
	}, nil
}

func (a *app) Run() {
	if a.config.IsDev {
		fmt.Printf("Listening on port :%s\n", a.config.Port)
		err := http.ListenAndServe(":"+a.config.Port, a.router)
		panic(err)
	} else {
		err := http.ListenAndServeTLS(":443", a.config.CertKey, a.config.PrivKey, a.router)
		panic(err)
		//TODO(suyash): Add redirect to https
	}
}
