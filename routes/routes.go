package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Placeholder: tambahkan route lain di sini nanti
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "HRIS API is running 🚀"})
	})
}
