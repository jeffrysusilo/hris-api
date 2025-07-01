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

	attendance := router.Group("/attendance")
	{
		attendance.POST("/checkin", controllers.CheckIn)
		attendance.POST("/checkout/:id", controllers.CheckOut)
		attendance.GET("/employee/:id", controllers.GetEmployeeAttendance)
	}

	payroll := router.Group("/payroll")
	{
		payroll.POST("/generate", controllers.GeneratePayroll)
		payroll.GET("/employee/:id", controllers.GetPayrollHistory)
	}
}
