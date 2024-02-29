package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var testDBPath = "./test.db"

func SetupTestDB() *gorm.DB {
	err := InitializeDB(testDBPath)
	testDb := GetDB()
	if err != nil {
		log.Fatalf("Error initializing test database: %v", err)
	}
	return testDb
}

func CleanupTestDB(testDb *gorm.DB) {
	if testDb != nil {
		testDb.Close()
		fmt.Println("database connection is closed")
	}
	if err := os.Remove(testDBPath); err != nil {
		log.Fatalf("Error cleaning up test database: %v", err)
	}
}
