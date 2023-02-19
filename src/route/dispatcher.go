package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/middlewares"
)

type Dispatcher struct {
	routes map[Route]fiber.Router
}

func NewDispatcher(mainRouter fiber.Router, routes ...Route) Dispatcher {
	routesMap := make(map[Route]fiber.Router, len(routes))

	for _, route := range routes {
		routesMap[route] = mainRouter.Group(fmt.Sprintf("%s", route))
	}

	return Dispatcher{
		routes: routesMap,
	}
}

func (this *Dispatcher) GetRouter(routeKey Route) fiber.Router {
	route, ok := this.routes[routeKey]
	if !ok {
		panic(fmt.Sprintf("missing router %s", routeKey))
	}

	return route
}

func (this *Dispatcher) SetMiddleware(routeKey Route, middlewares []middlewares.Middleware) {
	var newMiddlewares []any

	for _, middleware := range middlewares {
		newMiddlewares = append(newMiddlewares, func(c *fiber.Ctx) error {
			return middleware(c)
		})
	}

	this.routes[routeKey] = this.GetRouter(routeKey).Use(newMiddlewares...)
}
