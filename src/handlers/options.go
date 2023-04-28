package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HandlerConfig struct {
	withTransaction     bool
	requestMiddlewares  []fiber.Handler
	responseMiddlewares []fiber.Handler
}

type Option func(*HandlerConfig)

func WithTransaction(withTransaction bool) Option {
	return func(config *HandlerConfig) {
		config.withTransaction = withTransaction
	}
}

func WithRequestMiddlewares(requestMiddlewares ...fiber.Handler) Option {
	return func(config *HandlerConfig) {
		config.requestMiddlewares = requestMiddlewares
	}
}

func WithResponseMiddlewares(responseMiddlewares ...fiber.Handler) Option {
	return func(config *HandlerConfig) {
		config.responseMiddlewares = responseMiddlewares
	}
}
