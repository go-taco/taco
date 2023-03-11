package structs

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title string `json:"title"`
}

type BookFilter struct {
	IdIN  *[]uint `query:"id__in"`
	Title *string `query:"title"`
}
