package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"yourapp/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	workoutRoutes := r.Group("/workouts")
	{
		workoutRoutes.GET("/", controllers.GetWorkouts)
		workoutRoutes.GET("/:id", controllers.GetWorkoutByID)
		workoutRoutes.POST("/", controllers.CreateWorkout)
		workoutRoutes.PUT("/:id", controllers.UpdateWorkout)
		workoutRoutes.DELETE("/:id", controllers.DeleteWorkout)
	}

	return r
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	r := setupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}