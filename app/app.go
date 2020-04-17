package app

import (
	"github.com/gorilla/mux"
	"go-http-example/app/handler"
)

func LoadRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", handler.GetUser)
	return r
}
