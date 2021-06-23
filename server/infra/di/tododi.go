package infra

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kk103467/go_react_todo/server/domain/model"
	"github.com/kk103467/go_react_todo/server/domain/repository"
)

type todoRepo struct{}

func NewTodoRepo() repository.TodoRepo {
	return &todoRepo{}
}

func (tr *todoRepo) GetAll() ([]model.Todo, error) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("select * from todo_table;")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			text string
			is_completed bool
		)
		if err := rows.Scan(&id, &text, &is_completed); err != nil{
			log.Fatal(err)
		}
		log.Printf("id->%d, text->%s, is_completed->%t\n", id, text, is_completed)
	}

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
