package main

import (
	"fmt"
	"log" // for better error logging
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Exercise struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	WorkoutID uint
	Name      string
	Sets      int
	Reps      int
	Weight    float64
}

type Workout struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
		log.Fatalf("Failed to connect to database: %v", err) // log.Fatalf logs the error and calls os.Exit(1)
	}

	err = db.AutoMigrate(&Workout{}, &Exercise{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err) // Improved error handling
	}

	exercises := []Exercise{
		{Name: "Push Ups", Sets: 3, Reps: 15, Weight: 0},
		{Name: "Pull Ups", Sets: 3, Reps: 10, Weight: 0},
	}
	if err := CreateWorkout(db, "Morning Routine", 30, exercises); err != nil {
		log.Printf("Creating workout failed: %v\n", err) // Using log.Printf for coherent logging
		return
	}
	fmt.Println("Workout created successfully")
}

func CreateWorkout(db *gorm.DB, workoutName string, duration int, exercises []Exercise) error {
	workout := Workout{
		Name:      workoutName,
		Date:      time.Now(),
		Duration:  duration,
		Exercises: exercises,
	}

	if err := db.Create(&workout).Error; err != nil {
		return fmt.Errorf("error creating workout: %w", err) // Error wrapping for better error context
	}

	return nil
}