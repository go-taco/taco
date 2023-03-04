package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/serializers"
)

func SetRouteWithDocs[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, docsUrl, method string, name string, processor handlers.Processor[UrlParams, Payload, Response]) {
	handler := handlers.Handler[UrlParams, Payload, Response]{
		Url:       url,
		DocUrl:    docsUrl,
		Method:    method,
		Processor: processor,
		Router:    router,
		Name:      name,
	}

	handler.SetServer(base.GetServer())
	handler.Mount()
}

func SetRouteWithSerializer[UrlParams any, Payload any, Model any, Response any](base handlers.BaseHandler, router fiber.Router, url string, docsUrl, method string, name string, serializer serializers.Serializer[UrlParams, Payload, Model, Response]) {
	SetRouteWithDocs(
		base,
		router,
		url,
		"",
		method,
		name,
		serializers.SerializerToProcessor(serializer),
	)
}

func SetRoute[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, method string, name string, processor handlers.Processor[UrlParams, Payload, Response]) {
	SetRouteWithDocs(
		base,
		router,
		url,
		"",
		method,
		name,
		processor,
	)
}

func SetPost[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, processor handlers.Processor[UrlParams, Payload, Response]) {
	SetRoute(base, router, url, http.MethodPost, name, processor)
}

func SetGet[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, processor handlers.Processor[UrlParams, Payload, Response]) {
	SetRoute(base, router, url, http.MethodGet, name, processor)
}
