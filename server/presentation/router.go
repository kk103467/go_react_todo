package presentation

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MyMux(th TodoHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api", HomeHandler)
	r.HandleFunc("/api/view", th.ViewHandler)
	
	return r
}
