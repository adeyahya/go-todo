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

func (r *TodoRepository) List(limit int, cursor string) (*models.Paginated[models.Todo], error) {
	var rows *sql.Rows
	var err error
	if cursor == "" {
		rows, err = r.Query(`
			SELECT id, title, is_completed, created_at FROM todo
			ORDER BY created_at DESC
			LIMIT $1
		`, limit)
	} else {
		rows, err = r.Query(`
			SELECT id, title, is_completed, created_at FROM todo
			WHERE created_at < $1
			ORDER BY created_at DESC
			LIMIT $2
		`, cursor, limit)
	}

	if err != nil {
		return nil, err
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

	return &result, nil
}

func (r *TodoRepository) Get(id string) (*models.Todo, error) {
	row := r.QueryRow(`
		SELECT id, title, is_completed, created_at FROM todo where id = $1
	`, id)
	var todo models.Todo

	err := row.Scan(&todo.Id, &todo.Title, &todo.IsCompleted, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Create(title string) (*models.Todo, error) {
	id := uuid.NewString()
	_, err := r.Exec(`
		INSERT INTO todo(id, title, is_completed, created_at)
		values($1, $2, $3, $4)`,
		id, title, false, time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *TodoRepository) Update(id string, title *string, isCompleted *bool) (*models.Todo, error) {
	todo, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	if title != nil {
		todo.Title = *title
	}
	if isCompleted != nil {
		todo.IsCompleted = *isCompleted
	}
	_, err = r.Exec(`
		UPDATE todo SET title = $2, is_completed = $3
		WHERE id = $1`,
		todo.Id, todo.Title, todo.IsCompleted,
	)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
