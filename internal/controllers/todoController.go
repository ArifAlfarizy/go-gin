package controllers

import (
	"gin-test/internal/dto"
	"gin-test/internal/models"
	"gin-test/internal/services"

	"github.com/gin-gonic/gin"
)

// temporary, need to separate query and logic layers

func CreateTodo(c *gin.Context) {
	// Get data off req body

	// req is DTO type
	var req dto.CreateTodoRequest

	// Bind req json to req
	if err := c.BindJSON(&req); err != nil {
		c.JSON((400), gin.H{
			"error": err.Error(),
		})
		return
	}

	// Mapping, req is being used here
	todo := models.Todo{
		Title:     req.Title,
		Completed: req.Completed,
	}

	// Call query(service)
	todo, err := services.CreateTodo(todo)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	c.JSON(201, todo)
}
