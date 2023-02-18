package server

type DatabaseServer string

const (
	POSTGRES DatabaseServer = "postgres"
)

type ServerConfig struct {
	DatabaseConnections DatabaseConfig
	Docs                bool
	AsyncTask           bool
}

type DatabaseConfig struct {
	Server   DatabaseServer
	Host     string
	Name     string
	User     string
	Password string
}
