package main

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/yagobatista/taco-go-web-framework/example/handlers/book"
	"github.com/yagobatista/taco-go-web-framework/example/routes"
	"github.com/yagobatista/taco-go-web-framework/src/common_handlers"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

func main() {
	server.NewServer(&server.ServerConfig{
		Port: 8000,
		Handlers: []server.Handler{
			common_handlers.NewDbHealthCheckHandler(routes.PUBLIC),
			&book.BookHandler{},
		},
		Docs: true,
		DatabaseConnections: server.DatabaseConfig{
			Server:   server.POSTGRES,
			Host:     "localhost",
			Name:     "example",
			User:     "postgres",
			Password: "postgres",
			Port:     5432,
		},
		Routes: server.Router{
			routes.PUBLIC: nil,
			routes.INTERNAL: {
				basicauth.New(basicauth.Config{
					Users: map[string]string{"admin": "admin"},
				}),
			},
		},
	}).Start()
}
