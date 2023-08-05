package repositories

import (
	"database/sql"
	"time"

	"github.com/adeyahya/go-todo/models"
	"github.com/google/uuid"
)

type TodoRepository struct {
	*sql.DB
}

func (r *TodoRepository) List(limit int, cursor string) models.Paginated[models.Todo] {
	var rows *sql.Rows
	if cursor == "" {
		rows, _ = r.Query(`
			SELECT id, title, is_completed, created_at FROM todo
			ORDER BY created_at DESC
			LIMIT $1
		`, limit)
	} else {
		rows, _ = r.Query(`
			SELECT id, title, is_completed, created_at FROM todo
			WHERE created_at < $1
			ORDER BY created_at ASC
			LIMIT $2
		`, cursor, limit)
	}

	defer rows.Close()
	var todoList []models.Todo

	lastCreatedAt := time.Now()
	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.IsCompleted, &todo.CreatedAt)
		todoList = append(todoList, todo)
		lastCreatedAt = todo.CreatedAt
	}

	result := models.Paginated[models.Todo]{
		Cursor: lastCreatedAt.Format("2006-01-02T15:04:05.000000000Z"),
		Data:   &todoList,
	}

	return result
}

func (r *TodoRepository) Get(id string) (*models.Todo, error) {
	row := r.QueryRow("SELECT id, title, is_completed, created_at FROM todo where id = ?", id)
	var todo models.Todo

	err := row.Scan(&todo.Id, &todo.Title, &todo.IsCompleted, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Create(title string) (*models.Todo, error) {
	id := uuid.NewString()
	_, err := r.Exec("INSERT INTO todo(id, title, is_completed, created_at) values(?, ?, ?, ?)", id, title, false, time.Now())

	if err != nil {
		return nil, err
	}

	return r.Get(id)
}
