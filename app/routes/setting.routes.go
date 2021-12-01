package routes

import (
	"fmt"
	"todo/app/services"

	"github.com/gin-gonic/gin"
)

func RegisterSettingRoutes(e *gin.Engine, s services.Services) {
	e.GET("/health", func(c *gin.Context) {
		response := s.SettingService.GetHealth()
		fmt.Println(response)
		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}
	})
}
