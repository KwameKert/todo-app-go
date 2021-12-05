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

func RegisterTaskRoutes(e *gin.Engine, s services.Services) {
	e.POST("/tasks", func(c *gin.Context) {
		var req core.CreateTaskRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		response := s.TaskService.CreateTask(req)

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

	e.PUT("/tasks", func(c *gin.Context) {
		var req models.Task

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		response := s.TaskService.UpdateTask(req)

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)
	})

	e.GET("/tasks", func(c *gin.Context) {

		response := s.TaskService.FetchAllTasks()

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)

	})

	e.GET("/tasks/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		response := s.TaskService.GetTaskById(id)

		if response.Error {
			c.JSON(response.Code, gin.H{
				"message": response.Meta.Message,
			})
			return
		}

		c.JSON(response.Code, response.Meta)

	})
}
