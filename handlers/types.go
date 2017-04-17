package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

type Handler func(
	http.ResponseWriter,
	*http.Request,
	httprouter.Params,
	*Context,
)

type Context struct {
	DB *gorm.DB
}
