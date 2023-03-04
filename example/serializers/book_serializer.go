package serializers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

type CreateBookSerializer struct {
	serializers.ModelSerializer[struct{}, structs.BookCreatePayload, structs.Book, structs.BookResponse]
}

func (this CreateBookSerializer) ToModel(ctx context.Context, payload structs.BookCreatePayload) (structs.Book, error) {
	return structs.Book{
		Title: payload.Title,
	}, nil
}

func (this CreateBookSerializer) ToResponse(ctx context.Context, instance structs.Book) (structs.BookResponse, error) {
	return structs.BookResponse{
		ID:        instance.ID,
		CreatedAt: instance.CreatedAt,
		UpdatedAt: instance.UpdatedAt,
		Title:     instance.Title,
	}, nil
}
