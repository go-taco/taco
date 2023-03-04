package router

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

func NewCreateModelSerializerHandler[Payload any, Model any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, serializer serializers.Serializer[Payload, Model, Response]) {
	SetRouteWithSerializer(
		base,
		router,
		url,
		"/create",
		http.MethodPost,
		fmt.Sprintf("Create %s", name),
		serializer,
	)
}

func NewUpdateModelSerializerHandler[Payload any, Model any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, serializer serializers.Serializer[Payload, Model, Response]) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRouteWithSerializer(
		base,
		router,
		endpoint,
		"/update",
		http.MethodPatch,
		fmt.Sprintf("Update %s", name),
		serializer,
	)
}

func NewGetModelSerializerHandler[Payload any, Model any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, serializer serializers.Serializer[Payload, Model, Response]) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRouteWithSerializer(
		base,
		router,
		endpoint,
		"/detail",
		http.MethodGet,
		fmt.Sprintf("Detail %s", name),
		serializer,
	)
}

func NewListModelSerializerHandler[Filters any, Model any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, serializer serializers.Serializer[Filters, Model, Response]) {
	SetRouteWithSerializer(
		base,
		router,
		url,
		"/list",
		http.MethodGet,
		fmt.Sprintf("List %s", name),
		serializer,
	)
}
