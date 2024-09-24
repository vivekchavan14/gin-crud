package main

import (
	"test/initializers"
	"test/models"
)

func init() {
	initializers.GetEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Employee{})
}
