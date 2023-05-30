package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/database"
	"github.com/yagobatista/taco-go-web-framework/src/middlewares"
)

func connectionMiddleware(server *Server) middlewares.Middleware {
	return func(c *fiber.Ctx) error {
		ctx := database.SetConnectionToCtx(c.UserContext(), server.dbConnection.GetConnection())

		c.SetUserContext(ctx)

		return c.Next()
	}
}
