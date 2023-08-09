package models

import (
	"time"
)

type Todo struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	IsCompleted bool       `json:"isCompleted"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type TodoRequestDTO struct {
	Id    *string `json:"id"`
	Title string  `json:"title"`
}
