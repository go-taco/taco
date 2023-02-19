package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/invopop/jsonschema"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

type Processor[UrlParams any, Body any, Response any] func(ctx context.Context, UrlParams UrlParams, input Body) (Response, error)

type Handler[UrlParams any, Body any, Response any] struct {
	BaseHandler

	Url       string
	Method    string
	processor Processor[UrlParams, Body, Response]
	router    fiber.Router
	name      string
}

func (this *Handler[UrlParams, Body, Response]) Mount() {
	if this.Method == http.MethodPost {
		this.router.Post(this.Url, this.Process)
	}

	if this.Method == http.MethodGet {
		this.router.Get(this.Url, this.Process)
	}

	this.router.Get(this.Url+"/docs", this.Docs).
		Name(fmt.Sprintf("%s - %s", this.Method, this.name))
}

func (this *Handler[UrlParams, Body, Response]) Parse(c *fiber.Ctx) (urlParams UrlParams, body Body, err error) {
	err = c.ParamsParser(&urlParams)
	if err != nil {
		return
	}

	parser := c.BodyParser

	if this.Method == http.MethodGet {
		parser = c.QueryParser
	}

	err = parser(&body)
	if err != nil {
		return
	}

	v := validator.New()
	err = v.Struct(body)

	return
}

func (this *Handler[UrlParams, Body, Response]) Docs(c *fiber.Ctx) error {
	var payload Body
	var response Response

	expectedPayload, err := json.Marshal(jsonschema.Reflect(&payload))
	if err != nil {
		return err
	}

	expectedResponse, err := json.Marshal(jsonschema.Reflect(&response))
	if err != nil {
		return err
	}

	return c.Render("templates/docs-detail", fiber.Map{
		"Payload":  string(expectedPayload),
		"Response": string(expectedResponse),
		"Title":    fmt.Sprintf("%s - %s", this.Method, this.name),
		"Route":    strings.Replace(c.Route().Path, "/docs", "", 1),
	})
}

func (this *Handler[UrlParams, Body, Response]) Process(c *fiber.Ctx) error {
	urlParams, body, err := this.Parse(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	response, err := this.process(c.UserContext(), urlParams, body)
	if err != nil {
		return err
	}

	return c.Status(this.getStatus()).JSON(response)
}

func (this *Handler[UrlParams, Body, Response]) getStatus() int {
	if this.Method == http.MethodPost {
		return http.StatusCreated
	}

	return http.StatusOK

}

func (this *Handler[UrlParams, Body, Response]) process(requestCtx context.Context, urlParams UrlParams, input Body) (Response, error) {
	if this.Method == http.MethodGet {
		return this.processor(requestCtx, urlParams, input)
	}

	return server.RunWithTransaction(
		requestCtx,
		this.server,
		func(ctx context.Context) (Response, error) {
			return this.processor(ctx, urlParams, input)
		},
	)
}
