package presentation

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MyMux(th TodoHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api", HomeHandler)
	r.HandleFunc("/api/todos", th.ViewHandler).Methods("GET")
	r.HandleFunc("/api/todo", th.AddHandler).Methods("GET")
	
	return r
}
