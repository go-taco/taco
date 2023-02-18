package main

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/yagobatista/taco-go-web-framework/example/routes"
	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

func main() {
	server.NewServer(&server.ServerConfig{
		Port: 8000,
		Handlers: []server.Handler{
			handlers.NewDbHealthCheckHandler(routes.PUBLIC),
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
