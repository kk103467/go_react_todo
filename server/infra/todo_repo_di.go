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
	todo1.Text = "learn about react"
	todo1.IsCompleted = false

	todo2 := model.Todo{}
	todo2.Id = 2
	todo2.Text = "meet friend for lunch"
	todo2.IsCompleted = true

	return []model.Todo{todo1, todo2}, nil
}