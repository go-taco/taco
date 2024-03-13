package books

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/example/dtos"
	"github.com/yagobatista/taco-go-web-framework/example/models"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

type CreateBookSerializer struct {
	baseBookSerializer
	serializers.CreateModelSerializer[dtos.BookCreatePayload, models.Book, dtos.BookResponse]
}

type UpdateBookSerializer struct {
	baseBookSerializer
	serializers.UpdateModelSerializer[dtos.BookCreatePayload, models.Book, dtos.BookResponse]
}

type DetailBookSerializer struct {
	baseBookSerializer
	serializers.DetailModelSerializer[dtos.BookFilter, models.Book, dtos.BookResponse]
}

type ListBookSerializer struct {
	baseBookSerializer
	serializers.ListModelSerializer[dtos.BookFilter, models.Book, dtos.BookResponse]
}

type baseBookSerializer struct{}

func (this baseBookSerializer) CreateToModel(ctx context.Context, payload dtos.BookCreatePayload) (models.Book, error) {
	return models.Book{
		Title: payload.Title,
	}, nil
}

func (this baseBookSerializer) UpdateToModel(ctx context.Context, payload dtos.BookUpdatePayload) (models.Book, error) {
	return models.Book{
		Author: payload.Author,
	}, nil
}

func (this baseBookSerializer) ToResponse(ctx context.Context, instance models.Book) (dtos.BookResponse, error) {
	return dtos.BookResponse{
		ID:        instance.ID,
		CreatedAt: instance.CreatedAt,
		UpdatedAt: instance.UpdatedAt,
		Title:     instance.Title,
		Author:    instance.Author,
	}, nil
}
