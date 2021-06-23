package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/usecase"
)

type TodoHandler interface {
	ViewHandler(w http.ResponseWriter, r *http.Request)
	AddHandler(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	Usecase_field usecase.TodoUsecase
}

func NewTodoHandler(tu usecase.TodoUsecase) TodoHandler {
	return &todoHandler{
		Usecase_field: tu,
	}
}

func (th *todoHandler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := th.Usecase_field.GetAll()
	
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(todos)
}

func (th *todoHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo model.Todo
	todos, err := th.Usecase_field.AddTodo(newTodo)
	
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(todos)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gorillllla"))
}
