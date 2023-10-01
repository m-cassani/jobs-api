package config

import (
	"os"

	"github.com/m-cassani/jobs-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// Create Logger for SQLite
	logger := GetLogger("sqlite")

	// Define database path
	dbPath := "./db/main.db"

	// Check if database file already exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found. Creating...")

		// Create directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Database directory creation error: %v", err)
			return nil, err
		}

		// Create database file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Database file creation error: %v", err)
			return nil, err
		}
		file.Close()

	}

	// Create SQLite and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite opening error: %v", err)
		return nil, err
	}

	// Migrate Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite automigration error: %v", err)
		return nil, err
	}

	// Return the database
	return db, nil
}
