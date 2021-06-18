package usecase

import (
	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/domain/repository"
)

type TodoUsecase interface{
	GetAll() ([]model.Todo, error)
}

type todoUsecase struct{
	Repo repository.TodoRepo
}

func NewTodoUsecase(tr repository.TodoRepo) TodoUsecase {
	return &todoUsecase{
		Repo: tr,
	}
}

func (tu *todoUsecase) GetAll() ([]model.Todo, error) {
	todos, err := tu.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}