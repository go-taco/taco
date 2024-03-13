package setup

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	project_configs "github.com/yagobatista/taco-go-web-framework/example/configs"
	"github.com/yagobatista/taco-go-web-framework/example/handlers/book"
	"github.com/yagobatista/taco-go-web-framework/example/routes"
	"github.com/yagobatista/taco-go-web-framework/src/common_handlers"
	"github.com/yagobatista/taco-go-web-framework/src/middlewares"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

func GetServerConfig() server.ServerConfig {
	return server.ServerConfig{
		Handlers: []server.Handler{
			common_handlers.NewDbHealthCheckHandler(routes.PUBLIC),
			&book.BookHandler{},
		},
		Routes: server.Router{
			routes.PUBLIC: nil,
			routes.INTERNAL: {
				basicauth.New(basicauth.Config{
					Users: map[string]string{"admin": "admin"},
				}),
				middlewares.NewProjectConfigs[project_configs.Configs](),
			},
		},
	}
}
