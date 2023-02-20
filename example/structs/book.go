package structs

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title string `json:"title"`
}

type BookFilter struct {
	Title string `query:"title"`
}
