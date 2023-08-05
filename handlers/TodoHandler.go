package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adeyahya/go-todo/repositories"
)

type TodoHandler struct {
	*repositories.TodoRepository
}

func (handler *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	todoList := handler.TodoRepository.List(nil, nil)
	json.NewEncoder(w).Encode(todoList)
}
