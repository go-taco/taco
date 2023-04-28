package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
)

func SetRouteWithDocs[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, docsUrl, method string, name string, processor handlers.Processor[UrlParams, Payload, Response], options ...handlers.Option) {
	handler := handlers.Handler[UrlParams, Payload, Response]{
		Url:       url,
		DocUrl:    docsUrl,
		Method:    method,
		Processor: processor,
		Router:    router,
		Name:      name,
	}

	var config handlers.HandlerConfig
	for _, option := range options {
		option(&config)
	}

	handler.SetConfig(config)

	handler.SetServer(base.GetServer())
	handler.Mount()
}

func SetRoute[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, method string, name string, processor handlers.Processor[UrlParams, Payload, Response], options ...handlers.Option) {
	SetRouteWithDocs(
		base,
		router,
		url,
		"",
		method,
		name,
		processor,
		options...,
	)
}

func SetPost[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, processor handlers.Processor[UrlParams, Payload, Response], options ...handlers.Option) {
	SetRoute(base, router, url, http.MethodPost, name, processor, options...)
}

func SetGet[UrlParams any, Payload any, Response any](base handlers.BaseHandler, router fiber.Router, url string, name string, processor handlers.Processor[UrlParams, Payload, Response], options ...handlers.Option) {
	SetRoute(base, router, url, http.MethodGet, name, processor, options...)
}
