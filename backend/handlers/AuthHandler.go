package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	utils "github.com/adeyahya/go-todo/core"
	"github.com/adeyahya/go-todo/core/errors"
	"github.com/adeyahya/go-todo/models"
	"github.com/adeyahya/go-todo/repositories"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	*repositories.UserRepository
	*validator.Validate
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserRequestDTO
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	json.Unmarshal(reqBody, &userRequest)
	err = h.Validate.Struct(userRequest)
	if err != nil {
		errors.BadRequest(w, err)
		return
	}
	user, err := h.UserRepository.Create(userRequest.Email, userRequest.Name, userRequest.Password)
	userReponse := models.UserResponseDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	json.NewEncoder(w).Encode(userReponse)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequestDTO
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errors.InternalServerError(w, err)
		return
	}
	json.Unmarshal(reqBody, &loginRequest)
	err = h.Validate.Struct(loginRequest)
	if err != nil {
		errors.BadRequest(w, err)
		return
	}
	user, err := h.UserRepository.GetByEmail(loginRequest.Email)
	if err != nil || user == nil {
		errors.NotFoundError(w)
		return
	}
	isPasswordMatch := utils.ComparePassword(user.Password, loginRequest.Password)
	if !isPasswordMatch {
		errors.UnauthorizedError(w, "Invalid Password")
		return
	}
}
