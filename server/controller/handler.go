package presentation

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/usecase"
)

type textStruct struct {
	Text string `json:"text"`
}

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (th *todoHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	// handle POST request; INSERT
	headerContentType := r.Header.Get("Content-type")
	if headerContentType != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var t textStruct
	if err := dec.Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newTodo model.Todo
	newTodo.Text = t.Text
	newTodo.IsCompleted = false

	todos, err := th.Usecase_field.AddTodo(newTodo)
	if err != nil {
		log.Fatal(err)
	}

	// respond to the client; SELECT
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gorillllla"))
}
