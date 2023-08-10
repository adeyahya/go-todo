package models

import "time"

type User struct {
	Id               string
	Name             string
	Email            string
	Password         string
	IsEmailConfirmed bool
	EmailConfirmedAt *time.Time
	CreatedAt        time.Time
}

type UserResponseDTO struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=72"`
}

type LoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=72"`
}
