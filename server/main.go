package main

import (
	"log"
	"net/http"

	// "github.com/kk103467/go_react_todo/server/infra"
	"github.com/kk103467/go_react_todo/server/presentation"
)

func main() {
	r := presentation.MyMux()

	// todoRepo := infra.NewTodoRepo()
	// todoUsecase := usecase.NewTodoUsecase(todoRepo)
	// todoHandler := presentation.NewTodoHandler(todoUsecase)
	

	log.Fatal(http.ListenAndServe(":8000", r))
}
