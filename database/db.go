package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "os"
)

var database *gorm.DB

func InitializeDatabaseConnection() {
    databaseName := os.Getenv("DATABASE_NAME")
    if databaseName == "" {
        databaseName = "test.db"
    }

    var err error
    database, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

    fmt.Println("Database connection successfully established.")
}

func GetDatabase() *gorm.DB {
    return database
}

func main() {
    InitializeDatabaseConnection()

    dbConnection := GetDatabase()
}