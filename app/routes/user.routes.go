package routes

import (
	"net/http"
	"todo/app/models"
	"todo/app/services"

	"github.com/gin-gonic/gin"

	//	"todo/app/utils"
	"strconv"
	"todo/core"
	//	log "github.com/sirupsen/logrus"
)

func RegisterUserRoutes(e *gin.Engine, s services.Services) {
	e.POST("/users", func(c *gin.Context) {
		var req core.CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		response := s.UserService.CreateUser(req)

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

	e.PUT("/users", func(c *gin.Context) {
		var req models.User

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		response := s.UserService.UpdateUser(req)

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

	e.GET("/users", func(c *gin.Context) {

		response := s.UserService.FetchUsers()

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

	e.GET("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		response := s.UserService.GetUser(id)
		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}
		c.JSON(response.Code, response.Meta)
	})

	e.DELETE("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		response := s.UserService.DeleteUser(id)
		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}
		c.JSON(response.Code, response.Meta)
	})

	e.GET("/users/:id/tasks", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		response := s.TaskService.FetchAllUserTasks(id)
		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}
		c.JSON(response.Code, response.Meta)
	})
}
