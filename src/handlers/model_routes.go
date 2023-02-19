package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewCreateModelHandler[Model any](router fiber.Router, url string, name string) {
	SetRoute(router, url, http.MethodPost, name, CreateModel[Model])
}

func NewUpdateModelHandler[Model any](router fiber.Router, url string, name string) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRoute(router, endpoint, http.MethodPatch, name, UpdateModel[Model])
}

func NewGetModelHandler[Model any](router fiber.Router, url string, name string) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRoute(router, endpoint, http.MethodPatch, name, UpdateModel[Model])
}

func NewListModelHandler[Filters any, Model any](router fiber.Router, url string, name string) {
	SetRoute(router, url, http.MethodGet, name, ListModel[Filters, Model])
}
