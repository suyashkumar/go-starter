package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/suyashkumar/go-starter/config"
	"github.com/suyashkumar/go-starter/log"
	"github.com/suyashkumar/go-starter/routes"
)

func main() {
	log.Configure()

	r := routes.Build()

	p := fmt.Sprintf(":%s", config.Get(config.Port))

	// TODO: clean up config bool checks
	if config.Get(config.IsDev) == "true" {
		// Serve without SSL
		logrus.WithField("port", p).Info("Serving without SSL")
		err := http.ListenAndServe(p, r)
		logrus.Fatal(err)
	} else {
		// Serve with SSL
		logrus.Info("Serving with SSL")
		err := http.ListenAndServeTLS(
			p,
			config.Get(config.CertKey),
			config.Get(config.PrivKey),
			r,
		)
		// TODO: reroute http requests to https
		logrus.Fatal(err)
	}

}
