package service

import (
	"encoding/json"
	"net/http"

	"github.com/adeyahya/go-todo/repository"
)

type Todo struct {
	Repository *repository.Todo
}

func (s *Todo) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todos, _ := s.Repository.FindMany()
	s.Repository.Insert("Hello")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
