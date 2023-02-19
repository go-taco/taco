package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func getFiberConfig(serverConfig *ServerConfig) (fiberConfig fiber.Config) {

	if serverConfig.Docs {
		eng := html.New("./", ".html")
		eng.Reload(true)
		fiberConfig.Views = eng
	}

	return
}
