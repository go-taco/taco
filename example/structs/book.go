package structs

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Title string `json:"title"`
}

type BookCreatePayload struct {
	Title string `json:"title"`
}

type BookResponse struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `json:"title"`
}

type BookFilter struct {
	IdIN  *[]uint `query:"id__in"`
	Title *string `query:"title"`
}
