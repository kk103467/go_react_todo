package presentation

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MyMux(th TodoHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api", HomeHandler)
	r.HandleFunc("/api/view", th.ViewHandler)
	r.HandleFunc("/api/add", th.AddHandler)
	
	return r
}
