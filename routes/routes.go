package routes

import (
	"github.com/gin-gonic/gin"
	"hris-api/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "HRIS API is running 🚀"})
	})

	employee := router.Group("/employees")
	{
		employee.POST("", controllers.CreateEmployee)
		employee.GET("", controllers.GetAllEmployees)
		employee.GET(":id", controllers.GetEmployeeByID)
		employee.PUT(":id", controllers.UpdateEmployee)
		employee.DELETE(":id", controllers.DeleteEmployee)
	}
}
