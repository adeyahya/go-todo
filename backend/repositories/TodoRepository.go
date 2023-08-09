package repositories

import (
	"database/sql"
	"time"

	utils "github.com/adeyahya/go-todo/core"
	"github.com/adeyahya/go-todo/models"
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
		SELECT id, title, is_completed, completed_at, created_at FROM todo where id = $1
	`, id)
	var todo models.Todo

	err := row.Scan(&todo.Id, &todo.Title, &todo.IsCompleted, &todo.CompletedAt, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Create(id *string, title string) (*models.Todo, error) {
	var err error
	var _id string

	if id == nil {
		_id, err = utils.GenerateId()
	} else {
		_id = *id
	}

	if err != nil {
		return nil, err
	}
	_, err = r.Exec(`
		INSERT INTO todo(id, title, is_completed, created_at)
		values($1, $2, $3, $4)`,
		_id, title, false, time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return r.Get(_id)
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

	if *isCompleted == true {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}
	_, err = r.Exec(`
		UPDATE todo SET title = $2, is_completed = $3, completed_at = $4
		WHERE id = $1`,
		todo.Id, todo.Title, todo.IsCompleted, todo.CompletedAt,
	)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
