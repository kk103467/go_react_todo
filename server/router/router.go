package router

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/kk103467/go_react_todo/server/handlers"
)

func MyMux() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	return r
}
