package repositories

import (
	"database/sql"

	"github.com/adeyahya/go-todo/models"
)

type TodoRepository struct {
	*sql.DB
}

func (r *TodoRepository) List(limit *int, cursor *string) models.Paginated[models.Todo] {
	rows, _ := r.Query("SELECT id, title, is_completed, created_at FROM todo")
	defer rows.Close()
	var todoList []models.Todo

	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.IsCompleted, nil)
		todoList = append(todoList, todo)
	}

	result := models.Paginated[models.Todo]{
		Cursor: nil,
		Data:   &todoList,
	}

	return result
}

func (r *TodoRepository) Get(id string) (*models.Todo, error) {
	todo := models.Todo{}
	return &todo, nil
}

func (r *TodoRepository) Create(title string) (*models.Todo, error) {
	todo := models.Todo.New(models.Todo{}, "Hello")
	return &todo, nil
}
