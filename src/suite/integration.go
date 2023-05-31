package suite

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/stretchr/testify/suite"
	"github.com/yagobatista/taco-go-web-framework/src/database"
	"github.com/yagobatista/taco-go-web-framework/src/server"
)

type IntegrationSuite struct {
	suite.Suite

	mu sync.Mutex

	tables []any

	server       *server.Server
	serverConfig server.ServerConfig

	Ctx context.Context

	Client Client
}

func (this *IntegrationSuite) SetupSuite() {
	this.loadEnv()
	this.createDB()
	this.migrateDB()
}

func (this *IntegrationSuite) SetupTest() {
	this.server = server.NewServer(this.serverConfig)

	this.Client = NewClient(this.server, &this.Suite)

	conn := this.server.GetConnection().GetConnection()

	this.Ctx = context.Background()

	tx := conn.WithContext(this.Ctx).Begin()

	this.Ctx = database.SetConnectionToCtx(this.Ctx, tx)
}

func (this *IntegrationSuite) RollbackTest() {
	conn := database.GetConnectionFromCtx(this.Ctx)
	if conn == nil {
		return
	}

	conn.Rollback()
}

func (this *IntegrationSuite) SetServerConfig(config server.ServerConfig) {
	this.serverConfig.Handlers = config.Handlers
	this.serverConfig.MainMiddlewares = config.MainMiddlewares
	this.serverConfig.Routes = config.Routes
}

func (this *IntegrationSuite) SetTables(tables []any) {
	this.tables = tables
}

func (this *IntegrationSuite) loadEnv() {
	err := cleanenv.ReadConfig("../.test.env", &this.serverConfig)
	if err == nil {
		return
	}

	err = cleanenv.ReadConfig("../../.test.env", &this.serverConfig)
	if err == nil {
		return
	}

	err = cleanenv.ReadEnv(&this.serverConfig)
	if err != nil {
		this.FailNow("failed to load database envs")
	}
}

func (this *IntegrationSuite) createDB() {
	this.mu.Lock()
	defer this.mu.Unlock()

	dbName := this.serverConfig.DatabaseConfig.Name
	if dbName == "postgres" || dbName == "" {
		this.FailNow("DB_NAME can not be postgres ou empty")
	}

	dbName, _ = strings.CutSuffix(dbName, "_test")
	dbName = fmt.Sprintf("%s_test", dbName)

	this.serverConfig.DatabaseConfig.Name = dbName

	err := database.CreateDB(dbName, database.DatabaseConfig{
		Server:   database.POSTGRES,
		Host:     "localhost",
		Name:     "postgres",
		User:     "postgres",
		Password: "postgres",
		Port:     5432,
	})
	this.Require().NoError(err)
}

func (this *IntegrationSuite) migrateDB() {
	this.mu.Lock()
	defer this.mu.Unlock()

	conn := database.NewDatabaseConnection(this.serverConfig.DatabaseConfig).GetConnection()

	err := conn.AutoMigrate(this.tables...)
	this.Require().NoError(err, "failed to migrate database")
}
