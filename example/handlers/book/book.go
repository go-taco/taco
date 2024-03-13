package book

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/yagobatista/taco-go-web-framework/example/dtos"
	"github.com/yagobatista/taco-go-web-framework/example/models"
	"github.com/yagobatista/taco-go-web-framework/example/routes"
	"github.com/yagobatista/taco-go-web-framework/example/serializers/books"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/route"
	"github.com/yagobatista/taco-go-web-framework/src/router"
)

type BookHandler struct {
}

func (this *BookHandler) Routes(d route.Dispatcher) {
	group := d.GetRouter(routes.INTERNAL).Group("book")
	serializerModelGroup := d.GetRouter(routes.INTERNAL).Group("book-serializer-model")
	serializerGroup := d.GetRouter(routes.INTERNAL).Group("book-serializer")

	router.SetPost(
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
		handlers.WithResponseMiddlewares(func(c *fiber.Ctx) error {
			log.Printf("response status %d", c.Response().StatusCode())
			return nil
		}, func(c *fiber.Ctx) error {
			log.Printf("response body %s", c.Response().Body())
			return nil
		},
		),
	)

	router.NewModelHandler[dtos.BookFilter, models.Book](
		group,
		"",
		"Book model",
	)

	router.NewCreateModelSerializerHandler[dtos.BookCreatePayload, models.Book, dtos.BookResponse](
		serializerGroup,
		"",
		"Book serializer create",
		books.CreateBookSerializer{},
	)
	router.NewUpdateModelSerializerHandler[dtos.BookUpdatePayload, models.Book, dtos.BookResponse](
		serializerGroup,
		"",
		"Book serializer update",
		books.UpdateBookSerializer{},
	)
	router.NewDetailModelSerializerHandler[dtos.BookFilter, models.Book, dtos.BookResponse](
		serializerGroup,
		"",
		"Book serializer detail",
		books.DetailBookSerializer{},
	)
	router.NewListModelSerializerHandler[dtos.BookFilter, models.Book, dtos.BookResponse](
		serializerGroup,
		"",
		"Book serializer list",
		books.ListBookSerializer{},
	)
	router.NewModelSerializerHandler[dtos.BookCreatePayload, dtos.BookUpdatePayload, struct{}, dtos.BookFilter, models.Book, dtos.BookResponse](
		serializerModelGroup,
		"",
		"Book serializer model",
		books.BookSerializer{},
	)
}
