package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/database"
)

type Processor[UrlParams any, Body any, Response any] func(ctx context.Context, UrlParams UrlParams, input Body) (Response, error)

type Handler[UrlParams any, Payload any, Response any] struct {
	Url    string
	DocUrl string

	Method string
	Router fiber.Router

	Processor Processor[UrlParams, Payload, Response]
	Name      string

	disableTransaction  bool
	requestMiddlewares  []fiber.Handler
	responseMiddlewares []fiber.Handler
}

func (this *Handler[UrlParams, Payload, Response]) Mount() {
	var completeHandler []fiber.Handler

	completeHandler = append(completeHandler, this.requestMiddlewares...)
	completeHandler = append(completeHandler, this.getResponseMiddlewares()...)
	completeHandler = append(completeHandler, this.Process)

	this.Router.Add(this.Method, this.Url, completeHandler...)

	docsUrl := fmt.Sprintf("%s%s/docs", this.Url, this.DocUrl)
	docRouteName := fmt.Sprintf("%s - %s", this.Method, this.Name)

	this.Router.Get(docsUrl, this.Docs).Name(docRouteName)
}

func (this *Handler[UrlParams, Payload, Response]) getResponseMiddlewares() (parserResponseMiddlewares []fiber.Handler) {
	for _, middleware := range this.responseMiddlewares {
		parserResponseMiddlewares = append(parserResponseMiddlewares, func(c *fiber.Ctx) error {
			c.Next()
			return middleware(c)
		})
	}
	return
}

func (this *Handler[UrlParams, Payload, Response]) SetConfig(config HandlerConfig) {
	this.disableTransaction = !config.withTransaction
	this.requestMiddlewares = config.requestMiddlewares
	this.responseMiddlewares = config.responseMiddlewares
}

func (this *Handler[UrlParams, Payload, Response]) Parse(c *fiber.Ctx) (urlParams UrlParams, body Payload, err error) {
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

func (this *Handler[UrlParams, Payload, Response]) Process(c *fiber.Ctx) error {
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

func (this *Handler[UrlParams, Payload, Response]) getStatus() int {
	if this.Method == http.MethodPost {
		return http.StatusCreated
	}

	return http.StatusOK
}

func (this *Handler[UrlParams, Payload, Response]) process(requestCtx context.Context, urlParams UrlParams, payload Payload) (Response, error) {
	if this.Method == http.MethodGet || this.disableTransaction {
		return this.Processor(requestCtx, urlParams, payload)
	}

	return database.RunWithTransaction(
		requestCtx,
		func(ctx context.Context) (Response, error) {
			return this.Processor(ctx, urlParams, payload)
		},
	)
}
