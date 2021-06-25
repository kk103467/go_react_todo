package main

import (
	"log"
	"net/http"

	"github.com/kk103467/go_react_todo/server/controller"
	"github.com/kk103467/go_react_todo/server/infrastructure"
	"github.com/kk103467/go_react_todo/server/usecase"
	"github.com/rs/cors"
)

func main() {
	infrastructure.ConnectDB()
	defer infrastructure.Database.Close()

	todoRepo := infrastructure.NewTodoRepo()
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := controller.NewTodoHandler(todoUsecase)

	r := controller.MyMux(todoHandler)
	corsHandler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8000", corsHandler))
}
