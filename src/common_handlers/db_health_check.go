package common_handlers

import (
	"context"

	"github.com/yagobatista/taco-go-web-framework/src/database"
	"github.com/yagobatista/taco-go-web-framework/src/route"
	"github.com/yagobatista/taco-go-web-framework/src/router"
)

type DbHealthCheckHandler struct {
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
		group,
		"/ping",
		"health check route",
		this.Ping,
	)
}

func (this *DbHealthCheckHandler) Ping(ctx context.Context, urlParams struct{}, payload struct{}) (string, error) {
	conn := database.GetConnectionFromCtx(ctx)

	db, err := conn.DB()
	if err != nil {
		return "", nil
	}

	err = db.Ping()
	if err != nil {
		return "", err
	}

	return "pong", nil
}
