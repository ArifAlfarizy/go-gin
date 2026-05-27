package routes

import (
	"gin-test/internal/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todo")
	{
		todoGroup.POST("/", controllers.CreateTodo)
	}
}
