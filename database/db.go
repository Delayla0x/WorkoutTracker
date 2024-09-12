package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDatabase() {
	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		dbName = "test.db"
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	fmt.Println("Database connection successfully established.")
}

func GetDB() *gorm.DB {
	return DB
}

func main() {
	InitDatabase()

	db := GetDB()
}