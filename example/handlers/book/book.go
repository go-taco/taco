package book

import (
	"github.com/yagobatista/taco-go-web-framework/example/routes"
	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/route"
)

type BookHandler struct {
	handlers.BaseHandler
}

func (this *BookHandler) Routes(d route.Dispatcher) {
	group := d.GetRouter(routes.INTERNAL).Group("book")

	handlers.NewModelHandler[structs.BookFilter, structs.Book](
		this.BaseHandler,
		group,
		"",
		"Book model",
	)
}
