package presentation

import (
	"net/http"

	"github.com/kk103467/go_react_todo/server/usecase"
)

type TodoHandler interface{
	ViewHandler(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct{
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
		return
	}
	w.Write([]byte(todos[0].Text))
	return
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gorillllla"))
}
