package services

import (
	"gin-test/config"
	"gin-test/internal/models"

	"gorm.io/gorm"
)

// Query to database

func GetAllTodos() ([]models.Todo, error) { // Expecting result todo models or an error
	var todos []models.Todo

	result := config.DB.Find(&todos)

	return todos, result.Error
}

func GetTodoById(id uint) (models.Todo, error) {
	var todo models.Todo
	result := config.DB.First(&todo, id)
	return todo, result.Error
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	result := config.DB.Create(&todo)

	return todo, result.Error
}

func UpdateTodo(id uint, todo models.Todo) (models.Todo, error) {
	config.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(todo)
	var updated models.Todo
	result := config.DB.First(&updated, id)
	return updated, result.Error
}

func DeleteTodo(id uint) error { // Find models.todo, with id: then delete
	result := config.DB.Delete(&models.Todo{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
