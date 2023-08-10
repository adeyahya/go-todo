package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/adeyahya/go-todo/core/errors"
	"github.com/adeyahya/go-todo/models"
	"github.com/adeyahya/go-todo/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	*repositories.TodoRepository
	*validator.Validate
}

func (handler *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	limitString := queryParams.Get("limit")
	limit, _ := strconv.Atoi(limitString)
	if limit <= 0 {
		limit = 10
	}
	cursor := queryParams.Get("cursor")
	todoList, err := handler.TodoRepository.List(limit, cursor)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(todoList)
}

func (handler *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todoRequest models.TodoRequestDTO
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &todoRequest)
	todo, err := handler.TodoRepository.Create(todoRequest.Id, todoRequest.Title)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (handler *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo, err := handler.TodoRepository.Get(id)
	if err != nil {
		errors.NotFoundError(w)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (handler *TodoHandler) Done(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todo, err := handler.TodoRepository.Get(id)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	isCompleted := true
	todo, err = handler.TodoRepository.Update(id, nil, &isCompleted)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (handler *TodoHandler) Undone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todo, err := handler.TodoRepository.Get(id)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	isCompleted := false
	todo, err = handler.TodoRepository.Update(id, nil, &isCompleted)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(todo)
}
