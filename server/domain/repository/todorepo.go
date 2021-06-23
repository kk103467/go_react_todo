package repository

import "github.com/kk103467/go_react_todo/server/domain/model"

type TodoRepo interface {
	GetAll()([]model.Todo, error)
	AddTodo()([]model.Todo, error)
}