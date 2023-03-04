package books

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/example/structs"
)

type bookSerializer struct{}

func (this bookSerializer) ToModel(ctx context.Context, payload structs.BookCreatePayload) (structs.Book, error) {
	return structs.Book{
		Title: payload.Title,
	}, nil
}

func (this bookSerializer) ToResponse(ctx context.Context, instance structs.Book) (structs.BookResponse, error) {
	return structs.BookResponse{
		ID:        instance.ID,
		CreatedAt: instance.CreatedAt,
		UpdatedAt: instance.UpdatedAt,
		Title:     instance.Title,
	}, nil
}
