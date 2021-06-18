package presentation

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MyMux(th TodoHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/view", th.ViewHandler)
	
	return r
}
