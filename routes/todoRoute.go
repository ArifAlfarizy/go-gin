package routes

import (
	"gin-test/internal/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todo")
	{
		todoGroup.GET("/", controllers.GetAllTodos)
		todoGroup.GET("/:id", controllers.GetTodoById)
		todoGroup.POST("/", controllers.CreateTodo)
		todoGroup.PATCH("/:id", controllers.UpdateTodo)
		todoGroup.DELETE("/:id", controllers.DeleteTodo)
	}
}
