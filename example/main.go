package main

import "github.com/yagobatista/taco-go-web-framework/src/server"

func main() {
	server.NewServer(&server.ServerConfig{
		Port: 8000,
	}).Start()
}
