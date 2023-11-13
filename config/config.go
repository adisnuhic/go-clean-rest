package config

import (
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

// DBConnection connection to a database
type DBConnection struct {
	DBDialect         string
	DBConnection      string
	DbMaxIdleConns    int
	DbMaxOpenConns    int
	DbConnMaxLifetime int
	DbLogging         bool
}

// JWT configuration
type JWTConf struct {
	Secret string
}

// Env conf
type Env struct {
	Env string
}

// AppConfig application configuration
type AppConfig struct {
	DBConnections map[string]DBConnection
	JWTConf
	Env
}

// Load app configuration
func Load() *AppConfig {
	gotenv.Load()

	// defaults
	maxIdleConnections := 10
	maxOpenConnections := 100
	maxConnectionsLifetime := 20

	// customs
	if os.Getenv("DB_MAX_IDLE_CONNECTIONS") != "" {
		maxIdleConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	}

	if os.Getenv("DB_MAX_OPEN_CONNECTIONS") != "" {
		maxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	}

	if os.Getenv("DB_CONNECTION_MAX_LIFETIME_MINUTES") != "" {
		maxConnectionsLifetime, _ = strconv.Atoi(os.Getenv("DB_CONNECTION_MAX_LIFETIME_MINUTES"))
	}

	return &AppConfig{
		DBConnections: map[string]DBConnection{
			"development": {
				DBDialect:         "mysql",
				DBConnection:      os.Getenv("DB_DEV_CONNECTION"),
				DbMaxIdleConns:    maxIdleConnections,
				DbMaxOpenConns:    maxOpenConnections,
				DbConnMaxLifetime: maxConnectionsLifetime,
				DbLogging:         true,
			},
			"production": {
				DBDialect:         "mysql",
				DBConnection:      os.Getenv("DB_PROD_CONNECTION"),
				DbMaxIdleConns:    maxIdleConnections,
				DbMaxOpenConns:    maxOpenConnections,
				DbConnMaxLifetime: maxConnectionsLifetime,
				DbLogging:         true,
			},
		},
		JWTConf: JWTConf{
			Secret: os.Getenv("JWT_SECRET"),
		},
		Env: Env{
			Env: os.Getenv("ENV"),
		},
	}
}
