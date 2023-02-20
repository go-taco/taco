package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetRouteWithDocs[UrlParams any, Body any, Response any](base BaseHandler, router fiber.Router, url string, docsUrl, method string, name string, processor Processor[UrlParams, Body, Response]) Handler[UrlParams, Body, Response] {
	handler := Handler[UrlParams, Body, Response]{
		Url:       url,
		DocUrl:    docsUrl,
		Method:    method,
		processor: processor,
		router:    router,
		name:      name,
	}

	handler.SetServer(base.GetServer())
	handler.Mount()

	return handler
}

func SetRoute[UrlParams any, Body any, Response any](base BaseHandler, router fiber.Router, url string, method string, name string, processor Processor[UrlParams, Body, Response]) {
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

func SetPost[UrlParams any, Body any, Response any](base BaseHandler, router fiber.Router, url string, name string, processor Processor[UrlParams, Body, Response]) {
	SetRoute(base, router, url, http.MethodPost, name, processor)
}

func SetGet[UrlParams any, Body any, Response any](base BaseHandler, router fiber.Router, url string, name string, processor Processor[UrlParams, Body, Response]) {
	SetRoute(base, router, url, http.MethodGet, name, processor)
}
