package main

import (
	"github.com/yagobatista/taco-go-web-framework/example/cmd/setup"
	"github.com/yagobatista/taco-go-web-framework/src/database"
)

func main() {
	// migration
	dbName := "example"

	err := database.CreateDB(dbName, database.DatabaseConfig{
		Server:   database.POSTGRES,
		Host:     "localhost",
		Name:     "postgres",
		User:     "postgres",
		Password: "postgres",
		Port:     5432,
	})
	if err != nil {
		panic(err)
	}

	conn := database.NewDatabaseConnection(database.DatabaseConfig{
		Server:   database.POSTGRES,
		Host:     "localhost",
		Name:     dbName,
		User:     "postgres",
		Password: "postgres",
		Port:     5432,
	}).GetConnection()

	err = conn.Migrator().AutoMigrate(setup.GetModelsRegistry()...)
	if err != nil {
		panic(err)
	}
}
