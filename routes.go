package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"yourapp/controllers"
)

func initializeRouter() *gin.Engine {
	router := gin.Default()

	workoutGroup := router.Group("/workouts")
	{
		workoutGroup.GET("/", controllers.FetchAllWorkouts)
		workoutGroup.GET("/:id", controllers.FetchWorkoutByID)
		workoutGroup.POST("/", controllers.AddNewWorkout)
		workoutGroup.PUT("/:id", controllers.ModifyWorkout)
		workoutGroup.DELETE("/:id", controllers.RemoveWorkout)
	}

	return router
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}

	router := initializeRouter()

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	router.Run(":" + serverPort)
}