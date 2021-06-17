package main

import (
	"log"
	"net/http"

	"github.com/kk103467/go_react_todo/server/router"
)

func main() {
	r := router.MyMux()

	log.Fatal(http.ListenAndServe(":8000", r))
}
