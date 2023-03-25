package router

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/model"
)

func NewCreateModelHandler[Model any](base handlers.BaseHandler, router fiber.Router, url string, name string) {
	SetRouteWithDocs(
		base,
		router,
		url,
		"/create",
		http.MethodPost,
		fmt.Sprintf("Create %s", name),
		model.CreateModel[Model],
	)
}

func NewUpdateModelHandler[Model any](base handlers.BaseHandler, router fiber.Router, url string, name string) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRouteWithDocs(
		base,
		router,
		endpoint,
		"/update",
		http.MethodPatch,
		fmt.Sprintf("Update %s", name),
		model.UpdateModel[Model],
	)
}

func NewGetModelHandler[Model any](base handlers.BaseHandler, router fiber.Router, url string, name string) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRouteWithDocs(
		base,
		router,
		endpoint,
		"/detail",
		http.MethodGet,
		fmt.Sprintf("Detail %s", name),
		model.GetModel[Model],
	)
}

func NewListModelHandler[Filters any, Model any](base handlers.BaseHandler, router fiber.Router, url string, name string) {
	SetRouteWithDocs(
		base,
		router,
		url,
		"/list",
		http.MethodGet,
		fmt.Sprintf("List %s", name),
		model.ListModel[Filters, Model],
	)
}

func NewModelHandler[Filters any, Model any](base handlers.BaseHandler, router fiber.Router, url string, name string) {
	NewCreateModelHandler[Model](base, router, url, name)
	NewUpdateModelHandler[Model](base, router, url, name)
	NewGetModelHandler[Model](base, router, url, name)
	NewListModelHandler[Filters, Model](base, router, url, name)
}
