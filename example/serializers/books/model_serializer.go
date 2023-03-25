package books

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

type BookSerializer struct {
	serializers.ModelSerializer[
		structs.BookFilter,
		structs.BookCreatePayload,
		structs.BookUpdatePayload,
		struct{},
		structs.Book,
		structs.BookResponse,
	]
}

func (this BookSerializer) CreateToModel(ctx context.Context, payload structs.BookCreatePayload) (structs.Book, error) {
	return structs.Book{
		Title: payload.Title,
	}, nil
}

func (this BookSerializer) UpdateToModel(ctx context.Context, payload structs.BookUpdatePayload) (structs.Book, error) {
	return structs.Book{
		Author: payload.Author,
	}, nil
}

func (this BookSerializer) ToResponse(ctx context.Context, instance structs.Book) (structs.BookResponse, error) {
	return structs.BookResponse{
		ID:        instance.ID,
		CreatedAt: instance.CreatedAt,
		UpdatedAt: instance.UpdatedAt,
		Title:     instance.Title,
		Author:    instance.Author,
	}, nil
}
