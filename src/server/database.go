package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yagobatista/taco-go-web-framework/src/middlewares"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseServer string

const (
	POSTGRES DatabaseServer = "postgres"
)

type DatabaseConfig struct {
	Disabled bool
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
	if config.Disabled {
		return &DatabaseConnection{}
	}

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

func RunWithTransaction[Result any](ctx context.Context, server *Server, f func(ctx context.Context) (Result, error)) (Result, error) {
	var result Result

	err := server.dbConnection.conn.Transaction(func(tx *gorm.DB) (err error) {
		newCtx := SetToCtx(ctx, tx)
		result, err = f(newCtx)
		return err
	})

	return result, err
}

func connectionMiddleware(server *Server) middlewares.Middleware {
	return func(c *fiber.Ctx) error {
		ctx := SetToCtx(c.UserContext(), server.dbConnection.conn)

		c.SetUserContext(ctx)

		return c.Next()
	}
}

func GetFromCtx(ctx context.Context) *gorm.DB {
	return ctx.Value("conn").(*gorm.DB)
}

func SetToCtx(ctx context.Context, conn *gorm.DB) context.Context {
	return context.WithValue(ctx, "conn", conn)
}
