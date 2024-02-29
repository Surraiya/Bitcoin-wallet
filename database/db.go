package database

import (
	"bitcoin-wallet/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitializeDB(dbPath string) error {
	var err error
	db, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Transaction{})
	log.Println("Connected to the database.")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
