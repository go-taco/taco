package book

import (
	"github.com/yagobatista/taco-go-web-framework/example/routes"
	"github.com/yagobatista/taco-go-web-framework/example/serializers/books"
	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/route"
	"github.com/yagobatista/taco-go-web-framework/src/router"
)

type BookHandler struct {
	handlers.BaseHandler
}

func (this *BookHandler) Routes(d route.Dispatcher) {
	group := d.GetRouter(routes.INTERNAL).Group("book")
	serializerGroup := d.GetRouter(routes.INTERNAL).Group("book-serializer")

	router.NewModelHandler[structs.BookFilter, structs.Book](
		this.BaseHandler,
		group,
		"",
		"Book model",
	)

	router.NewCreateModelSerializerHandler[structs.BookCreatePayload, structs.Book, structs.BookResponse](
		this.BaseHandler,
		serializerGroup,
		"",
		"Book serializer model",
		books.CreateBookSerializer{},
	)
	router.NewUpdateModelSerializerHandler[structs.BookCreatePayload, structs.Book, structs.BookResponse](
		this.BaseHandler,
		serializerGroup,
		"",
		"Book serializer model",
		books.UpdateBookSerializer{},
	)
	router.NewGetModelSerializerHandler[structs.BookCreatePayload, structs.Book, structs.BookResponse](
		this.BaseHandler,
		serializerGroup,
		"",
		"Book serializer model",
		books.GetBookSerializer{},
	)
}
