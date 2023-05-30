package database

import "fmt"

func CreateDB(dbName string, config DatabaseConfig) error {
	globalConnection := NewDatabaseConnection(config).GetConnection()

	dbAlreadyExists := globalConnection.Exec("SELECT 1 FROM pg_database WHERE datname = ?", dbName).RowsAffected == 1
	if dbAlreadyExists {
		return nil
	}

	createDbQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	return globalConnection.Exec(createDbQuery).Error
}
