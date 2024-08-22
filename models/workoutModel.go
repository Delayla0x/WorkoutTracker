package main

import (
    "fmt"
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
        panic("Failed to connect to database: " + err.Error())
    }

    err = db.AutoMigrate(&Workout{}, &Exercise{})
    if err != nil {
        panic("Database migration failed: " + err.Error())
    }

    exercises := []Exercise{
        {Name: "Push Ups", Sets: 3, Reps: 15, Weight: 0},
        {Name: "Pull Ups", Sets: 3, Reps: 10, Weight: 0},
    }
    if err := CreateWorkout(db, "Morning Routine", 30, exercises); err != nil {
        fmt.Printf("Creating workout failed: %v\n", err)
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
        return fmt.Errorf("error creating workout: %w", err)
    }

    return nil
}