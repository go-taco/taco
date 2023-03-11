package books

import (
	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

type BookSerializer struct {
	bookSerializer
	serializers.CreateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
	serializers.UpdateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
	serializers.DetailModelSerializer[structs.BookFilter, structs.Book, structs.BookResponse]
	serializers.ListModelSerializer[structs.BookFilter, structs.Book, structs.BookResponse]
}

type CreateBookSerializer struct {
	bookSerializer
	serializers.CreateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
}

type UpdateBookSerializer struct {
	bookSerializer
	serializers.UpdateModelSerializer[structs.BookCreatePayload, structs.Book, structs.BookResponse]
}

type DetailBookSerializer struct {
	bookSerializer
	serializers.DetailModelSerializer[structs.BookFilter, structs.Book, structs.BookResponse]
}

type ListBookSerializer struct {
	bookSerializer
	serializers.ListModelSerializer[structs.BookFilter, structs.Book, structs.BookResponse]
}
