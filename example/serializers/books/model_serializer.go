package books

import (
	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

type CreateBookSerializer struct {
	bookSerializer
	serializers.CreateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
}

type UpdateBookSerializer struct {
	bookSerializer
	serializers.UpdateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
}

type GetBookSerializer struct {
	bookSerializer
	serializers.GetModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
}
