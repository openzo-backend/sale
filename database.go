package main

import (
	"fmt"

	"github.com/tanush-128/openzo_backend/sale/config"
	"github.com/tanush-128/openzo_backend/sale/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectToDB(cfg *config.Config) (*gorm.DB, error) {

	if cfg.MODE == "production" {
		dsn := cfg.DB_URL

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to open database connection: %w", err)
		}
		return db, nil

	}
	db, err := gorm.Open(
		sqlite.Open("test.db"),

		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	db.Migrator().AutoMigrate(&models.Sale{})
	db.Migrator().AutoMigrate(&models.OrderItem{})
	db.Migrator().AutoMigrate(&models.Customer{})

	return db, nil
}
