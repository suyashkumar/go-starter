package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/suyashkumar/go-starter/handlers"
)

func injectContext(h handlers.Handler, ctx *handlers.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h(w, r, ps, ctx)
	}
}

func attachRoutes(r *httprouter.Router, ctx *handlers.Context) {
	r.GET("/hello", injectContext(handlers.Hello, ctx))
}
