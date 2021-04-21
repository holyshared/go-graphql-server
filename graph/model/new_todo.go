package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}
