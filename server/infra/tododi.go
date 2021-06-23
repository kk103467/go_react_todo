package infra

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/domain/repository"
)

type todoRepo struct{}

func NewTodoRepo() repository.TodoRepo {
	return &todoRepo{}
}

func (tr *todoRepo) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	rows, err := Database.Query("select * from todo_table;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Id, &todo.Text, &todo.IsCompleted); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	// To be modified; error handling
	return todos, nil
}

func (tr *todoRepo) AddTodo() ([]model.Todo, error) {
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
