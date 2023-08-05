package models

type Paginated[T any] struct {
	Cursor *string `json:"cursor"`
	Data   *[]T    `json:"data"`
}
