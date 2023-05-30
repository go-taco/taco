package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/configs"
)

func NewProjectConfigs[T any]() Middleware {
	cfg, err := configs.LoadEnvs[T]()
	if err != nil {
		log.Panicf("failed to load env var. Error: %v", err)
	}

	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		ctx = configs.SetToCtx(ctx, cfg)
		c.SetUserContext(ctx)
		return nil
	}
}
