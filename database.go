package database

import (
	"fmt"
	"log"
	"time"

	"github.com/hmly-tech/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	DB *gorm.DB
}

func Connect(cfg config.Config, model interface{}) DatabaseConnection {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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

	db.AutoMigrate(model)

	return DatabaseConnection{db}
}
