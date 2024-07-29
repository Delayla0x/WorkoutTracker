package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"yourapp/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	workoutRoutes := r.Group("/workouts")
	{
		workoutRoutes.GET("/", controllers.GetWorkouts)
		workoutRoutes.GET("/:id", controllers.GetWorkoutByID)
		workoutRoutes.POST("/", controllers.CreateWorkout)
		workoutRoutes.PUT("/:id", controllers.UpdateWorkout)
		workoutRoutes.DELETE("/:id", controllers.DeleteWorkout)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}