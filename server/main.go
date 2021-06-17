package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kk103467/go_react_todo/server/handlers"
)

func main()  {
	r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler)

	// Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}