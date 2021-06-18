package main

import (
	"log"
	"net/http"

	// "github.com/kk103467/go_react_todo/server/infra"
	"github.com/kk103467/go_react_todo/server/presentation/router"
)

func main() {
	r := router.MyMux()

	// todoRepo := infra.NewTodoRepo()

	

	log.Fatal(http.ListenAndServe(":8000", r))
}
