package config

import (
	"os"

	"github.com/felipefernandesdev/api_opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		//create database file and directory
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Create directory ./db error: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("Create file ./db/main.db error: %v", err)
			return nil, err
		}

		file.Close()
	}

	//Create Database and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	//migrate and schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite AutoMigrate error: %v", err)
		return nil, err
	}

	//return database
	return db, nil
}
