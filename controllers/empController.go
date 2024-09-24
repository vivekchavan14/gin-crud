package controllers

import (
	"errors"
	"log"
	"test/initializers"
	"test/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetEmployee(ctx *gin.Context) {
	var emp []models.Employee
	initializers.DB.Find(&emp)
	ctx.JSON(200, gin.H{
		"employees": emp,
	})
}

func PostEmployee(ctx *gin.Context) {
	var body struct {
		Name  string
		Email string
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to bind request body"})
		return
	}

	employee := models.Employee{Name: body.Name, Email: body.Email}

	result := initializers.DB.Create(&employee)

	if result.Error != nil {
		log.Println("Employee creation failed:", result.Error)
		ctx.JSON(500, gin.H{"error": "Failed to create employee"})
		return
	}
	ctx.JSON(200, gin.H{
		"emp": employee,
	})
}

func GetEmpByID(ctx *gin.Context) {
	var employee models.Employee
	id := ctx.Param("id")
	// Query the database for the first employee that matches the ID
	if err := initializers.DB.First(&employee, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "Employee not found"})
		} else {
			ctx.JSON(500, gin.H{"error": "Database error"})
		}
		return
	}

	// Return the employee data as a JSON response
	ctx.JSON(200, employee)
}

func DeleteEmployee(ctx *gin.Context) {
	var employee models.Employee
	id := ctx.Param("id")
	if err := initializers.DB.Delete(&employee, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(400, gin.H{"error": "Employee not found"})
		} else {
			ctx.JSON(500, gin.H{"error": "Database error"})
		}
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "Successfully deleted",
	})
}

func UpdateEmployee(ctx *gin.Context) {
	var employee models.Employee
	var body struct {
		Name  string
		Email string
	}
	ctx.Bind(&body)
	id := ctx.Param("id")
	initializers.DB.First(&employee, id)
	if err := initializers.DB.Model(&employee).Updates(models.Employee{Name: body.Name, Email: body.Email}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(400, gin.H{"error": "Employee not found"})
		} else {
			ctx.JSON(500, gin.H{"error": "Database error"})
		}
		return
	}

	ctx.JSON(200, gin.H{
		"employee": employee,
	})
}
