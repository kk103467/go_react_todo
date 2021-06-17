package router

import(
	"net/http"
	"github.com/gorilla/mux"

	"github.com/kk103467/go_react_todo/server/handlers"
)
func MyMux() http.Handler {
	r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler)
	return r
}