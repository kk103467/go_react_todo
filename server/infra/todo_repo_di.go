package infra

import (
	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/domain/repository"
)

type todoRepo struct {}

func NewTodoRepo() repository.TodoRepo {
	return &todoRepo{}
}

func (tr *todoRepo) GetAll() ([]model.Todo, error) {
	todo1 := model.Todo{}
	todo1.Id = 1
	todo1.Text = "First todo"
	todo1.Completed = false

	todo2 := model.Todo{}
	todo2.Id = 2
	todo2.Text = "Second todo done"
	todo2.Completed = true

	return []model.Todo{todo1, todo2}, nil
}