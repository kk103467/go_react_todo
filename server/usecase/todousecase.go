package usecase

import (
	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/domain/repository"
)

type TodoUsecase interface{
	GetAll() ([]model.Todo, error)
	AddTodo() ([]model.Todo, error)
}

type todoUsecase struct{
	Repo_field repository.TodoRepo
}

func NewTodoUsecase(tr repository.TodoRepo) TodoUsecase {
	return &todoUsecase{
		Repo_field: tr,
	}
}

func (tu *todoUsecase) GetAll() ([]model.Todo, error) {
	todos, err := tu.Repo_field.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tu *todoUsecase) AddTodo() ([]model.Todo, error) {
	todos, err := tu.Repo_field.AddTodo()
	if err != nil {
		return nil, err
	}
	return todos, nil
}