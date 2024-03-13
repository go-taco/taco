package main

import (
	"github.com/yagobatista/taco-go-web-framework/example/cmd/setup"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

func main() {
	server.NewServer(setup.GetServerConfig()).Start()
}
