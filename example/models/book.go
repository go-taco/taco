package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`

	AvailableCopies uint64 `json:"available_copies"`
}
