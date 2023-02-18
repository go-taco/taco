package handlers

import "github.com/yagobatista/taco-go-web-framework/src/server"

type BaseHandler struct {
	server *server.Server
}

func (this *BaseHandler) SetServer(server *server.Server) {
	this.server = server
}

func (this *BaseHandler) GetServer() *server.Server {
	return this.server
}
