package book

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
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
	serializerModelGroup := d.GetRouter(routes.INTERNAL).Group("book-serializer-model")
	serializerGroup := d.GetRouter(routes.INTERNAL).Group("book-serializer")

	router.SetPost(
		this.BaseHandler,
		group,
		":id/buy",
		"Buy book",
		this.BuyBook,
		handlers.WithTransaction(false),
		handlers.WithRequestMiddlewares(
			basicauth.New(basicauth.Config{
				Users: map[string]string{"admin2": "admin2"},
			}),
		),
	)

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
		"Book serializer create",
		books.CreateBookSerializer{},
	)
	router.NewUpdateModelSerializerHandler[structs.BookUpdatePayload, structs.Book, structs.BookResponse](
		this.BaseHandler,
		serializerGroup,
		"",
		"Book serializer update",
		books.UpdateBookSerializer{},
	)
	router.NewDetailModelSerializerHandler[structs.BookFilter, structs.Book, structs.BookResponse](
		this.BaseHandler,
		serializerGroup,
		"",
		"Book serializer detail",
		books.DetailBookSerializer{},
	)
	router.NewListModelSerializerHandler[structs.BookFilter, structs.Book, structs.BookResponse](
		this.BaseHandler,
		serializerGroup,
		"",
		"Book serializer list",
		books.ListBookSerializer{},
	)
	router.NewModelSerializerHandler[structs.BookCreatePayload, structs.BookUpdatePayload, struct{}, structs.BookFilter, structs.Book, structs.BookResponse](
		this.BaseHandler,
		serializerModelGroup,
		"",
		"Book serializer model",
		books.BookSerializer{},
	)
}
