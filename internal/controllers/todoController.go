package controllers

import (
	"errors"
	"gin-test/internal/dto"
	"gin-test/internal/models"
	"gin-test/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Use http instead of code error

func GetAllTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, todos)

}

func GetTodoById(c *gin.Context) {
	// Get id from param
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := services.GetTodoById(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)

}

func CreateTodo(c *gin.Context) {
	// Get data off req body

	// req is DTO type
	var req dto.CreateTodoRequest

	// Bind req json to req
	if err := c.BindJSON(&req); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusCreated, todo)
}
