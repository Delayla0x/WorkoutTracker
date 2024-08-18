package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Name   string
	Sets   int
	Reps   int
	Weight float64
}

type Workout struct {
	gorm.Model
	Name      string
	Date      time.Time
	Duration  int
	Exercises []Exercise `gorm:"foreignKey:WorkoutID"`
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Workout{}, &Exercise{})
	if err != nil {
		panic("failed to migrate database")
	}
}