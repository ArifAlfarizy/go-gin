package services

import (
	"gin-test/config"
	"gin-test/internal/models"
)

// Query to database


func GetAllTodos() ([]models.Todo, error) { // Expecting result todo models or an error
	var todos []models.Todo

	result := config.DB.Find(&todos)

	return todos, result.Error
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	result := config.DB.Create(&todo)

	return todo, result.Error
}

func DeleteTodo(id uint) error {	// Find models.todo, with id: then delete
	result := config.DB.Delete(&models.Todo{}, id)

	return result.Error
}