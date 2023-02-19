package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (this *Server) buildDocs(serverConfig ServerConfig) {
	if !serverConfig.Docs {
		return
	}

	documentedRoutes := make(map[string]string)

	for _, route := range this.app.GetRoutes() {
		if !strings.Contains(route.Path, "/docs") {
			continue
		}

		name, found := documentedRoutes[route.Path]
		if found && name != "" {
			continue
		}

		documentedRoutes[route.Path] = route.Name
	}

	this.app.Use("docs", func(c *fiber.Ctx) error {
		return c.Render("templates/docs", fiber.Map{
			"Routes": documentedRoutes,
		})
	})
}
