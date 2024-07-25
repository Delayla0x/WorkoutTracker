package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	loadEnv()
	db := initDB()
	defer db.Close()

	router := gin.Default()
	configureRoutes(router)

	port := getEnvWithDefault("PORT", "8080")
	router.Run(fmt.Sprintf(":%s", port))
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func configureRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}