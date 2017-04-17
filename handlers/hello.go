package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx *Context) {
	fmt.Fprintf(w, "Hello, world")
}
