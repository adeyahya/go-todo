package repository

import (
	"database/sql"

	"github.com/adeyahya/go-todo/model"
	"github.com/google/uuid"
)

type Todo struct {
	Db *sql.DB
}

func (r *Todo) FindMany() ([]model.Todo, error) {
	rows, err := r.Db.Query("SELECT id, title, description, is_completed from todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.IsCompleted)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *Todo) GetById(id uuid.UUID) (*model.Todo, error) {
	row := r.Db.QueryRow("SELECT id, title, description, is_completed FROM todo where id = ?", id)
	var todo model.Todo

	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.IsCompleted)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *Todo) Insert(title string) (*model.Todo, error) {
	id := uuid.New()
	_, err := r.Db.Exec("INSERT INTO todo(id, title, is_completed) values(?, ?, ?)", id, title, false)

	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}
