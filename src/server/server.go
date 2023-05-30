package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/imdario/mergo"
	"github.com/yagobatista/taco-go-web-framework/src/configs"
	"github.com/yagobatista/taco-go-web-framework/src/database"
	"github.com/yagobatista/taco-go-web-framework/src/middlewares"
	"github.com/yagobatista/taco-go-web-framework/src/route"
)

type Shutdown interface {
	Shutdown() error
}

type Handler interface {
	Routes(d route.Dispatcher)
}

type Router map[route.Route][]middlewares.Middleware

type ServerConfig struct {
	DatabaseConfig database.DatabaseConfig
	DisableDocs    bool `env:"DISABLE_DOCS"`
	AsyncTask      bool
	Port           int `env:"PORT"`

	Handlers []Handler

	MainMiddlewares []middlewares.Middleware

	Routes Router
}

type Server struct {
	dbConnection *database.DatabaseConnection
	app          *fiber.App

	configs ServerConfig

	services []Shutdown

	routes []route.Route

	port int
}

func NewServer(config ServerConfig) *Server {
	var server Server

	server.setConfigs(config)

	conn := database.NewDatabaseConnection(server.configs.DatabaseConfig)

	app := fiber.New(getFiberConfig(server.configs))

	server.dbConnection = conn
	server.port = server.configs.Port
	server.app = app

	for key := range config.Routes {
		server.routes = append(server.routes, key)
	}

	server.loadRoutes()
	server.buildDocs()

	return &server
}

func (this *Server) GetConnection() *database.DatabaseConnection {
	return this.dbConnection
}

func (this *Server) GetFiberApp() *fiber.App {
	return this.app
}

func (this *Server) setConfigs(config ServerConfig) {
	this.configs = config

	cfg, err := configs.LoadEnvs[ServerConfig]()
	if err != nil {
		log.Printf("failed to load env file. Error: %v", err)
		return
	}

	err = mergo.Merge(&this.configs, cfg)
	if err != nil {
		log.Panic(err)
	}
}

func (this *Server) loadRoutes() {
	handlers := this.configs.Handlers

	mainRouter := this.setMainRouter(this.configs.MainMiddlewares)

	routerDispatcher := route.NewDispatcher(
		mainRouter,
		this.routes...,
	)

	for route, middleware := range this.configs.Routes {
		if middleware == nil {
			continue
		}
		routerDispatcher.SetMiddleware(route, middleware)
	}

	for _, handler := range handlers {
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
