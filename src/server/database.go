package server

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseServer string

const (
	POSTGRES DatabaseServer = "postgres"
)

type DatabaseConfig struct {
	Server   DatabaseServer
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type DatabaseConnection struct {
	conn *gorm.DB
}

var availableServer = map[DatabaseServer]func(config DatabaseConfig) gorm.Dialector{
	POSTGRES: postgresDialector,
}

func NewDatabaseConnection(config DatabaseConfig) *DatabaseConnection {
	dialectorFunc, ok := availableServer[config.Server]
	if !ok {
		panic("missing server type")
	}

	db, err := gorm.Open(dialectorFunc(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &DatabaseConnection{
		conn: db,
	}
}

func postgresDialector(config DatabaseConfig) gorm.Dialector {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)
	return postgres.Open(dsn)
}
