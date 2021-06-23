package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/kk103467/go_react_todo/server/infra/di"
	"github.com/kk103467/go_react_todo/server/presentation"
	"github.com/kk103467/go_react_todo/server/usecase"
)

func main() {
	todoRepo := di.NewTodoRepo()
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := presentation.NewTodoHandler(todoUsecase)

	r := presentation.MyMux(todoHandler)
	corsHandler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8000", corsHandler))
}
