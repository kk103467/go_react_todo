package main

import (
	"log"
	"net/http"

	"github.com/kk103467/go_react_todo/server/infra"
	"github.com/kk103467/go_react_todo/server/usecase"
	"github.com/kk103467/go_react_todo/server/presentation"
)

func main() {
	todoRepo := infra.NewTodoRepo()
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := presentation.NewTodoHandler(todoUsecase)

	r := presentation.MyMux(todoHandler)
	
	log.Fatal(http.ListenAndServe(":3000", r))
}
