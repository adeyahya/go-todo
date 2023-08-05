package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/adeyahya/go-todo/models"
	"github.com/adeyahya/go-todo/repositories"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	*repositories.TodoRepository
}

func (handler *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	limitString := queryParams.Get("limit")
	limit, _ := strconv.Atoi(limitString)
	if limit <= 0 {
		limit = 10
	}
	cursor := queryParams.Get("cursor")
	todoList := handler.TodoRepository.List(limit, cursor)
	json.NewEncoder(w).Encode(todoList)
}

func (handler *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todoRequest models.TodoRequestDTO
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &todoRequest)
	todo, _ := handler.TodoRepository.Create(todoRequest.Title)
	json.NewEncoder(w).Encode(todo)
}

func (handler *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo, _ := handler.TodoRepository.Get(id)
	json.NewEncoder(w).Encode(todo)
}
