package database

import (
	"log"
	"time"

	"github.com/hmlylabs/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	DB *gorm.DB
}

func Connect[T any](cfg config.Config, model T) DatabaseConnection {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	db.AutoMigrate(&model)

	return DatabaseConnection{db}
}
