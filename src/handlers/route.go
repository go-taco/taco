package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetRoute[UrlParams any, Body any, Response any](router fiber.Router, url string, method string, name string, processor Processor[UrlParams, Body, Response]) {
	handler := Handler[UrlParams, Body, Response]{
		Url:       url,
		Method:    method,
		processor: processor,
		router:    router,
		name:      name,
	}

	handler.Mount()
}

func SetPost[UrlParams any, Body any, Response any](router fiber.Router, url string, name string, processor Processor[UrlParams, Body, Response]) {
	SetRoute(router, url, http.MethodPost, name, processor)
}

func SetGet[UrlParams any, Body any, Response any](router fiber.Router, url string, name string, processor Processor[UrlParams, Body, Response]) {
	SetRoute(router, url, http.MethodGet, name, processor)
}
