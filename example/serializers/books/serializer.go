package books

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

type CreateBookSerializer struct {
	baseBookSerializer
	serializers.CreateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
}

type UpdateBookSerializer struct {
	baseBookSerializer
	serializers.UpdateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
}

type DetailBookSerializer struct {
	baseBookSerializer
	serializers.DetailModelSerializer[structs.BookFilter, structs.Book, structs.BookResponse]
}

type ListBookSerializer struct {
	baseBookSerializer
	serializers.ListModelSerializer[structs.BookFilter, structs.Book, structs.BookResponse]
}

type baseBookSerializer struct{}

func (this baseBookSerializer) CreateToModel(ctx context.Context, payload structs.BookCreatePayload) (structs.Book, error) {
	return structs.Book{
		Title: payload.Title,
	}, nil
}

func (this baseBookSerializer) UpdateToModel(ctx context.Context, payload structs.BookUpdatePayload) (structs.Book, error) {
	return structs.Book{
		Author: payload.Author,
	}, nil
}

func (this baseBookSerializer) ToResponse(ctx context.Context, instance structs.Book) (structs.BookResponse, error) {
	return structs.BookResponse{
		ID:        instance.ID,
		CreatedAt: instance.CreatedAt,
		UpdatedAt: instance.UpdatedAt,
		Title:     instance.Title,
		Author:    instance.Author,
	}, nil
}
