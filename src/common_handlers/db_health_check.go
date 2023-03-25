package common_handlers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/handlers"
	"github.com/yagobatista/taco-go-web-framework/src/route"
	"github.com/yagobatista/taco-go-web-framework/src/router"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

type DbHealthCheckHandler struct {
	handlers.BaseHandler

	route route.Route
}

func NewDbHealthCheckHandler(route route.Route) *DbHealthCheckHandler {
	return &DbHealthCheckHandler{
		route: route,
	}
}

func (this *DbHealthCheckHandler) Routes(d route.Dispatcher) {
	group := d.GetRouter(this.route).Group("health")

	router.SetGet(
		this.BaseHandler,
		group,
		"/ping",
		"health check route",
		this.Ping,
	)
}

func (this *DbHealthCheckHandler) Ping(ctx context.Context, urlParams struct{}, payload struct{}) (string, error) {
	conn := server.GetConnectionFromCtx(ctx)

	db, err := conn.DB()
	if err != nil {
		return "", nil
	}

	err = db.Ping()
	if err != nil {
		return "", nil
	}

	return "pong", nil
}
