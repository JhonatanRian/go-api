package config

import (
	"os"

	"github.com/JhonatanRian/go-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/database.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Creating database...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	db, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to migrate database: %v", err)
		return nil, err
	}

	return db, nil
}
