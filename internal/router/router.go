package router

import (
	"github.com/gorilla/mux"
)

func Middleware() *mux.Router {
	router := mux.NewRouter()
	return router
}
