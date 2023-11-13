package db

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/adisnuhic/go-stanard-layout/config"
	"github.com/jinzhu/gorm"

	// Initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
	// Initialize mysql migrate
	"github.com/golang-migrate/migrate/v4"
	mysqlmigrate "github.com/golang-migrate/migrate/v4/database/mysql"

	// Initialize mysql migrate source file
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// db app data store
var dbStore Store

// Init initialize db
func Init(cfg *config.AppConfig) {
	env := os.Getenv("ENV")
	dbStore = initDB(cfg.DBConnections[env])
	runMigrate(dbStore)
}

// Connection get databse connection
func Connection() Store {
	return dbStore
}

// Close database connection
func Close() error {
	if dbStore != nil {
		return dbStore.DB().Close()
	}
	return nil
}

// initDB init database connection
func initDB(dbConn config.DBConnection) Store {

	if dbConn.DBDialect == "" || dbConn.DBConnection == "" {
		return nil
	}

	// open DB connection
	myDB, err := gorm.Open(dbConn.DBDialect, dbConn.DBConnection)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	// ping database
	if err := myDB.DB().Ping(); err != nil {
		log.Error().Msg(err.Error())
	}

	// SetMaxIdleConns sets maximum number of connections in the idle connection pool
	maxConn := dbConn.DbMaxIdleConns
	myDB.DB().SetMaxIdleConns(maxConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database
	maxConn = dbConn.DbMaxOpenConns
	myDB.DB().SetMaxOpenConns(maxConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxConn = dbConn.DbConnMaxLifetime
	duration := time.Minute * time.Duration(maxConn)
	myDB.DB().SetConnMaxLifetime(duration)

	// Enable Logger, show detailed log
	myDB.LogMode(dbConn.DbLogging)

	log.Print("initialized API database successfully")

	return myDB

}

// executes migrations against database
func runMigrate(store Store) {
	driver, err := mysqlmigrate.WithInstance(store.DB(), &mysqlmigrate.Config{})
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql", driver)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	migrateErr := m.Up()
	if migrateErr != nil {
		log.Printf("Unable to migrate: %v", migrateErr)
	}

	log.Print("migrations executed successfully")
}
