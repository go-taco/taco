package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/docs"
)

func (this *Server) buildDocs() {
	if this.configs.DisableDocs {
		return
	}

	this.app.Use("docs", func(c *fiber.Ctx) error {
		return docs.RenderMainDoc(c, this.app)
	})
}
