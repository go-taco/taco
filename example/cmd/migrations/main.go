package main

import (
	"github.com/yagobatista/taco-go-web-framework/example/structs"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

func main() {
	// migration
	conn := server.NewDatabaseConnection(server.DatabaseConfig{
		Server:   server.POSTGRES,
		Host:     "localhost",
		Name:     "example",
		User:     "postgres",
		Password: "postgres",
		Port:     5432,
	}).GetConnection()

	err := conn.Migrator().AutoMigrate([]any{
		&structs.Book{},
	}...)
	if err != nil {
		panic(err)
	}
}
