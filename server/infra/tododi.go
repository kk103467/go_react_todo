package infra

import (
	"fmt"
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

func (tr *todoRepo) AddTodo(newTodo model.Todo) ([]model.Todo, error) {
	newTodo.Id = 2
	newTodo.Text = "learn about react"
	newTodo.IsCompleted = false

	query := fmt.Sprintf("INSERT INTO todo_table VALUES(%d, '%s', %t);", newTodo.Id, newTodo.Text, newTodo.IsCompleted)
	log.Println(query)

	rows, err := Database.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return []model.Todo{newTodo}, nil
}
