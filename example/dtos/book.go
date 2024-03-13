package dtos

import "time"

type BookCreatePayload struct {
	Title string `json:"title"`
}

type BookUpdatePayload struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
}

type BookFilter struct {
	IdIN  *[]uint `query:"id__in"`
	Title *string `query:"title"`
}
