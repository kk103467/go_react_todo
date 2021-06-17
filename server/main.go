package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("gorillllla"))
}

func main()  {
	r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)

	// Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}