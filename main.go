package main

import (
	"test/controllers"
	"test/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.GetEnv()
	initializers.ConnectDB()
}

func main() {

	router := gin.Default()
	router.GET("/", controllers.GetEmployee)
	router.POST("/create", controllers.PostEmployee)
	router.GET("/:id", controllers.GetEmpByID)
	router.DELETE("/:id", controllers.DeleteEmployee)
	router.PUT("/:id", controllers.UpdateEmployee)
	router.Run()
}
