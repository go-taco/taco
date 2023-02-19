package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yagobatista/taco-go-web-framework/src/middlewares"
	"github.com/yagobatista/taco-go-web-framework/src/route"
)

type Shutdown interface {
	Shutdown() error
}

type Handler interface {
	SetServer(server *Server)
	Routes(d route.Dispatcher)
}

type Router map[route.Route][]middlewares.Middleware

type ServerConfig struct {
	DatabaseConnections DatabaseConfig
	Docs                bool
	AsyncTask           bool
	Port                int

	Handlers []Handler

	MainMiddlewares []middlewares.Middleware

	Routes Router
}

type Server struct {
	dbConnection *DatabaseConnection
	app          *fiber.App

	services []Shutdown

	routes []route.Route

	port int
}

func NewServer(config *ServerConfig) *Server {
	conn := NewDatabaseConnection(config.DatabaseConnections)

	app := fiber.New(getFiberConfig(config))

	server := &Server{
		dbConnection: conn,
		port:         config.Port,
		app:          app,
	}

	for key := range config.Routes {
		server.routes = append(server.routes, key)
	}

	server.loadRoutes(*config)
	server.buildDocs(*config)

	return server
}

func (this *Server) loadRoutes(serverConfig ServerConfig) {
	handlers := serverConfig.Handlers

	mainRouter := this.setMainRouter(serverConfig.MainMiddlewares)

	routerDispatcher := route.NewDispatcher(
		mainRouter,
		this.routes...,
	)

	for route, middleware := range serverConfig.Routes {
		if middleware == nil {
			continue
		}
		routerDispatcher.SetMiddleware(route, middleware)
	}

	for _, handler := range handlers {
		handler.SetServer(this)
		handler.Routes(routerDispatcher)
	}
}

func (this *Server) setMainRouter(mainMiddlewares []middlewares.Middleware) fiber.Router {
	mainRouter := this.app.Group("")

	mainRouter.Use(
		recover.New(recover.Config{
			EnableStackTrace: true,
		}),
	)

	mainMiddlewares = append(mainMiddlewares, connectionMiddleware(this))

	for _, middleware := range mainMiddlewares {
		mainRouter.Use(func(c *fiber.Ctx) error {
			return middleware(c)
		})
	}

	return mainRouter
}

func (this *Server) Start() {
	defer this.shutdown()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	serverShutdown := make(chan struct{})

	go func() {
		<-quit

		err := this.app.Shutdown()
		if err != nil {
			panic(err)
		}

		serverShutdown <- struct{}{}
	}()

	err := this.app.Listen(fmt.Sprintf(":%d", this.port))
	if err != nil {
		panic(err)
	}

	<-serverShutdown
}

func (this *Server) shutdown() {
	for _, service := range this.services {
		// TODO: log error
		service.Shutdown()
	}
}
