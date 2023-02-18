package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type Shutdown interface {
	Shutdown() error
}

type ServerConfig struct {
	DatabaseConnections DatabaseConfig
	Docs                bool
	AsyncTask           bool
	Port                int
}

type Server struct {
	dbConnection *DatabaseConnection
	app          *fiber.App

	services []Shutdown

	port int
}

func NewServer(config *ServerConfig) *Server {
	conn := NewDatabaseConnection(config.DatabaseConnections)

	return &Server{
		dbConnection: conn,
		port:         config.Port,
	}
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
		// log error
		service.Shutdown()
	}
}
