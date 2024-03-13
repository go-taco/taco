package books

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/example/dtos"
	"github.com/yagobatista/taco-go-web-framework/example/models"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

type BookSerializer struct {
	serializers.ModelSerializer[
		dtos.BookFilter,
		dtos.BookCreatePayload,
		dtos.BookUpdatePayload,
		struct{},
		models.Book,
		dtos.BookResponse,
	]
}

func (this BookSerializer) CreateToModel(ctx context.Context, payload dtos.BookCreatePayload) (models.Book, error) {
	return models.Book{
		Title: payload.Title,
	}, nil
}

func (this BookSerializer) UpdateToModel(ctx context.Context, payload dtos.BookUpdatePayload) (models.Book, error) {
	return models.Book{
		Author: payload.Author,
	}, nil
}

func (this BookSerializer) ToResponse(ctx context.Context, instance models.Book) (dtos.BookResponse, error) {
	return dtos.BookResponse{
		ID:        instance.ID,
		CreatedAt: instance.CreatedAt,
		UpdatedAt: instance.UpdatedAt,
		Title:     instance.Title,
		Author:    instance.Author,
	}, nil
}
