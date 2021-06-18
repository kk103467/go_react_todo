package presentation

import (
	"net/http"

	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/usecase"
)

type TodoHandler interface{
	ViewHandler(w http.ResponseWriter, r *http.Request) ([]model.Todo, error)
}

type todoHandler struct{
	Usecase_field usecase.TodoUsecase
}

func NewTodoHandler(tu usecase.TodoUsecase) TodoHandler {
	return &todoHandler{
		Usecase_field: tu,
	}
}

func (th *todoHandler) ViewHandler(w http.ResponseWriter, r *http.Request) ([]model.Todo, error) {
	todos, err := th.Usecase_field.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gorillllla"))
}
