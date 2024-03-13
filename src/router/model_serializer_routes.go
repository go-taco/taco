package router

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

func NewCreateModelSerializerHandler[Payload any, Model any, Response any](router fiber.Router, url string, name string, serializer serializers.CreateSerializerInterface[Payload, Model, Response]) {
	SetRouteWithDocs(
		router,
		url,
		"/create",
		http.MethodPost,
		fmt.Sprintf("Create %s", name),
		serializers.SerializerToCreateProcessor(serializer),
	)
}

func NewUpdateModelSerializerHandler[Payload any, Model any, Response any](router fiber.Router, url string, name string, serializer serializers.UpdateSerializerInterface[Payload, Model, Response]) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRouteWithDocs(
		router,
		endpoint,
		"/update",
		http.MethodPatch,
		fmt.Sprintf("Update %s", name),
		serializers.SerializerToUpdateProcessor(serializer),
	)
}

func NewDetailModelSerializerHandler[Filter, Model any, Response any](router fiber.Router, url string, name string, serializer serializers.DetailSerializerInterface[Filter, Model, Response]) {
	endpoint := fmt.Sprintf("%s/:id", url)
	SetRouteWithDocs(
		router,
		endpoint,
		"/detail",
		http.MethodGet,
		fmt.Sprintf("Detail %s", name),
		serializers.SerializerToDetailProcessor(serializer),
	)
}

func NewListModelSerializerHandler[Filter any, Model any, Response any](router fiber.Router, url string, name string, serializer serializers.ListSerializerInterface[Filter, Model, Response]) {
	SetRouteWithDocs(
		router,
		url,
		"/list",
		http.MethodGet,
		fmt.Sprintf("List %s", name),
		serializers.SerializerToListProcessor(serializer),
	)
}

func NewModelSerializerHandler[CreatePayload any, UpdatePayload any, DetailQueryParams any, Filter any, Model any, Response any](router fiber.Router, url string, name string, serializer serializers.ModelSerializerInterface[CreatePayload, UpdatePayload, DetailQueryParams, Filter, Model, Response]) {
	NewCreateModelSerializerHandler[CreatePayload, Model, Response](router, url, name, serializer)
	NewUpdateModelSerializerHandler[UpdatePayload, Model, Response](router, url, name, serializer)
	NewDetailModelSerializerHandler[DetailQueryParams, Model, Response](router, url, name, serializer)
	NewListModelSerializerHandler[Filter, Model, Response](router, url, name, serializer)
}
