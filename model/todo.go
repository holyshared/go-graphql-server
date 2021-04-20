package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}
