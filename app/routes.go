package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/go-starter/handlers"
)

func attachRoutes(r *httprouter.Router) {
	r.GET("/", handlers.Index)
	r.GET("/hello", handlers.Hello)

	r.ServeFiles("/static/*filepath", http.Dir("public/static"))
}
