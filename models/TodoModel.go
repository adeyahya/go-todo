package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (s Todo) New(title string) Todo {
	return Todo{
		Id:          uuid.NewString(),
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}
}
