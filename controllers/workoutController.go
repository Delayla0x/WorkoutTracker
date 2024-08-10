package main

import (
    "fmt"
    "net/http"
    "os"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/subosito/gotenv"
)

func init() {
    gotenv.Load()
}

type Workout struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}

var db *gorm.DB
var err error

var workoutCache = struct {
    sync.RWMutex
    workouts   []Workout
    lastUpdate time.Time
}{}

const cacheDuration = 5 * time.Minute 

func initDB() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbName := os.Getenv("DB_NAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbSSLMode := os.Getenv("DB_SSLMODE")

    dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, dbPort, dbUser, dbName, dbSSLMode, dbPassword)
    fmt.Println(dbUri)
    db, err = gorm.Open("postgres", dbUri)
    if err != nil {
        fmt.Println(err)
    }
    db.AutoMigrate(&Workout{})
}

func createWorkout(c *gin.Context) {
    var workout Workout
    if err := c.BindJSON(&workout); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&workout).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    workoutCache.Lock()
    workoutCache.lastUpdate = time.Time{}
    workoutCache.Unlock()

    c.JSON(http.StatusOK, workout)
}

func getWorkouts(c *gin.Context) {
    workoutCache.RLock()
    cacheExpired := time.Since(workoutCache.lastUpdate) > cacheDuration
    if !cacheExpired && len(workoutCache.workouts) > 0 {
        c.JSON(http.StatusOK, workoutCache.workouts)
        workoutCache.RUnlock()
        return
    }
    workoutCache.RUnlock()

    var workouts []Workout
    if err := db.Find(&workouts).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    workoutCache.Lock()
    workoutCache.workouts = workouts
    workoutCache.lastUpdate = time.Now()
    workoutCache.Unlock()

    c.JSON(http.StatusOK, workouts)
}

func getWorkoutByID(c *gin.Context) {
    id := c.Param("id")
    var workout Workout
    if err := db.Where("id = ?", id).First(&workout).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, workout)
    }
}

func updateWorkout(c *gin.Context) {
    id := c.Param("id")
    var workout Workout
    if db.Where("id = ?", id).First(&workout).RecordNotFound() {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    var input Workout
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Model(&workout).Updates(input)

    workoutCache.Lock()
    workoutCache.lastUpdate = time.Time{}
    workoutCache.Unlock()

    c.JSON(http.StatusOK, workout)
}

func deleteWorkout(c *gin.Context) {
    id := c.Param("id")
    if err := db.Where("id = ?", id).Delete(&Workout{}).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    workoutCache.Lock()
    workoutCache.lastUpdate = time.Time{}
    workoutCache.Unlock()

    c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}

func main() {
    initDB()
    r := gin.Default()

    r.POST("/workouts", createWorkout)
    r.GET("/workouts", getWorkouts)
    r.GET("/workouts/:id", getWorkoutByID)
    r.PUT("/workouts/:id", updateWorkout)
    r.DELETE("/workouts/:id", deleteWorkout)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}