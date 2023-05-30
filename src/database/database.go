package database

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseServer string

const (
	POSTGRES DatabaseServer = "POSTGRES"
)

var availableServer = map[DatabaseServer]func(config DatabaseConfig) gorm.Dialector{
	POSTGRES: postgresDialector,
}

type DatabaseConfig struct {
	Disabled bool           `env:"DB_DISABLED"`
	Server   DatabaseServer `env:"DB_SERVER"`
	Host     string         `env:"DB_HOST"`
	Port     int            `env:"DB_PORT"`
	Name     string         `env:"DB_NAME"`
	User     string         `env:"DB_USER"`
	Password string         `env:"DB_PASSWORD"`
}

type DatabaseConnection struct {
	conn *gorm.DB
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

func (this *DatabaseConnection) GetConnection() *gorm.DB {
	return this.conn
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

func RunWithTransaction[Result any](ctx context.Context, f func(ctx context.Context) (Result, error)) (result Result, err error) {
	globalConnection := GetConnectionFromCtx(ctx)

	err = globalConnection.Transaction(func(tx *gorm.DB) (err error) {
		newCtx := SetConnectionToCtx(ctx,
			tx.WithContext(ctx),
		)

		result, err = f(newCtx)
		return err
	})

	return
}

func GetConnectionFromCtx(ctx context.Context) *gorm.DB {
	return ctx.Value("conn").(*gorm.DB)
}

func SetConnectionToCtx(ctx context.Context, conn *gorm.DB) context.Context {
	return context.WithValue(ctx, "conn", conn)
}
