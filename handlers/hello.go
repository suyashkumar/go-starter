package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Hello, world")
}
